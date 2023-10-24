package timecard

import (
	"github.com/blocky/timekeeper/internal/entry"
	"github.com/blocky/timekeeper/internal/tap"
)

type Timecard struct {
	tap.Tap
}

func MakeTimecard(tap tap.Tap) Timecard {
	return Timecard{tap}
}

// func (t *Timecard) ListEntries() ([]entry.Entry, error) {
// 	bytes, err := t.ReadAll()
// 	if err != nil {
// 		return nil, fmt.Errorf("could not list entries: %s", err)
// 	}
// 	var entries []entry.Entry
// 	err := json.Unmarshal()
// }

func (t *Timecard) WriteEntry(entry entry.Entry) error {
	_, err := t.WriteLine(entry.String())
	if err != nil {
		return err
	}
	return nil
}
