package operatools

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Reconciler interface {
	client.Reader
	Process() error
}

func Reconcile(r Reconciler, req ctrl.Request, ctx context.Context, object client.Object) (ctrl.Result, error) {
	if err := r.Get(ctx, req.NamespacedName, object); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if err := r.Process(); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}
