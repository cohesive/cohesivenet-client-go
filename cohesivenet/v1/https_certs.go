package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *Client) UpdateHttpsCerts(cert, key string) (HttpsCertResponse, error) {

	certFile, err := ioutil.ReadFile(cert)
	if err != nil {
		return HttpsCertResponse{}, err
	}

	keyFile, err2 := ioutil.ReadFile(key)
	if err2 != nil {
		return HttpsCertResponse{}, err
	}

	cf := string(certFile)
	kf := string(keyFile)

	https := HttpsCert{
		Cert: cf,
		Key:  kf,
	}

	rb, err := json.Marshal(https)
	if err != nil {
		return HttpsCertResponse{}, err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/system/ssl/keypair", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return HttpsCertResponse{}, err
	}

	//body, err := c.doRequest(req)
	c.doRequest(req)
	if err != nil {
		return HttpsCertResponse{}, err
	}

	req2, err := http.NewRequest("PUT", fmt.Sprintf("%s/system/ssl/install", c.HostURL), nil)
	if err != nil {
		return HttpsCertResponse{}, err
	}
	body2, err := c.doRequest(req2)
	if err != nil {
		return HttpsCertResponse{}, err
	}

	httpsResponse := HttpsCertResponse{}
	err = json.Unmarshal([]byte(body2), &httpsResponse)
	if err != nil {
		return HttpsCertResponse{}, err
	}

	return httpsResponse, nil
}
