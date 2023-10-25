package timecard

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/blocky/timekeeper/internal/entry"
	"github.com/blocky/timekeeper/internal/tap"
)

type Timecard struct {
	tap.Tap
}

func MakeTimecard(tap tap.Tap) Timecard {
	return Timecard{tap}
}

func (t Timecard) ReadEntries() ([]entry.Entry, error) {
	var entries []entry.Entry

	dec := json.NewDecoder(t.NewReader())
	for {
		var e entry.Entry
		if err := dec.Decode(&e); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}
	return entries, nil
}

func (t *Timecard) WriteEntry(entry entry.Entry) error {
	bytes, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("could not write entry: %s", err)
	}

	_, err = t.Write(bytes)
	if err != nil {
		return err
	}

	_, err = t.Write([]byte("\n"))
	if err != nil {
		return err
	}
	return nil
}
