package main

import "github.com/sbstjn/predown/cmd"

var (
	version string
	name    = "predown"
)

func main() {
	cmd.Name = name

	if version != "" {
		cmd.Version = version
	} else {
		cmd.Version = "0.0.0-dev"
	}

	cmd.Execute()
}
