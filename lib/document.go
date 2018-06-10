package lib

import (
	"bytes"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

const (
	blackfridayOptions = blackfriday.CommonExtensions | blackfriday.AutoHeadingIDs
)

// Document stores the built markdown document in memory
type Document struct {
	Data *bytes.Buffer
}

// Markdown return the document as Markdown
func (d *Document) Markdown() []byte {
	return d.format(ToMarkdown)
}

// HTML returns the document as HTML
func (d *Document) HTML() []byte {
	return d.format(ToHTML)
}

func (d *Document) format(f Formatter) []byte {
	return f(d)
}
