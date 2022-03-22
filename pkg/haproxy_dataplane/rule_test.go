package haproxy_dataplane

import (
	"fmt"
	"testing"
)

func TestRuleGetList(t *testing.T) {
	cli := New("http://172.17.0.2:5555")

	var rules []Rule

	rules, err := cli.RuleGetList("http")

	fmt.Println(rules)

	if err != nil {
		t.Errorf("Received %v", err)
	}
}

func TestRuleCreateOrUpdate(t *testing.T) {
	cli := New("http://172.17.0.2:5555")

	acl := Acl{
		Name:      "frontend3",
		Criterion: "hdr(host)",
		Index:     0,
		Value:     "-i test3.com",
		Frontend:  "http",
	}

	rule := Rule{
		Backend:   "backend1",
		Acl:       acl.Name,
		Index:     0,
		Condition: "if",
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

	err = cli.RuleCreateOrUpdate(rule)
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
