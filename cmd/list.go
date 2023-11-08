package cmd

import (
	"github.com/spf13/cobra"
)

var ListAll bool
var ListNumberOfLatestEntries uint
var ListPretty bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list things",
}

func init() {
	rootCmd.AddCommand(listCmd)
}
