/*
Copyright 2021 Filecoin Infra Team.

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

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	lotusv1alpha1 "github.com/coryschwartz/lotus-operator/api/v1alpha1"
)

// MinerReconciler reconciles a Miner object
type MinerReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=lotus.filecoin.io,resources=miners,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=lotus.filecoin.io,resources=miners/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=lotus.filecoin.io,resources=miners/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Miner object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.0/pkg/reconcile
// +kubebuilder:rbac:groups=lotus.filecoin.io,resources=miners,verbs=get;list;watch;create;patch;delete
// +kubebuilder:rbac:groups=lotus.filecoin.io,resources=miners/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=lotus.filecoin.io,resources=miners/finalizers,verbs=update
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;
func (r *MinerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = r.Log.WithValues("miner", req.NamespacedName)
	miner := lotusv1alpha1.Miner{}
	if err := r.Get(ctx, req.NamespacedName, &miner); err != nil {
		r.Log.Error(err, "can't get miner for reconsiliation.")
		return ctrl.Result{}, err
	}
	r.Log.Info("reconciled.")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MinerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&lotusv1alpha1.Miner{}).
		Complete(r)
}
