package internal

import (
	"context"
	"crypter/pkg/crypter/encrypt"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt a value with key",
	RunE: func(cmd *cobra.Command, _ []string) error {
		var cfg encrypt.Config
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

		return encrypt.Encrypt(context.Background(), cmd.OutOrStdout(), cfg)
	},
}

func init() {
	encryptCmd.PersistentFlags().StringP("key-file", "k", "", "path to the public key")
	encryptCmd.PersistentFlags().StringP("key-string", "s", "", "value of the public key")
	encryptCmd.PersistentFlags().StringP("value", "v", "", "value")
	if err := encryptCmd.MarkPersistentFlagRequired("value"); err != nil {
		logrus.WithError(err).Fatal("failed mark flag as required")
	}
	cobra.OnInitialize(func() {
		fillWithEnvVars(encryptCmd.Flags())
	})
	rootCmd.AddCommand(encryptCmd)
}
