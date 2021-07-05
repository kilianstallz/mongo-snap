package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var restoreCommand = &cobra.Command{
	Use:   "restore",
	Short: "Restore a dump.",
	Run: func(cmd *cobra.Command, args []string) {
		uri, _ := cmd.Flags().GetString("uri")
		dir, _ := cmd.Flags().GetString("dir")

		restoreArgs := []string{
			"--uri=" + uri,
			"--dir=" + dir,
		}

		dumpCmd := exec.Command("mongorestore", restoreArgs...)
		dumpCmd.Stdout = os.Stdout
		dumpCmd.Stderr = os.Stderr

		err := dumpCmd.Run()
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(restoreCommand)
	restoreCommand.PersistentFlags().String("dir", "dump", "The directory to restore from.")
	restoreCommand.PersistentFlags().String("uri", "mongodb://localhost:27017", "The mongodb connection string.")
}
