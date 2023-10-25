package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/blocky/timekeeper/internal/tap"
	"github.com/blocky/timekeeper/internal/timecard"
)

var listCmd = &cobra.Command{
	Use:   "list [file]",
	Short: "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		listEntries(filename)
	},
}

func init() {
	addCmd.AddCommand(listCmd)
}

func listEntries(filename string) {
	tap, err := tap.MakeTap(filename)
	check(err)

	t := timecard.MakeTimecard(tap)
	entries, err := t.ReadEntries()
	check(err)

	for _, entry := range entries {
		fmt.Printf("%+v\n", entry)
	}
}
