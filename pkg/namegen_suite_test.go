package pkg_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNamegen(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Namegen Suite")
}
