package main

import (
    "log"
	"fmt"
	"os"
	"context"
	"cohesivenet"
)


// func retryRequest(callapi func() (interface{}, interface{}, error), attempts int, sleep int) (interface{}, error) {
//     resp, _, err := callapi()
    
//     if err != nil {
//         if _, reason := checkHttpError(err); reason == "timeout" {
//             time.Sleep(time.Duration(sleep) * time.Second)
//             return retryRequest(callapi, (attempts - 1), sleep)
//         } else {
//             return false, err
//         }
//     }

//     return resp, nil
// }


// 1. set topology name (can be done on licensing or PUT config)
// 2. Uploadlicense
// 3. set license parameters. wait
// 4. generate keyset
// 5. set peering id

func setupController(vns3 *cohesivenet.VNS3Client, ctx context.Context) (*cohesivenet.ConfigDetail, error) {
    // set_topology_req := cohesivenet.NewUpdateConfigRequest()
    // set_topology_req.SetTopologyName("TEST Topo")
    // set_topology_req.SetControllerName("Ctrl 1")
    // api_set_topology_req := vns3.ConfigurationApi.PutConfigRequest(ctx)
    // api_set_topology_req = api_set_topology_req.UpdateConfigRequest(*set_topology_req)
    // log.Println("Setting VNS3 Config")
    // _, _, err := vns3.ConfigurationApi.PutConfig(api_set_topology_req)
    // if err != nil {
    //     log.Printf("ERROR: Configuration: PutConfig: %v\n", err)
    //     return nil, err
    // }

    // license := "/Users/benplatta/code/cohesive/vns3-functional-testing/test-assets/license.txt"
    // licenseFile, err := os.Open(license)
    // if err != nil {
    //     log.Printf("ERROR: Failed to read file: %v\n", license)
    //     return nil, fmt.Errorf("ERROR: Failed to read file: %v\n", license)
    // }

    // log.Println("Uploading License")
    // upload_req := vns3.ConfigurationApi.UploadLicenseRequest(ctx)
    // upload_req = upload_req.Body(licenseFile)
    // upload_resp, _, err := upload_req.Execute()
    // if err != nil {
    //     log.Printf("ERROR: Upload License Error: %+v\n", err)
    //     return nil, fmt.Errorf("ERROR: Upload License Error: %+v\n", err)
    // } else {
    //     d, _ := json.Marshal(upload_resp.Response)
    //     log.Printf("License uploaded %v", string(d))
    // }

    // license_params := cohesivenet.NewSetLicenseParametersRequest()
    // license_params.SetDefault(true)
    // set_params_req := vns3.ConfigurationApi.PutSetLicenseParametersRequest(ctx)
    // set_params_req = set_params_req.PutLicenseParametersRequest(*license_params)
    // log.Println("Setting license params")
    // params_resp, _, err := vns3.ConfigurationApi.PutSetLicenseParameters(set_params_req)
    // if err != nil {
    //     return nil, fmt.Errorf("ERROR: Set License params Error: %+v\n", err)
    // } else {
    //     log.Println("Set license params")
    //     d, _ := json.Marshal(params_resp.Response)
    //     log.Printf(string(d))
    // }

    // log.Printf("Waiting for VNS3 reboot")
    // _, err = waitForDown(vns3, 180, 2)
    // if err != nil {
    //     return nil, fmt.Errorf("ERROR: VNS3 Reboot: %+v\n", err)
    // }
    // log.Printf("VNS3 Down. Polling API for ready")
    configDetail, err := vns3.ConfigurationApi.WaitForApi(&ctx, 60*5, 3, 10)
    
    // if err != nil {
    //     return nil, fmt.Errorf("ERROR: VNS3 Polling: %+v\n", err)
    // }

    _, err = vns3.ConfigurationApi.WaitForKeyset(&ctx, 60, 2, false)
    if err != nil {
        return nil, fmt.Errorf("ERROR: Keyset Polling (1): %+v\n", err)
    }

    apiKeysetRequest := vns3.ConfigurationApi.PutKeysetRequest(ctx)
    apiKeysetRequest = apiKeysetRequest.UpdateKeysetRequest(*cohesivenet.NewUpdateKeysetRequest("test"))
    keysetDetail, _, err := apiKeysetRequest.Execute()
    if err != nil {
        return nil, fmt.Errorf("ERROR: Keyset Put: %+v\n", err)
    } else {
        d, _ := keysetDetail.Response.MarshalJSON()
        log.Println("Keyset Detail:", string(d))
    }
    _, err = vns3.ConfigurationApi.WaitForKeyset(&ctx, 60*5, 2, true)
    if err != nil {
        return nil, fmt.Errorf("ERROR: Keyset Polling. Is Present (2): %+v\n", err)
    }

    return configDetail, nil
}

func main() {
    // auth := context.WithValue(context.Background(), cohesivenet.ContextBasicAuth, cohesivenet.BasicAuth{
    //     UserName: "api",
    //     Password: "testtest1234",
    // })

    // token := "eb0431ded2351eff780f30a81d01786941a8af58975e7d9c4b0cc1b0fcc337f8"
    // auth := context.WithValue(context.Background(), cohesivenet.ContextAccessToken, token)
    host := "18.212.252.129"
    ps := "i-00d0d0483ff899734"
    // config := cohesivenet.NewConfigurationWithAuth("3.222.246.246", cohesivenet.ContextAccessToken, token)
    config := cohesivenet.NewConfigurationWithAuth(host, cohesivenet.ContextBasicAuth, cohesivenet.BasicAuth{
        UserName: "api",
        Password: ps,
    })
    vns3 := cohesivenet.NewVNS3Client(config, cohesivenet.ClientParams{
        Timeout: 3,
        TLS: false,
    })

    // start := time.Now()
    ctx := context.Background()
    // req := vns3.ConfigurationApi.GetConfigRequest(ctx)
    // r, _, err := vns3.ConfigurationApi.GetConfig(req)
    // time.Sleep(5 * time.Second)
    // elapsed := (time.Now()).Sub(start)
    // var timeout float64 = 3
    // log.Println(elapsed)
    // log.Println(elapsed.Seconds() >= timeout)


    _, err :=  vns3.ConfigurationApi.WaitForApi(&ctx, 30, 2, 5)
    if err != nil {
        log.Printf("VNS3 connectivity error: %+v", err)
        os.Exit(1)
    }

    configDetail, setupErr := setupController(vns3, ctx)

    if setupErr != nil {
        log.Printf("VNS3 Setup error: %+v", setupErr)
    } else {
        d, _ := configDetail.MarshalJSON()
        log.Printf("Setup success: %v", string(d))
    }

    // req := vns3.RoutingApi.GetRoutesRequest(context.Background())
    // r, _, err := vns3.RoutingApi.GetRoutes(req)

    // if err != nil {
    //     failed, message := checkHttpError(err)
    //     log.Printf("Failed %v: %v\n", failed, message)
    // }

    // if setupErr != nil {
    //     log.Printf("Failed setup: %v\n", setupErr)
    // }
}