package ctx

import (
	"context"
	"os"
	"os/signal"
)

type Ctx struct {
	context.Context
	context.CancelFunc
}

// Interrupt - Does the whole program process interuption, usefull for testing.
// Not so much useful otherwise for now.
func (c *Ctx) Interrupt() error {
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		return err
	}

	if err := p.Signal(os.Interrupt); err != nil {
		return err
	}

	return nil
}

// Stop - Closes the context altogether
func (c *Ctx) Stop() {
	c.CancelFunc()
}

// New - Creates new context from a parent context
func New(c context.Context) *Ctx {
	ctx, stop := signal.NotifyContext(c, os.Interrupt)
	return &Ctx{ctx, stop}
}
