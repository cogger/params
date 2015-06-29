package params

import (
	"bytes"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
)

var _ = Describe("ParamContext", func() {
	It("should add params to the context", func() {
		req, err := http.NewRequest("GET", "/", &bytes.Buffer{})
		Expect(err).ToNot(HaveOccurred())

		ctx := Add(func(req *http.Request) map[string]string {
			return map[string]string{"test": "test"}
		})(context.Background(), req)

		Expect(Get(ctx, "test")).To(Equal("test"))
	})

	It("should panic when no params context is added", func() {
		ctx := context.Background()
		Expect(func() {
			Get(ctx, "test")
		}).To(Panic())
	})
})
