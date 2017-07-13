package diamonds_test

import (
	. "github.com/harryganz/diamonds"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewParametersFromYaml", func() {
	table := []struct {
		input    []byte
		expected Parameters
		context  string
		outcome  string
	}{
		{[]byte(`shape: Round`), NewParameters(RoundShape), "when all fields are passed in", "produces Parameters with all passed in fields"},
		{[]byte(``), NewParameters(NoneShape), "when no fields are passed in", "produces Parameters with default fields"},
	}

	for _, entry := range table {
		Context(entry.context, func() {
			It(entry.outcome, func() {
				Expect(NewParametersFromYaml(entry.input)).To(Equal(entry.expected))
			})
		})
	}
})
