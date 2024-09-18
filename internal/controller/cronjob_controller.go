/*
Copyright 2024.

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

package controller

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	batchv1 "github.com/wxquare/simple-operator/api/v1"
)

// CronJobReconciler reconciles a CronJob object
type CronJobReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=batch.example.com,resources=cronjobs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=batch.example.com,resources=cronjobs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=batch.example.com,resources=cronjobs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CronJob object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.0/pkg/reconcile
func (r *CronJobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// TODO(user): your logic here
	// log := log.FromContext(ctx)

	// TODO(user): your logic here
	log.Info("my operator Reconcile called", "time", time.Now().Local().Hour())

	// var cronJob v1.CronJob
	// if err := r.Get(ctx, req.NamespacedName, &cronJob); err != nil {
	// 	log.Error(err, "unable to fetch CronJob")
	// 	return ctrl.Result{RequeueAfter: 10 * time.Second}, client.IgnoreNotFound(err)
	// }
	// // startTime := cronJob.Spec.Start
	// // endTime := cronJob.Spec.End
	// // replicas := cronJob.Spec.Replicas
	// log.Info("my operator Reconcile called", time.Now().Local().Hour())
	// log.Info(fmt.Sprintf("currentTIme: %d", currenHour), startTime, endTime, replicas)
	// if currenHour >= startTime && currenHour < endTime {
	// 	for _,deployment := cronJob.Spec.Deployments {
	// 		err := r.Get(ctx,types.NamespacedName{
	// 			Name: deploy.Name,
	// 			Namespace:deploy.Namespace,
	// 		},deployment)
	// 		if err != nil {
	// 			return ctrl.Result{},err
	// 		}

	// 		if deployment.Sepc.Replicas != replicas {

	// 		}
	// 	}
	// }
	return ctrl.Result{RequeueAfter: 10 * time.Second}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CronJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&batchv1.CronJob{}).
		Complete(r)
}
