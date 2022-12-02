package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func (c *Client) UpdateHttpsCertByValue(cert, key string) (HttpsCertResponse, error) {

	cf := string(cert)
	kf := string(key)

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

func (c *Client) UpdateHttpsCertsByFilepath(cert_file, key_file string) (HttpsCertResponse, error) {

	certFile, err := ioutil.ReadFile(cert_file)
	if err != nil {
		return HttpsCertResponse{}, err
	}

	keyFile, err2 := ioutil.ReadFile(key_file)
	if err2 != nil {
		return HttpsCertResponse{}, err
	}
	log.Println(certFile)
	log.Println(keyFile)

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
