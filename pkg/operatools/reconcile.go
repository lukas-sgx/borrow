package operatools

import (
	"context"
	"fmt"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Reconciler interface {
	client.Reader
	Process(conciliator *Conciliator) error
}

func Reconcile(r Reconciler, c client.Client, req ctrl.Request, ctx context.Context, object client.Object) (ctrl.Result, error) {
	if err := r.Get(ctx, req.NamespacedName, object); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	conciliator := newConciliator(ctrl.Log, object)

	conciliator.Log.Info(fmt.Sprintf("Process %s", req.Name))
	if err := r.Process(conciliator); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}
