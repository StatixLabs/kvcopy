package auth_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"clifig/pkg/auth"
)

var _ = Describe("Basic", func() {
	Context("when given a file", func() {
		It("converts contents of file map[string]string", func() {
			username, password := auth.Basic("username", "password")
			Expect(username).To(Equal("username"))
			Expect(password).To(Equal("password"))
		})
	})
})
