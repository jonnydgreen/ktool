package ktoolgrid

import (
	"log"

	ui "github.com/gizak/termui/v3"
)

// InitGrid initialises the grid for use in our program
func InitGrid() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize ktool: %v", err)
	}
}
