package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(
			"%s %s %s %s\n",
			Name,
			"v"+Version,
			runtime.GOOS+"/"+runtime.GOARCH,
			runtime.Version(),
		)
	},
}
