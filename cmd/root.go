package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mongo-snap",
	Short: "A tool to snapshot and restore mongo databases.",
}

func Execute() {
	rootCmd.Execute()
}
