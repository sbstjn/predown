package lib

import (
	"bytes"
	"fmt"
	"path"
	"text/template"
)

// Builder stores the configuration and parses the
type Builder struct {
	Config Config
}

// Parse template to Markdown file
func (b *Builder) Parse(file string) (*Document, error) {
	data := Document{&bytes.Buffer{}}
	name := path.Base(file)

	tmpl, err := template.New(name).ParseFiles(file)

	if err != nil {
		return nil, fmt.Errorf("Unable to parse template: %s\n\n%s", file, err.Error())
	}

	if tmpl.Execute(data.Data, b.Config) != nil {
		return nil, fmt.Errorf("Unable to render template: %s\n\n%s", file, err.Error())
	}

	return &data, nil
}

// NewBuilder returns a Builder with config
func NewBuilder(config Config) Builder {
	return Builder{config}
}
