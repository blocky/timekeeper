#!/bin/bash

set -e

TIMECARD="$HOME/timecard.json"
UPLOADS="$HOME/.timecard-uploads"
QUEUE_ENTRIES="$HOME/queue-entries.sh"

YEAR=$(date +%Y)

echo "Enter Month"
MONTH_STR=$(gum choose --selected $(date +%B) --height 12 "January" "February" "March" "April" "May" "June" "July" "August" "September" "October" "November" "December")
MONTH=$(date --date="$(printf "01 %s" $MONTH_STR)" +"%m")

gum style --foreground 212 --padding "0 2" "* $MONTH_STR"

echo "Enter Day"
DAY=$(date +%d)
DAY=$(gum filter --indicator ">" --height 10 --placeholder "today is $DAY" {1..31})

gum style --foreground 212 --padding "0 2" "* $DAY"

MILITARY_TIMES=$(echo -e {00..23}{00,15,30,45}'\n' | sed 's/ //g')

echo "Enter Start Time (military time)"
START=$(gum filter --indicator ">" --height 10 $(echo "$MILITARY_TIMES"))

echo "Enter Stop Time (military time)"
STOP=$(gum filter --indicator ">" --height 10 $(echo "$MILITARY_TIMES"))

gum style --foreground 212 --padding "0 2" "* $START->$STOP"

TASK_ID=$(echo $TASK | cut -d':' -f1)
TASK_NAME=$(echo $TASK | cut -d':' -f2)

echo "Select Task"
TASKS=$(tk list task --key-value)
TASK=$(echo "$TASKS" | gum filter --height 10)

TASK_ID=$(echo $TASK | cut -d':' -f1)
TASK_NAME=$(echo $TASK | cut -d':' -f2 | tr -d \')

gum style --foreground 212 --padding "0 2" "* $TASK_NAME"

echo "Enter Details"
DETAILS=$(gum input --placeholder "I did a thing!")

gum style --foreground 212 --padding "0 2" "* $DETAILS"

echo "Writing to Entry Queue"

echo "tk add entry -t=$TIMECARD -u=$UPLOADS" \
	"-m=$MONTH -d=$DAY --start=$START --stop=$STOP" \
	"--task-id=$TASK_ID --details='$DETAILS'" >> $QUEUE_ENTRIES
