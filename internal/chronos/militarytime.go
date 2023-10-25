package chronos

import (
	"encoding/json"
	"fmt"
)

type MilitaryTime string

func MakeMilitaryTime(s string) (MilitaryTime, error) {
	if len(s) == 3 {
		return MilitaryTime("0" + s), nil

	} else if len(s) < 3 || 4 < len(s) {
		return "", fmt.Errorf("could not create military time (HHMM) from: %s", s)
	}
	return MilitaryTime(s), nil
}

func (mt *MilitaryTime) UnmarshalJSON(bytes []byte) error {
	var s string
	err := json.Unmarshal(bytes, &s)
	if err != nil {
		return fmt.Errorf("could not unmarshal military time: %s", err)
	}

	mt2, err := MakeMilitaryTime(s)
	if err != nil {
		return fmt.Errorf("could not unmarshal military time: %s", err)
	}
	*mt = mt2
	return nil
}
