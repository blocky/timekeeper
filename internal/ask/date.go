package ask

import (
	"github.com/blocky/timekeeper/internal/chronos"
)

func AskDate() (chronos.Date, error) {
	year, _ := chronos.MakeYear(2023)

	month, err := AskMonth("what month? (integer)")
	if err != nil {
		return chronos.Date{}, err
	}

	day, err := AskDay("what day? (integer)")
	if err != nil {
		return chronos.Date{}, err
	}

	start, err := AskMilitaryTime("start hour? (HourMinutes)")
	if err != nil {
		return chronos.Date{}, err
	}

	stop, err := AskMilitaryTime("stop hour? (HourMinutes)")
	if err != nil {
		return chronos.Date{}, err
	}

	date := chronos.Date{
		Year:  year,
		Month: month,
		Day:   day,
		Start: start,
		Stop:  stop,
	}
	return date, nil
}
