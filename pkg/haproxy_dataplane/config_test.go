package haproxy_dataplane

import (
	"testing"
)

func TestSaveConfig(t *testing.T) {
	cli := New("http://172.17.0.2:5555")

	config := HaproxyConfig{
		Version: 1,
		Data: HaproxyConfigData{
			Frontends: []Frontend{
				{
					Name:    "frontend1",
					Backend: "pepe",
					Host:    "pepe.com",
				},
			},
			Backends: []Backend{
				{
					Name: "pepe",
					Servers: []Server{
						{
							Name:    "pepe1",
							Port:    8080,
							Address: "10.0.0.1",
						}, {
							Name:    "pepe2",
							Port:    8080,
							Address: "10.0.0.2",
						},
					},
				},
			},
		},
	}

	err := cli.SaveConfig(config)

	if err != nil {
		t.Errorf("Received %v", err)
	}
}

func TestSuccess(t *testing.T) {
	var val bool

	val = success(200)
	if val == false {
		t.Errorf("200: Received %v", val)
	}
	val = success(202)
	if val == false {
		t.Errorf("202: Received %v", val)
	}
	val = success(402)
	if val == true {
		t.Errorf("402: Received %v", val)
	}
}
