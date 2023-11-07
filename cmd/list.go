package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/blocky/timekeeper/internal/entry"
	"github.com/blocky/timekeeper/internal/tap"
	"github.com/blocky/timekeeper/internal/timecard"
)

var ListAll bool
var ListNumberOfLatestEntries uint
var ListPretty bool

var listCmd = &cobra.Command{
	Use:   "list [timecard]",
	Short: "list entries",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		listEntries(filename)
	},
}

func init() {
	listCmd.Flags().BoolVarP(&ListAll, "all", "a", false, "list all entries")
	listCmd.Flags().UintVarP(&ListNumberOfLatestEntries, "list-latest-entries", "n", 1, "list latest number of entries")
	listCmd.Flags().BoolVarP(&ListPretty, "pretty", "p", false, "list pretty-fied entry JSON")

	rootCmd.AddCommand(listCmd)
}

func listEntries(filename string) {
	tap, err := tap.MakeTap(filename)
	check(err)

	t := timecard.MakeTimecard(tap)
	entries, err := t.ReadEntries()
	check(err)

	if !ListAll {
		var list []entry.Entry

		n := ListNumberOfLatestEntries
		for i := len(entries) - 1; i > -1 && n > 0; i-- {
			list = append(list, entries[i])
			n--
		}
		entries = list
	}

	var bytes []byte
	if ListPretty {
		bytes, err = json.MarshalIndent(entries, "", "")
	} else {
		bytes, err = json.Marshal(entries)
	}
	check(err)
	fmt.Printf("%s\n", string(bytes))
}
