package target_test

import (
	"clifig/pkg/target"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("parameters", func() {
	Context("When given a map and a prefix", func() {
		testMap := map[string]string{"key": "value"}
		prefix := "/test/value/"
		It("will call parameters and it will convert it to /prefix/name, value, type", func() {
			name := prefix + "key"
			value := "value"
			paramType := "String"
			Expect(target.ParseParameter(testMap, prefix)).To(ContainElements([]target.ParameterStoreInput{{Name: name, Value: value, ParamType: paramType}}))
		})
	})
	Context("When given a map with sensative value, and a prefix", func() {
		testMap := map[string]string{"*key": "value"}
		prefix := "/test/value/"
		It("will call parameters and it will convert it to /prefix/name, value, type", func() {
			name := prefix + "key"
			value := "value"
			paramType := "SecureString"
			Expect(target.ParseParameter(testMap, prefix)).To(ContainElements([]target.ParameterStoreInput{{Name: name, Value: value, ParamType: paramType}}))
		})
	})
	Context("When given multiple values in the map", func() {
		testMap := map[string]string{"*secret": "password", "jedi": "blackandwhite", "buddy": "whiteandgrey"}
		prefix := "/test/value/"
		It("the values shouldn't be all the same.", func() {
			Expect(target.ParseParameter(testMap, prefix)).To(ContainElements(
				[]target.ParameterStoreInput{
					{Name: prefix + "secret", Value: "password", ParamType: "SecureString"},
					{Name: prefix + "jedi", Value: "blackandwhite", ParamType: "String"},
					{Name: prefix + "buddy", Value: "whiteandgrey", ParamType: "String"},
				}))
		})
	})
})
