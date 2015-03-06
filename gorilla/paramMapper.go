package gorilla

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Mapper maps gorilla params from request
func Mapper(req *http.Request) map[string]string {
	return mux.Vars(req)
}
