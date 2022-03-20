package haproxy_dataplane

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RuleList struct {
	Version int    `json:"_version"`
	Data    []Rule `json:"data"`
}

type Rule struct {
	Backend   string `json:"name"`
	Condition string `json:"cond"`
	Acl       string `json:"cond_test"`
	Index     int    `json:"index"`
	Frontend  string `json:"-"`
}

func (c *Client) RuleCreate(rule Rule) error {
	url := fmt.Sprintf(
		URL_RULE_ADD,
		c.Host,
		c.Transaction.Id,
		rule.Frontend,
	)
	rule.Index = 0

	j, err := json.Marshal(rule)
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

func (c *Client) RuleGetList(frontend string) ([]Rule, error) {
	url := fmt.Sprintf(
		URL_RULE_GET,
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

	var data RuleList
	err = json.Unmarshal(ret, &data)
	if err != nil {
		return nil, err
	}

	return data.Data, nil
}

func (c *Client) RuleUpdate(rule Rule) error {
	url := fmt.Sprintf(
		URL_RULE_REPLACE,
		c.Host,
		rule.Index,
		c.Transaction.Id,
		rule.Frontend,
	)

	j, err := json.Marshal(rule)
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

func (c *Client) RuleCreateOrUpdate(rule Rule) error {
	var rules []Rule
	rules, err := c.RuleGetList(rule.Frontend)
	if err != nil {
		return err
	}

	for _, elem := range rules {
		if elem.Acl == rule.Acl {
			rule.Index = elem.Index
			err = c.RuleUpdate(rule)
			if err != nil {
				return err
			}
			return nil
		}
	}

	err = c.RuleCreate(rule)
	if err != nil {
		return err
	}
	return nil
}
