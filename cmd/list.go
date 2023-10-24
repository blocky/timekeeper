package cmd

// import (
// 	"os"

// 	"github.com/spf13/cobra"

// 	"github.com/blocky/timekeeper/internal/timecard"
// )

// var listCmd = &cobra.Command{
// 	Use:   "list [file]",
// 	Short: "",
// 	Args:  cobra.ExactArgs(1),
// 	Run: func(cmd *cobra.Command, args []string) {
// 		filename := args[0]
// 		listEntries(filename)
// 	},
// }

// func init() {
// 	addCmd.AddCommand(listCmd)
// }

// func listEntries(filename string) {
// 	f, err := os.Open(filename)
// 	check(err)
// 	defer f.Close()

// 	t := timecard.MakeTimecard(f)

// }
