package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func (c *Client) UpdateSslCerts(cert, key string) (SSLCertResponse, error) {

	certFile, err := ioutil.ReadFile(cert)
	if err != nil {
		log.Fatal(err)
	}

	keyFile, err2 := ioutil.ReadFile(key)
	if err2 != nil {
		log.Fatal(err)
	}

	cf := string(certFile)
	kf := string(keyFile)

	ssl := SSLCert{
		Cert: cf,
		Key:  kf,
	}

	rb, err := json.Marshal(ssl)
	if err != nil {
		log.Println(err)
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/system/ssl/keypair", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		log.Println(err)
	}

	//body, err := c.doRequest(req)
	c.doRequest(req)
	if err != nil {
		log.Println(err)
	}

	req2, err := http.NewRequest("PUT", fmt.Sprintf("%s/system/ssl/install", c.HostURL), nil)
	if err != nil {
		log.Println(err)
	}
	body2, err := c.doRequest(req2)
	if err != nil {
		log.Println(err)
	}

	sslResponse := SSLCertResponse{}
	errUnmarshal := json.Unmarshal([]byte(body2), &sslResponse)
	if err != nil {
		log.Println(errUnmarshal)
	}

	return sslResponse, nil
}