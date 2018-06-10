package lib

import (
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

// Config holds all data from the configuration file
type Config interface{}

// Load tries to load a configuration file
func Load(file string) (interface{}, error) {
	var data interface{}

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("Unable to find configuration: %s", file)
	}

	if _, err := toml.Decode(string(content), &data); err != nil {
		return nil, fmt.Errorf("Cannot parse TOML document: %s", file)
	}

	return data, nil
}
