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

	corev1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logger "sigs.k8s.io/controller-runtime/pkg/log"

	blakeleadiov1alpha1 "github.com/blakelead/aks-route-controller/api/v1alpha1"
)

// RouteTableReconciler reconciles a RouteTable object
type RouteTableReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=blakelead.io,resources=routetables,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=blakelead.io,resources=routetables/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=blakelead.io,resources=routetables/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the RouteTable object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *RouteTableReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logger.FromContext(ctx)

	node := &corev1.Node{}

	err := r.Get(ctx, req.NamespacedName, node)
	if err != nil {
		log.Error(err, "Failed to get Node")
		return ctrl.Result{}, err
	}

	log.Info("Get Node", "Node.Name", node.Name, "node.Status.Addresses", node.Status.Addresses, "Node.Spec.PodCIDR", node.Spec.PodCIDR)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RouteTableReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&blakeleadiov1alpha1.RouteTable{}).
		Owns(&corev1.Node{}).
		Complete(r)
}
