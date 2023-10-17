package ask

import "fmt"

func AskDetails() string {
	details := Ask("what did you do?")
	fmt.Printf("details: %s\n", details)
	return details
}
