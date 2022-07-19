package cohesivenet

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetEndpoint(endpointID string) (*Endpoints, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/status/ipsec", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	endpoints := Endpoints{}
	err = json.Unmarshal(body, &endpoints)
	if err != nil {
		return nil, err
	}

	return &endpoints, nil
}
