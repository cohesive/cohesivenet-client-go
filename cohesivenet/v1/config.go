package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetConfig() (ConfigResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/config", c.HostURL), nil)
	if err != nil {
		return ConfigResponse{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return ConfigResponse{}, err
	}
	response := ConfigResponse{}
	err = json.Unmarshal([]byte(body), &response)

	if err != nil {
		return ConfigResponse{}, err
	}

	return response, nil
}
