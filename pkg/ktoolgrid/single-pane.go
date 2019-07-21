package ktoolgrid

import (
	"time"

	ui "github.com/gizak/termui/v3"
)

// SinglePane sets up a single Pane
func SinglePane(content interface{}) (*ui.Grid, <-chan ui.Event, <-chan time.Time, error) {
	// InitGrid()

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0, content),
	)

	ui.Render(grid)

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C

	return grid, uiEvents, ticker, nil
}
