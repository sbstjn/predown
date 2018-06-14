package main

import (
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

func hasDataFlag() bool {
	if flagData == "" {
		return false
	}

	return true
}

func getDataContent(file string) (string, error) {
	content, err := ioutil.ReadFile(file)

	if err != nil {
		return "", fmt.Errorf("Failed to access data file: %s", file)
	}

	return string(content), nil
}

func parseDataContent(content string) (interface{}, error) {
	var data interface{}

	if _, err := toml.Decode(content, &data); err != nil {
		return nil, fmt.Errorf("Failed to parse data from content")
	}

	return data, nil
}

func parseData() error {
	if !hasDataFlag() {
		return nil
	}

	var content string
	var err error

	if content, err = getDataContent(flagData); err != nil {
		return err
	}

	if data, err = parseDataContent(content); err != nil {
		return err
	}

	return nil
}
