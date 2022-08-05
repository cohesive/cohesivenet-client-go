# Example usage of the client
Eventually we will have a nice CLI using this library but for the time being here is an example on how to use it

## Installing
You'll need to run `go get` for the dependencies in go.mod file

## Running
You can run the script cncli.go with `go run cncli.go`

## About Authenticating

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
