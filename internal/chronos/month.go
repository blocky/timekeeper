package chronos

import (
	"fmt"
)

type Month int

func MakeMonth(m int) (Month, error) {
	if 0 >= m || m > 12 {
		return 0, fmt.Errorf("month is invalid: '%d'", m)
	}
	return Month(m), nil
}

func (m Month) String() string {
	if m < 10 {
		return fmt.Sprintf("0%d", m)
	}
	return fmt.Sprintf("%d", m)
}
