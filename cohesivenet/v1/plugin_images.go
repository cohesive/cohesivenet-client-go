package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

/*
func (c *Client) GetContainerNetwork() (ContainerSystem, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/container_system", c.HostURL), nil)
	if err != nil {
		log.Println(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		log.Println(err)
	}

	containerResponse := ContainerSystem{}
	errUnmarshal := json.Unmarshal([]byte(body), &containerResponse)
	if err != nil {
		log.Println(errUnmarshal)

	}
	return containerResponse, nil
}
*/
/*
func (c *Client) GetRoute(routeId string) (RouteResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/routes", c.HostURL), nil)
	if err != nil {
		log.Println(err)
	}
	body, err := c.doRequest(req)
	if err != nil {
		log.Println(err)
	}

	simpleJson := SimplifyJson(string(body))
	routeResponse := RouteResponse{}
	errUnmarshal := json.Unmarshal([]byte(simpleJson), &routeResponse)
	if err != nil {
		log.Println(errUnmarshal)
	}

	returnRoutes := RouteResponse{}
	for _, r := range returnRoutes.Routes {
		if r.ID == routeId {
			routeResponse.Routes = []Route{r}
		}
	}

	return routeResponse, nil
}

func (c *Client) GetRoutes() (RouteResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/routes", c.HostURL), nil)
	if err != nil {
		log.Println(err)
	}
	body, err := c.doRequest(req)
	if err != nil {
		log.Println(err)
	}

	simpleJson := SimplifyJson(string(body))
	routeResponse := RouteResponse{}
	errUnmarshal := json.Unmarshal([]byte(simpleJson), &routeResponse)

	if err != nil {
		log.Println(errUnmarshal)
	}

	return routeResponse, nil
}
*/
func (c *Client) CreateImage(image *PluginImage) (ImageResponse, error) {

	rb, err := json.Marshal(image)
	if err != nil {
		log.Println(err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/container_system/images", c.HostURL), strings.NewReader(string(rb)))

	if err != nil {
		log.Println(err)
	}

	body, err2 := c.doRequest(req)
	if err2 != nil {
		log.Println(err2)
	}

	createImageResponse := CreateImageResponse{}
	errUnmarshal := json.Unmarshal([]byte(body), &createImageResponse)
	if err != nil {
		log.Println(errUnmarshal)
	}
	uuid := createImageResponse.Import_uuid

	getReq, err := http.NewRequest("GET", fmt.Sprintf("%s/container_system/images", c.HostURL), nil)
	if err != nil {
		log.Println(err)
	}

	getBody, err := c.doRequest(getReq)
	if err != nil {
		log.Println(err)
	}

	simpleJson := SimplifyImageJson(string(getBody), uuid)

	newImage := ImageResponse{}
	errUnmarshall := json.Unmarshal([]byte(simpleJson), &newImage)
	if err != nil {
		log.Println(errUnmarshall)
	}

	return newImage, err
}

/*
func (c *Client) DeleteImage() error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/routes/%s", c.HostURL, routeId), nil)
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
*/

func SimplifyImageJson(input string, uuid string) string {
	var result map[string]interface{}
	json.Unmarshal([]byte(input), &result)

	var response = result["response"].(map[string]interface{})
	var peers = response["images"].([]interface{})
	var list string = "{\"response\" : [ "
	var lastPeer interface{} = nil

	for _, peer := range peers {
		if response["id"] == uuid {
			lastPeer = peer
		}
		if lastPeer != nil {
			list += "{ "
			values := lastPeer.(map[string]interface{})
			list += "\"id\":\"" + unmarshalImageString(values, "id") + "\","
			list += "\"image_name\":\"" + unmarshalImageString(values, "image_name") + "\","
			list += "\"status\":\"" + unmarshalImageString(values, "status") + "\","
			list += "\"status_msg\":\"" + unmarshalImageString(values, "status_msg") + "\","
			list += "\"import_id\":\"" + unmarshalImageString(values, "import_id") + "\","
			list += "\"created\":\"" + unmarshalImageString(values, "created") + "\"," // date
			list += "\"description\":\"" + unmarshalImageString(values, "description") + "\","
			list += "\"tag_name\":\"" + unmarshalImageString(values, "tag_name") + "\","
			list += "\"import_uuid\":\"" + uuid + "\","
			list += " },"
		}
	}
	//list = strings.TrimSuffix(list, ",") //remove last comma
	list += "]}"
	return list
}

// handles missing properties and null values for strings
func unmarshalImageString(values map[string]interface{}, name string) string {
	if val, ok := values[name]; ok {
		if val != nil {
			return fmt.Sprintf("%v", val)
		}
	}
	return ""
}
