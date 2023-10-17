package ask

import (
	"github.com/blocky/timekeeper/internal/chronos"
)

func AskDay(question string) (chronos.Day, error) {
	d, err := AskInt(question)
	if err != nil {
		return 0, err
	}
	day, err := chronos.MakeDay(d)
	if err != nil {
		return 0, err
	}
	return day, nil
}
