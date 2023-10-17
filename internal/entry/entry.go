package entry

import (
	"fmt"

	"github.com/blocky/timekeeper/internal/chronos"
	"github.com/blocky/timekeeper/internal/task"
)

type Entry struct {
	Date    chronos.Date
	Task    task.Task
	Details string
}

func (e Entry) String() string {
	return fmt.Sprintf("## %s %s %s %s\n", e.Date, e.Task.Project, e.Task.ID, e.Details)
}

func MakeEntry(
	date chronos.Date,
	task task.Task,
	details string,
) Entry {
	return Entry{
		Date:    date,
		Task:    task,
		Details: details,
	}
}
