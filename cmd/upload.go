package cmd

import (
	"github.com/blocky/timekeeper/internal/tap"
	"github.com/blocky/timekeeper/internal/timecard"
	"github.com/blocky/timekeeper/internal/upload"
	"github.com/spf13/cobra"
)

var DryRun bool
var UploadAll bool
var UploadNumberOfLatestEntries uint
var Verbose bool

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload entries",
	Run: func(cmd *cobra.Command, args []string) {
		uploadEntries()
	},
}

func init() {
	f := uploadCmd.Flags()

	f.BoolVarP(&DryRun, "dry-run", "d", false, "do a dry run of uploads")
	f.BoolVarP(&UploadAll, "all", "a", false, "list all entries")
	f.UintVarP(&UploadNumberOfLatestEntries, "list-latest-entries", "n", 1, "list latest number of entries")
	f.BoolVarP(&Verbose, "verbose", "v", false, "verbose mode on uploader")

	rootCmd.AddCommand(uploadCmd)
}

func uploadEntries() {
	timecardTap, err := tap.MakeAppendingTap(TimecardFilepath)
	check(err)

	uploadTap, err := tap.MakeCreatingTap(TimecardUploadFilepath)
	check(err)

	timecardConfig := timecard.MakeTimecard(timecardTap)
	entries, err := timecardConfig.ReadEntries()
	check(err)

	uploader := upload.MakeUploader(uploadTap, DryRun, Verbose)
	err = uploader.ReadInConfig()
	check(err)

	if UploadAll {
		err = uploader.UploadAll(entries)

	} else {
		err = uploader.UploadNumberOfLatestEntries(
			entries,
			UploadNumberOfLatestEntries,
		)
	}
	check(err)

	err = uploader.UpdateConfig()
	check(err)
}
