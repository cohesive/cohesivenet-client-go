package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const identityVpnEndpoint = "/identity/vpn/"

func (c *Client) GetIdentityVpn() (IdentityVpn, error) {

	req, errNewRequest := http.NewRequest("GET", fmt.Sprintf("%s%s", c.HostURL, identityVpnEndpoint), nil)
	if errNewRequest != nil {
		return IdentityVpn{}, fmt.Errorf("error creating request: %w", errNewRequest)
	}

	body, errDoRequest := c.doRequest(req)
	if errDoRequest != nil {
		return IdentityVpn{}, fmt.Errorf("error doing request: %w", errDoRequest)
	}

	if body == nil {
		return IdentityVpn{}, fmt.Errorf("empty response body")
	}

	newIdentityVpn := IdentityVpn{}
	errUnmarshal := json.Unmarshal([]byte(body), &newIdentityVpn)
	if errUnmarshal != nil {
		return IdentityVpn{}, fmt.Errorf("error unmarshaling response: %w", errUnmarshal)
	}

	return newIdentityVpn, nil
}

func (c *Client) CreateIdentityVpn(identityVpn NewIdentityVpn) (*IdentityVpn, error) {

	rb, errMarshal := json.Marshal(identityVpn)
	if errMarshal != nil {
		return &IdentityVpn{}, fmt.Errorf("error marshaling alert: %w", errMarshal)
	}

	req, errNewRequest := http.NewRequest("PUT", fmt.Sprintf("%s%s", c.HostURL, identityVpnEndpoint), strings.NewReader(string(rb)))
	if errNewRequest != nil {
		return &IdentityVpn{}, fmt.Errorf("error creating request: %w", errNewRequest)
	}

	body, errDoRequest := c.doRequest(req)
	if errDoRequest != nil {
		return &IdentityVpn{}, fmt.Errorf("error doing request: %w", errDoRequest)
	}

	if body == nil {
		return &IdentityVpn{}, fmt.Errorf("empty response body")
	}
	newIdentityVpn := IdentityVpn{}
	errUnmarshal := json.Unmarshal([]byte(body), &newIdentityVpn)
	if errUnmarshal != nil {
		return &IdentityVpn{}, fmt.Errorf("error unmarshaling response: %w", errUnmarshal)
	}
	log.Printf("Created new alert: %+v", newIdentityVpn)

	return &newIdentityVpn, nil
}

func (c *Client) UpdateIdentityVpn(identityVpn NewIdentityVpn) (*IdentityVpn, error) {
	rb, errMarshal := json.Marshal(identityVpn)
	if errMarshal != nil {
		return &IdentityVpn{}, fmt.Errorf("error marshaling alert: %w", errMarshal)
	}

	req, errNewRequest := http.NewRequest("PUT", fmt.Sprintf("%s%s", c.HostURL, identityVpnEndpoint), strings.NewReader(string(rb)))
	if errNewRequest != nil {
		return &IdentityVpn{}, fmt.Errorf("error creating request: %w", errNewRequest)
	}

	body, errDoRequest := c.doRequest(req)
	if errDoRequest != nil {
		return &IdentityVpn{}, fmt.Errorf("error doing request: %w", errDoRequest)
	}

	if body == nil {
		return &IdentityVpn{}, fmt.Errorf("empty response body")
	}

	newIdentityVpn := IdentityVpn{}
	errUnmarshal := json.Unmarshal([]byte(body), &newIdentityVpn)
	if errUnmarshal != nil {
		return &IdentityVpn{}, fmt.Errorf("error unmarshaling response: %w", errUnmarshal)
	}
	log.Printf("Updated alert: %+v", newIdentityVpn)

	return &newIdentityVpn, nil
}

func (c *Client) DeleteIdentityVpn() error {
	req, errNewRequest := http.NewRequest("DELETE", fmt.Sprintf("%s%s", c.HostURL, identityVpnEndpoint), nil)
	if errNewRequest != nil {
		return fmt.Errorf("error creating request: %w", errNewRequest)
	}

	_, errDoRequest := c.doRequest(req)
	if errDoRequest != nil {
		return fmt.Errorf("error doing request: %w", errDoRequest)
	}

	return nil
}
