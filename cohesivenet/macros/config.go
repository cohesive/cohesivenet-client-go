package macros

import (
	"fmt"
	"context"
	"os"
    "time"
    "strings"
    cn "github.com/cohesive/cohesivenet-client-go/cohesivenet"
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
    _, err = vns3.ConfigurationApi.WaitForDown(&ctx, 60*5, 2)
    if err != nil {
        return nil, fmt.Errorf("ERROR: VNS3 Reboot: %+v", err)
    }
    vns3.Log.Debug("VNS3 Down. Polling API for ready")
	// 60*5
    return vns3.ConfigurationApi.WaitForApi(&ctx, waitTimeout, 2, 10)
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


func TrySetKeyset(vns3 *cn.VNS3Client, keysetParams cn.SetKeysetParamsRequest, keysetTimeout int) (*cn.KeysetStatus, bool, error) {
    ctx := context.Background()

    keysetDetail, _, connecterr := vns3.ConfigurationApi.GetKeyset(vns3.ConfigurationApi.GetKeysetRequest(ctx))
    if connecterr != nil {
        return nil, false, fmt.Errorf("Set keyset failed. Error: %v", connecterr.Error())
    }

    keysetStatus := keysetDetail.GetResponse()
    var inProgress bool = keysetStatus.GetInProgress()
    var keysetPresent bool = keysetStatus.GetKeysetPresent()
    if keysetPresent {
        vns3.Log.Debug("Keyset already present. Skipping.")
        return &keysetStatus, false, nil
    }

    if !inProgress {
        vns3.Log.Debug("Keyset not set and not in progress. Setting.")
        apiKeysetRequest := vns3.ConfigurationApi.PutKeysetRequest(ctx)
        apiKeysetRequest = apiKeysetRequest.SetKeysetParamsRequest(keysetParams)
        keysetDetail, _, err := apiKeysetRequest.Execute()
        if err != nil {
            return nil, false, fmt.Errorf("ERROR: Keyset Put Failure. %+v", err)
        } else {
            d, _ := keysetDetail.Response.MarshalJSON()
            vns3.Log.Debug(fmt.Sprintf("Set Keyset Response: %+v", string(d)))
        }
    }

    vns3.Log.Debug("Waiting for keyset")
    keysetStatus2, err := vns3.ConfigurationApi.WaitForKeyset(&ctx, keysetTimeout, 2, true)
    if err != nil {
        return nil, true, fmt.Errorf("ERROR: Keyset Polling Failure. %+v", err)
    }

    return keysetStatus2, true, nil
}


func FetchKeysetFromSource(vns3 *cn.VNS3Client, source string, token string, waitTimeout int) (*cn.KeysetStatus, error) {
    // """fetch_keyset_from_source Put keyset by providing source controller to download keyset. This
    // contains logic that handles whether or not fetching from the source fails, typically due
    // to a firewall or routing issue in the underlay network (e.g. security groups and route tables).

    // Pseudo-logic:
    //     PUT new keyset request to fetch from remote controller
    //     if keyset exists or already in progress, fail immediately as its unexpected
    //     if PUT succees:
    //         wait:
    //             if a new successful put returns: that indicates failure to download from source. return 400
    //             if timeout: that indicates controller is rebooting, return wait for keyset
    //             if keyset already exists, wait to ensure keyset  exists, then return keyset details

    // Arguments:
    //     source {str} - host controller to fetch keyset from
    //     token {str} - secret token used when generating keyset
    //     wait_timeout {float} - timeout for waiting for keyset and while polling for download failure (default: 1 min)
    //     allow_exists {bool} - If true and keyset already exists, DONT throw exception
    // """
    // sleep_time = 2.0
    // failure_error_str = (
    //     "Failed to fetch keyset for source. This typically due to a misconfigured "
    //     "firewall or routing issue between source and client controllers."
    // )

    ctx := context.Background()
    keysetParams := cn.NewSetKeysetParamsRequest(token)
    keysetParams.SetSource(source)
    apiKeysetRequest := vns3.ConfigurationApi.PutKeysetRequest(ctx)
    apiKeysetRequest = apiKeysetRequest.SetKeysetParamsRequest(*keysetParams)
    fetchFailureMessage := `Failed to fetch keyset for source. This typically due to a misconfigured firewall 
    or routing issue between source and client controllers.`

    // initial put keyset request
    keysetDetail, httpResp, err := vns3.ConfigurationApi.PutKeyset(apiKeysetRequest)
    if err != nil {
        // handle in progress?
        vns3.Log.Error(fmt.Sprintf("Failed to set keyset: %+v", err.Error()))
        return nil, err
    }

    vns3.Log.Info(fmt.Sprintf("Status code: %v", httpResp.StatusCode))
    keysetStatus, hasResponse := keysetDetail.GetResponseOk()
    if !hasResponse {
        keysetDetail, _, err := vns3.ConfigurationApi.GetKeyset(vns3.ConfigurationApi.GetKeysetRequest(ctx))
        if err != nil {
            return nil, fmt.Errorf("Empty set keyset response. Error: %v", err.Error())
        }

        keysetStatus := keysetDetail.GetResponse()
        if keysetStatus.GetKeysetPresent() {
            return nil, fmt.Errorf("keyset already exists")
        }

        return nil, fmt.Errorf("PUT keyset returned null response")
    }

    initTime := keysetStatus.GetStartedAtI()
    sleep := 2
    pollStart := time.Now()
    elapsed := time.Now().Sub(pollStart).Seconds()
    for elapsed < float64(waitTimeout) {
        vns3.Log.Debug("Sending set keyset request check")
        keysetDetail, _, err := vns3.ConfigurationApi.PutKeyset(apiKeysetRequest)

        if err != nil {
            isUnavailable, _ := cn.CheckHttpErrorUnavailable(err)
            // for nested api and keyset waits. Technically we can go over waitTimeout as we dont recalc
            // remainingTime after api wait
            remainingTime := waitTimeout - int(elapsed)

            if isUnavailable {
                // this likely means controller is rebooting
                vns3.Log.Debug("Keyset fetch: connection error. VNS3 is likely rebooting. Waiting for API.")
                vns3.ConfigurationApi.WaitForApi(&ctx, remainingTime, 2, 10)
                keysetStatus, _ := vns3.ConfigurationApi.WaitForKeyset(&ctx, remainingTime, 2, true)
                return keysetStatus, nil
            } else {
                apiError, isApiError := err.(*cn.GenericAPIError)
                if isApiError {
                    message := strings.ToLower(apiError.GetErrorMessage())
                    if strings.Contains(message, "in progress") {
                        vns3.Log.Debug("Keyset fetch in progress")
                        // keyset generation in progress but could fail - so we retry
                        // pass - fall through to sleep
                    } else if strings.Contains(message, "already exists") {
                        vns3.Log.Debug("Keyset already exists. Waiting for VNS3 reboot and keyset")
                        // keyset already exists, wait for vns3 reboot and keyset
                        _, err = vns3.ConfigurationApi.WaitForDown(&ctx, 180, 2)
                        if err != nil {
                            return nil, fmt.Errorf("VNS3 failed to go down for reboot: %v", err.Error())
                        }

                        remainingTime = remainingTime - 180
                        vns3.ConfigurationApi.WaitForApi(&ctx, remainingTime, 2, 5)
                        keysetStatus, _ := vns3.ConfigurationApi.WaitForKeyset(&ctx, remainingTime, 2, true)
                        return keysetStatus, nil
                    } else {
                        // unknown error
                        return nil, apiError
                    }
                } else {
                    return nil, fmt.Errorf("Unexpected PUT keyset error: %v", err.Error())
                }
            }
        }

        keysetStatus, hasStatus := keysetDetail.GetResponseOk()
        if hasStatus {
            curStartTime := keysetStatus.GetStartedAtI()
            if curStartTime != initTime {
                return nil, fmt.Errorf(fetchFailureMessage)
            }
        }

        time.Sleep(time.Duration(sleep) * time.Second)
        elapsed = time.Now().Sub(pollStart).Seconds()
    }

    return nil, fmt.Errorf("Timeout fetching keyset from source %v", source)

    // try:
    //     put_response = client.config.put_keyset(**{"source": source, "token": token})
    // except ApiException as e:
    //     if allow_exists and ("keyset already exists" in e.get_error_message().lower()):
    //         Logger.info("Keyset already exists.", host=client.host_uri)
    //         return client.config.try_get_keyset()

    //     Logger.info(
    //         "Failed to fetch keyset: %s" % e.get_error_message(),
    //         host=client.host_uri,
    //     )
    //     raise e
    // except UrlLib3ConnExceptions:
    //     raise ApiException(
    //         status=HTTPStatus.SERVICE_UNAVAILABLE,
    //         reason="Controller unavailable. It is likely rebooting. Try client.sys_admin.wait_for_api().",
    //     )

    // if not put_response.response:
    //     keyset_data = client.config.get_keyset()
    //     if keyset_data.response and keyset_data.response.keyset_present:
    //         raise ApiException(status=400, reason="Keyset already exists.")
    //     raise ApiException(status=500, reason="Put keyset returned None.")

    // start_time = put_response.response.started_at_i
    // Logger.info(message="Keyset downloading from source.", start_time=start_time)
    // polling_start = time.time()
    // while time.time() - polling_start <= wait_timeout:
    //     try:
    //         duplicate_call_resp = client.config.put_keyset(
    //             **{"source": source, "token": token}
    //         )
    //     except UrlLib3ConnExceptions:
    //         Logger.info(
    //             "API call timeout. Controller is likely rebooting. Waiting for keyset.",
    //             wait_timeout=wait_timeout,
    //             source=source,
    //         )
    //         client.sys_admin.wait_for_api(timeout=wait_timeout, wait_for_reboot=True)
    //         return client.config.wait_for_keyset(timeout=wait_timeout)
    //     except ApiException as e:
    //         duplicate_call_error = e.get_error_message()

    //         if duplicate_call_error == "Keyset already exists.":
    //             keyset_data = client.config.try_get_keyset()
    //             if not keyset_data:
    //                 Logger.info(
    //                     "Keyset exists. Waiting for reboot.",
    //                     wait_timeout=wait_timeout,
    //                     source=source,
    //                 )
    //                 client.sys_admin.wait_for_api(
    //                     timeout=wait_timeout, wait_for_reboot=True
    //                 )
    //                 return client.config.wait_for_keyset()
    //             return keyset_data

    //         if duplicate_call_error == "Keyset setup in progress.":
    //             # this means download is in progress, but might fail. Wait and retry
    //             time.sleep(sleep_time)
    //             continue

    //         # Unexpected ApiException
    //         raise e

    //     # If there is a new start time for keyset generation, that indicates a failure
    //     new_start_resp = duplicate_call_resp.response
    //     new_start = new_start_resp.started_at_i if new_start_resp else None
    //     if (new_start and start_time) and (new_start != start_time):
    //         Logger.error(failure_error_str, source=source)
    //         raise ApiException(status=HTTPStatus.BAD_REQUEST, reason=failure_error_str)

    //     time.sleep(sleep_time)
    // raise CohesiveSDKException("Timeout while waiting for keyset.")

    // return nil, nil
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
    keysetStatus, _, err := TrySetKeyset(vns3, setupRequest.KeysetParams, setupRequest.KeysetTimeout)
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