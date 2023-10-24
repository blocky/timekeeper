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
	return fmt.Sprintf("## %s %s %s %s", e.Date, e.Task.Project, e.Task.ID, e.Details)
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

// func (e *Entry) Marshal() ([]byte, error) {
// 	return []byte(e.String()), nil
// }

// func (e *Entry) Unmarshal(bytes []byte) error {
// 	pattern := `^##\s(?P<date>[\d]{4}-[\d]{2}-[\d]{2}:[\d]{4}-[\d]{4})\s(?P<project>[[:alnum:]]{24})\s(?P<id>[[:alnum:]]{24})\s(?P<details>.*)`
// 	r := regexp.MustCompile(pattern)
// 	matches := r.FindStringSubmatch()

// 	if len(matches) != 5 {
// 		return fmt.Errorf("key:'%s' does not match regex", bytes)
// 	}

// }
