package cmd

import "github.com/0x19/atomicfs/internal/core/ctx"

var (
	globalCtx *ctx.Ctx
)

// ExecuteRootContext - Executes the root command with the global context
func ExecuteRootContext(c *ctx.Ctx) error {
	globalCtx = c
	return rootCmd.ExecuteContext(c)
}
