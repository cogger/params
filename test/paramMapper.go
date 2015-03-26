package test

import "net/http"

//Mapper maps predefined params
func Mapper(values map[string]string) func(*http.Request) map[string]string {
	return func(req *http.Request) map[string]string {
		return values
	}
}
