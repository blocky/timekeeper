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

func (d Date) String() string {
	return fmt.Sprintf("%d-%s-%s:%s-%s", d.Year, d.Month, d.Day, d.Start, d.Stop)
}
