package source_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"clifig/pkg/source"
)

var _ = Describe("File", func() {
	Context("when given a file", func() {
		It("converts contents of file map[string]string", func() {
			Expect(source.File("source.txt")).To(Equal(map[string]string{"key": "value", "key1": "value1"}))
		})
	})
})
