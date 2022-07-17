package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *Client) GetConfig() (ConfigResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/config", c.HostURL), nil)
	if err != nil {
		log.Println(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		log.Println(err)
	}
	response := ConfigResponse{}
	errUnmarshal := json.Unmarshal([]byte(body), &response)

	if err != nil {
		log.Println(errUnmarshal)
	}

	return response, nil
}
