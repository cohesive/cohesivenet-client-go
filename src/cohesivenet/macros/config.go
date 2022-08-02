package macros

import (
	"fmt"
	"context"
	"os"
	cohesivenet "github.com/cohesive/cohesivenet-client-go"
)

type SetupRequest struct {
	TopologyName string
	ControllerName string
	LicenseParams *cohesivenet.SetLicenseParametersRequest
	LicenseFile string
	PeerId int
	KeysetParams cohesivenet.SetKeysetParamsRequest
	WaitTimeout int
	KeysetTimeout int
}


func SetupController(vns3 *cohesivenet.VNS3Client, setupRequest SetupRequest) (*cohesivenet.ConfigDetail, error) {
	ctx := context.Background()
    set_topology_req := cohesivenet.NewUpdateConfigRequest()
    set_topology_req.SetTopologyName(setupRequest.TopologyName)
    set_topology_req.SetControllerName(setupRequest.ControllerName)
    api_set_topology_req := vns3.ConfigurationApi.PutConfigRequest(ctx)
    api_set_topology_req = api_set_topology_req.UpdateConfigRequest(*set_topology_req)
    vns3.log.Debug("Setting VNS3 Config")
    _, _, err := vns3.ConfigurationApi.PutConfig(api_set_topology_req)
    if err != nil {
        vns3.log.Debug(fmt.Sprintg("ERROR: Configuration: PutConfig: %v\n", err))
        return nil, err
    }

    licenseFile, err := os.Open(setupRequest.LicenseFile)
    if err != nil {
        vns3.log.Error(fmt.Sprintf("ERROR: Failed to read file: %v\n", setupRequest.LicenseFile))
        return nil, fmt.Errorf("ERROR: Failed to read file: %v\n", setupRequest.LicenseFile)
    }

    vns3.log.Debug("Uploading License")
    upload_req := vns3.ConfigurationApi.UploadLicenseRequest(ctx)
    upload_req = upload_req.Body(licenseFile)
    upload_resp, _, err := upload_req.Execute()
    if err != nil {
        vns3.log.Error(fmt.Sprintf("ERROR: Upload License Error: %+v\n", err))
        return nil, fmt.Errorf("ERROR: Upload License Error: %+v\n", err)
    } else {
        d, _ := upload_resp.MarshalJSON()
        vns3.log.Error(fmt.Sprintf("License uploaded %v", string(d)))
    }

	var licenseRequest *cohesivenet.SetLicenseParametersRequest
	if setupRequest.LicenseParams != nil {
		licenseRequest = setupRequest.LicenseParams
	} else {
		licenseRequest = cohesivenet.NewSetLicenseParametersRequest()
		licenseRequest.SetDefault(true)
	}

    setParamsReq := vns3.ConfigurationApi.PutSetLicenseParametersRequest(ctx)
    setParamsReq = setParamsReq.PutLicenseParametersRequest(*licenseRequest)
    vns3.log.Debug(fmt.Sprintf("Setting license params"))
    params_resp, _, err := vns3.ConfigurationApi.PutSetLicenseParameters(setParamsReq)
    if err != nil {
        return nil, fmt.Errorf("ERROR: Set License params Error: %+v\n", err)
    } else {
        d, _ := params_resp.Response.MarshalJSON()
        vns3.log.Debug(fmt.Sprintf("Set license params %+v", string(d)))
    }

    vns3.log.Debug("Waiting for VNS3 reboot")
    _, err = waitForDown(vns3, 180, 2)
    if err != nil {
        return nil, fmt.Errorf("ERROR: VNS3 Reboot: %+v\n", err)
    }
    vns3.log.Debug("VNS3 Down. Polling API for ready")
	// 60*
    configDetail, err := vns3.ConfigurationApi.WaitForApi(&ctx, setupRequest.WaitTimeout, 2, 10)
    
    // if err != nil {
    //     return nil, fmt.Errorf("ERROR: VNS3 Polling: %+v\n", err)
    // }

    _, err = vns3.ConfigurationApi.WaitForKeyset(&ctx, 30, 2, false)
    if err != nil {
        return nil, fmt.Errorf("ERROR: Keyset Polling (1): %+v\n", err)
    }

    apiKeysetRequest := vns3.ConfigurationApi.PutKeysetRequest(ctx)
    apiKeysetRequest = apiKeysetRequest.SetKeysetParamsRequest(*setupRequest.KeysetParams)
    keysetDetail, _, err := apiKeysetRequest.Execute()
    if err != nil {
        return nil, fmt.Errorf("ERROR: Keyset Put: %+v\n", err)
    } else {
        d, _ := keysetDetail.Response.MarshalJSON()
        vns3.log.Info(fmt.Sprintf("Keyset Detail: %+v", string(d)))
    }
	// 60* 5
    _, err = vns3.ConfigurationApi.WaitForKeyset(&ctx, setupRequest.KeysetTimeout, 2, true)
    if err != nil {
        return nil, fmt.Errorf("ERROR: Keyset Polling. Is Present (2): %+v\n", err)
    }

    return configDetail, nil
}