package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/ianhecker/timekeeper/timekeep"
)

var addEntryCmd = &cobra.Command{
	Use:   "add-entry [file]",
	Short: "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		run(filename)
	},
}

func init() {
	rootCmd.AddCommand(addEntryCmd)
}

func run(filename string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	check(err)
	defer f.Close()

	tk := timekeep.MakeTimekeeper(f)
	entry, err := tk.MakeEntry()
	check(err)

	fmt.Printf("time entry is:\n%s", entry)

	err = tk.WriteEntry(entry)
	check(err)
}
