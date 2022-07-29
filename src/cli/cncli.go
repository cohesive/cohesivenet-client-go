package main

import (
	"flag"
	"fmt"
    "time"
	// "os"
	"context"
	"cohesivenet"
    "encoding/json"
)


func pollvns3(vns3 cohesivenet.VNS3Client, timeout int, sleep int) bool {
    start := time.Now()

    for (time.Now()).Sub(start).Seconds() < float64(timeout) {
        time.Sleep(time.Duration(sleep) * time.Second)
    }

    return false
}


// 1. set topology name (can be done on licensing or PUT config)
// 2. Uploadlicense
// 3. set license parameters
// 4. generate keyset
// 5. set peering id

func setupController(vns3 cohesivenet.VNS3Client, ctx context.Context) bool {
    // license := "X"
    set_topology_req := cohesivenet.NewUpdateConfigRequest()
    set_topology_req.SetTopologyName("TEST Topo")
    set_topology_req.SetControllerName("Ctrl 1")
    config_resp := vns3.ConfigurationApi.PutConfig(set_topology_req)
    

}

func main() {
    flag.Parse()

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
        Timeout: 10,
        TLS: false,
    })

    start := time.Now()
    req := vns3.ConfigurationApi.GetConfigRequest(context.Background())
    r, _, err := vns3.ConfigurationApi.GetConfig(req)
    time.Sleep(5 * time.Second)
    elapsed := (time.Now()).Sub(start)
    var timeout float64 = 3
    fmt.Println(elapsed)
    fmt.Println(elapsed.Seconds() >= timeout)

    // req := vns3.RoutingApi.GetRoutesRequest(context.Background())
    // r, _, err := vns3.RoutingApi.GetRoutes(req)

    if err != nil {
        fmt.Printf("Error is not nil")
        fmt.Printf("%+v\n", err)
    } else {
        d, _ := json.Marshal(r.Response)
        fmt.Printf("Heres the response\n")
        fmt.Printf(string(d))
    }
}