# params 

[![GoDoc](https://godoc.org/gopkg.in/cogger/params.v1?status.png)](http://godoc.org/gopkg.in/cogger/params.v1)  
[![Build Status](https://travis-ci.org/cogger/params.svg?branch=master)](https://travis-ci.org/cogger/params)  
[![Coverage Status](https://coveralls.io/repos/cogger/params/badge.svg?branch=master)](https://coveralls.io/r/cogger/params?branch=master)  
[![License](http://img.shields.io/:license-apache-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)


params addes generic way to get params associated with a request

## Installation

The import path for the package is *gopkg.in/cogger/params.v1*.

To install it, run:

    go get gopkg.in/cogger/params.v1

## Usage
~~~ go
// main.go
package main

import (
	"net/http"
	"log"
	"gopkg.in/cogger/cogger.v1"
	"gopkg.in/cogger/params.v1"
	"gopkg.in/cogger/params.v1/gorilla"
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
	fooHandler.AddContext(params.Add(gorilla.Mapper))
	fooHandler.SetTimeout(3 * time.Second)
	fooHandler.SetHandler(foo)

 	r := mux.NewRouter()
	r.Methods("GET").Path("/{key}/").Handler(base.GetCompositeHandler)
  	http.ListenAndServe(":8080", r)
}

~~~
