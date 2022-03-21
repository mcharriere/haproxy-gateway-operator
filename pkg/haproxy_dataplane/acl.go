package haproxy_dataplane

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AclList struct {
	Version int   `json:"_version"`
	Data    []Acl `json:"data"`
}

type Acl struct {
	Name      string `json:"acl_name"`
	Criterion string `json:"criterion"`
	Index     int    `json:"index"`
	Value     string `json:"value"`
	Frontend  string `json:"-"`
}

func (c *Client) AclCreate(acl Acl) error {
	url := fmt.Sprintf(
		URL_ACL_ADD,
		c.Host,
		c.Transaction.Id,
		acl.Frontend,
	)
	acl.Index = 0

	j, err := json.Marshal(acl)
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

func (c *Client) AclGetList(frontend string) ([]Acl, error) {
	url := fmt.Sprintf(
		URL_ACL_GET,
		c.Host,
		frontend,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	ret, err := c.send(req)
	if err != nil {
		return nil, err
	}

	var data AclList
	err = json.Unmarshal(ret, &data)
	if err != nil {
		return nil, err
	}

	return data.Data, nil
}

func (c *Client) AclUpdate(acl Acl) error {
	url := fmt.Sprintf(
		URL_ACL_REPLACE,
		c.Host,
		acl.Index,
		c.Transaction.Id,
		acl.Frontend,
	)

	j, err := json.Marshal(acl)
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

func (c *Client) AclCreateOrUpdate(acl Acl) error {
	acls, err := c.AclGetList(acl.Frontend)
	if err != nil {
		return err
	}

	for _, elem := range acls {
		if elem.Name == acl.Name {
			acl.Index = elem.Index
			err = c.AclUpdate(acl)
			if err != nil {
				return err
			}
			return nil
		}
	}

	err = c.AclCreate(acl)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) AclDelete(acl Acl) error {
	acls, err := c.AclGetList(acl.Frontend)
	if err != nil {
		return err
	}

	for _, elem := range acls {
		if elem.Name == acl.Name {
			acl.Index = elem.Index

			url := fmt.Sprintf(
				URL_ACL_DELETE,
				c.Host,
				acl.Index,
				c.Transaction.Id,
				acl.Frontend,
			)

			req, err := http.NewRequest("DELETE", url, nil)
			if err != nil {
				return err
			}

			_, err = c.send(req)
			if err != nil {
				return err
			}

			break
		}
	}

	return nil
}
