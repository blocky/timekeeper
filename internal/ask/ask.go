package ask

import (
	"fmt"
	"strconv"
)

func Ask(question string) string {
	answer := prompt(question)
	return answer
}

func AskInt(question string) (int, error) {
	answer := Ask(question)
	i, err := strconv.Atoi(answer)
	if err != nil {
		return 0, fmt.Errorf("given non-int: '%s'", answer)
	}
	return i, nil
}
