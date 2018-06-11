package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/sbstjn/markdownfmt/markdown"
	"github.com/spf13/cobra"
)

func parseArguments(args []string) error {
	if len(args) < 1 || len(args) > 2 {
		return fmt.Errorf("Invalid number of arguments")
	}

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

var command = &cobra.Command{
	Short: "Preprocess Markdown templates as Go templates",
	Run: func(cmd *cobra.Command, args []string) {
		var result []byte
		var err error

		// Check for valid number of arguments
		if err := parseArguments(args); err != nil {
			abort(err.Error())
		}

		// Check if input file can be read and parsed
		if err := parseIn(); err != nil {
			abort(err.Error())
		}

		// Check if data file can be read and parsed (optional)
		if err := parseData(); err != nil {
			abort(err.Error())
		}

		// Check if wrapper file can read and parsed (optional)
		if err := parseWrapper(); err != nil {
			abort(err.Error())
		}

		// Merge template with data
		if result, err = merge(templateIn, data); err != nil {
			abort(err.Error())
		}

		// Format Markdown
		if result, err = markdown.Process("", result, nil); err != nil {
			abort(err.Error())
		}

		// If --wrap is used, wrap result in wrapper
		if flagWrap != "" {
			wrapperData := map[string]interface{}{
				"Content": string(result),
				"Data":    data,
			}

			if result, err = merge(templateWrapper, wrapperData); err != nil {
				abort(err.Error())
			}
		}

		// If no output file is defined, just write result to stdout
		if fileOut == "" {
			fmt.Println(string(result))
		} else {
			if _, err := ioutil.ReadFile(fileOut); err == nil {
				abort("File already exists: %s", fileOut)
			}

			if err := os.MkdirAll(filepath.Dir(fileOut), os.ModePerm); err != nil {
				abort("Failed to create needed folders for file: %s", fileOut)
			}

			if err := ioutil.WriteFile(fileOut, result, 0644); err != nil {
				abort("Failed to write file to destination: %s", fileOut)
			}
		}
	},
}
