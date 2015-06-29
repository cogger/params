package test

import (
	"bytes"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParamMapper", func() {
	It("should return the map provided", func() {

		req, err := http.NewRequest("GET", "/", &bytes.Buffer{})
		Expect(err).ToNot(HaveOccurred())

		m := map[string]string{"test": "test"}
		Expect(Mapper(m)(req)).To(Equal(m))
	})
})
