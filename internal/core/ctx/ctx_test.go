package ctx_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/0x19/atomicfs/internal/core/ctx"
	"github.com/stretchr/testify/assert"
)

func TestCtx(t *testing.T) {
	pa := assert.New(t)
	tctx := ctx.New(context.TODO())
	pa.IsType(tctx, &ctx.Ctx{})
}

func TestCtxWithInterrupt(t *testing.T) {
	pa := assert.New(t)
	tctx := ctx.New(context.TODO())
	pa.IsType(tctx, &ctx.Ctx{})

	go func() {
		time.Sleep(500 * time.Millisecond)
		tctx.Interrupt()
	}()

	select {
	case <-time.After(time.Second):
		log.Fatal("failure to exit the context in one second")
	case <-tctx.Done():
		return
	}
}

func TestCtxWithStop(t *testing.T) {
	pa := assert.New(t)
	tctx := ctx.New(context.TODO())
	pa.IsType(tctx, &ctx.Ctx{})

	go func() {
		time.Sleep(500 * time.Millisecond)
		tctx.Stop()
	}()

	select {
	case <-time.After(time.Second):
		log.Fatal("failure to exit the context in one second")
	case <-tctx.Done():
		return
	}
}

func TestNewCorruptedCtx(t *testing.T) {
	pa := assert.New(t)
	tctx := ctx.New(context.TODO())
	pa.IsType(tctx, &ctx.Ctx{})

	go func() {
		time.Sleep(500 * time.Millisecond)
		tctx.Stop()
	}()

	select {
	case <-time.After(300 * time.Millisecond):
		return
	case <-tctx.Done():
		log.Fatal("context not closed within the second")
	}
}
