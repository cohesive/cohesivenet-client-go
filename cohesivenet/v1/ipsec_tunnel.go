package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (c *Client) GetTunnel(endpointId int, remoteSubnet string, tunnelId string) ([]NewTunnel, error) {

	endId := strconv.Itoa(endpointId)
	tunId, _ := strconv.Atoi(tunnelId)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/ipsec/endpoints/%s", c.HostURL, endId), nil)
	if err != nil {
		return []NewTunnel{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return []NewTunnel{}, err
	}

	simpleJson := SimplifyTunnelJson(string(body))

	newTunnel := []NewTunnel{}
	err = json.Unmarshal([]byte(simpleJson), &newTunnel)
	if err != nil {
		return []NewTunnel{}, err
	}

	returnTunnel := []NewTunnel{}
	for _, t := range newTunnel {
		if t.RemoteSubnet == remoteSubnet && t.ID == tunId {
			returnTunnel = []NewTunnel{t}
			break
		}
	}

	return returnTunnel, err
}

func (c *Client) CreateTunnel(endpointId string, remoteSubnet string, tunnel *Tunnel) ([]NewTunnel, error) {
	rb, err := json.Marshal(tunnel)
	if err != nil {
		return []NewTunnel{}, err

	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/ipsec/endpoints/%s/tunnels", c.HostURL, endpointId), strings.NewReader(string(rb)))
	if err != nil {
		return []NewTunnel{}, err

	}

	body, err := c.doRequest(req)
	if err != nil {
		return []NewTunnel{}, err

	}

	simpleJson := SimplifyTunnelJson(string(body))

	newTunnel := []NewTunnel{}
	err = json.Unmarshal([]byte(simpleJson), &newTunnel)
	if err != nil {
		return []NewTunnel{}, err

	}

	returnTunnel := []NewTunnel{}
	for _, t := range newTunnel {
		if strings.EqualFold(t.RemoteSubnet, remoteSubnet) {
			returnTunnel = []NewTunnel{t}
			break
		}
	}

	return returnTunnel, err

}

func (c *Client) UpdateTunnel(endpointId int, tunnelId string, remoteSubnet string, tunnel *Tunnel) (UpdateTunnelResponse, error) {

	endId := strconv.Itoa(endpointId)

	rb, err := json.Marshal(tunnel)
	if err != nil {
		return UpdateTunnelResponse{}, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/ipsec/endpoints/%s/tunnels/%s", c.HostURL, endId, tunnelId), strings.NewReader(string(rb)))
	if err != nil {
		return UpdateTunnelResponse{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return UpdateTunnelResponse{}, err
	}

	newTunnel := UpdateTunnelResponse{}
	err = json.Unmarshal([]byte(body), &newTunnel)
	if err != nil {
		return UpdateTunnelResponse{}, err
	}
	return newTunnel, err
}

func (c *Client) DeleteTunnel(endpointId int, tunnelId string) error {

	endId := strconv.Itoa(endpointId)

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/ipsec/endpoints/%s/tunnels/%s", c.HostURL, endId, tunnelId), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return err
}

func SimplifyTunnelJson(input string) string {
	var result map[string]interface{}
	json.Unmarshal([]byte(input), &result)

	var response = result["response"].(map[string]interface{})
	var tunnels = response["tunnels"].(map[string]interface{})
	var list string = "["
	for _, tunnel := range tunnels {
		list += "{"
		values := tunnel.(map[string]interface{})
		list += "\"id\":" + formatTunnelFloatValue(values, "id") + ","
		list += "\"local_subnet\":\"" + unmarshalTunnelString(values, "local_subnet") + "\","
		list += "\"remote_subnet\":\"" + unmarshalTunnelString(values, "remote_subnet") + "\","
		list += "\"description\":\"" + unmarshalTunnelString(values, "description") + "\","
		list += "\"enabled\":" + strconv.FormatBool(values["enabled"].(bool)) + ","
		list += "\"endpoint_id\": " + formatTunnelFloatValue(values, "endpoint_id") + ","
		list += "\"ping_ipaddress\":\"" + unmarshalTunnelString(values, "ping_ipaddress") + "\","
		list += "\"ping_interface\":\"" + unmarshalTunnelString(values, "ping_interface") + "\","
		list += "\"ping_interval\": " + formatTunnelFloatValue(values, "ping_interval") + ","
		list = strings.TrimSuffix(list, ",") //remove last comma
		list += "},"

	}
	list = strings.TrimSuffix(list, ",") //remove last comma
	list += "]"
	return list
}

// handles missing properties and null values for strings
func unmarshalTunnelString(values map[string]interface{}, name string) string {
	if val, ok := values[name]; ok {
		if val != nil {
			return fmt.Sprintf("%v", val)
		}
	}
	return ""
}

func formatTunnelFloatValue(values map[string]interface{}, key string) string {
	value, ok := values[key]
	if ok && value != nil {
		if floatValue, ok := value.(float64); ok && floatValue != 0 {
			return strconv.FormatFloat(floatValue, 'f', -1, 64)
		}
	}
	return "null"
}
