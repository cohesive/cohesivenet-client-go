package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const alertsEndpoint = "/alert/"

func (c *Client) GetAlert(alertId string) (Alert, error) {

	req, errNewRequest := http.NewRequest("GET", fmt.Sprintf("%s%s%s", c.HostURL, alertsEndpoint, alertId), nil)
	if errNewRequest != nil {
		return Alert{}, fmt.Errorf("error creating request: %w", errNewRequest)
	}

	body, errDoRequest := c.doRequest(req)
	if errDoRequest != nil {
		return Alert{}, fmt.Errorf("error doing request: %w", errDoRequest)
	}

	if body == nil {
		return Alert{}, fmt.Errorf("empty response body")
	}

	newAlert := Alert{}
	errUnmarshal := json.Unmarshal([]byte(body), &newAlert)
	if errUnmarshal != nil {
		return Alert{}, fmt.Errorf("error unmarshaling response: %w", errUnmarshal)
	}

	return newAlert, nil
}

func (c *Client) CreateAlert(alert NewAlert) (*Alert, error) {

	rb, errMarshal := json.Marshal(alert)
	if errMarshal != nil {
		return &Alert{}, fmt.Errorf("error marshaling alert: %w", errMarshal)
	}

	req, errNewRequest := http.NewRequest("POST", fmt.Sprintf("%s%s", c.HostURL, alertsEndpoint), strings.NewReader(string(rb)))
	if errNewRequest != nil {
		return &Alert{}, fmt.Errorf("error creating request: %w", errNewRequest)
	}

	body, errDoRequest := c.doRequest(req)
	if errDoRequest != nil {
		return &Alert{}, fmt.Errorf("error doing request: %w", errDoRequest)
	}

	if body == nil {
		return &Alert{}, fmt.Errorf("empty response body")
	}
	newAlert := Alert{}
	errUnmarshal := json.Unmarshal([]byte(body), &newAlert)
	if errUnmarshal != nil {
		return &Alert{}, fmt.Errorf("error unmarshaling response: %w", errUnmarshal)
	}
	log.Printf("Created new alert: %+v", newAlert)

	return &newAlert, nil
}

func (c *Client) UpdateAlert(alertId string, alert NewAlert) error {
	rb, errMarshal := json.Marshal(alert)
	if errMarshal != nil {
		return fmt.Errorf("error marshaling alert: %w", errMarshal)
	}

	req, errNewRequest := http.NewRequest("PUT", fmt.Sprintf("%s%s%s", c.HostURL, alertsEndpoint, alertId), strings.NewReader(string(rb)))
	if errNewRequest != nil {
		return fmt.Errorf("error creating request: %w", errNewRequest)
	}

	body, errDoRequest := c.doRequest(req)
	if errDoRequest != nil {
		return fmt.Errorf("error doing request: %w", errDoRequest)
	}

	if body == nil {
		return fmt.Errorf("empty response body")
	}

	newAlert := Alert{}
	errUnmarshal := json.Unmarshal([]byte(body), &newAlert)
	if errUnmarshal != nil {
		return fmt.Errorf("error unmarshaling response: %w", errUnmarshal)
	}
	log.Printf("Updated alert: %+v", newAlert)

	return nil
}

func (c *Client) DeleteAlert(alertId string) error {
	req, errNewRequest := http.NewRequest("DELETE", fmt.Sprintf("%s%s%s", c.HostURL, alertsEndpoint, alertId), nil)
	if errNewRequest != nil {
		return fmt.Errorf("error creating request: %w", errNewRequest)
	}

	_, errDoRequest := c.doRequest(req)
	if errDoRequest != nil {
		return fmt.Errorf("error doing request: %w", errDoRequest)
	}

	return nil
}
