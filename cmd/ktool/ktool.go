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

	// command := cmd.NewDefaultKubectlCommand()

	// // TODO: once we switch everything over to Cobra commands, we can go back to calling
	// // cliflag.InitFlags() (by removing its pflag.Parse() call). For now, we have to set the
	// // normalize func and add the go flag set by hand.
	// pflag.CommandLine.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)
	// pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	// // cliflag.InitFlags()
	// logs.InitLogs()
	// defer logs.FlushLogs()

	// Run our command in a separate goroutine
	go func() {
		pods.WatchPods()
	}()

	cmd.Cleanup(exitChan)
}
