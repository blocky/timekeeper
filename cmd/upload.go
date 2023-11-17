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

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload entries",
	Run: func(cmd *cobra.Command, args []string) {
		uploadEntries()
	},
}

func init() {
	uploadCmd.Flags().BoolVarP(&DryRun, "dry-run", "d", true, "do a dry run of uploads")
	uploadCmd.Flags().BoolVarP(&UploadAll, "all", "a", false, "list all entries")
	uploadCmd.Flags().UintVarP(&UploadNumberOfLatestEntries, "list-latest-entries", "n", 1, "list latest number of entries")

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

	uploader := upload.MakeUploader(uploadTap, DryRun)
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
