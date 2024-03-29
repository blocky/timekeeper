package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Use the interactive mode
var Interactive bool

//go:embed configs/tasks.json
var TasksJSON []byte

// Timecard filepath
var TimecardFilepath string

// Timecard upload config
var TimecardUploadFilepath string

var rootCmd = &cobra.Command{
	Short: "",
	Long:  `A tool to record & upload work hours`,
}

func init() {
	f := rootCmd.PersistentFlags()

	f.BoolVarP(&Interactive, "interactive", "i", false, "use interactive prompts")
	f.StringVarP(&TimecardFilepath, "timecard", "t", "timecard.json", "timecard filepath")
	f.StringVarP(&TimecardUploadFilepath, "uploads", "u", ".timecard-uploads", "timecard upload config filepath")
}

func Execute() {
	err := rootCmd.Execute()
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		os.Exit(1)
	}
}
