/*
Copyright 2020 xutao1989103.

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
	"log"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1alpha1 "github.com/xutao1989103/oam-go/api/v1alpha1"
)

// ApplicationConfigurationReconciler reconciles a ApplicationConfiguration object
type ApplicationConfigurationReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=apps.oam.io,resources=applicationconfigurations,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps.oam.io,resources=applicationconfigurations/status,verbs=get;update;patch

func (r *ApplicationConfigurationReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("applicationconfiguration", req.NamespacedName)

	ctx := context.Background()
	obja := &appsv1alpha1.ApplicationConfiguration{}

	if err := r.Get(ctx, req.NamespacedName, obja); err != nil {
		log.Println(err, "unable to fetch New Object")
	} else {
		log.Println("fetch New Object:", obja.Spec.Foo)
	}

	return ctrl.Result{}, nil
}

func (r *ApplicationConfigurationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1alpha1.ApplicationConfiguration{}).
		Complete(r)
}
