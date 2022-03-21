package haproxy_dataplane

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Transaction struct {
	Version int    `json:"_version"`
	Id      string `json:"id"`
	Status  string `json:"status"`
}

func (c *Client) StartTransaction() error {
	version, err := c.GetVersion()
	if err != nil {
		return err
	}

	url := fmt.Sprintf(URL_TRANSACTION_START, c.Host, *version)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	ret, err := c.send(req)
	if err != nil {
		return err
	}

	var data Transaction
	err = json.Unmarshal(ret, &data)
	if err != nil {
		return err
	}

	c.Transaction = data
	return nil
}

func (c *Client) CommitTransaction() error {
	url := fmt.Sprintf(URL_TRANSACTION_COMMIT, c.Host, c.Transaction.Id)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}

	_, err = c.send(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteTransaction() error {
	url := fmt.Sprintf(URL_TRANSACTION_DELETE, c.Host, c.Transaction.Id)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	_, err = c.send(req)
	if err != nil {
		return err
	}

	return nil
}
func (c *Client) GetVersion() (*int, error) {
	url := fmt.Sprintf(URL_VERSION_GET, c.Host)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	ret, err := c.send(req)
	if err != nil {
		return nil, err
	}

	var data int
	err = json.Unmarshal(ret, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
