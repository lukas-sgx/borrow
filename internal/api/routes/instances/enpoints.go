package instances

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/lukas-sgx/borrow/api/v1alpha1"
	"github.com/lukas-sgx/borrow/pkg/operatools"
)

type Endpoints struct {
	Conciliator *operatools.Conciliator
}

func NewEndpoints(conciliator *operatools.Conciliator) *Endpoints {
	operatools.GetClient(conciliator, "microservice-sample", "sre-system", &v1alpha1.Microservice{})

	return &Endpoints{
		Conciliator: conciliator,
	}
}

func (e *Endpoints) Add(response http.ResponseWriter, req *http.Request) {
	response.Header().Set("Content-type", "application/json")
	response.WriteHeader(http.StatusOK)
}

func (e *Endpoints) Delete(response http.ResponseWriter, req *http.Request) {
	response.Header().Set("Content-type", "application/json")
	response.WriteHeader(http.StatusOK)
}

func (e *Endpoints) Get(response http.ResponseWriter, req *http.Request) {
	path := strings.Split(req.URL.EscapedPath(), "/")

	if len(path) != 2 {
		internalServerError(response, errors.New("malformed request"))
		return
	}
	instanceName := path[1]

	obj, ok := e.Conciliator.Object.(*v1alpha1.Microservice)
	if !ok {
		internalServerError(response, errors.New("v1alpha1.Microservice: object nil"))
		return
	}

	data, err := json.Marshal(getInstanceInfo(obj, instanceName))
	if err != nil {
		internalServerError(response, err)
		return
	}

	response.Header().Set("Content-type", "application/json")
	response.WriteHeader(http.StatusOK)

	if _, err := response.Write(data); err != nil {
		internalServerError(response, err)
	}
}

func internalServerError(response http.ResponseWriter, err error) {
	response.Header().Set("Content-type", "application/json")
	response.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(response).Encode(map[string]string{"error": err.Error()}); err != nil {
		internalServerError(response, err)
	}
}
