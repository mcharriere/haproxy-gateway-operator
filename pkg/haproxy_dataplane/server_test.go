package haproxy_dataplane

import (
	"fmt"
	"testing"
)

func TestServerGetList(t *testing.T) {
	cli := New("http://172.17.0.2:5555")

	var servers []Server

	servers, err := cli.ServerGetList("backend1")

	fmt.Println(servers)

	if err != nil {
		t.Errorf("Received %v", err)
	}
}

func TestServerCreateOrUpdate(t *testing.T) {
	cli := New("http://172.17.0.2:5555")

	server := Server{
		Backend: "backend1",
		Name:    "server1",
		Address: "10.0.0.1:8080",
	}

	err := cli.StartTransaction()
	if err != nil {
		t.Errorf("Received %v", err)
		return
	}

	err = cli.ServerCreateOrUpdate(server)
	if err != nil {
		t.Errorf("Received %v", err)
		return
	}

	err = cli.CommitTransaction()
	if err != nil {
		t.Errorf("Received %v", err)
		return
	}
}
