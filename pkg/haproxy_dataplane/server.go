package haproxy_dataplane

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ServerList struct {
	Version int      `json:"_version"`
	Data    []Server `json:"data"`
}

type Server struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Backend string `json:"-"`
}

func (c *Client) ServerCreate(server Server) error {
	url := fmt.Sprintf(
		URL_SERVER_ADD,
		c.Host,
		c.Transaction.Id,
		server.Backend,
	)

	j, err := json.Marshal(server)
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

func (c *Client) ServerGetList(backend string) ([]Server, error) {
	url := fmt.Sprintf(
		URL_SERVER_GET,
		c.Host,
		backend,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	ret, err := c.send(req)
	if err != nil {
		return nil, err
	}

	var data ServerList
	err = json.Unmarshal(ret, &data)
	if err != nil {
		return nil, err
	}

	return data.Data, nil
}

func (c *Client) ServerUpdate(server Server) error {
	url := fmt.Sprintf(
		URL_SERVER_REPLACE,
		c.Host,
		server.Name,
		c.Transaction.Id,
		server.Backend,
	)

	j, err := json.Marshal(server)
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

func (c *Client) ServerCreateOrUpdate(server Server) error {
	servers, err := c.ServerGetList(server.Backend)
	if err != nil {
		return err
	}

	for _, elem := range servers {
		if elem.Name == server.Name {
			err = c.ServerUpdate(server)
			if err != nil {
				return err
			}
			return nil
		}
	}

	err = c.ServerCreate(server)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ServerDelete(server Server) error {
	servers, err := c.ServerGetList(server.Backend)
	if err != nil {
		return err
	}

	for _, elem := range servers {
		if elem.Name == server.Name {
			url := fmt.Sprintf(
				URL_SERVER_DELETE,
				c.Host,
				server.Name,
				c.Transaction.Id,
				server.Backend,
			)

			req, err := http.NewRequest("DELETE", url, nil)
			if err != nil {
				return err
			}

			_, err = c.send(req)
			if err != nil {
				return fmt.Errorf("server %s: %+v", server.Name, err)
			}

			break
		}
	}

	return nil
}
