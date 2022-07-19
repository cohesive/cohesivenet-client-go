package cohesivenet

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
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

	simpleJson := simplifyJson(string(body))

	firewallResponse := FirewallResponse{}
	errUnmarshal := json.Unmarshal([]byte(simpleJson), &firewallResponse)

	if err != nil {
		log.Println(errUnmarshal)
	}

	return firewallResponse, nil
}

func (c *Client) CreateFirewallRules(rule *FirewallRule) (FirewallResponse, error) {
	rb, err := json.Marshal(rule)
	if err != nil {
		log.Println(err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/firewall/rules", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		log.Println(err)
	}

	_, err2 := c.doRequest(req)
	if err != nil {
		log.Println(err2)
	}

	req2, err := http.NewRequest("GET", fmt.Sprintf("%s/firewall/rules", c.HostURL), nil)
	if err != nil {
		log.Println(err)
	}

	body, err := c.doRequest(req2)
	if err != nil {
		log.Println(err)
	}

	simpleJson := simplifyJson(string(body))
	newRules := FirewallResponse{}
	errUnmarshal := json.Unmarshal([]byte(simpleJson), &newRules)
	if err != nil {
		log.Println(errUnmarshal)
	}

	returnRules := FirewallResponse{}
	for _, r := range newRules.FirewallRules {
		if strings.EqualFold(r.Rule, rule.Rule) {
			returnRules.FirewallRules = []FirewallRule{r}

		}
	}

	return returnRules, nil
}

func (c *Client) DeleteRule(ruleId string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/firewall/rules/%s", c.HostURL, ruleId), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	if string(body) == "Deleted route" {
		return err
	}

	return nil
}

func simplifyJson(input string) string {
	var result map[string]interface{}
	json.Unmarshal([]byte(input), &result)
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

	return list
}
