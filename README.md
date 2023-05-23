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

	tk add-entry ~/timecard.md
