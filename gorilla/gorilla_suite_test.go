package gorilla

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGorilla(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gorilla Suite")
}
