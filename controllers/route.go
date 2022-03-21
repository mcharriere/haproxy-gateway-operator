/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"fmt"

	ho "github.com/mcharriere/giantswarm-task/api/v1alpha1"
	hocli "github.com/mcharriere/giantswarm-task/pkg/haproxy_dataplane"
)

func RouteCreateOrUpdate(route ho.Route) error {
	cli := hocli.New("http://172.17.0.2:5555")

	acl := hocli.Acl{
		Name:      route.Name,
		Criterion: "hdr(host)",
		Value:     fmt.Sprintf("-i %s", route.Spec.Host),
		Frontend:  "http",
	}

	rule := hocli.Rule{
		Backend:   route.Name,
		Acl:       acl.Name,
		Condition: "if",
		Frontend:  "http",
	}

	backend := hocli.Backend{
		Name: route.Name,
	}

	server := hocli.Server{
		Backend: backend.Name,
		Name:    "service",
		Address: "10.0.0.1:8080",
	}

	err := cli.StartTransaction()
	if err != nil {
		return err
	}

	err = cli.AclCreateOrUpdate(acl)
	if err != nil {
		return err
	}

	err = cli.RuleCreateOrUpdate(rule)
	if err != nil {
		return err
	}

	err = cli.BackendCreateOrUpdate(backend)
	if err != nil {
		return err
	}

	err = cli.ServerCreateOrUpdate(server)
	if err != nil {
		return err
	}

	err = cli.CommitTransaction()
	if err != nil {
		return err
	}

	return nil
}

func RouteDelete(route string) error {
	cli := hocli.New("http://172.17.0.2:5555")

	acl := hocli.Acl{
		Name:     route,
		Frontend: "http",
	}

	rule := hocli.Rule{
		Backend:  route,
		Acl:      route,
		Frontend: "http",
	}

	backend := hocli.Backend{
		Name: route,
	}

	server := hocli.Server{
		Backend: route,
		Name:    "service",
	}

	err := cli.StartTransaction()
	if err != nil {
		return err
	}

	err = cli.RuleDelete(rule)
	if err != nil {
		return err
	}

	err = cli.AclDelete(acl)
	if err != nil {
		return err
	}

	err = cli.ServerDelete(server)
	if err != nil {
		return err
	}

	err = cli.BackendDelete(backend)
	if err != nil {
		return err
	}

	err = cli.CommitTransaction()
	if err != nil {
		return err
	}

	return nil
}
