package cohesivenet

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *Client) GetFirewallRules() (FirewallResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/firewall/rules", c.HostURL), nil)
	if err != nil {
		log.Println(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		log.Println(err)
	}

	firewallResponse := FirewallResponse{}
	errUnmarshal := json.Unmarshal([]byte(body), &firewallResponse)

	if err != nil {
		log.Println(errUnmarshal)
	}

	return firewallResponse, nil
}
