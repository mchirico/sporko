/*


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

	mice "github.com/mchirico/sporko/pkg/apis/groupie/v1alpha1"
)

// MouseReconciler reconciles a Mouse object
type MouseReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=,resources=mice,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=,resources=mice/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=,resources=mice/finalizers,verbs=update

func (r *MouseReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("mouse", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *MouseReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mice.Mouse{}).
		Complete(r)
}
