package cohesivenet

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
