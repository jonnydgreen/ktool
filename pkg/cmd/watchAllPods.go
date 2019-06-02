package cmd

import (
	"github.com/projectjudge/ktool/pkg/pods"
	"github.com/spf13/cobra"
)

// WatchAllPodsCmd watches all pods in k8s cluster
var WatchAllPodsCmd = &cobra.Command{
	Use:   "a",
	Short: "Watch all the pods in a k8s cluster",
	Long:  "Watch all the pods in a k8s cluster",
	Run: func(cmd *cobra.Command, args []string) {
		pods.WatchPods()
	},
}
