package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPredown(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Predown Suite")
}
