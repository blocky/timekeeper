package main

import (
	"github.com/blocky/timekeeper/cmd"
)

var Version string

func main() {
	cmd.Version = Version
	cmd.Execute()
}
