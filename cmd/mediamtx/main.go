// main executable.
package main

import (
	"os"

	"github.com/bluenviron/mediamtx/internal/core"
)

func main() {
	s, ok := core.New(nil)
	if !ok {
		os.Exit(1)
	}

	s.Wait()
}
