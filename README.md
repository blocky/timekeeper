# Timekeeper
A small set of tools to automate timekeeping for Clockify while recording a
backup set of hours in markdown

## Set Up Timecard
Create an empty file of your choice for time keeping. I recommened:

	touch ~/timecard.md

This file will hold the automated markdown

I also recommend that you place this into version control. A basic repo will do

## Fetch Tasks
You will need to fetch the different epics, time off, and meetings tasks

	make tasks

This will create a file that will be embedded into the timekeep binary

## Install the Timekeeper Binary
Check into the repo and build & install a binary. This will also automatically
embed the necessary Clockify tasks as JSON into the binary

	make install

## Source Bash Scripts
You will need to source the Bash scripts so that you can update Clockify

	vi ~/.bash_profile
	export PATH="/home/USER/timekeeper/sh:$PATH"

## Clockify-cli
These tools leverage Clockify-cli to send hours to Clockify
https://github.com/lucassabreu/clockify-cli

You will need to download and install the binary

You will need to setup the config

	clockify-cli config

You will need to set the following values in the .clockify-cli.yaml

	allow-incomplete: true
	token: <your-clockify-token>

## Test Out Logging Your Hours
This tool allows you to input your hours retrospectively. You will need to know:
1) The Year, Month, and Day in the format: YYYY-MM-DD
2) Starting hour & minutes in military time: HHMM
3) Stopping hours & minutes in military time: HHMM
4) The Blocky Epic you worked on. This is a Clockify Task

The binary will prompt you for input and then output your time entry

	tk add-entry ~/timecard.md

The markdown should follow the format

	## YYYY-MM-DD:HHMM-HMMM CLOCKIFY-PROJECT CLOCKIFY-TASK DESCRIPTION
	## <year>-<month>-<day>:<starting-hour-and-minutes>-<ending-hour-and-minues> clockify-project clockify-task description

## Pushing Hours to Clockify
All your hours are stored in your local markdown file. Currently, the markdown
line

	---

This is used to denote the beginning of the hours that should be logged. In the
future, the tool will be able to tell what hours it has & has not logged

For now, place --- after lines you have sent to Clockify

	## my old hours
	---
	## hours to send to clockify

This script will grep the hours after the ---

	last-weeks-hours timecard.md

Now, the scripts will parse your hours and bulk update Clockify

	last-weeks-hours timecard.md | bulk-update-hours
