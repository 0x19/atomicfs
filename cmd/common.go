package cmd

import (
	"os"
	"path"
	"strings"

	"github.com/0x19/atomicfs/internal/core/ctx"
	"github.com/0x19/atomicfs/pkg/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	cfgFile   string
	dataPath  string
	globalCtx *ctx.Ctx
)

const flagKeystoreFile = "keystore"

// ExecuteRootContext - Executes the root command with the global context
func ExecuteRootContext(c *ctx.Ctx) error {
	globalCtx = c
	return rootCmd.ExecuteContext(c)
}

func SetupDefaults(labels []zapcore.Field) error {
	options := []zap.Option{
		zap.WithCaller(true),
		zap.AddStacktrace(zapcore.ErrorLevel),
	}

	_, err := logger.New(viper.GetString("log_level"), viper.GetString("environment"), true, options, labels)
	if err != nil {
		zap.L().Error("failure to init new logger", zap.Error(err))
		os.Exit(1)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("yaml")

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile(path.Join(dataPath, "config.yaml"))
	}

	if err := viper.ReadInConfig(); err != nil {
		zap.L().Error("failure to read configuration file", zap.Error(err))
		os.Exit(1)
	}

	// Just notification that we know that file is generally changed, nothing else...
	viper.OnConfigChange(func(e fsnotify.Event) {
		zap.L().Info("Configuration file changed! Reloading...")
	})

	viper.WatchConfig()

	zap.L().Info("Successfully loaded configuration file.", zap.String("path", viper.ConfigFileUsed()))

	if viper.GetString("data_path") == "" {
		viper.SetDefault("data_path", func() string {
			home, _ := os.UserHomeDir()
			return path.Join(home, "atomicfs")
		}())
	}

	return nil
}

func addKeystoreFlag(cmd *cobra.Command) {
	cmd.Flags().String(flagKeystoreFile, "", "Absolute path to the encrypted keystore file")
	cmd.MarkFlagRequired(flagKeystoreFile)
}
