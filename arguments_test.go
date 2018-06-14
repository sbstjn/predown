package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Arguments", func() {
	Describe("Parse Input and Output files", func() {
		Context("With one argument", func() {
			args := []string{"example.md"}
			in, out := parseInputOutputFilesFromArguments(args)

			It("should find input", func() {
				Expect(in).To(Equal("example.md"))
			})

			It("should have empty output", func() {
				Expect(out).To(Equal(""))
			})
		})

		Context("With two arguments", func() {
			args := []string{"example.md", "output.md"}
			in, out := parseInputOutputFilesFromArguments(args)

			It("should find input", func() {
				Expect(in).To(Equal("example.md"))
			})

			It("should output", func() {
				Expect(out).To(Equal("output.md"))
			})
		})
	})
})
