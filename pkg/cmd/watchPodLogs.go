package cmd

import (
	"github.com/projectjudge/ktool/pkg/views"
	"github.com/spf13/cobra"
)

// WatchPodLogsCommand watches a pod logs in k8s cluster
var WatchPodLogsCommand = &cobra.Command{
	Use:   "l",
	Short: "Watch a pods logs",
	Long:  "Watch a pods logs",
	Run: func(cmd *cobra.Command, args []string) {
		views.WatchPodLogs()
	},
}
