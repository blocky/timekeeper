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

// func (d *Date) Unmarshal(bytes []byte) error {
// 	pattern := `^(?P<year>[\d]{4})-(?P<month>[\d]{2})-(?P<day>[\d]{2}):(?P<start>[\d]{4})-(?P<stop>[\d]{4})`
// 	r := regexp.MustCompile(pattern)
// 	matches := r.FindStringSubmatch(string(bytes))

// 	if len(matches) != 6 {
// 		return fmt.Errorf("key:'%s' does not match regex", bytes)
// 	}

// 	yearStr := matches[r.SubexpIndex("year")]
// 	year, err := strconv.Atoi(yearStr)
// 	if err != nil {
// 		return fmt.Errorf("could not convert year to int: %s", err)
// 	}

// 	monthStr := matches[r.SubexpIndex("month")]
// 	monthInt, err := strconv.Atoi(monthStr)
// 	if err != nil {
// 		return fmt.Errorf("could not convert month to int: %s", err)
// 	}

// 	month, err := MakeMonth(monthInt)
// 	if err != nil {
// 		return err
// 	}

// }
