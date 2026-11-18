package cohesivenetv1

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Username   string
	Password   string

	// used to synchronize client requests i.e. create plugin image
	ReqLock sync.Mutex
}

// NewClient creates a client. When insecureSkipVerify is true the server's TLS
// certificate is NOT verified; it should only be enabled by
// explicit user opt-in. Each client gets its own transport; http.DefaultTransport
// is never modified.
func NewClient(username, password, token, hostUrl *string, insecureSkipVerify bool) (*Client, error) {
	deref := func(s *string) string {
		if s == nil {
			return ""
		}
		return *s
	}

	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: insecureSkipVerify}

	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second, Transport: transport},
		HostURL:    deref(hostUrl),
		Token:      deref(token),
		Username:   deref(username),
		Password:   deref(password),
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	if len(c.Token) != 0 {
		req.Header.Add("Api-Token", c.Token)
	} else {
		req.SetBasicAuth(c.Username, c.Password)
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

func (c *Client) doRequestWithTextHeader(req *http.Request) ([]byte, error) {
	if len(c.Token) != 0 {
		req.Header.Add("Api-Token", c.Token)
	} else {
		req.SetBasicAuth(c.Username, c.Password)
	}
	req.Header.Add("Content-Type", "text/plain")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
