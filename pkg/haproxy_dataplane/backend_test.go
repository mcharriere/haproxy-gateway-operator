package haproxy_dataplane

import (
	"fmt"
	"testing"
)

func TestBackendGetList(t *testing.T) {
	cli := New("http://172.17.0.2:5555")

	var backends []Backend

	backends, err := cli.BackendGetList()

	fmt.Println(backends)

	if err != nil {
		t.Errorf("Received %v", err)
	}
}

func TestBackendCreateOrUpdate(t *testing.T) {
	cli := New("http://172.17.0.2:5555")

	backend := Backend{
		Name: "backend1",
	}

	err := cli.StartTransaction()
	if err != nil {
		t.Errorf("Received %v", err)
		return
	}

	err = cli.BackendCreateOrUpdate(backend)
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
