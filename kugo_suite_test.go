package kugo_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestKugo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kugo Suite")
}
