package chronos

import (
	"encoding/json"
	"fmt"
)

type Year int

func MakeYear(y int) (Year, error) {
	if y < 0 {
		return 0, fmt.Errorf("year is invalid: %d", y)
	}
	return Year(y), nil
}

func (y *Year) UnmarshalJSON(bytes []byte) error {
	var i int
	err := json.Unmarshal(bytes, &i)
	if err != nil {
		return fmt.Errorf("could not unmarshal year", err)
	}

	y2, err := MakeYear(i)
	if err != nil {
		return err
	}
	*y = y2
	return nil
}
