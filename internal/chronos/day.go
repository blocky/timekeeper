package chronos

import (
	"fmt"
)

type Day int

func MakeDay(d int) (Day, error) {
	if 0 >= d || d > 31 {
		return 0, fmt.Errorf("day is invalid: '%d'", d)
	}
	return Day(d), nil
}

func (d Day) String() string {
	if d < 10 {
		return fmt.Sprintf("0%d", d)
	}
	return fmt.Sprintf("%d", d)
}
