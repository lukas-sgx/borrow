package operatools

import (
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetClient(conciliator *Conciliator, name string, namespace string, object client.Object) {
	ressource := types.NamespacedName{
		Name:      name,
		Namespace: namespace,
	}

	if err := conciliator.Client.Get(conciliator.Ctx, ressource, object); err != nil {
		conciliator.Log.Error(err, "")
		return
	}

	conciliator.Object = object
}
