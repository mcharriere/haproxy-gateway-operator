package haproxy_dataplane

import (
	"fmt"
	"testing"
)

func TestAclGetList(t *testing.T) {
	cli := New("http://172.17.0.2:5555")

	var acls []Acl

	acls, err := cli.AclGetList("http")

	fmt.Println(acls)

	if err != nil {
		t.Errorf("Received %v", err)
	}
}

func TestAclCreate(t *testing.T) {
	cli := New("http://172.17.0.2:5555")

	acl := Acl{
		Name:      "frontend3",
		Criterion: "hdr(host)",
		Index:     0,
		Value:     "-i cholo.com",
		Frontend:  "http",
	}

	err := cli.StartTransaction()
	if err != nil {
		t.Errorf("Received %v", err)
	}

	err = cli.AclCreate(acl)
	if err != nil {
		t.Errorf("Received %v", err)
	}

	err = cli.CommitTransaction()
	if err != nil {
		t.Errorf("Received %v", err)
	}
}

func TestAclCreateOrUpdate(t *testing.T) {
	cli := New("http://172.17.0.2:5555")

	acl := Acl{
		Name:      "frontend2",
		Criterion: "hdr(host)",
		Value:     "-i coling.com",
		Frontend:  "http",
	}

	err := cli.StartTransaction()
	if err != nil {
		t.Errorf("Received %v", err)
		return
	}

	err = cli.AclCreateOrUpdate(acl)
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
