package target_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"clifig/pkg/target"
)

var _ = Describe("Env", func() {
	Context("when given a map", func() {
		It("displays key=value strings", func() {
			Expect(target.Env(map[string]string{"key": "value"})).To(Equal("key=value\n"))
		})
		It("displays multiple key=value strings", func() {
			Expect(target.Env(map[string]string{"key": "value", "key1": "value1"})).To(ContainSubstring("key=value\n"), ContainSubstring("key1=value1\n"))
		})
	})
})
