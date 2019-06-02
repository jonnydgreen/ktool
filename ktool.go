package main

import (
	"math/rand"
	"time"

	"github.com/projectjudge/ktool/pkg/cmd"
	"github.com/projectjudge/ktool/pkg/pods"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	exitChan := cmd.RegisterCleanup()

	// Run our command in a separate goroutine
	go func() {
		pods.WatchPods()
	}()

	cmd.Cleanup(exitChan)
}
