package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (c *Client) GetRoute(routeId string) (RouteResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/routes", c.HostURL), nil)
	if err != nil {
		return RouteResponse{}, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return RouteResponse{}, err
	}

	simpleJson := SimplifyJson(string(body))
	routeResponse := RouteResponse{}
	err = json.Unmarshal([]byte(simpleJson), &routeResponse)
	if err != nil {
		return RouteResponse{}, err
	}

	returnRoutes := RouteResponse{}
	for _, r := range routeResponse.Routes {
		if r.ID == routeId {
			returnRoutes.Routes = []Route{r}
		}
	}

	return routeResponse, nil
}

func (c *Client) GetRoutes() (RouteResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/routes", c.HostURL), nil)
	if err != nil {
		return RouteResponse{}, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return RouteResponse{}, err
	}

	simpleJson := SimplifyJson(string(body))
	routeResponse := RouteResponse{}
	errUnmarshal := json.Unmarshal([]byte(simpleJson), &routeResponse)

	if err != nil {
		return RouteResponse{}, errUnmarshal
	}

	return routeResponse, nil
}

func (c *Client) CreateRoute(routeList []*Route) error {
	i := 0
	for _, rt := range routeList {
		route := Route{
			Cidr:        rt.Cidr,
			Description: rt.Description,
			Interface:   rt.Interface,
			Gateway:     rt.Gateway,
			Tunnel:      rt.Tunnel,
			Advertise:   rt.Advertise,
			Metric:      rt.Metric,
		}
		rb, err := json.Marshal(route)
		if err != nil {
			return err
		}
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/routes", c.HostURL), strings.NewReader(string(rb)))
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

func (c *Client) UpdateRoute(routeList []*Route) error {

	c.DeleteRoute()
	err := c.CreateRoute(routeList)

	return err
}

func (c *Client) DeleteRoute() error {

	req1, err := http.NewRequest("GET", fmt.Sprintf("%s/routes", c.HostURL), nil)
	if err != nil {
		return err
	}
	routes, err := c.doRequest(req1)
	if err != nil {
		return err
	}

	simpleJson := SimplifyJson(string(routes))
	routeResponse := RouteResponse{}
	err = json.Unmarshal([]byte(simpleJson), &routeResponse)
	if err != nil {
		return err
	}

	for _, r := range routeResponse.Routes {
		routeId := r.ID

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
