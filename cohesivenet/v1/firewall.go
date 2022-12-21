package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (c *Client) GetFirewallRules() (FirewallResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/firewall/rules", c.HostURL), nil)
	if err != nil {
		return FirewallResponse{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return FirewallResponse{}, err
	}

	simpleJson := simplifyJson(string(body))

	firewallResponse := FirewallResponse{}
	err = json.Unmarshal([]byte(simpleJson), &firewallResponse)

	if err != nil {
		return FirewallResponse{}, err
	}

	return firewallResponse, nil
}

func (c *Client) CreateFirewallRules(ruleList []*FirewallRule) error {
	i := 0
	for _, rl := range ruleList {
		rule := FirewallRule{
			ID:   rl.ID,
			Rule: rl.Rule,
		}

		rb, err := json.Marshal(rule)
		if err != nil {
			return err
		}
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/firewall/rules", c.HostURL), strings.NewReader(string(rb)))
		if err != nil {
			return err
		}

		_, err = c.doRequest(req)
		if err != nil {
			return err
		}

		i++
	}

	return nil
}
func (c *Client) UpdateRules(ruleList []*FirewallRule) error {
	var rulesList []*FirewallRule
	rules := FirewallResponse{}
	rules, _ = c.GetFirewallRules()
	r := getRulesData(rules)

	for _, rule := range r {
		rle := rule.(map[string]interface{})
		rule := FirewallRule{
			ID:   rle["id"].(string),
			Rule: rle["script"].(string),
		}

		rulesList = append(rulesList, &rule)
	}

	c.DeleteRules(rulesList)
	c.CreateFirewallRules(ruleList)

	return nil
}

func (c *Client) DeleteRules(ruleList []*FirewallRule) error {
	i := 0
	for _, rl := range ruleList {
		rule := FirewallRule{
			ID:   rl.ID,
			Rule: rl.Rule,
		}

		rb, err := json.Marshal(rule)
		if err != nil {
			return err
		}
		req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/firewall/rules", c.HostURL), strings.NewReader(string(rb)))
		if err != nil {
			return err
		}

		_, err = c.doRequest(req)
		if err != nil {
			return err
		}

		i++

	}

	return nil
}

func simplifyJson(input string) string {
	var result map[string]interface{}
	json.Unmarshal([]byte(input), &result)
	var rules = result["response"].([]interface{})
	var list string = "{\"response\" : [ "
	i := 0
	for _, rule := range rules {
		list += "{ "
		values := rule.([]interface{})
		desc := values[0].(string)
		id := strconv.Itoa(int(values[1].(float64)))
		list += "\"id\":\"" + id + "\", \"rule\":\"" + desc + "\"" + " },"
	}
	i++
	list = list[:len(list)-1]
	list += "]}"

	return list
}

func getRulesData(ruleResponse FirewallResponse) []interface{} {
	routes := make([]interface{}, len(ruleResponse.FirewallRules))

	i := 0
	for _, rt := range ruleResponse.FirewallRules {
		row := make(map[string]interface{})
		row["id"] = rt.ID
		row["script"] = rt.Rule
		routes[i] = row
		i++
	}

	return routes

}
