package entry

import (
	"encoding/json"
	"fmt"

	"github.com/blocky/timekeeper/internal/chronos"
	"github.com/blocky/timekeeper/internal/task"
)

type Entry struct {
	Date    chronos.Date `json:"date"`
	Task    task.Task    `json:"task"`
	Details string       `json:"details"`
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

func (e *Entry) UnmarshalJSON(bytes []byte) error {
	type Alias Entry
	err := json.Unmarshal(bytes, &struct {
		*Alias
	}{
		Alias: (*Alias)(e),
	})
	if err != nil {
		return fmt.Errorf("could not unmarshal entry: %s", err)
	}
	return nil
}
