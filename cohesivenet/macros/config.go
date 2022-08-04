package macros

import (
	"fmt"
	"context"
	"os"
    cn "cohesivenet"
)

type SetupRequest struct {
	TopologyName string
	ControllerName string
	LicenseParams *cn.SetLicenseParametersRequest
	LicenseFile string
	PeerId int32
	KeysetParams cn.SetKeysetParamsRequest
	WaitTimeout int
	KeysetTimeout int
}

func LicenseController(vns3 *cn.VNS3Client, licensePath string, licenseParams *cn.SetLicenseParametersRequest, waitTimeout int) (*cn.ConfigDetail, error) {
    ctx := context.Background()

    licenseFile, err := os.Open(licensePath)
    if err != nil {
        vns3.Log.Error(fmt.Sprintf("ERROR: Failed to read file: %v", licensePath))
        return nil, fmt.Errorf("ERROR: Failed to read file: %v", licensePath)
    }

    vns3.Log.Debug("Uploading License")
    upload_req := vns3.ConfigurationApi.UploadLicenseRequest(ctx)
    upload_req = upload_req.Body(licenseFile)
    upload_resp, _, err := upload_req.Execute()
    if err != nil {
        vns3.Log.Error(fmt.Sprintf("ERROR: Upload License Error: %+v", err))
        return nil, fmt.Errorf("ERROR: Upload License Error: %+v", err)
    } else {
        d, _ := upload_resp.MarshalJSON()
        vns3.Log.Error(fmt.Sprintf("License uploaded %v", string(d)))
    }

	var licenseRequest *cn.SetLicenseParametersRequest
	if licenseParams != nil {
		licenseRequest = licenseParams
	} else {
		licenseRequest = cn.NewSetLicenseParametersRequest(true)
	}

    setParamsReq := vns3.ConfigurationApi.PutSetLicenseParametersRequest(ctx)
    setParamsReq = setParamsReq.PutLicenseParametersRequest(*licenseRequest)
    vns3.Log.Debug(fmt.Sprintf("Setting license params"))
    params_resp, _, err := vns3.ConfigurationApi.PutSetLicenseParameters(setParamsReq)
    if err != nil {
        return nil, fmt.Errorf("ERROR: Set License params Error: %+v", err)
    } else {
        d, _ := params_resp.Response.MarshalJSON()
        vns3.Log.Debug(fmt.Sprintf("Set license params %+v", string(d)))
    }

    vns3.Log.Debug("Waiting for VNS3 reboot")
    _, err = vns3.ConfigurationApi.WaitForDown(&ctx, 180, 2)
    if err != nil {
        return nil, fmt.Errorf("ERROR: VNS3 Reboot: %+v", err)
    }
    vns3.Log.Debug("VNS3 Down. Polling API for ready")
	// 60*5
    return vns3.ConfigurationApi.WaitForApi(&ctx, waitTimeout, 2, 10)
}


func TrySetKeyset(vns3 *cn.VNS3Client, keysetParams cn.SetKeysetParamsRequest, keysetTimeout int) (*cn.KeysetStatus, error) {
    ctx := context.Background()
    keysetStatus, err := vns3.ConfigurationApi.WaitForKeyset(&ctx, 30, 2, false)
    if err != nil {
        return nil, fmt.Errorf("ERROR: Init keyset read failure. %+v", err)
    }

    var inProgress bool = keysetStatus.GetInProgress()
    var keysetPresent bool = keysetStatus.GetKeysetPresent()
    if keysetPresent {
        vns3.Log.Debug("Keyset already present. Skipping.")
        return keysetStatus, nil
    }

    if !inProgress {
        vns3.Log.Debug("Keyset not set and not in progress. Setting.")
        apiKeysetRequest := vns3.ConfigurationApi.PutKeysetRequest(ctx)
        apiKeysetRequest = apiKeysetRequest.SetKeysetParamsRequest(keysetParams)
        keysetDetail, _, err := apiKeysetRequest.Execute()
        if err != nil {
            return nil, fmt.Errorf("ERROR: Keyset Put Failure. %+v", err)
        } else {
            d, _ := keysetDetail.Response.MarshalJSON()
            vns3.Log.Debug(fmt.Sprintf("Set Keyset Response: %+v", string(d)))
        }
    }

	// 60* 5
    vns3.Log.Debug("Waiting for keyset")
    keysetStatus, err = vns3.ConfigurationApi.WaitForKeyset(&ctx, keysetTimeout, 2, true)
    if err != nil {
        return nil, fmt.Errorf("ERROR: Keyset Polling Failure. %+v", err)
    }

    return keysetStatus, nil
}

func TrySetPeering(vns3 *cn.VNS3Client, peerId int32) (*cn.PeersDetail, error) {
    ctx := context.Background()
    get_request := vns3.PeeringApi.GetPeeringStatusRequest(ctx)
    peeringResponse, _, err := vns3.PeeringApi.GetPeeringStatus(get_request)
    if err != nil {
        return nil, fmt.Errorf("Failed to get peering status %+v", err)
    }

    peeringStatus := peeringResponse.GetResponse()
    if peeringStatus.GetPeered() {
        currentId := peeringStatus.GetId()
        if currentId != peerId {
            return &peeringStatus, fmt.Errorf("Incorrect peer Id found. Expecting %+v, found %+v", peerId, currentId)
        }
        return &peeringStatus, nil
    }

    peer_request := cn.NewPeerSelfRequest(peerId)
    api_peer_request := vns3.PeeringApi.PutSelfPeeringIdRequest(ctx)
    api_peer_request = api_peer_request.PeerSelfRequest(*peer_request)
    peeringResponse, _, err = api_peer_request.Execute()
    if err != nil {
        return nil, fmt.Errorf("Failed to set peer id: %+v", err)
    }

    peeringStatus = peeringResponse.GetResponse()
    return &peeringStatus, nil
}

func SetupController(vns3 *cn.VNS3Client, setupRequest SetupRequest) (*cn.ConfigDetail, error) {
	ctx := context.Background()

    initConfig, err :=  vns3.ConfigurationApi.WaitForApi(&ctx, 30, 1, 5)
    if err != nil {
        return nil, fmt.Errorf("Failed to connect to VNS3: %+v", err)
    }

    set_topology_req := cn.NewUpdateConfigRequest()
    set_topology_req.SetTopologyName(setupRequest.TopologyName)
    set_topology_req.SetControllerName(setupRequest.ControllerName)
    api_set_topology_req := vns3.ConfigurationApi.PutConfigRequest(ctx)
    api_set_topology_req = api_set_topology_req.UpdateConfigRequest(*set_topology_req)
    vns3.Log.Debug("Setting VNS3 Config")
    _, _, err = vns3.ConfigurationApi.PutConfig(api_set_topology_req)
    if err != nil {
        vns3.Log.Debug(fmt.Sprintf("ERROR: Configuration: PutConfig: %v", err))
        return nil, fmt.Errorf("ERROR: Configuration: PutConfig: %v", err)
    }

    runtimeConfig := initConfig.GetResponse()

    if !runtimeConfig.GetLicensed() {
        _, err = LicenseController(vns3, setupRequest.LicenseFile, setupRequest.LicenseParams, setupRequest.WaitTimeout)
        if err != nil {
            return nil, fmt.Errorf("Failed to license controller: %+v", err)
        }
    } else {
        vns3.Log.Debug("VNS3 already licensed. Skipping.")
    }

    vns3.Log.Debug("Setting keyset")
    keysetStatus, err := TrySetKeyset(vns3, setupRequest.KeysetParams, setupRequest.KeysetTimeout)
    if err != nil {
        return nil, fmt.Errorf("Generate Keyset error: %+v", err)
    } else {
        keysetData, _ := keysetStatus.MarshalJSON()
        vns3.Log.Debug(fmt.Sprintf("Keysetstatus: %+v", string(keysetData)))
    }
 
    vns3.Log.Debug("Setting peering")
    _, err = TrySetPeering(vns3, setupRequest.PeerId)
    if err != nil {
        return nil, fmt.Errorf("Set peering error: %+v", err)
    } else {
        keysetData, _ := keysetStatus.MarshalJSON()
        vns3.Log.Debug(fmt.Sprintf("Keysetstatus: %+v", string(keysetData)))
    }

    configDetail, _, err := vns3.ConfigurationApi.GetConfig(vns3.ConfigurationApi.GetConfigRequest(ctx))
    return configDetail, nil
}