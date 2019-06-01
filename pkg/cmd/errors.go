package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// CheckError prints err to stderr and exits with code 1 if err is not nil. Otherwise, it is a
// no-op.
func CheckError(err error) {
	if err != nil {
		if err != context.Canceled {
			fmt.Fprintf(os.Stderr, fmt.Sprintf("An error occurred: %v\n", err))
		}
		os.Exit(1)
	}
}

// Exit prints msg (with optional args), plus a newline, to stderr and exits with code 1.
func Exit(code int, msg string, args ...interface{}) {
	exitMessage := fmt.Sprintf("Received code: %d | %s | Exiting gracefully...\n", code, msg)
	if msg == "" {
		exitMessage = fmt.Sprintf("Received code: %d | Exiting gracefully...\n", code)
	}
	fmt.Fprintf(os.Stderr, exitMessage, args...)
	os.Exit(code)
}

// RegisterCleanup cleans up when call to cancel is initiated
func RegisterCleanup() chan int {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	exitChan := make(chan int)
	go func() {
		for sig := range signalChan {
			switch sig {
			case syscall.SIGHUP:
				fmt.Println("\nReceived SIGHUP, triggering cleanup...")
			case syscall.SIGINT:
				fmt.Println("\nReceived SIGINT, triggering cleanup...")
			case syscall.SIGTERM:
				fmt.Println("\nReceived SIGTERM, triggering cleanup...")
			case syscall.SIGQUIT:
				fmt.Println("\nReceived SIGQUIT, triggering cleanup...")
			default:
				fmt.Println("\nReceived Unknown signal")
				exitChan <- 1
			}
			exitChan <- 0
			exitChan <- 1
			exitChan <- 2
			exitChan <- 3
		}
	}()
	return exitChan
}

// Cleanup Cleans up after ourself
func Cleanup(exitChan chan int) {
	code := <-exitChan
	Exit(code, "")
}
