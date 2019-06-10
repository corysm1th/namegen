package pkg_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/corysm1th/namegen/pkg"
)

var _ = Describe("Namegen", func() {
	It("Returns a random name of things", func() {
		c := Case(Snake)
		things := Things(c)
		Expect(things).To(MatchRegexp("\\D?"))
	})
})
