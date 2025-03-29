package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	VERSION = "1.0.0"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("caflax v%s\n", VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
