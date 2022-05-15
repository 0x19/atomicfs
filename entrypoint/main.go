/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"context"

	"github.com/0x19/atomicfs/cmd"
	"github.com/0x19/atomicfs/internal/core/ctx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	// GitCommit for more information how to set and use GitCommit
	// https://blog.alexellis.io/inject-build-time-vars-golang/
	GitCommit  string
	GitVersion string
)

func main() {
	viper.SetDefault("git_commit", GitCommit)
	viper.SetDefault("git_version", GitVersion)

	globalCtx := ctx.New(context.Background())

	if err := cmd.ExecuteRootContext(globalCtx); err != nil {
		zap.L().Fatal("command execute error", zap.Error(err))
		return
	}
}
