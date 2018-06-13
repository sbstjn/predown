package main

import (
	"fmt"
	"html/template"
	"os"
	"runtime"
)

var (
	version = "0.0.0-dev"

	flagData string
	flagWrap string

	data interface{}

	templateWrapper *template.Template
	templateIn      *template.Template

	fileIn  string
	fileOut string
)

func abort(format string, a ...interface{}) {
	write("[Error] "+format+"\n", a...)
	os.Exit(1)
}

func write(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", a...)
}

func init() {
	command.PersistentFlags().StringVar(&flagData, "data", "", "Path to TOML data file")
	command.PersistentFlags().StringVar(&flagWrap, "wrap", "", "Path to wrapper file")
}

func main() {
	command.Version = version

	command.SetVersionTemplate(fmt.Sprintf(
		"predown v{{ .Version }} %s %s\n",
		runtime.GOOS+"/"+runtime.GOARCH,
		runtime.Version(),
	))

	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
