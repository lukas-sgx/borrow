package instances

import (
	"fmt"
	"net/http"
	"strings"
)

func Add(response http.ResponseWriter, req *http.Request) {
	response.Header().Add("Content-type", "application/json")
	response.WriteHeader(http.StatusBadRequest)
}

func Delete(response http.ResponseWriter, req *http.Request) {
	response.Header().Add("Content-type", "application/json")
	response.WriteHeader(http.StatusCreated)
}

func Get(response http.ResponseWriter, req *http.Request) {
	path := strings.Split(req.URL.EscapedPath(), "/")

	if len(path) == 2 {
		instanceName := path[1]

		fmt.Println(instanceName)
	}

	response.Header().Add("Content-type", "application/json")
	response.WriteHeader(http.StatusCreated)
}
