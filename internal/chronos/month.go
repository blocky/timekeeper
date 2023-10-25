package chronos

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Month int

func MakeMonth(m int) (Month, error) {
	if 0 >= m || m > 12 {
		return 0, fmt.Errorf("month is invalid: '%d'", m)
	}
	return Month(m), nil
}

func MakeMonthFromString(s string) (Month, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("could not convert:'%s' to int", s)
	}
	return MakeMonth(i)
}

func (m Month) String() string {
	if m < 10 {
		return fmt.Sprintf("0%d", m)
	}
	return fmt.Sprintf("%d", m)
}

func (m *Month) UnmarshalJSON(bytes []byte) error {
	var i int
	err := json.Unmarshal(bytes, &i)
	if err != nil {
		return fmt.Errorf("could not unmarshal month: %s", err)
	}

	m2, err := MakeMonth(i)
	if err != nil {
		return fmt.Errorf("could not unmarshal month: %s", err)
	}
	*m = m2
	return nil
}
