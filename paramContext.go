package params

import (
	"errors"
	"net/http"

	"golang.org/x/net/context"
)

type paramKey struct{}

//ErrNoParams denotes that the param context was not added to handler
var ErrNoParams = errors.New("params context not added")

//Add adds the params context creator to a handler
func Add(mapper func(*http.Request) map[string]string) func(context.Context, *http.Request) context.Context {
	return func(ctx context.Context, req *http.Request) context.Context {
		return context.WithValue(ctx, paramKey{}, mapper(req))
	}
}

//Get a param from a context
func Get(ctx context.Context, key string) string {
	params, ok := ctx.Value(paramKey{}).(map[string]string)
	if !ok {
		panic(ErrNoParams)
	}

	return params[key]
}
