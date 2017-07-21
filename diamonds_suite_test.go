package diamonds_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDiamonds(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Diamonds Suite")
}
