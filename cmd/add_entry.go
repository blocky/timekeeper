package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/blocky/timekeeper/internal/ask"
	"github.com/blocky/timekeeper/internal/chronos"
	"github.com/blocky/timekeeper/internal/entry"
	"github.com/blocky/timekeeper/internal/tap"
	"github.com/blocky/timekeeper/internal/task"
	"github.com/blocky/timekeeper/internal/timecard"
)

var Year int
var Month int
var Day int
var Start string
var Stop string

var TaskID string
var Details string

var current time.Time = time.Now()

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
	Run: func(cmd *cobra.Command, args []string) {
		addEntry()
	},
}

func init() {
	f := addEntryCmd.Flags()

	f.IntVarP(&Year, "year", "y", current.Year(), "year")
	f.IntVarP(&Month, "month", "m", int(current.Month()), "month")
	f.IntVarP(&Day, "day", "d", current.Day(), "day")

	f.StringVar(&Start, "start", "0000", "start (military time)")
	f.StringVar(&Stop, "stop", "0000", "stop (military time)")

	f.StringVar(&TaskID, "task-id", "", "task id from clockify")
	f.StringVar(&Details, "details", "", "details of your work")

	addCmd.AddCommand(addEntryCmd)
}

func addEntry() {
	var tasks task.Tasks
	err := json.Unmarshal(TasksJSON, &tasks)
	check(err)

	tap, err := tap.MakeAppendingTap(TimecardFilepath)
	check(err)

	var e entry.Entry

	if Interactive {
		e, err = ask.AskEntry(tasks)
		check(err)

	} else {
		date, err := chronos.MakeDateFromRaw(Year, Month, Day, Start, Stop)
		check(err)

		taskMap := task.MakeTaskMap(tasks)

		task, err := taskMap.GetTask(TaskID)
		check(err)

		e = entry.MakeEntry(date, task, Details)
	}

	fmt.Printf("time entry is: %+v\n", e)

	t := timecard.MakeTimecard(tap)
	err = t.WriteEntry(e)
	check(err)
}
