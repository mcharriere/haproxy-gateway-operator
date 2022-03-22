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
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logger "sigs.k8s.io/controller-runtime/pkg/log"

	ho "github.com/mcharriere/giantswarm-task/api/v1alpha1"
)

// RouteReconciler reconciles a Route object
type RouteReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=*,resources=pods,verbs=get;list;watch
//+kubebuilder:rbac:groups=haproxy-opeartor.my.domain,resources=routes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=haproxy-opeartor.my.domain,resources=routes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=haproxy-opeartor.my.domain,resources=routes/finalizers,verbs=update

func (r *RouteReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logger.FromContext(ctx)

	log.Info("reconciling route")

	var haproxyInstances corev1.PodList
	if err := r.List(ctx, &haproxyInstances, client.MatchingLabels{"app": "haproxy-operator"}); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	var route ho.Route
	if err := r.Get(ctx, req.NamespacedName, &route); err != nil {
		if err := RouteDelete(haproxyInstances, req.Name); err != nil {
			return ctrl.Result{}, fmt.Errorf("could not delete route: %+v", err)
		}
		log.Info("deleted route", "route", req.NamespacedName)
		return ctrl.Result{}, nil
	}

	log.Info("adding new route", "route", req.NamespacedName)
	if err := RouteCreateOrUpdate(haproxyInstances, route); err != nil {
		return ctrl.Result{}, fmt.Errorf("could not create route: %+v", err)
	}

	log.Info("reconciled route")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RouteReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ho.Route{}).
		Complete(r)
}
