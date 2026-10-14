package instances

import (
	"slices"

	"github.com/lukas-sgx/borrow/api/v1alpha1"
)

type Instance struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

func isInstanceExist(mslist []string, name string) bool {
	return slices.Contains(mslist, name)
}

func getInstanceInfo(Microservice *v1alpha1.Microservice, name string) *Instance {
	if isInstanceExist(Microservice.Status.MsList, name) {
		return &Instance{
			Name:      name,
			Namespace: "sre-system",
		}
	}

	return nil
}
