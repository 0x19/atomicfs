package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/0x19/atomicfs/internal/core/cli/prompt"
	"github.com/0x19/atomicfs/pkg/wallet"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/spf13/cobra"
)

func init() {
	walletCmd.AddCommand(walletNewAccountCmd)

	addKeystoreFlag(walletPrintPrivateKeyCmd)
	walletCmd.AddCommand(walletPrintPrivateKeyCmd)

	walletCmd.AddCommand(walletBalanceCmd)

	rootCmd.AddCommand(walletCmd)
}

var walletCmd = &cobra.Command{
	Use:          "wallet",
	Short:        "Manage blockchain accounts and keys",
	SilenceUsage: true,
}

var walletNewAccountCmd = &cobra.Command{
	Use:          "new",
	Short:        "Creates a new account with a new set of a elliptic-curve private and public keys.",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		secret, err := prompt.PromptForSecret("Please enter wallet password")
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

var walletPrintPrivateKeyCmd = &cobra.Command{
	Use:          "dump",
	Short:        "Unlocks keystore file and dumps the keystore information (id, address, private and public keys)",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ksFile, _ := cmd.Flags().GetString(flagKeystoreFile)
		secret, err := prompt.PromptForSecret("Please enter a password to decrypt the wallet")
		if err != nil {
			return err
		}

		keyJson, err := ioutil.ReadFile(ksFile)
		if err != nil {
			return err
		}

		key, err := keystore.DecryptKey(keyJson, secret)
		if err != nil {
			return err
		}

		spew.Dump(key)
		return nil
	},
}

var walletBalanceCmd = &cobra.Command{
	Use:          "balance",
	Short:        "Get account balance",
	SilenceUsage: true,
}
