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
	corev1 "k8s.io/api/core/v1"
)

func RouteCreateOrUpdate(instances corev1.PodList, route ho.Route) error {

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
		Address: fmt.Sprintf("%s:%d", route.Spec.Backend.Service, route.Spec.Backend.Port),
	}

	for _, instance := range instances.Items {
		url := fmt.Sprintf("http://%s:5555", instance.Status.PodIP)
		cli := hocli.New(url)

		err := cli.StartTransaction()
		if err != nil {
			return err
		}

		err = cli.AclCreateOrUpdate(acl)
		if err != nil {
			cli.DeleteTransaction()
			return err
		}

		err = cli.RuleCreateOrUpdate(rule)
		if err != nil {
			cli.DeleteTransaction()
			return err
		}

		err = cli.BackendCreateOrUpdate(backend)
		if err != nil {
			cli.DeleteTransaction()
			return err
		}

		err = cli.ServerCreateOrUpdate(server)
		if err != nil {
			cli.DeleteTransaction()
			return err
		}

		err = cli.CommitTransaction()
		if err != nil {
			return err
		}
	}

	return nil
}

func RouteDelete(instances corev1.PodList, route string) error {

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

	for _, instance := range instances.Items {
		url := fmt.Sprintf("http://%s:5555", instance.Status.PodIP)
		cli := hocli.New(url)

		err := cli.StartTransaction()
		if err != nil {
			return err
		}

		err = cli.RuleDelete(rule)
		if err != nil {
			cli.DeleteTransaction()
			return err
		}

		err = cli.AclDelete(acl)
		if err != nil {
			cli.DeleteTransaction()
			return err
		}

		err = cli.ServerDelete(server)
		if err != nil {
			cli.DeleteTransaction()
			return err
		}

		err = cli.BackendDelete(backend)
		if err != nil {
			cli.DeleteTransaction()
			return err
		}

		err = cli.CommitTransaction()
		if err != nil {
			return err
		}

	}

	return nil
}
