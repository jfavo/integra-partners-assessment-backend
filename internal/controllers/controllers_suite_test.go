package controllers_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	format.TruncatedDiff = false
	RunSpecs(t, "Controllers Suite")
}
