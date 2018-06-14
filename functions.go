package main

import (
	"bytes"
	"html/template"
	"path"
)

func parseTemplateFile(file string) (*template.Template, error) {
	name := path.Base(file)
	return template.New(name).ParseFiles(file)
}

func parseWrapper() error {
	if flagWrap == "" {
		return nil
	}

	tpl, err := parseTemplateFile(flagWrap)
	templateWrapper = tpl

	return err
}

func parseIn() error {
	tpl, err := parseTemplateFile(fileIn)
	templateIn = tpl

	return err
}

func merge(template *template.Template, data interface{}) ([]byte, error) {
	var result = bytes.Buffer{}
	err := template.Execute(&result, data)

	return result.Bytes(), err
}
