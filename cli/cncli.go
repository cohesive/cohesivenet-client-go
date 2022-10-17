package main

import (
    "log"
    "os"
    "context"
    // "fmt"
    cn "github.com/cohesive/cohesivenet-client-go/cohesivenet"

	// "github.com/cohesive/cohesivenet-client-go/cohesivenet/macros"
    // clientv1 "github.com/cohesive/cohesivenet-client-go/cohesivenet/v1"
)


// 1. set topology name (can be done on licensing or PUT config)
// 2. Uploadlicense
// 3. set license parameters. wait
// 4. generate keyset
// 5. set peering id



func main() {
    // Example basic auth
    // auth := context.WithValue(context.Background(), cn.ContextBasicAuth, cn.BasicAuth{
    //     UserName: "api",
    //     Password: "password",
    // })

    // Example token auth
    // token := "tokentokentoken"
    // auth := context.WithValue(context.Background(), cn.ContextAccessToken, token)
    host := "34.231.116.245"
    ps := "i-09e5074ff34cdc9a6"

    // config := cn.NewConfigurationWithAuth("3.222.246.246", cn.ContextAccessToken, token)
    config := cn.NewConfigurationWithAuth(host, cn.ContextBasicAuth, cn.BasicAuth{
        UserName: "api",
        Password: ps,
    })
    vns3 := cn.NewVNS3Client(config, cn.ClientParams{
        Timeout: 10,
        TLS: false,
    })

    ctx := context.Background()
    _, err :=  vns3.ConfigurationApi.WaitForApi(&ctx, 60*10, 2, 2)
    if err != nil {
        log.Printf("VNS3 not available: %v", err)
        os.Exit(1)
    }

    // start := time.Now()
    // req := vns3.ConfigurationApi.GetConfigRequest(ctx)
    // r, _, err := vns3.ConfigurationApi.GetConfig(req)
    // time.Sleep(5 * time.Second)
    // elapsed := (time.Now()).Sub(start)
    // var timeout float64 = 3
    // log.Println(elapsed)
    // log.Println(elapsed.Seconds() >= timeout)

    // newLink := cn.NewCreateLinkRequest(2, "BensAutoOpenpvn")
    // confBytes, err := os.ReadFile("openvpn.conf")
    // if err != nil {
    //     // fmt.Print(err)
    //     os.Exit(1)
    // }

    // newLink.SetConf(string(confBytes))
    // newLink.SetDescription("I hope this works")

    // req := vns3.OverlayNetworkApi.CreateLink(ctx)
    // req = req.CreateLinkRequest(*newLink)
    // detail, _, err := vns3.OverlayNetworkApi.CreateLinkExecute(req)

    req := vns3.OverlayNetworkApi.DeleteLinkTagRequest(ctx, "lnk-2-172_31_1_4", "test")
    detail, _, err := vns3.OverlayNetworkApi.DeleteLinkTag(req)

    if err != nil {
        apiError, isApiError := err.(*cn.GenericAPIError)
        var errorMessage string
        if isApiError {
            errorMessage = apiError.GetErrorMessage()
        }

        if errorMessage == "" {
            errorMessage = string(err.Error())
        }
        log.Printf("CLI Failure: %v", errorMessage)
        os.Exit(1)
    }

    // token := "testtest"
    // source := "10.0.1.87"
    // waitTimeoutSecs := 60*10
    // vns3.Log.Info("Running fetch keyset")
    // keysetDetail, err := macros.FetchKeysetFromSource(vns3, source, token, waitTimeoutSecs)
    // if err != nil {
    //     apiError, hasApiError := err.(*cn.GenericAPIError)
    //     vns3.Log.Error(fmt.Sprintf("CLI ERROR: %+v", err))
    //     if hasApiError {
    //         vns3.Log.Error(apiError.GetErrorMessage())
    //     }
    //     return
    // }

    // detail := *keysetDetail
    data, _ := detail.MarshalJSON()
    log.Printf("CLI success: %v", string(data))
    
    // setupReq := macros.SetupRequest{
    //     TopologyName: "Terraform",
    //     ControllerName: "Ctrl 1",
    //     LicenseParams: cn.NewSetLicenseParametersRequest(true),
    //     LicenseFile: "/Users/benplatta/code/cohesive/vns3-functional-testing/test-assets/license.txt",
    //     PeerId: 1,
    //     KeysetParams: cn.SetKeysetParamsRequest{
    //         Token: "token",
    //     },
    //     WaitTimeout: 60*5,
    //     KeysetTimeout: 60*5,
    // }

    // configDetail, setupErr := macros.SetupController(vns3, setupReq)

    // if setupErr != nil {
    //     log.Printf("VNS3 Setup error: %+v", setupErr)
    // } else {
    //     c := *configDetail
    //     d, _ := c.MarshalJSON()
    //     log.Printf("Setup success: %v", string(d))
    // }

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