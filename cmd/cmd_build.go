package cmd

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/sbstjn/predown/lib"
	"github.com/spf13/cobra"
)

var (
	cfgFormat string
)

func init() {
	rootCmd.AddCommand(buildCmd)
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build files",
	Run: func(cmd *cobra.Command, args []string) {
		requireValidConfiguration()
		builder := lib.NewBuilder(Config)

		var files []string
		var err error

		if files, err = lib.Files(cfgTemplates); err != nil {
			abort("Unable to get templates: %s", err)
		}

		for _, file := range files {
			var doc *lib.Document
			var err error

			srcFile := path.Join(cfgTemplates, file)
			dstFile := path.Join(cfgDestination, file)

			if doc, err = builder.Parse(srcFile); err != nil {
				abort(err.Error())
			}

			if err := os.MkdirAll(filepath.Dir(dstFile), os.ModePerm); err != nil {
				abort("Failed to create needed folders for file: %s", dstFile)
			}

			if err := ioutil.WriteFile(dstFile, doc.Markdown(), 0644); err != nil {
				abort("Failed to write file to destination: %s", dstFile)
			}

			write("Done: %s -> %s", srcFile, dstFile)
		}
	},
}
