package cohesivenet

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func (c *Client) GetEndpoints() (EndpointResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/status/ipsec", c.HostURL), nil)
	if err != nil {
		log.Println(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		log.Println(err)
	}

	endpoints := EndpointResponse{}
	err = json.Unmarshal(body, &endpoints)
	if err != nil {
		log.Println(err)
	}

	return endpoints, nil
}

func (c *Client) CreateEndpoints(endpoints EndpointResponse) (*EndpointResponse, error) {
	rb, err := json.Marshal(endpoints)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/ipsec/endpoints", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		log.Println(err)
	}

	newEndpoint := EndpointResponse{}
	err = json.Unmarshal(body, &endpoints)
	if err != nil {
		log.Println(err)
	}

	return &newEndpoint, nil
}
