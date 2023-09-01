package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (c *Client) GetTrafficPair(endpointId int, remoteSubnet string, trafficPairId string) ([]TrafficPair, error) {

	endId := strconv.Itoa(endpointId)
	tpId, _ := strconv.Atoi(trafficPairId)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/ipsec/endpoints/%s", c.HostURL, endId), nil)
	if err != nil {
		return []TrafficPair{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return []TrafficPair{}, err
	}

	simpleJson := SimplifyTrafficPairJson(string(body))

	newTrafficPair := []TrafficPair{}
	err = json.Unmarshal([]byte(simpleJson), &newTrafficPair)
	if err != nil {
		return []TrafficPair{}, err
	}

	returnedTrafficPair := []TrafficPair{}
	for _, t := range newTrafficPair {
		if strings.EqualFold(t.Remote_Subnet, remoteSubnet) && t.ID == tpId {
			returnedTrafficPair = []TrafficPair{t}
			break
		}
	}

	return returnedTrafficPair, err
}

func (c *Client) CreateTrafficPair(endpointId string, remoteSubnet string, trafficPair *TrafficPair) (CreateTrafficPairResponse, error) {
	rb, err := json.Marshal(trafficPair)
	if err != nil {
		return CreateTrafficPairResponse{}, err

	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/ipsec/endpoints/%s/traffic_pairs", c.HostURL, endpointId), strings.NewReader(string(rb)))
	if err != nil {
		return CreateTrafficPairResponse{}, err

	}

	body, err := c.doRequest(req)
	if err != nil {
		return CreateTrafficPairResponse{}, err

	}

	newTrafficPair := CreateTrafficPairResponse{}
	err = json.Unmarshal([]byte(body), &newTrafficPair)
	if err != nil {
		return CreateTrafficPairResponse{}, err

	}

	return newTrafficPair, err

}

func (c *Client) UpdateTrafficPair(endpointId int, trafficPairId string, trafficPair *TrafficPair) (CreateTrafficPairResponse, error) {

	endId := strconv.Itoa(endpointId)

	rb, err := json.Marshal(trafficPair)
	if err != nil {
		return CreateTrafficPairResponse{}, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/ipsec/endpoints/%s/traffic_pairs/%s", c.HostURL, endId, trafficPairId), strings.NewReader(string(rb)))
	if err != nil {
		return CreateTrafficPairResponse{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return CreateTrafficPairResponse{}, err
	}
	newTrafficPair := CreateTrafficPairResponse{}
	err = json.Unmarshal([]byte(body), &newTrafficPair)
	if err != nil {
		return CreateTrafficPairResponse{}, err

	}

	return newTrafficPair, err
}

func (c *Client) DeleteTrafficPair(endpointId int, trafficPairId string) error {

	endId := strconv.Itoa(endpointId)

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/ipsec/endpoints/%s/traffic_pairs/%s", c.HostURL, endId, trafficPairId), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return err
}

func SimplifyTrafficPairJson(input string) string {
	var result map[string]interface{}
	json.Unmarshal([]byte(input), &result)

	var response = result["response"].(map[string]interface{})
	var trafficPairs = response["traffic_pairs"].(map[string]interface{})
	var list string = "["
	for _, trafficPair := range trafficPairs {
		list += "{"
		values := trafficPair.(map[string]interface{})
		list += "\"id\":" + formatTrafficPairFloatValue(values, "id") + ","
		list += "\"local_subnet\":\"" + unmarshalTrafficPairString(values, "local_subnet") + "\","
		list += "\"remote_subnet\":\"" + unmarshalTrafficPairString(values, "remote_subnet") + "\","
		list += "\"description\":\"" + unmarshalTrafficPairString(values, "description") + "\","
		list += "\"enabled\":" + strconv.FormatBool(values["enabled"].(bool)) + ","
		list += "\"endpoint_id\": " + formatTrafficPairFloatValue(values, "endpoint_id") + ","
		list += "\"ping_ipaddress\":\"" + unmarshalTrafficPairString(values, "ping_ipaddress") + "\","
		list += "\"ping_interface\":\"" + unmarshalTrafficPairString(values, "ping_interface") + "\","
		list += "\"ping_interval\": " + formatTrafficPairFloatValue(values, "ping_interval") + ","
		list = strings.TrimSuffix(list, ",") //remove last comma
		list += "},"

	}
	list = strings.TrimSuffix(list, ",") //remove last comma
	list += "]"
	return list
}

// handles missing properties and null values for strings
func unmarshalTrafficPairString(values map[string]interface{}, name string) string {
	if val, ok := values[name]; ok {
		if val != nil {
			return fmt.Sprintf("%v", val)
		}
	}
	return ""
}

func formatTrafficPairFloatValue(values map[string]interface{}, key string) string {
	value, ok := values[key]
	if ok && value != nil {
		if floatValue, ok := value.(float64); ok && floatValue != 0 {
			return strconv.FormatFloat(floatValue, 'f', -1, 64)
		}
	}
	return "null"
}
