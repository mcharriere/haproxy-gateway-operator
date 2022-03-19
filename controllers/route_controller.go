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

//+kubebuilder:rbac:groups=haproxy-opeartor.my.domain,resources=routes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=haproxy-opeartor.my.domain,resources=routes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=haproxy-opeartor.my.domain,resources=routes/finalizers,verbs=update

func (r *RouteReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logger.FromContext(ctx)

	log.Info("reconciling route")

	haproxyConfigName := client.ObjectKey{Name: "haproxyconfig-sample", Namespace: req.Namespace}
	var haproxyConfig ho.HaproxyConfig
	if err := r.Get(ctx, haproxyConfigName, &haproxyConfig); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	var route ho.Route
	if err := r.Get(ctx, req.NamespacedName, &route); err != nil {
		if i, is_there := find(haproxyConfig.Spec.Data.Frontends, req.Name); is_there {
			haproxyConfig.Spec.Data.Frontends = remove(haproxyConfig.Spec.Data.Frontends, i)
			log.Info("deleted route")
		} else {

			log.Error(err, "unable to fetch route. cleanup!")
			return ctrl.Result{}, client.IgnoreNotFound(err)
		}
	} else {

		frontend := ho.HaproxyConfigFrontend{
			Name:    route.Name,
			Backend: route.Name,
			Host:    route.Spec.Host,
		}

		if i, is_there := find(haproxyConfig.Spec.Data.Frontends, frontend.Name); is_there {
			log.Info("patching existing frontend", "frontend", frontend.Name)
			haproxyConfig.Spec.Data.Frontends[i] = frontend
		} else {
			log.Info("adding new frontend", "frontend", frontend.Name)
			haproxyConfig.Spec.Data.Frontends = append(haproxyConfig.Spec.Data.Frontends, frontend)
		}
	}

	if err := r.Update(ctx, &haproxyConfig); err != nil {
		return ctrl.Result{}, fmt.Errorf("could not write ReplicaSet: %+v", err)
	}

	log.Info("get haproxyconfig", "version", haproxyConfig.Spec.Version)

	log.Info("reconciled route")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RouteReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ho.Route{}).
		Complete(r)
}

func find(f []ho.HaproxyConfigFrontend, val string) (int, bool) {
	for i, frontend := range f {
		if frontend.Name == val {
			return i, true
		}
	}
	return -1, false
}

func remove(f []ho.HaproxyConfigFrontend, i int) []ho.HaproxyConfigFrontend {
	copy(f[i:], f[i+1:])
	return f[:len(f)-1]
}
