package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"https://github.com/fsmile2/ckm8/version"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version of current ckm8cli binary.",
	Run:   runVersion,
}

func init() {
	RootCmd.AddCommand(versionCmd)
}

func runVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("Version %v %s\nBuilt at %s\n", version.Version, version.GitHash, version.Timestamp)
}
