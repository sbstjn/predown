package lib

import (
	"github.com/sbstjn/markdownfmt/markdown"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

// Formatter is an abstract handler to format Documents
type Formatter func(d *Document) []byte

// ToMarkdown formats a Document into a Markdown document
func ToMarkdown(d *Document) []byte {
	output, _ := markdown.Process("", d.Data.Bytes(), nil)

	return output
}

// ToHTML formats a Document into a HTML document
func ToHTML(d *Document) []byte {
	return blackfriday.Run(
		d.Markdown(),
		blackfriday.WithExtensions(blackfridayOptions),
	)
}
