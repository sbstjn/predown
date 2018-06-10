package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/sbstjn/predown/lib"
	"github.com/spf13/cobra"
)

var (
	// Name stores the name of the binary
	Name string
	// Version stores the version of the binary
	Version string
	// Config hold all data from the .toml file
	Config lib.Config

	cfgFile        string
	cfgDestination string
	cfgTemplates   string
)

var rootCmd = &cobra.Command{
	Use:   Name,
	Short: "Preprocess Markdown templates as Go templates",
}

func abort(format string, a ...interface{}) {
	write("[Error] "+format+"\n", a...)
	os.Exit(1)
}

func write(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", a...)
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default: privacy.toml)")
	rootCmd.PersistentFlags().StringVar(&cfgDestination, "destination", "", "Path to store generated files")
	rootCmd.PersistentFlags().StringVar(&cfgTemplates, "templates", "", "Prefix for template lookup")
}

func requireValidConfiguration() {
	// Config file is always required
	if Config == nil {
		abort("Unable to load config file %s", cfgFile)
	}

	// Output flag is always required
	if cfgDestination == "" {
		abort("Please configure destination using --destination")
	} else {
		if _, err := os.Stat(cfgDestination); os.IsNotExist(err) {
			abort("Configured destination does not exist: %s", cfgDestination)
		}

		if handle, err := os.Stat(cfgDestination); err == nil && !handle.IsDir() {
			abort("Configured destination is not a directory: %s", cfgDestination)
		}
	}

	// Templates flag is always required
	if cfgTemplates == "" {
		abort("Please configure prefix for template lookup using --templates")
	} else {
		if _, err := os.Stat(cfgTemplates); os.IsNotExist(err) {
			abort("Configured template prefix does not exist: %s", cfgTemplates)
		}

		if handle, err := os.Stat(cfgTemplates); err == nil && !handle.IsDir() {
			abort("Configured template prefix is not a directory: %s", cfgTemplates)
		}
	}

	write("Using config file %s", cfgFile)
	write("Using templates prefix %s", cfgTemplates)
	write("Using destination path %s", cfgDestination)
	write("")
}

func initConfig() {
	if cfgFile == "" {
		cfgFile = path.Join("privacy.toml")
	}

	Config, _ = lib.Load(cfgFile)
}

// Execute handles the CLI command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
