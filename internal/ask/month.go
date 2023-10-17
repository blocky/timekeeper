package ask

import (
	"github.com/blocky/timekeeper/internal/chronos"
)

func AskMonth(question string) (chronos.Month, error) {
	m, err := AskInt(question)
	if err != nil {
		return 0, err
	}
	month, err := chronos.MakeMonth(m)
	if err != nil {
		return 0, err
	}
	return month, nil
}
