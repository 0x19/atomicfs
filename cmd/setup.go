package cmd

import (
	"github.com/0x19/atomicfs/internal/setup"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setupCmd)
}

var setupCmd = &cobra.Command{
	Use:          "setup",
	Short:        "AtomicFS setup. Should be ran only one time",
	SilenceUsage: true, // Should not display usage if error occurs
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := setup.Run(globalCtx); err != nil {
			return err
		}
		return nil
	},
}
