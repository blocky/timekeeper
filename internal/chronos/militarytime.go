package chronos

import (
	"fmt"
)

type MilitaryTime string

func MakeMilitaryTime(s string) (MilitaryTime, error) {
	if len(s) != 4 {
		return "", fmt.Errorf("provide 4 digits (HourMinutes)")
	}
	return MilitaryTime(s), nil
}
