# Timekeeper
A small set of tools to automate timekeeping for Clockify while recording a
backup set of hours in markdown

## Set Up Timecard
Create an empty file of your choice for time keeping. I recommend setting up a
version control system for your hours and a symbolic link:

```bash
mkdir timecards
touch timecards/timecard.json
ln -s /home/USER/timecards/timecard.json /home/USER/
```

This file will hold your hour entries in JSON1

## Clockify-cli
These tools leverage Clockify-cli to do things like send hours to Clockify or
fetch list of active epics

```bash
https://github.com/lucassabreu/clockify-cli
```

1. You will need to download and install the binary
2. You will need to setup a config

```bash
clockify-cli config
```

3. You will need to set the following values in the `.clockify-cli.yaml`

```yaml
allow-incomplete: true
token: <your-clockify-token>
```

## Fetch Tasks
You will need to fetch the different epics from clockify. These include epics
for work, time off, and meetings. This make command will fetch all the pertinent
epics and combine them into one list for the Golang tool to use

```bash
make tasks
```

This list will be embedded into the Golang binary - so it is important to update
this list regurlary for active epics

## Install the Timekeeper Binary
Check into the repo and build & install a binary

> This will automatically embed current epics from Clockify into the binary

```bash
make install
```

## Logging your Hours
This tool allows you to input your hours retrospectively. You will need to know:

1. Month
2. Day
2. Starting hour & minutes in military time: HHMM
	- 8:00am -> 0800
3. Stopping hours & minutes in military time: HHMM
	- 5:00pm -> 1700
4. Name of the Blocky Epic you worked on
	- Epics will be listed and you can select the correct one
5. Brief description of your work

The binary will prompt you for input and then append your time entry to the end
of the file

```bash
tk add entry ~/timecard.json
```

## Listing your Hours
You can list your hours from your timecard file. Some examples are:

```bash
tk list --all ~/timecard.json
tk list --all --pretty ~/timecard.json
tk list -n 3 ~/timecard.json
tk list -n 3 --pretty ~/timecard.json
```

## Upload Hours to Clockify
Currenty, there is no system for recognizing which hours have or have not been
uploaded to Clockify

> Uploading hours that have already been uploaded will cause duplication in
Clockify! - please be aware of what hours you are uploading to prevent
large-scale duplication!

To upload your most recent entry to Clockify:

```bash
tk list -n 1 ~/timecard.json
tk upload -n 1 ~/timecard.json
```
