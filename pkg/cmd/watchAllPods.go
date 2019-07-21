package cmd

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/projectjudge/ktool/pkg/config"
	"github.com/projectjudge/ktool/pkg/ktoolgrid"
	"github.com/projectjudge/ktool/pkg/pods"
	"github.com/projectjudge/ktool/pkg/utils"
	"github.com/spf13/cobra"
)

// WatchAllPodsCmd watches all pods in k8s cluster
var WatchAllPodsCmd = &cobra.Command{
	Use:   "a",
	Short: "Watch all the pods in a k8s cluster",
	Long:  "Watch all the pods in a k8s cluster",
	Run: func(cmd *cobra.Command, args []string) {
		ktoolgrid.InitGrid()
		defer ui.Close()

		// Get kubeconfig
		kubeconfig, _ := config.KubeConfig()

		// Setup pane to render
		pane := widgets.NewParagraph()
		pane.Text = pods.GetPods(kubeconfig)
		pane.Title = "Context | All pods"

		// Setup grid
		grid, uiEvents, ticker, err := ktoolgrid.SinglePane(pane)
		utils.ErrorHandler(err)

		for {
			select {
			case e := <-uiEvents:
				switch e.ID {
				case "q", "<C-c>":
					return
				case "<Resize>":
					payload := e.Payload.(ui.Resize)
					grid.SetRect(0, 0, payload.Width, payload.Height)
					ui.Clear()
					ui.Render(grid)
				}
			case <-ticker:
				// Get pods
				pane.Text = pods.GetPods(kubeconfig)
				ui.Render(grid)
			}
		}
	},
}
