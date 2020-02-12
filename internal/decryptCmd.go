package internal

import (
	"context"
	"crypter/pkg/crypter/decrypt"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt a file with key",
	RunE: func(cmd *cobra.Command, _ []string) error {
		var cfg decrypt.Config
		var err error
		if cfg.Value, err = cmd.Flags().GetString("value"); err != nil {
			return fmt.Errorf("failed to get %v flag value. %w", "value", err)
		}

		keyString, err := cmd.Flags().GetString("key-string")
		if err != nil {
			return fmt.Errorf("failed to get %v flag value. %w", "key-string", err)
		}

		if keyString == "" {
			keyFile, err := cmd.Flags().GetString("key-file")
			if err != nil {
				return fmt.Errorf("failed to get %v flag value. %w", "key-file", err)
			}
			if keyFile == "" {
				return fmt.Errorf("either %s or %s option should be provided", "key-file", "key-string")
			}
			content, err := ioutil.ReadFile(keyFile)
			if err != nil {
				return fmt.Errorf("failed to read data from %v. %w", keyFile, err)
			}
			keyString = string(content)
		}

		cfg.Key = keyString

		return decrypt.Decrypt(context.Background(), cmd.OutOrStdout(), cfg)
	},
}

func init() {
	decryptCmd.PersistentFlags().StringP("key-file", "k", "", "path to the public key")
	decryptCmd.PersistentFlags().StringP("key-string", "s", "", "value of the public key")
	decryptCmd.PersistentFlags().StringP("value", "v", "", "value")
	if err := decryptCmd.MarkPersistentFlagRequired("value"); err != nil {
		logrus.WithError(err).Fatal("failed mark flag as required")
	}
	cobra.OnInitialize(func() {
		fillWithEnvVars(decryptCmd.Flags())
	})
	rootCmd.AddCommand(decryptCmd)
}
