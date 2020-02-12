package internal

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"os"
	"strings"
)

const envPrefix = rootCmdName

func fillWithEnvVars(flagSet *pflag.FlagSet) {
	flagSet.VisitAll(func(flag *pflag.Flag) {
		envKey := buildEnvKey(envPrefix, flag.Name)
		value := os.Getenv(envKey)
		if len(value) > 0 {
			err := flagSet.Set(flag.Name, value)
			if err != nil {
				logrus.WithError(err).Fatalf("failed to set %v for flag %+v", value, flagSet)
			}
		}
	})
}

func buildEnvKey(prefix string, name string) string {
	key := prefix + "_" + name
	key = strings.ReplaceAll(key, "-", "_")
	return strings.ToUpper(key)
}
