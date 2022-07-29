package cohesivenet

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func (c *Client) GetRoutes() (RouteResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/routes", c.HostURL), nil)
	if err != nil {
		log.Println(err)
	}
	body, err := c.doRequest(req)
	if err != nil {
		log.Println(err)
	}

	routeResponse := RouteResponse{}
	//var routeResponse map[string]interface{}

	errUnmarshal := json.Unmarshal([]byte(body), &routeResponse)

	if err != nil {
		log.Println(errUnmarshal)
	}

	return routeResponse, nil
}

func (c *Client) CreateRoute(route *Route) (*RouteResponse, error) {

	rb, err := json.Marshal(route)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/routes", c.HostURL), strings.NewReader(string(rb)))

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		log.Println(err)
	}

	newRoute := RouteResponse{}
	err = json.Unmarshal(body, &newRoute)
	if err != nil {
		log.Println(err)
	}

	return &newRoute, nil
}
