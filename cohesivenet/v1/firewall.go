package cohesivenetv1

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

func (c *Client) CreateFirewallRules(ruleList []*FirewallRule) ([]*FirewallResponse, error) {
	i := 0
	for _, rl := range ruleList {
		rule := FirewallRule{
			ID:   rl.ID,
			Rule: rl.Rule,
		}

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

		i++
	}

	req2, err := http.NewRequest("GET", fmt.Sprintf("%s/firewall/rules", c.HostURL), nil)
	if err != nil {
		log.Println(err)
	}

	body, err := c.doRequest(req2)
	if err != nil {
		log.Println(err)
	}

	//simpleJson := simplifyJson(string(body))
	//newRules := FirewallResponse{}
	//errUnmarshal := json.Unmarshal([]byte(simpleJson), &newRules)

	//if err != nil {
	//	log.Println(errUnmarshal)
	//}

	//returnRules := FirewallResponse{}
	//for _, r := range newRules.FirewallRules {
	//			if strings.EqualFold(r.Rule, rule.Rule) {
	//		returnRules.FirewallRules = []FirewallRule{r}
	//
	//	}
	//}
	returnRules := []*FirewallResponse{}
	errUnmarshal := json.Unmarshal([]byte(body), &returnRules)
	if err != nil {
		log.Println(errUnmarshal)
	}

	return returnRules, nil
}

func (c *Client) DeleteRule() error {
	req1, err := http.NewRequest("GET", fmt.Sprintf("%s/firewall/rules", c.HostURL), nil)
	if err != nil {
		log.Println(err)
	}

	rules, err := c.doRequest(req1)
	if err != nil {
		log.Println(err)
	}

	simpleJson := simplifyJson(string(rules))

	firewallResponse := FirewallResponse{}
	errUnmarshal := json.Unmarshal([]byte(simpleJson), &firewallResponse)

	if err != nil {
		log.Println(errUnmarshal)
	}

	for _, r := range firewallResponse.FirewallRules {
		ruleId := r.ID

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
