package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (c *Client) GetEbgpPeer(endpointId string, ebgpPeerId string) (EbgpPeer, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/ipsec/endpoints/%s/ebgp_peers/%s", c.HostURL, endpointId, ebgpPeerId), nil)
	if err != nil {
		log.Println(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		log.Println(err)
	}

	simplePeer := SimplifyGetEbpgJson(string(body))

	ebgPeer := EbgpPeer{}
	err = json.Unmarshal([]byte(simplePeer), &ebgPeer)
	if err != nil {
		log.Println(err)
	}

	return ebgPeer, nil
}

/*
func (c *Client) GetEbgpPeers() (EndpointResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/status/ipsec", c.HostURL), nil)
	if err != nil {
		log.Println(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		log.Println(err)
	}

	endpoints := EndpointResponse{}
	err = json.Unmarshal(body, &endpoints)
	if err != nil {
		log.Println(err)
	}

	return endpoints, nil
}
*/
func (c *Client) CreateEbgpPeer(endpointId string, ebgp_peer *EbgpPeer) (*EbgpPeer, error) {

	rb, err := json.Marshal(ebgp_peer)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/ipsec/endpoints/%s/ebgp_peers", c.HostURL, endpointId), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		log.Println(err)
	}

	simplePeer := SimplifyEbpgJson(string(body))

	newPeer := EbgpPeer{}
	err = json.Unmarshal([]byte(simplePeer), &newPeer)
	if err != nil {
		log.Println(err)
	}

	return &newPeer, nil
}

func (c *Client) UpdateEbgpPeer(endpointId string, ebgpPeerId string, ebgp_peer *EbgpPeer) (*EbgpPeer, error) {

	rb, err := json.Marshal(ebgp_peer)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/ipsec/endpoints/%s/ebgp_peers/%s", c.HostURL, endpointId, ebgpPeerId), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	_, err2 := c.doRequest(req)
	if err2 != nil {
		log.Println(err2)
	}

	getReq, err := http.NewRequest("GET", fmt.Sprintf("%s/ipsec/endpoints/%s/ebgp_peers/%s", c.HostURL, endpointId, ebgpPeerId), nil)
	if err != nil {
		log.Println(err)
	}

	body, err := c.doRequest(getReq)
	if err != nil {
		log.Println(err)
	}

	ebgPeer := EbgpPeer{}
	err = json.Unmarshal(body, &ebgPeer)
	if err != nil {
		log.Println(err)
	}

	return &ebgPeer, nil
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

func SimplifyEbpgJson(input string) string {
	var result map[string]interface{}
	json.Unmarshal([]byte(input), &result)

	var response = result["response"].(map[string]interface{})
	var peers = response["bgp_peers"].(map[string]interface{})
	var list string = " { "
	var lastPeer interface{} = nil

	for _, peer := range peers {
		lastPeer = peer
	}
	if lastPeer != nil {
		//list += "{ "
		values := lastPeer.(map[string]interface{})
		list += "\"id\":\"" + strconv.Itoa(int(values["id"].(float64))) + "\","
		list += "\"ipaddress\":\"" + UnmarshalString(values, "ipaddress") + "\","
		list += "\"asn\": " + strconv.Itoa(int(values["asn"].(float64))) + ","
		list += "\"local_asn_alias\":" + strconv.Itoa(int(values["local_asn_alias"].(float64))) + ","
		list += "\"keepalive_interval\":" + strconv.Itoa(int(values["keepalive_interval"].(float64))) + ","
		list += "\"hold_time\":\"" + strconv.Itoa(int(values["hold_time"].(float64))) + "\","
		list += "\"bgp_password\":\"" + UnmarshalString(values, "bgp_password") + "\","
		list += "\"access_list\":\"" + UnmarshalString(values, "access_list") + "\","
		list += "\"add_network_distance\":" + strconv.FormatBool(values["add_network_distance"].(bool)) + ","
		list += "\"add_network_distance_direction\":\"" + UnmarshalString(values, "add_network_distance_direction") + "\","
		list += "\"add_network_distance_hops\":" + strconv.Itoa(int(values["add_network_distance_hops"].(float64))) + ""
		//	list += " },"
	}
	//list = strings.TrimSuffix(list, ",") //remove last comma
	list += "}"
	return list
}

func SimplifyGetEbpgJson(input string) string {
	var result map[string]interface{}
	json.Unmarshal([]byte(input), &result)

	var values = result["response"].(map[string]interface{})
	//var peers = response["bgp_peers"].(map[string]interface{})
	var list string = " { "

	list += "\"id\":\"" + strconv.Itoa(int(values["id"].(float64))) + "\","
	list += "\"ipaddress\":\"" + UnmarshalString(values, "ipaddress") + "\","
	list += "\"asn\": " + strconv.Itoa(int(values["asn"].(float64))) + ","
	list += "\"local_asn_alias\":" + strconv.Itoa(int(values["local_asn_alias"].(float64))) + ","
	list += "\"keepalive_interval\":" + strconv.Itoa(int(values["keepalive_interval"].(float64))) + ","
	list += "\"hold_time\":\"" + strconv.Itoa(int(values["hold_time"].(float64))) + "\","
	list += "\"bgp_password\":\"" + UnmarshalString(values, "bgp_password") + "\","
	list += "\"access_list\":\"" + UnmarshalString(values, "access_list") + "\","
	list += "\"add_network_distance\":" + strconv.FormatBool(values["add_network_distance"].(bool)) + ","
	list += "\"add_network_distance_direction\":\"" + UnmarshalString(values, "add_network_distance_direction") + "\","
	list += "\"add_network_distance_hops\":" + strconv.Itoa(int(values["add_network_distance_hops"].(float64))) + ""
	//	list += " },
	//list = strings.TrimSuffix(list, ",") //remove last comma
	list += "}"
	return list
}

// handles missing properties and null values for strings
func UnmarshalString(values map[string]interface{}, name string) string {
	if val, ok := values[name]; ok {
		if val != nil {
			return fmt.Sprintf("%v", val)
		}
	}
	return ""
}
