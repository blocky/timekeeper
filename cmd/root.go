package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "",
	Long:  ``,
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
