# Golang Cohesivenet Client Library

Cohesive networks VNS3 provides complete control of your network's addressing, routes, rules and edge enabling a secure, connected and reactive cloud network.

- API version: 6.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://support.cohesive.net](https://support.cohesive.net)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import cohesivenet "github.com/GIT_USER_ID/GIT_REPO_ID"
```

## Authenticating with VNS3

**With API Password**

```go
import (
	"fmt"
	"context"
	"cohesivenet"
    "encoding/json"
)

auth := context.WithValue(context.Background(), cohesivenet.ContextBasicAuth, cohesivenet.BasicAuth{
     UserName: "api",
     Password: "mypassword",
})

// Create config with VNS3 host
config := cohesivenet.NewConfiguration("3.222.246.246")
vns3 := cohesivenet.NewVNS3Client(config, cohesivenet.ClientParams{
    Timeout: 10,
    TLS: false,
})

// Create request with auth context
req := vns3.ConfigurationApi.GetConfig(auth)
data, _, err := vns3.ConfigurationApi.GetConfigExecute(req)

if err != nil {
    fmt.Printf("Error is not nil")
    fmt.Printf("%+v\n", err)    
} else {
    output, _ := json.Marshal(data.Response)
    fmt.Printf("Heres the response\n")
    fmt.Printf(string(output))
}
```


**With API Token** 


```go
import (
	"fmt"
	"context"
	"cohesivenet"
    "encoding/json"
)

token := "1b0431ded2357eff780530a81d01786941a8af53975e7d9c4b0cc1q0fcc337fx"
auth := context.WithValue(context.Background(), cohesivenet.ContextAccessToken, token)
config := cohesivenet.NewConfiguration("3.222.246.246")
vns3 := cohesivenet.NewVNS3Client(config, cohesivenet.ClientParams{
    Timeout: 10,
    TLS: false,
})
req := vns3.ConfigurationApi.GetConfig(auth)
data, _, err := vns3.ConfigurationApi.GetConfigExecute(req)

if err != nil {
    fmt.Printf("Error is not nil")
    fmt.Printf("%+v\n", err)    
} else {
    output, _ := json.Marshal(data.Response)
    fmt.Printf("Heres the response\n")
    fmt.Printf(string(output))
}
```

## VNS3 Client API Categories

```golang
vns3 := cohesivenet.NewVNS3Client(config, cohesivenet.ClientParams{
    Timeout: 10,
    TLS: true,
})
vns3.AccessApi = (*AccessApiService)(&vns3.common)
// vns3.BGPApi currently not supported
vns3.ConfigurationApi
vns3.FirewallApi
vns3.IPsecApi
// vns3.InterfacesApi currently not supported
// vns3.MonitoringAlertingApi currently not supported
// vns3.NetworkEdgePluginsApi currently not supported
// vns3.OverlayNetworkApi currently not supported
vns3.PeeringApi
vns3.RoutingApi
// vns3.SnapshotsApi currently not supported
// vns3.SystemAdministrationApi = currently not supported
```

## About the Library