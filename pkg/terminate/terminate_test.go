package terminate_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/0x19/atomicfs/pkg/terminate"
)

func TestTerminate(t *testing.T) {
	go func() {
		// Just sleep for a bit and do nothing... Mimics running application.
		time.Sleep(200 * time.Millisecond)

		p, err := os.FindProcess(os.Getpid())
		if err != nil {
			log.Fatal(err)
		}

		// On a Unix-like system, pressing Ctrl+C on a keyboard sends a
		// SIGINT signal to the process of the program in execution.
		//
		// This example simulates that by sending a SIGINT signal to itself.
		if err := p.Signal(os.Interrupt); err != nil {
			log.Fatal(err)
		}
	}()

	// If terminate hangs, well that's the failed test for now...
	terminate.Wait()
}
