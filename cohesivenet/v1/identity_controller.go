package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const identityControllerEndpoint = "/identity/controller/"

func (c *Client) GetIdentityController() (IdentityController, error) {

	req, errNewRequest := http.NewRequest("GET", fmt.Sprintf("%s%s", c.HostURL, identityControllerEndpoint), nil)
	if errNewRequest != nil {
		return IdentityController{}, fmt.Errorf("error creating request: %w", errNewRequest)
	}

	body, errDoRequest := c.doRequest(req)
	if errDoRequest != nil {
		return IdentityController{}, fmt.Errorf("error doing request: %w", errDoRequest)
	}

	if body == nil {
		return IdentityController{}, fmt.Errorf("empty response body")
	}

	newIdentityController := IdentityController{}
	errUnmarshal := json.Unmarshal([]byte(body), &newIdentityController)
	if errUnmarshal != nil {
		return IdentityController{}, fmt.Errorf("error unmarshaling response: %w", errUnmarshal)
	}

	return newIdentityController, nil
}

func (c *Client) CreateIdentityController(identityController NewIdentityController) (*IdentityController, error) {

	rb, errMarshal := json.Marshal(identityController)
	if errMarshal != nil {
		return &IdentityController{}, fmt.Errorf("error marshaling alert: %w", errMarshal)
	}

	req, errNewRequest := http.NewRequest("PUT", fmt.Sprintf("%s%s", c.HostURL, identityControllerEndpoint), strings.NewReader(string(rb)))
	if errNewRequest != nil {
		return &IdentityController{}, fmt.Errorf("error creating request: %w", errNewRequest)
	}

	body, errDoRequest := c.doRequest(req)
	if errDoRequest != nil {
		return &IdentityController{}, fmt.Errorf("error doing request: %w", errDoRequest)
	}

	if body == nil {
		return &IdentityController{}, fmt.Errorf("empty response body")
	}
	newIdentityController := IdentityController{}
	errUnmarshal := json.Unmarshal([]byte(body), &newIdentityController)
	if errUnmarshal != nil {
		return &IdentityController{}, fmt.Errorf("error unmarshaling response: %w", errUnmarshal)
	}
	log.Printf("Created new alert: %+v", newIdentityController)

	return &newIdentityController, nil
}

func (c *Client) UpdateIdentityController(identityController NewIdentityController) (*IdentityController, error) {
	rb, errMarshal := json.Marshal(identityController)
	if errMarshal != nil {
		return &IdentityController{}, fmt.Errorf("error marshaling alert: %w", errMarshal)
	}

	req, errNewRequest := http.NewRequest("PUT", fmt.Sprintf("%s%s", c.HostURL, identityControllerEndpoint), strings.NewReader(string(rb)))
	if errNewRequest != nil {
		return &IdentityController{}, fmt.Errorf("error creating request: %w", errNewRequest)
	}

	body, errDoRequest := c.doRequest(req)
	if errDoRequest != nil {
		return &IdentityController{}, fmt.Errorf("error doing request: %w", errDoRequest)
	}

	if body == nil {
		return &IdentityController{}, fmt.Errorf("empty response body")
	}

	newIdentityController := IdentityController{}
	errUnmarshal := json.Unmarshal([]byte(body), &newIdentityController)
	if errUnmarshal != nil {
		return &IdentityController{}, fmt.Errorf("error unmarshaling response: %w", errUnmarshal)
	}
	log.Printf("Updated alert: %+v", newIdentityController)

	return &newIdentityController, nil
}

func (c *Client) DeleteIdentityController() error {
	req, errNewRequest := http.NewRequest("DELETE", fmt.Sprintf("%s%s", c.HostURL, identityControllerEndpoint), nil)
	if errNewRequest != nil {
		return fmt.Errorf("error creating request: %w", errNewRequest)
	}

	_, errDoRequest := c.doRequest(req)
	if errDoRequest != nil {
		return fmt.Errorf("error doing request: %w", errDoRequest)
	}

	return nil
}
