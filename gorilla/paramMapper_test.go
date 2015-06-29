package gorilla

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type test struct {
	Title string
	Route string
	Path  string
	Map   map[string]string
}

var _ = Describe("ParamMapper", func() {
	tests := []test{
		test{
			Title: "basic route",
			Route: "/{key}/",
			Path:  "/key1/",
			Map: map[string]string{
				"key": "key1",
			},
		},
		test{
			Title: "slighty less basic route",
			Route: "/{key}/{other}/",
			Path:  "/key1/other1/",
			Map: map[string]string{
				"key":   "key1",
				"other": "other1",
			},
		},
	}

	It("should return the map from the request", func() {
		for _, t := range tests {
			func(t test) {
				r := mux.NewRouter()
				r.HandleFunc(t.Route, func(writer http.ResponseWriter, req *http.Request) {
					m := Mapper(req)
					Expect(m).To(Equal(t.Map), t.Title)
				})
				req, err := http.NewRequest("GET", t.Path, &bytes.Buffer{})
				Expect(err).ToNot(HaveOccurred())

				recorder := httptest.NewRecorder()
				r.ServeHTTP(recorder, req)
			}(t)
		}
	})
})
