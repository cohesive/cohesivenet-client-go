package main

import (
	"flag"
	"fmt"
	// "os"
	"context"
	"cohesivenet"
    "encoding/json"
)

func main() {
    flag.Parse()

    auth := context.WithValue(context.Background(), cohesivenet.ContextBasicAuth, cohesivenet.BasicAuth{
        UserName: "api",
        Password: "XXXX",
    })
    config := cohesivenet.NewConfiguration("3.222.246.246")
    vns3 := cohesivenet.NewVNS3Client(config, cohesivenet.ClientParams{
        Timeout: 10,
        TLS: false,
    })
    req := vns3.ConfigurationApi.GetConfig(auth)
    r, _, err := vns3.ConfigurationApi.GetConfigExecute(req)

    if err != nil {
        fmt.Printf("Error is not nil")
        fmt.Printf("%+v\n", err)    
    } else {
        d, _ := json.Marshal(r.Response)
        fmt.Printf("Heres the response\n")
        fmt.Printf(string(d))
    }
}