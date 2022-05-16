package cmd

import (
	"fmt"

	"github.com/0x19/atomicfs/internal/core/cli/prompt"
	"github.com/0x19/atomicfs/pkg/wallet"
	"github.com/spf13/cobra"
)

func init() {
	walletCmd.AddCommand(walletNewAccountCmd)
	rootCmd.AddCommand(walletCmd)
}

var walletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Manage blockchain accounts and keys",
}

var walletNewAccountCmd = &cobra.Command{
	Use:          "new",
	Short:        "Creates a new account with a new set of a elliptic-curve private and public keys.",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		secret, err := prompt.PromptForSecret("Please enter a password used for wallet encryption")
		if err != nil {
			return err
		}

		acc, err := wallet.NewKeystoreAccount(dataPath, secret)
		if err != nil {
			return err
		}

		fmt.Printf("New account created: %s\n", acc.Hex())
		fmt.Printf("Saved in: %s\n", wallet.GetKeystoreDirPath(dataPath))
		return nil
	},
}
