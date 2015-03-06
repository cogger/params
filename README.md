# gae [![GoDoc](https://godoc.org/github.com/cogger/gae?status.png)](http://godoc.org/github.com/cogger/gae)

params addes generic way to get params associated with a request

## Usage
~~~ go
// main.go
package main

import (
	"net/http"
	"log"
	"github.com/cogger/cogger"
	"github.com/cogger/params"
	"github.com/cogger/params/gorilla"
	"golang.org/x/net/context"
	"github.com/gorilla/mux"
)

func foo(ctx context.Context, w http.ResponseWriter, r *http.Request) int{
	key := params.Get(ctx, "key")
	log.Println(key)
	return http.StatusOK
}

func main() {
	fooHandler := cogger.NewHandler()
	fooHandler.AddContext(params.Add(gorrilla.Mapper))
	fooHandler.SetTimeout(3 * time.Second)
	fooHandler.SetHandler(foo)

 	r := mux.NewRouter()
	r.Methods("GET").Path("/{key}/").Handler(base.GetCompositeHandler)
  	http.ListenAndServe(":8080", r)
}

~~~
