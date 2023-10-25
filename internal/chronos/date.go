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

func (d Date) StartDateAndTime() string {
	return fmt.Sprintf("%d-%d-%d %s", d.Year, d.Month, d.Day, d.Start.TimeWithColon())
}

func (d Date) StopDateAndTime() string {
	return fmt.Sprintf("%d-%d-%d %s", d.Year, d.Month, d.Day, d.Stop.TimeWithColon())
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
