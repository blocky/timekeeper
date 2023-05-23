# Timekeeper
A small set of tools to automate timekeeping for Clockify while recording a
backup set of hours in markdown

## Set Up Timecard
Create an empty file of your choice for time keeping. I recommened:

	touch ~/timecard.md

This file will hold the automated markdown

I also recommend that you place this into version control. A basic repo will do

## Install the Timekeeper Binary
Check into the repo and build & install a binary

	make install

## Source Bash Scripts
You will need to source the Bash scripts so that you can update Clockify

	vi ~/.bash_profile
	export PATH="/home/USER/timekeeper/sh:$PATH"

## Test Out Logging Your Hours

	tk ~/timecard.md
