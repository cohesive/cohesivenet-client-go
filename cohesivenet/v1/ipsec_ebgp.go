package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (c *Client) GetEbgpPeer(endpointId string, ebgpPeerId string) (EbgpPeer, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/ipsec/endpoints/%s/ebgp_peers/%s", c.HostURL, endpointId, ebgpPeerId), nil)
	if err != nil {
		return EbgpPeer{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return EbgpPeer{}, err
	}

	simplePeer := SimplifyGetEbpgJson(string(body))

	ebgPeer := EbgpPeer{}
	err = json.Unmarshal([]byte(simplePeer), &ebgPeer)
	if err != nil {
		return EbgpPeer{}, err
	}

	return ebgPeer, nil
}

func (c *Client) CreateEbgpPeer(endpointId string, ebgp_peer *EbgpPeer) (*EbgpPeer, error) {

	rb, err := json.Marshal(ebgp_peer)
	if err != nil {
		return &EbgpPeer{}, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/ipsec/endpoints/%s/ebgp_peers", c.HostURL, endpointId), strings.NewReader(string(rb)))
	if err != nil {
		return &EbgpPeer{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return &EbgpPeer{}, err
	}

	simplePeer := SimplifyEbpgJson(string(body))

	newPeer := EbgpPeer{}
	err = json.Unmarshal([]byte(simplePeer), &newPeer)
	if err != nil {
		return &EbgpPeer{}, err
	}

	return &newPeer, nil
}

func (c *Client) UpdateEbgpPeer(endpointId string, ebgpPeerId string, ebgp_peer *EbgpPeer) (*EbgpPeer, error) {

	requestBody, err := json.Marshal(ebgp_peer)
	if err != nil {
		return &EbgpPeer{}, err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/ipsec/endpoints/%s/ebgp_peers/%s", c.HostURL, endpointId, ebgpPeerId), strings.NewReader(string(requestBody)))
	if err != nil {
		return &EbgpPeer{}, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return &EbgpPeer{}, err
	}

	simplePeer := SimplifyEbpgJson(string(body))

	newPeer := EbgpPeer{}
	err = json.Unmarshal([]byte(simplePeer), &newPeer)
	if err != nil {
		return &EbgpPeer{}, err
	}
	return &newPeer, nil
}

func (c *Client) DeleteEbgpPeer(endpointId string, ebgpPeerId string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/ipsec/endpoints/%s/ebgp_peers/%s", c.HostURL, endpointId, ebgpPeerId), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	if string(body) == "Deleted order" {
		return err
	}

	return nil
}

func SimplifyEbpgJson(jsonInput string) string {
	var result map[string]interface{}
	json.Unmarshal([]byte(jsonInput), &result)
	var response = result["response"].(map[string]interface{})
	var peers = response["bgp_peers"].(map[string]interface{})
	var list string = " { "
	var lastPeer interface{} = nil

	for _, peer := range peers {
		lastPeer = peer
	}
	if lastPeer != nil {
		values := lastPeer.(map[string]interface{})
		list += "\"id\":" + formatFloatValue(values, "id") + ","
		list += "\"ipaddress\":\"" + UnmarshalString(values, "ipaddress") + "\","
		list += "\"asn\":" + formatFloatValue(values, "asn") + ","
		list += "\"local_asn_alias\":" + formatFloatValue(values, "local_asn_alias") + ","
		list += "\"keepalive_interval\":" + formatFloatValue(values, "keepalive_interval") + ","
		list += "\"hold_time\":" + formatFloatValue(values, "hold_time") + ","
		list += "\"bgp_password\":\"" + UnmarshalString(values, "bgp_password") + "\","
		list += "\"access_list\":\"" + strings.ReplaceAll(strings.ReplaceAll(values["access_list"].(string), "\r", ""), "\n", ",") + "\","
		list += "\"add_network_distance\":" + strconv.FormatBool(values["add_network_distance"].(bool)) + ","
		list += "\"add_network_distance_direction\":\"" + UnmarshalString(values, "add_network_distance_direction") + "\","
		list += "\"add_network_distance_hops\":" + formatFloatValue(values, "add_network_distance_hops") + ""
	}
	list += "}"
	return list
}

func SimplifyGetEbpgJson(jsonInput string) string {
	var result map[string]interface{}
	json.Unmarshal([]byte(jsonInput), &result)
	var values = result["response"].(map[string]interface{})
	var list string = " { "
	list += "\"id\":" + formatFloatValue(values, "id") + ","
	list += "\"ipaddress\":\"" + UnmarshalString(values, "ipaddress") + "\","
	list += "\"asn\":" + formatFloatValue(values, "asn") + ","
	list += "\"local_asn_alias\":" + formatFloatValue(values, "local_asn_alias") + ","
	list += "\"keepalive_interval\":" + formatFloatValue(values, "keepalive_interval") + ","
	list += "\"hold_time\":" + formatFloatValue(values, "hold_time") + ","
	list += "\"bgp_password\":\"" + UnmarshalString(values, "bgp_password") + "\","
	list += "\"access_list\":\"" + strings.ReplaceAll(strings.ReplaceAll(values["access_list"].(string), "\r", ""), "\n", ",") + "\","
	list += "\"add_network_distance\":" + strconv.FormatBool(values["add_network_distance"].(bool)) + ","
	list += "\"add_network_distance_direction\":\"" + UnmarshalString(values, "add_network_distance_direction") + "\","
	list += "\"add_network_distance_hops\":" + formatFloatValue(values, "add_network_distance_hops") + ""
	list += " }"
	return list
}

func formatFloatValue(values map[string]interface{}, key string) string {
	value, ok := values[key]
	if ok && value != nil {
		if floatValue, ok := value.(float64); ok && floatValue != 0 {
			return strconv.FormatFloat(floatValue, 'f', -1, 64)
		}
	}
	return "null"
}

func UnmarshalString(values map[string]interface{}, key string) string {
	value, ok := values[key]
	if ok && value != nil {
		return value.(string)
	}
	return ""
}
