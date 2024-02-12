package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const webhookEndpoint = "/webhook/"

func (c *Client) GetWebhook(webhookId string) (Webhook, error) {

	req, errNewRequest := http.NewRequest("GET", fmt.Sprintf("%s%s%s", c.HostURL, webhookEndpoint, webhookId), nil)
	if errNewRequest != nil {
		return Webhook{}, fmt.Errorf("error creating request: %w", errNewRequest)
	}

	body, errDoRequest := c.doRequest(req)
	if errDoRequest != nil {
		return Webhook{}, fmt.Errorf("error doing request: %w", errDoRequest)
	}

	if body == nil {
		return Webhook{}, fmt.Errorf("empty response body")
	}

	newWebhook := Webhook{}
	errUnmarshal := json.Unmarshal([]byte(body), &newWebhook)
	if errUnmarshal != nil {
		return Webhook{}, fmt.Errorf("error unmarshaling response: %w", errUnmarshal)
	}

	return newWebhook, nil
}

func (c *Client) CreateWebhook(webhook NewWebhook) (Webhook, error) {

	webhook.Body, _ = url.QueryUnescape(webhook.Body)

	rb, errMarshal := json.Marshal(webhook)
	if errMarshal != nil {
		return Webhook{}, fmt.Errorf("error marshaling alert: %w", errMarshal)
	}

	req, errNewRequest := http.NewRequest("POST", fmt.Sprintf("%s%s", c.HostURL, webhookEndpoint), strings.NewReader(string(rb)))
	if errNewRequest != nil {
		return Webhook{}, fmt.Errorf("error creating request: %w", errNewRequest)
	}

	body, errDoRequest := c.doRequest(req)
	log.Println(string(body))
	if errDoRequest != nil {
		return Webhook{}, fmt.Errorf("error doing request: %w", errDoRequest)
	}

	if body == nil {
		return Webhook{}, fmt.Errorf("empty response body")
	}
	newWebhook := Webhook{}
	errUnmarshal := json.Unmarshal([]byte(body), &newWebhook)
	if errUnmarshal != nil {
		return Webhook{}, fmt.Errorf("error unmarshaling response: %w", errUnmarshal)
	}
	log.Printf("Created new webhook: %+v", newWebhook)

	return newWebhook, nil
}

func (c *Client) UpdateWebhook(webhookId string, webhook NewWebhook) error {

	webhook.Body, _ = url.QueryUnescape(webhook.Body)

	rb, errMarshal := json.Marshal(webhook)
	if errMarshal != nil {
		return fmt.Errorf("error marshaling webhook: %w", errMarshal)
	}

	req, errNewRequest := http.NewRequest("PUT", fmt.Sprintf("%s%s%s", c.HostURL, webhookEndpoint, webhookId), strings.NewReader(string(rb)))
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

	newWebhook := Webhook{}
	errUnmarshal := json.Unmarshal([]byte(body), &newWebhook)
	if errUnmarshal != nil {
		return fmt.Errorf("error unmarshaling response: %w", errUnmarshal)
	}
	log.Printf("Updated webhook: %+v", newWebhook)

	return nil
}

func (c *Client) DeleteWebhook(webhookId string) error {
	req, errNewRequest := http.NewRequest("DELETE", fmt.Sprintf("%s%s%s", c.HostURL, webhookEndpoint, webhookId), nil)
	if errNewRequest != nil {
		return fmt.Errorf("error creating request: %w", errNewRequest)
	}

	_, errDoRequest := c.doRequest(req)
	if errDoRequest != nil {
		return fmt.Errorf("error doing request: %w", errDoRequest)
	}

	return nil
}
