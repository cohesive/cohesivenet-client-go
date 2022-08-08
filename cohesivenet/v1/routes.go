package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

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

func (c *Client) CreateRoute(route *Route) (RouteResponse, error) {

	rb, err := json.Marshal(route)
	if err != nil {
		log.Println(err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/routes", c.HostURL), strings.NewReader(string(rb)))

	if err != nil {
		log.Println(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		log.Println(err)
	}

	simpleJson := SimplifyJson(string(body))

	newRoute := RouteResponse{}
	errUnmarshal := json.Unmarshal([]byte(simpleJson), &newRoute)
	if err != nil {
		log.Println(errUnmarshal)
	}

	return newRoute, err
}

func (c *Client) DeleteRoute(routeId string) error {
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
}

func SimplifyJson(input string) string {
	var result map[string]interface{}
	json.Unmarshal([]byte(input), &result)

	var routes = result["response"].(map[string]interface{})
	var list string = "{\"response\" : [ "
	for _, route := range routes {
		list += "{ "
		values := route.(map[string]interface{})
		list += "\"id\":\"" + strconv.Itoa(int(values["id"].(float64))) + "\","
		list += "\"description\":\"" + unmarshalString(values, "description") + "\","
		list += "\"advertise\": " + strconv.FormatBool(values["advertise"].(bool)) + ","
		list += "\"enabled\":" + strconv.FormatBool(values["enabled"].(bool)) + ","
		list += "\"editable\":" + strconv.FormatBool(values["editable"].(bool)) + ","
		list += "\"cidr\":\"" + unmarshalString(values, "cidr") + "\","
		list += "\"gateway\":\"" + unmarshalString(values, "gateway") + "\","
		list += "\"netmask\":\"" + unmarshalString(values, "netmask") + "\","
		list += "\"table\":\"" + unmarshalString(values, "table") + "\","
		list += "\"interface\":\"" + unmarshalString(values, "interface") + "\","
		list += "\"metric\":" + strconv.Itoa(int(values["metric"].(float64))) + ""
		list += " },"
	}
	list = strings.TrimSuffix(list, ",") //remove last comma
	list += "]}"
	return list
}

// handles missing properties and null values for strings
func unmarshalString(values map[string]interface{}, name string) string {
	if val, ok := values[name]; ok {
		if val != nil {
			return fmt.Sprintf("%v", val)
		}
	}
	return ""
}
