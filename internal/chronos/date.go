package chronos

import (
	"encoding/json"
	"fmt"
)

type Date struct {
	Year  Year         `json:"year"`
	Month Month        `json:"month"`
	Day   Day          `json:"day"`
	Start MilitaryTime `json:"start"`
	Stop  MilitaryTime `json:"stop"`
}

func MakeDate(
	year Year,
	month Month,
	day Day,
	start MilitaryTime,
	stop MilitaryTime,
) Date {
	return Date{
		Year:  year,
		Month: month,
		Day:   day,
		Start: start,
		Stop:  stop,
	}
}

func MakeDateFromRaw(
	yearInt int,
	monthInt int,
	dayInt int,
	startStr string,
	stopStr string,
) (Date, error) {
	var msg = "error creating date: %s"

	year, err := MakeYear(yearInt)
	if err != nil {
		return Date{}, fmt.Errorf(msg, err)
	}

	month, err := MakeMonth(monthInt)
	if err != nil {
		return Date{}, fmt.Errorf(msg, err)
	}

	day, err := MakeDay(dayInt)
	if err != nil {
		return Date{}, fmt.Errorf(msg, err)
	}

	start, err := MakeMilitaryTime(startStr)
	if err != nil {
		return Date{}, fmt.Errorf(msg, err)
	}

	stop, err := MakeMilitaryTime(stopStr)
	if err != nil {
		return Date{}, fmt.Errorf(msg, err)
	}

	return MakeDate(year, month, day, start, stop), nil
}

func (d Date) StartDateAndTime() string {
	return fmt.Sprintf("%d-%s-%s %s", d.Year, d.Month, d.Day, d.Start.TimeWithColon())
}

func (d Date) StopDateAndTime() string {
	return fmt.Sprintf("%d-%s-%s %s", d.Year, d.Month, d.Day, d.Stop.TimeWithColon())
}

func (d *Date) UnmarshalJSON(bytes []byte) error {
	type Alias Date
	err := json.Unmarshal(bytes, &struct {
		*Alias
	}{
		Alias: (*Alias)(d),
	})
	if err != nil {
		return fmt.Errorf("could not unmarshal date: %s", err)
	}
	return nil
}
