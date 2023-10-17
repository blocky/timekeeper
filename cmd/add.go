package cmd

import (
	_ "embed"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add to the timecard",
}

func init() {
	rootCmd.AddCommand(addCmd)
}
