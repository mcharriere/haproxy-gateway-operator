package haproxy_dataplane

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Host   string
	Client *http.Client
}

func New(host string) *Client {
	return &Client{
		Host:   host,
		Client: &http.Client{},
	}
}

func (c *Client) send(r *http.Request) ([]byte, error) {
	r.SetBasicAuth("admin", "securePassword")

	resp, err := c.Client.Do(r)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if !success(resp.StatusCode) {
		fmt.Println(resp.StatusCode)
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

func success(code int) bool {
	if code < 300 {
		return true
	}
	return false
}
