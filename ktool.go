package main

import (
	"math/rand"
	"time"

	"github.com/projectjudge/ktool/pkg/cmd"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	exitChan := cmd.RegisterCleanup()

	// Run our commands in a separate goroutine
	go func() {
		cmd.Execute()
	}()

	cmd.Cleanup(exitChan)
}
