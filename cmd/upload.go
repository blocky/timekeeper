package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/blocky/timekeeper/internal/entry"
	"github.com/blocky/timekeeper/internal/tap"
	"github.com/blocky/timekeeper/internal/timecard"
	"github.com/spf13/cobra"
)

var UploadAll bool
var UploadNumberOfLatestEntries uint

var uploadCmd = &cobra.Command{
	Use:   "upload [file]",
	Short: "upload entries",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		uploadEntries(filename)
	},
}

func init() {
	uploadCmd.Flags().BoolVarP(&UploadAll, "all", "a", false, "list all entries")
	uploadCmd.Flags().UintVarP(&UploadNumberOfLatestEntries, "list-latest-entries", "n", 1, "list latest number of entries")

	rootCmd.AddCommand(uploadCmd)
}

func uploadEntries(filename string) {
	tap, err := tap.MakeTap(filename)
	check(err)

	t := timecard.MakeTimecard(tap)
	entries, err := t.ReadEntries()
	check(err)

	if !UploadAll {
		var list []entry.Entry

		n := UploadNumberOfLatestEntries
		for i := len(entries) - 1; i > -1 && n > 0; i-- {
			list = append(list, entries[i])
			n--
		}
		entries = list
	}
	for _, e := range entries {
		format := "clockify-cli manual --interactive=0 --project='%s' --task='%s' --when='%s' --when-to-close='%s' --description='%s'\n"
		out := BashExec(format, e.Task.Project, e.Task.ID, e.Date.StartDateAndTime(), e.Date.StopDateAndTime(), e.Details)
		fmt.Println(out)
	}
}

func BashExec(format string, a ...any) string {
	cmd := fmt.Sprintf(format, a...)

	out, err := exec.Command("bash", "-c", cmd).Output()
	check(err)
	return strings.TrimRight(string(out), "\n")
}
