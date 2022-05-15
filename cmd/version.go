package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the AtomicFS version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(
			"AtomicFS -- Version: %s -- Git Commit: %s \n",
			viper.GetString("git_version"),
			viper.GetString("git_commit"),
		)
	},
}
