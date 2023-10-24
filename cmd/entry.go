package cmd

import (
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/blocky/timekeeper/internal/ask"
	"github.com/blocky/timekeeper/internal/tap"
	"github.com/blocky/timekeeper/internal/task"
	"github.com/blocky/timekeeper/internal/timecard"
)

//go:embed configs/tasks.json
var tasksJSON []byte

var entryCmd = &cobra.Command{
	Use:   "entry [file]",
	Short: "Add a time entry to the timecard",
	Long: `Add work done in an epic during some day during a specific time

Information required:
# -------------------
Month:
Day:
StartTime:
FinishTime:
Task (Jira Epic):
Task Description:
#-------------------
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		addEntry(tasksJSON, filename)
	},
}

func init() {
	addCmd.AddCommand(entryCmd)
}

func addEntry(tasksJSON []byte, filename string) {
	var tasks task.Tasks
	err := json.Unmarshal(tasksJSON, &tasks)
	check(err)

	tap, err := tap.MakeTap(filename)
	check(err)
	defer tap.Close()

	entry, err := ask.AskEntry(tasks)
	check(err)

	fmt.Printf("time entry is:\n%s", entry)

	t := timecard.MakeTimecard(tap)
	err = t.WriteEntry(entry)
	check(err)
}