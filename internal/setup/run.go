package setup

import (
	"github.com/0x19/atomicfs/internal/core/ctx"
	"go.uber.org/zap"
)

func Run(c *ctx.Ctx) error {
	zap.L().Info("Starting up AtomicFS setup...")
	<-c.Done()
	return nil
}
