package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetContainerNetwork() (ContainerSystem, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/container_system", c.HostURL), nil)
	if err != nil {
		return ContainerSystem{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return ContainerSystem{}, err
	}

	containerResponse := ContainerSystem{}
	err = json.Unmarshal([]byte(body), &containerResponse)
	if err != nil {
		return ContainerSystem{}, err
	}

	return containerResponse, nil
}
