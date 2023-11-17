package upload

import (
	"encoding/json"
	"fmt"

	"github.com/blocky/timekeeper/internal/bash"
	"github.com/blocky/timekeeper/internal/entry"
	"github.com/blocky/timekeeper/internal/tap"
)

type Uploader struct {
	tap.Tap
	uploads Uploads
	dryRun  bool
	verbose bool
}

func MakeUploader(tap tap.Tap, dryRun bool, verbose bool) Uploader {
	return Uploader{
		tap,
		MakeUploads(),
		dryRun,
		verbose,
	}
}

func (u Uploader) ReadConfig() (Uploads, error) {
	bytes, err := u.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read upload file: %s", err)
	}

	var uploads Uploads
	err = json.Unmarshal(bytes, &uploads)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal uploads: %s", err)
	}
	return uploads, nil
}

func (u *Uploader) ReadInConfig() error {
	uploads, err := u.ReadConfig()
	if err != nil {
		return err
	}

	u.uploads = uploads
	return nil
}

func (u *Uploader) UpdateConfig() error {
	bytes, err := json.MarshalIndent(u.uploads, " ", " ")
	if err != nil {
		return fmt.Errorf("could not write entry: %s", err)
	}

	_, err = u.WriteFromBeginning(bytes)
	if err != nil {
		return err
	}

	return nil
}

func (u *Uploader) UploadAll(entries []entry.Entry) error {
	for _, e := range entries {
		err := u.Upload(e)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *Uploader) UploadNumberOfLatestEntries(
	entries []entry.Entry,
	n uint,
) error {
	var latest []entry.Entry
	for i := len(entries) - 1; i >= 0 && n > 0; i-- {
		latest = append(latest, entries[i])
		n--
	}
	return u.UploadAll(latest)
}

func (u *Uploader) Upload(e entry.Entry) error {
	uploaded := u.uploads[e.ID]
	if uploaded {
		if u.verbose {
			fmt.Printf("ID:'%s' is already uploaded\n", e.ID)
		}
		return nil
	}

	if u.dryRun {
		upload := fmt.Sprintf(bash.UploadFormat,
			e.Task.Project,
			e.Task.ID,
			e.Date.StartDateAndTime(),
			e.Date.StopDateAndTime(),
			e.Details,
		)
		fmt.Printf("would upload: %s", upload)
		return nil
	}

	output, err := bash.BashExec(
		bash.UploadFormat,
		e.Task.Project,
		e.Task.ID,
		e.Date.StartDateAndTime(),
		e.Date.StopDateAndTime(),
		e.Details,
	)
	if err != nil {
		return fmt.Errorf("bash cmd error: %s", err)
	}

	fmt.Println(output)
	u.uploads[e.ID] = true

	return nil
}
