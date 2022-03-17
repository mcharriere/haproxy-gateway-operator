package haproxy_dataplane

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

type HaproxyConfig struct {
	Version int               `json:"_version"`
	Data    HaproxyConfigData `json:"data"`
}

type HaproxyConfigData struct {
	Backends  []Backend  `json:"backends"`
	Frontends []Frontend `json:"frontends"`
}

type Backend struct {
	Name    string   `json:"name"`
	Servers []Server `json:"servers"`
}

type Server struct {
	Name    string `json:"name"`
	Port    int    `json:"port"`
	Address string `json:"address,omitempty"`
}

type Frontend struct {
	Name    string `json:"name"`
	Backend string `json:"backend"`
	Host    string `json:"host"`
}

func (c *Client) SaveConfig(config HaproxyConfig) error {
	tmpl, err := template.New("test").Parse(RAW_CONFIG)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, config.Data)

	// fmt.Println(buf.String())

	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/raw?skip_version=1", c.Host)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "text/plain; charset=UTF-8")

	_, err = c.send(req)
	return err
}
