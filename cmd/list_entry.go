package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/blocky/timekeeper/internal/entry"
	"github.com/blocky/timekeeper/internal/tap"
	"github.com/blocky/timekeeper/internal/timecard"
)

var listEntryCmd = &cobra.Command{
	Use:   "entry",
	Short: "list entries",
	Run: func(cmd *cobra.Command, args []string) {
		listEntries()
	},
}

func init() {
	listEntryCmd.Flags().BoolVarP(&ListAll, "all", "a", false, "list all entries")
	listEntryCmd.Flags().UintVarP(&ListNumberOfLatestEntries, "list-latest-entries", "n", 1, "list latest number of entries")
	listEntryCmd.Flags().BoolVarP(&ListPretty, "pretty", "p", false, "list as pretty-fied JSON")

	listCmd.AddCommand(listEntryCmd)
}

func listEntries() {
	tap, err := tap.MakeAppendingTap(TimecardFilepath)
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
	for _, e := range entries {

		if ListPretty {
			bytes, err = json.MarshalIndent(e, " ", " ")
		} else {
			bytes, err = json.Marshal(e)
		}
		check(err)

		fmt.Printf("%s\n", bytes)
	}
}
