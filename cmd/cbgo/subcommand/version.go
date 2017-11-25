package subcommand

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.0.1"

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of the cbgo command.",
	Long:  "Prints the version of the cbgo command.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}