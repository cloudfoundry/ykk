package ykk_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestYkk(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "YKK Suite")
}
