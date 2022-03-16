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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	haproxyopeartorv1alpha1 "github.com/mcharriere/giantswarm-task/api/v1alpha1"
)

// RouteReconciler reconciles a Route object
type RouteReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=haproxy-opeartor.my.domain,resources=routes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=haproxy-opeartor.my.domain,resources=routes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=haproxy-opeartor.my.domain,resources=routes/finalizers,verbs=update

func (r *RouteReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("route", req.NamespacedName)

	log.Info("reconciling route")

	var route haproxyopeartorv1alpha1.Route
	if err := r.Get(ctx, req.NamespacedName, &route); err != nil {
		log.Error(err, "unable to fetch route")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("reconciled route")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RouteReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&haproxyopeartorv1alpha1.Route{}).
		Complete(r)
}