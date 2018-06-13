package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sbstjn/markdownfmt/markdown"
	"github.com/spf13/cobra"
)

var command = &cobra.Command{
	Use:  "predown <in> [<out>]",
	Args: cobra.RangeArgs(1, 2),
	Example: `  predown template.md
  predown template.md --data data.toml
  predown template.md --data data.toml --wrap wrapper.frontmatter

  predown template.md output.md
  predown template.md output.md --data data.toml
  predown template.md output.md --data data.toml --wrap wrapper.frontmatter`,
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
