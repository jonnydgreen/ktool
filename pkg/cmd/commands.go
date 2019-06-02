package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ktool",
	Short: "Ktool is a simplified kubectl",
	Long:  "Ktool is a simplified kubectl",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(WatchAllPodsCmd)
	rootCmd.AddCommand(WatchPodLogsCommand)
}

// Execute executes cobra
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
