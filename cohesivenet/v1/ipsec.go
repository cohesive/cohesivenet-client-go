package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetEndpoint(endpointId string) (NewEndpoint, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/ipsec/endpoints/%s", c.HostURL, endpointId), nil)
	if err != nil {
		return NewEndpoint{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return NewEndpoint{}, err
	}

	endpoint := NewEndpoint{}
	err = json.Unmarshal(body, &endpoint)
	if err != nil {
		return NewEndpoint{}, err
	}

	return endpoint, nil
}

func (c *Client) GetEndpoints() (EndpointResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/status/ipsec", c.HostURL), nil)
	if err != nil {
		return EndpointResponse{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return EndpointResponse{}, err
	}

	endpoints := EndpointResponse{}
	err = json.Unmarshal(body, &endpoints)
	if err != nil {
		return EndpointResponse{}, err
	}

	return endpoints, nil
}

func (c *Client) CreateEndpoint(endpoint *Endpoint) (*NewEndpoint, error) {

	rb, err := json.Marshal(endpoint)
	if err != nil {
		return &NewEndpoint{}, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/ipsec/endpoints", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return &NewEndpoint{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return &NewEndpoint{}, err
	}

	newEndpoint := NewEndpoint{}
	err = json.Unmarshal(body, &newEndpoint)
	if err != nil {
		return &NewEndpoint{}, err
	}

	return &newEndpoint, nil
}

func (c *Client) UpdateEndpoint(endpointId string, endpoint *Endpoint) (*NewEndpoint, error) {

	rb, err := json.Marshal(endpoint)
	if err != nil {
		return &NewEndpoint{}, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/ipsec/endpoints/%s", c.HostURL, endpointId), strings.NewReader(string(rb)))
	if err != nil {
		return &NewEndpoint{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return &NewEndpoint{}, err
	}

	newEndpoint := NewEndpoint{}
	err = json.Unmarshal(body, &newEndpoint)
	if err != nil {
		return &NewEndpoint{}, err
	}

	return &newEndpoint, nil
}

func (c *Client) DeleteEndpoint(endpointId string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/ipsec/endpoints/%s", c.HostURL, endpointId), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	if string(body) == "Deleted order" {
		return err
	}

	return nil
}
