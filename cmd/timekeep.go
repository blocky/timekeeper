package cmd

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/blocky/timekeeper/timekeep"
)

//go:embed configs/tasks.json
var tasksJSON []byte

var addEntryCmd = &cobra.Command{
	Use:   "add-entry [file]",
	Short: "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		addEntry(tasksJSON, filename)
	},
}

func init() {
	rootCmd.AddCommand(addEntryCmd)
}

func addEntry(tasksJSON []byte, filename string) {
	var tasks []timekeep.Task
	err := json.Unmarshal(tasksJSON, &tasks)
	check(err)

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	check(err)
	defer f.Close()

	tk := timekeep.MakeTimekeeper(f, tasks)
	entry, err := tk.MakeEntry()
	check(err)

	fmt.Printf("time entry is:\n%s", entry)

	err = tk.WriteEntry(entry)
	check(err)
}
