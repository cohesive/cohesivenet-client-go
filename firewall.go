package cohesivenet

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	var rules = result["response"].([]interface{})
	var list string = "{\"response\" : [ "
	for _, rule := range rules {
		list += "{ "
		values := rule.([]interface{})
		desc := values[0].(string)
		id := strconv.Itoa(int(values[1].(float64)))
		list += "\"id\":\"" + id + "\", \"rule\":\"" + desc + "\"" + " },"
	}
	list = list[:len(list)-1]
	list += "]}"

	firewallResponse := FirewallResponse{}
	errUnmarshal := json.Unmarshal([]byte(list), &firewallResponse)

	if err != nil {
		log.Println(errUnmarshal)
	}

	return firewallResponse, nil
}

