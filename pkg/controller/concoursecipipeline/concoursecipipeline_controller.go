/*
Copyright 2019 TAKAISHI Ryo.

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

package concoursecipipeline

import (
	"context"
	"fmt"
	showksv1beta1 "github.com/cloudnativedaysjp/showks-concourseci-pipeline-operator/pkg/apis/showks/v1beta1"
	"github.com/cloudnativedaysjp/showks-concourseci-pipeline-operator/pkg/concourseci"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller")
var finalizerName = "finalizer.concourseci.showks.cloudnativedays.jp"

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new ConcourseCIPipeline Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	url := os.Getenv("CONCOURSECI_URL")
	target := os.Getenv("CONCOURSECI_TARGET")
	team := os.Getenv("CONCOURSECI_TEAM")
	username := os.Getenv("CONCOURSECI_USERNAME")
	password := os.Getenv("CONCOURSECI_PASSWORD")
	return add(mgr, newReconciler(mgr, concourseci.NewClient(url, target, team, username, password)))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager, ccClient concourseci.ConcourseCIClientInterface) reconcile.Reconciler {
	return &ReconcileConcourseCIPipeline{
		Client:   mgr.GetClient(),
		scheme:   mgr.GetScheme(),
		ccClient: ccClient,
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("concoursecipipeline-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to ConcourseCIPipeline
	err = c.Watch(&source.Kind{Type: &showksv1beta1.ConcourseCIPipeline{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create
	// Uncomment watch a Deployment created by ConcourseCIPipeline - change this for objects you create
	err = c.Watch(&source.Kind{Type: &appsv1.Deployment{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &showksv1beta1.ConcourseCIPipeline{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileConcourseCIPipeline{}

// ReconcileConcourseCIPipeline reconciles a ConcourseCIPipeline object
type ReconcileConcourseCIPipeline struct {
	client.Client
	scheme   *runtime.Scheme
	ccClient concourseci.ConcourseCIClientInterface
}

// Reconcile reads that state of the cluster for a ConcourseCIPipeline object and makes changes based on the state read
// and what is in the ConcourseCIPipeline.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  The scaffolding writes
// a Deployment as an example
// Automatically generate RBAC rules to allow the Controller to read and write Deployments
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=deployments/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=showks.cloudnativedays.jp,resources=concoursecipipelines,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=showks.cloudnativedays.jp,resources=concoursecipipelines/status,verbs=get;update;patch
func (r *ReconcileConcourseCIPipeline) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	instance := &showksv1beta1.ConcourseCIPipeline{}
	err := r.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	if instance.ObjectMeta.DeletionTimestamp.IsZero() {
		if err := r.setFinalizer(instance); err != nil {
			return reconcile.Result{}, err
		}
	} else {
		return r.runFinalizer(instance)
	}

	target := instance.Spec.Target
	pipeline := instance.Spec.Pipeline
	manifest := instance.Spec.Manifest

	err = r.ccClient.SetPipeline(target, pipeline, manifest)
	if err != nil {
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}
func (r *ReconcileConcourseCIPipeline) setFinalizer(instance *showksv1beta1.ConcourseCIPipeline) error {
	fmt.Println("setFinalizer")
	if !containsString(instance.ObjectMeta.Finalizers, finalizerName) {
		instance.ObjectMeta.Finalizers = append(instance.ObjectMeta.Finalizers, finalizerName)
		if err := r.Update(context.Background(), instance); err != nil {
			return err
		}
	}

	return nil
}

func (r *ReconcileConcourseCIPipeline) runFinalizer(instannce *showksv1beta1.ConcourseCIPipeline) (reconcile.Result, error) {
	fmt.Println("runFinalizer")
	if containsString(instannce.ObjectMeta.Finalizers, finalizerName) {
		if err := r.deleteExternalDependency(instannce); err != nil {
			return reconcile.Result{}, err
		}

		instannce.ObjectMeta.Finalizers = removeString(instannce.ObjectMeta.Finalizers, finalizerName)
		if err := r.Update(context.Background(), instannce); err != nil {
			return reconcile.Result{}, err
		}
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileConcourseCIPipeline) deleteExternalDependency(instance *showksv1beta1.ConcourseCIPipeline) error {
	fmt.Println("deleteExternalDependency")
	return r.ccClient.DestroyPipeline(instance.Spec.Target, instance.Spec.Pipeline)
}

func containsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func removeString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}
