package chronos

import (
	"fmt"
)

type Date struct {
	Year  int
	Month Month
	Day   Day
	Start MilitaryTime
	Stop  MilitaryTime
}

func (t Date) String() string {
	return fmt.Sprintf("%d-%s-%s:%s-%s", t.Year, t.Month, t.Day, t.Start, t.Stop)
}
