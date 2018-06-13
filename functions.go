package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"path"

	"github.com/BurntSushi/toml"
)

func parseArguments(args []string) error {
	fileIn = args[0]

	if len(args) == 2 {
		fileOut = args[1]
	}

	return nil
}

func parseData() error {
	if flagData == "" {
		return nil
	}

	content, err := ioutil.ReadFile(flagData)

	if err != nil {
		return fmt.Errorf("Failed to access data file: %s", flagData)
	}

	if _, err := toml.Decode(string(content), &data); err != nil {
		return fmt.Errorf("Failed to parse data file: %s", flagData)
	}

	return nil
}

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
