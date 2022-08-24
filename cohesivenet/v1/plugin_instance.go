package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func (c *Client) GetInstance(instanceUuid string) (InstanceResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/container_system/containers", c.HostURL), nil)
	if err != nil {
		log.Println(err)
	}
	body, err := c.doRequest(req)
	if err != nil {
		log.Println(err)
	}

	simpleJson := SimplifyInstanceJson(string(body), instanceUuid)

	instanceResponse := InstanceResponse{}
	errUnmarshal := json.Unmarshal([]byte(simpleJson), &instanceResponse)
	if err != nil {
		log.Println(errUnmarshal)
	}

	pluginInstance := InstanceResponse{}
	for _, i := range instanceResponse.Instances {
		if i.ID == instanceUuid {
			pluginInstance.Instances = []Instance{i}
		}
	}

	return pluginInstance, nil
}

/*
reverted to "container_system/containers" api, need to revist
func (c *Client) CreateInstance(instance *CreatePluginInstance) (CreatePluginInstanceResponse, error) {

	rb, err := json.Marshal(instance)
	if err != nil {
		log.Println(err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/plugin-instances", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		log.Println(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		log.Println(err)
		//log.Fatalf("call failed with '%s'\n", err)
	}

	newInstance := CreatePluginInstanceResponse{}
	errUnmarshal := json.Unmarshal([]byte(body), &newInstance)
	if err != nil {
		log.Println(errUnmarshal)
	}

	return newInstance, err
}
*/
func (c *Client) CreateInstance(instance *CreatePluginInstance) (CreatePluginInstanceResponse, error) {

	rb, err := json.Marshal(instance)
	if err != nil {
		log.Println(err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/container_system/containers", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		log.Println(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		log.Println(err)
	}

	newInstance := CreatePluginInstanceResponse{}
	errUnmarshal := json.Unmarshal([]byte(body), &newInstance)
	if err != nil {
		log.Println(errUnmarshal)
	}

	return newInstance, err
}

func (c *Client) DeleteInstance(instanceUuid string) error {
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/container_system/containers/%s", c.HostURL, instanceUuid), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}
	if string(body) == "Stopped" {
		return err
	}

	req2, err := http.NewRequest("DELETE", fmt.Sprintf("%s/container_system/containers/%s", c.HostURL, instanceUuid), nil)
	if err != nil {
		return err
	}

	body2, err := c.doRequest(req2)
	if err != nil {
		return err
	}

	if string(body2) == "Deleted" {
		return err
	}

	return nil
}

func SimplifyInstanceJson(input string, import_uuid string) string {
	var result map[string]interface{}
	json.Unmarshal([]byte(input), &result)

	var response = result["response"].(map[string]interface{})
	var instances = response["containers"].([]interface{})
	var list string = "{\"response\" : [ "

	for _, instance := range instances {
		list += "{ "
		values := instance.(map[string]interface{})
		list += "\"Id\":\"" + unmarshalInstanceString(values, "Id") + "\","
		list += "\"Image\":\"" + unmarshalInstanceString(values, "Image") + "\","
		list += "\"Hostname\":\"" + unmarshalInstanceString(values, "Hostname") + "\","
		list += "\"Ipaddress\":\"" + unmarshalInstanceString(values, "IPAddress") + "\","
		list += "\"Path\":\"" + unmarshalInstanceString(values, "Path") + "\","
		list += "\"Hostname\":\"" + values["Config"].(map[string]interface{})["Hostname"].(string) + "\","
		list += "\"Ipaddress\":\"" + values["NetworkSettings"].(map[string]interface{})["Networks"].(map[string]interface{})["cohesive.net"].(map[string]interface{})["IPAddress"].(string) + "\","
		list = strings.TrimSuffix(list, ",") //remove last comma
		list += " },"

	}
	list = strings.TrimSuffix(list, ",") //remove last comma
	list += "]}"
	return list
}

// handles missing properties and null values for strings
func unmarshalInstanceString(values map[string]interface{}, name string) string {
	if val, ok := values[name]; ok {
		if val != nil {
			return fmt.Sprintf("%v", val)
		}
	}
	return ""
}
