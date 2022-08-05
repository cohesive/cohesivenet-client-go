package main

import (
    "log"
    "os"
    "context"
    "github.com/cohesive/cohesivenet-client-go/cohesivenet"
	"github.com/cohesive/cohesivenet-client-go/cohesivenet/macros"
    // clientv1 "github.com/cohesive/cohesivenet-client-go/cohesivenet/v1"
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


func main() {
    // auth := context.WithValue(context.Background(), cohesivenet.ContextBasicAuth, cohesivenet.BasicAuth{
    //     UserName: "api",
    //     Password: "testtest1234",
    // })

    // token := "eb0431ded2351eff780f30a81d01786941a8af58975e7d9c4b0cc1b0fcc337f8"
    // auth := context.WithValue(context.Background(), cohesivenet.ContextAccessToken, token)
    host := "54.242.59.62"
    ps := "i-00c2b187fb7ce56eb"
    // config := cohesivenet.NewConfigurationWithAuth("3.222.246.246", cohesivenet.ContextAccessToken, token)
    config := cohesivenet.NewConfigurationWithAuth(host, cohesivenet.ContextBasicAuth, cohesivenet.BasicAuth{
        UserName: "api",
        Password: ps,
    })
    vns3 := cohesivenet.NewVNS3Client(config, cohesivenet.ClientParams{
        Timeout: 3,
        TLS: false,
    })

    ctx := context.Background()
    _, err :=  vns3.ConfigurationApi.WaitForApi(&ctx, 60*10, 2, 5)
    if err != nil {
        log.Printf("VNS3 not available")
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

    setupReq := macros.SetupRequest{
        TopologyName: "Terraform",
        ControllerName: "Ctrl 1",
        LicenseParams: cohesivenet.NewSetLicenseParametersRequest(true),
        LicenseFile: "/Users/benplatta/code/cohesive/vns3-functional-testing/test-assets/license.txt",
        PeerId: 1,
        KeysetParams: cohesivenet.SetKeysetParamsRequest{
            Token: "token",
        },
        WaitTimeout: 60*5,
        KeysetTimeout: 60*5,
    }

    configDetail, setupErr := macros.SetupController(vns3, setupReq)

    if setupErr != nil {
        log.Printf("VNS3 Setup error: %+v", setupErr)
    } else {
        c := *configDetail
        d, _ := c.MarshalJSON()
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