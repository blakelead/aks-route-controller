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

//+kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch
//+kubebuilder:rbac:groups=blakelead.io,resources=routetables,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=blakelead.io,resources=routetables/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=blakelead.io,resources=routetables/finalizers,verbs=update

// Reconcile watches for Node events and updates Route Table accordingly
func (r *RouteTableReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logger.FromContext(ctx)

	rt := &blakeleadiov1alpha1.RouteTable{}
	err := r.Get(ctx, req.NamespacedName, rt)
	if err != nil {
		log.Error(err, "failed to get RouteTable CR")
		return ctrl.Result{}, err
	}

	nodes := &corev1.NodeList{}
	err = r.List(ctx, nodes)
	if err != nil {
		log.Error(err, "failed to get Node list")
		return ctrl.Result{}, err
	}

	log.Info("get Route Table CR", "route.table", rt)
	log.Info("get Nodes", "node.list", nodes)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RouteTableReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&blakeleadiov1alpha1.RouteTable{}).
		Complete(r)
}
