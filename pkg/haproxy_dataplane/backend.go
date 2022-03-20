package haproxy_dataplane

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type BackendList struct {
	Version int       `json:"_version"`
	Data    []Backend `json:"data"`
}

type Backend struct {
	Name string `json:"name"`
}

func (c *Client) BackendCreate(backend Backend) error {
	url := fmt.Sprintf(
		URL_BACKEND_ADD,
		c.Host,
		c.Transaction.Id,
	)

	j, err := json.Marshal(backend)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	_, err = c.send(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) BackendGetList() ([]Backend, error) {
	url := fmt.Sprintf(
		URL_BACKEND_GET,
		c.Host,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	ret, err := c.send(req)
	if err != nil {
		return nil, err
	}

	var data BackendList
	err = json.Unmarshal(ret, &data)
	if err != nil {
		return nil, err
	}

	return data.Data, nil
}

func (c *Client) BackendUpdate(backend Backend) error {
	url := fmt.Sprintf(
		URL_BACKEND_REPLACE,
		c.Host,
		backend.Name,
		c.Transaction.Id,
	)

	j, err := json.Marshal(backend)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	_, err = c.send(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) BackendCreateOrUpdate(backend Backend) error {
	var backends []Backend
	backends, err := c.BackendGetList()
	if err != nil {
		return err
	}

	for _, elem := range backends {
		if elem.Name == backend.Name {
			err = c.BackendUpdate(backend)
			if err != nil {
				return err
			}
			return nil
		}
	}

	err = c.BackendCreate(backend)
	if err != nil {
		return err
	}
	return nil
}
