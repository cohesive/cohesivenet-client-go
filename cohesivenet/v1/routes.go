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
		id := strconv.Itoa(int(values["id"].(float64)))
		metric := strconv.Itoa(int(values["metric"].(float64)))
		list += "\"id\":\"" + id + "\","
		list += "\"description\":\"" + values["description"].(string) + "\","
		list += "\"advertise\": " + strconv.FormatBool(values["advertise"].(bool)) + ","
		list += "\"enabled\":" + strconv.FormatBool(values["enabled"].(bool)) + ","
		list += "\"editable\":" + strconv.FormatBool(values["editable"].(bool)) + ","
		list += "\"cidr\":\"" + values["cidr"].(string) + "\","
		list += "\"gateway\":\"" + values["gateway"].(string) + "\","
		list += "\"netmask\":\"" + values["netmask"].(string) + "\","
		list += "\"table\":\"" + values["table"].(string) + "\","
		list += "\"interface\":\"" + values["interface"].(string) + "\","
		list += "\"metric\":" + metric + ""
		list += " },"
	}
	list = list[:len(list)-1]
	list += "]}"

	return list
}
