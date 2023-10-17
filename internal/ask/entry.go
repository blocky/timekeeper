package ask

import (
	"github.com/blocky/timekeeper/internal/entry"
	"github.com/blocky/timekeeper/internal/task"
)

func AskEntry(tasks task.Tasks) (entry.Entry, error) {
	date, err := AskDate()
	if err != nil {
		return entry.Entry{}, err
	}

	task, err := AskTask(tasks)
	if err != nil {
		return entry.Entry{}, err
	}
	details := AskDetails()

	return entry.MakeEntry(date, task, details), nil
}
