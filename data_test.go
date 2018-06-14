package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Data", func() {
	Describe("hasDataFlag()", func() {
		Context("no flag set", func() {
			It("should be false", func() {
				Expect(hasDataFlag()).To(BeFalse())
			})
		})

		Context("flag set", func() {
			It("should be true", func() {
				flagData = "data.toml"

				Expect(hasDataFlag()).To(BeTrue())
			})
		})
	})

	Describe("getDataContent()", func() {
		Context("file available", func() {
			file := "data_test.go"
			data, err := getDataContent(file)

			It("should not error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should load file contents", func() {
				Expect(data).To(ContainSubstring("getDataContent()"))
			})
		})

		Context("file not found", func() {
			file := "invalid.go"
			data, err := getDataContent(file)

			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})

			It("should load file contents", func() {
				Expect(data).To(BeEmpty())
			})
		})
	})

	Describe("parseDataContent()", func() {
		Context("invalid TOML data", func() {
			content := `Foo += !"Bar"`
			data, err := parseDataContent(content)

			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})

			It("should not parse TOML data", func() {
				Expect(data).To(BeNil())
			})
		})

		Context("valid TOML data", func() {
			content := `Foo = "Bar"`
			data, err := parseDataContent(content)

			It("should not error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should parse TOML data", func() {
				Expect(data).To(HaveKey("Foo"))
			})
		})
	})
})
