package timecard

import (
	"bufio"
	"os"

	"github.com/blocky/timekeeper/internal/entry"
)

type Timecard struct {
	*bufio.Writer
}

func MakeTimecard(f *os.File) Timecard {
	return Timecard{bufio.NewWriter(f)}
}

func (t *Timecard) WriteEntry(entry entry.Entry) error {
	_, err := t.WriteString(entry.String())
	if err != nil {
		return err
	}
	t.Flush()
	return nil
}
