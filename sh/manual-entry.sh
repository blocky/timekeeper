#!/bin/sh

set -e

TIMECARD="$HOME/timecard.json"

YEAR=$(date +"%Y")

echo "Enter month"
MONTH=$(date +"%m")
MONTH=$(gum input --value $MONTH --placeholder "default $MONTH")

echo "Enter day"
DAY=$(date +"%d")
DAY=$(gum input --value $DAY --placeholder "default $DAY")

echo "Enter start time (military time)"
START=$(gum input --placeholder "0000")

echo "Enter stop time (military time)"
STOP=$(gum input --placeholder "0000")

echo "Select task"
TASKS=$(tk list task --key-value)
TASK=$(echo "$TASKS" | gum filter --limit 1)
TASK_ID=$(echo $TASK | cut -d':' -f1)
TASK_NAME=$(echo $TASK | cut -d':' -f2)

echo "Selected $TASK_NAME"

echo "Enter your details"
DETAILS=$(gum input --placeholder "I did a thing!")

tk add entry \
	--timecard=$TIMECARD \
	--year=$YEAR \
	--month="$MONTH" \
	--day="$DAY" \
	--start="$START" \
	--stop="$STOP" \
	--task-id="$TASK_ID" \
	--details="$DETAILS"
