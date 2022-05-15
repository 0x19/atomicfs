package terminate

import (
	"os"
	"os/signal"
	"syscall"
)

// New - Creates a new os signal and applies an notifier to shut down the channel once
// external party decides to terminate the signal.
func New() chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	return sig
}

// Wait - Instanciates a new termination signal and waits until channel is released
func Wait() {
	<-New()
}
