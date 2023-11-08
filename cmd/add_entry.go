package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/blocky/timekeeper/internal/ask"
	"github.com/blocky/timekeeper/internal/tap"
	"github.com/blocky/timekeeper/internal/task"
	"github.com/blocky/timekeeper/internal/timecard"
)

var addEntryCmd = &cobra.Command{
	Use:   "entry",
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
		addEntry()
	},
}

func init() {
	addCmd.AddCommand(addEntryCmd)
}

func addEntry() {
	var tasks task.Tasks
	err := json.Unmarshal(TasksJSON, &tasks)
	check(err)

	tap, err := tap.MakeAppendingTap(TimecardFilepath)
	check(err)

	entry, err := ask.AskEntry(tasks)
	check(err)

	fmt.Printf("time entry is: %+v\n", entry)

	t := timecard.MakeTimecard(tap)
	err = t.WriteEntry(entry)
	check(err)
}
