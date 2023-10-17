package ask

import (
	"github.com/blocky/timekeeper/internal/chronos"
)

func AskMilitaryTime(question string) (chronos.MilitaryTime, error) {
	answer := Ask(question)
	mt, err := chronos.MakeMilitaryTime(answer)
	if err != nil {
		return "", err
	}
	return mt, nil
}
