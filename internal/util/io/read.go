package io

import (
	"fmt"
	"io"
	"os"
)

var (
	ErrNoInput = fmt.Errorf("provide a file or pipe CSV into stdin (or use '-')")
)

func ReadInput(args []string) ([]byte, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to stat stdin: %w", err)
	}

	isStdinPipe := (fi.Mode()&os.ModeNamedPipe != 0) ||
		!fi.Mode().IsRegular() && !fi.Mode().IsDir()

	// Case 1 — data is piped into stdin OR stdin is redirected
	if isStdinPipe && fi.Mode()&os.ModeCharDevice == 0 {
		return io.ReadAll(os.Stdin)
	}

	// Case 2 — a filename was provided
	if len(args) > 0 {
		if args[0] == "-" {
			// Read from stdin explicitly via "-"
			return io.ReadAll(os.Stdin)
		}
		return os.ReadFile(args[0])
	}

	// Case 3 — neither pipe nor file → error
	return nil, ErrNoInput
}
