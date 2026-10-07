package operatools

import (
	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Conciliator struct {
	Log    logr.Logger
	Object client.Object
}

func newConciliator(log logr.Logger, object client.Object) *Conciliator {
	return &Conciliator{
		Log:    log,
		Object: object,
	}
}
