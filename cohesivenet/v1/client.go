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

	// used to synchroize client requests i.e. create plugin image
	ReqLock	   sync.Mutex
}

// NewClient -
func NewClient(username, password, token, hostUrl *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    *hostUrl,
		Token:      *token,
		Username:   *username,
		Password:   *password,
	}

	//If username or password not provided, return empty client
	if username == nil || password == nil {
		return &c, nil
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
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
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
