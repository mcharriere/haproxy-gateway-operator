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
	logger "sigs.k8s.io/controller-runtime/pkg/log"

	haproxyopeartorv1alpha1 "github.com/mcharriere/giantswarm-task/api/v1alpha1"
)

// HaproxyConfigReconciler reconciles a HaproxyConfig object
type HaproxyConfigReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=haproxy-opeartor.my.domain,resources=haproxyconfigs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=haproxy-opeartor.my.domain,resources=haproxyconfigs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=haproxy-opeartor.my.domain,resources=haproxyconfigs/finalizers,verbs=update

func (r *HaproxyConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logger.FromContext(ctx)

	log.Info("reconciling haproxyconfig")
	// TODO(user): your logic here

	log.Info("reconciled haproxyconfig")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *HaproxyConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&haproxyopeartorv1alpha1.HaproxyConfig{}).
		Complete(r)
}

// func HaproxyConfigPatch(haproxyopeartorv1alpha1.HaproxyConfig) haproxyopeartorv1alpha1.HaproxyConfig {
//
// }
