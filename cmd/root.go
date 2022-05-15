/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "atomicfs",
	Short: "Lorem ipsum dolor sit amet...",
	Long:  ``,
}

func init() {
	cobra.OnInitialize(func() {
		SetupDefaults([]zapcore.Field{
			zapcore.Field{
				Key:    "root",
				Type:   zapcore.StringType,
				String: "atomicfs",
			},
		})
	})

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is $HOME/atomicfs/config.yaml)")
	rootCmd.PersistentFlags().String("log-level", "", "The minimum level at which to emit log data.")
	rootCmd.PersistentFlags().Bool("pprof-enabled", true, "Boolean value indicating whenever we want to enable pprof server")
	rootCmd.PersistentFlags().String("pprof-addr", ":9000", "Listener address for the pprof server. Valid address example: {ip}:{port}")
	rootCmd.PersistentFlags().Bool("prometheus-enabled", true, "Boolean value indicating whenever we want to enable prometheus (metrics) server")
	rootCmd.PersistentFlags().String("prometheus-addr", ":9001", "Listener address for the prometheus server. Valid address example: {ip}:{port}")
	rootCmd.PersistentFlags().String("environment", "development", "Which environment should we start service in? Available: production, development.")

	_ = viper.BindPFlag("log_level", rootCmd.PersistentFlags().Lookup("log-level"))
	_ = viper.BindPFlag("pprof_server_enabled", rootCmd.PersistentFlags().Lookup("pprof-enabled"))
	_ = viper.BindPFlag("pprof_server_addr", rootCmd.PersistentFlags().Lookup("pprof-addr"))
	_ = viper.BindPFlag("prometheus_server_enabled", rootCmd.PersistentFlags().Lookup("prometheus-enabled"))
	_ = viper.BindPFlag("prometheus_server_addr", rootCmd.PersistentFlags().Lookup("prometheus-addr"))
	_ = viper.BindPFlag("environment", rootCmd.PersistentFlags().Lookup("environment"))
}
