package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var snapCommand = &cobra.Command{
	Use:   "snap",
	Short: "Get a snapshot (dump) of your database",
	Run: func(cmd *cobra.Command, args []string) {
		// uri, _ := cmd.Flags().GetString("uri")
		db, _ := cmd.Flags().GetString("db")
		uri, _ := cmd.Flags().GetString("uri")
		out, _ := cmd.Flags().GetString("out")
		exclude, _ := cmd.Flags().GetStringSlice("exclude")

		excludeArgs := []string{
			"--uri=" + uri,
			"--db=" + db,
			"--out=" + out,
		}

		for _, s := range exclude {
			excludeArgs = append(excludeArgs, "--excludeCollection="+s)
		}

		dumpCmd := exec.Command("mongodump", excludeArgs...)
		dumpCmd.Stdout = os.Stdout
		dumpCmd.Stderr = os.Stderr

		err := dumpCmd.Run()
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(snapCommand)
	snapCommand.PersistentFlags().String("db", "", "The mongodb database to get a snapshot from.")
	snapCommand.PersistentFlags().String("out", "dump", "The output directory.")
	snapCommand.PersistentFlags().String("uri", "mongodb://localhost:27017", "The mongodb connection string.")
	snapCommand.PersistentFlags().StringSlice("exclude", []string{}, "Collections to exclude")
}
