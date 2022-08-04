package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *Client) GetContainerNetwork() (ContainerSystem, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/container_system", c.HostURL), nil)
	if err != nil {
		log.Println(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		log.Println(err)
	}

	containerResponse := ContainerSystem{}
	errUnmarshal := json.Unmarshal([]byte(body), &containerResponse)
	if err != nil {
		log.Println(errUnmarshal)
	}

	return containerResponse, nil
}
