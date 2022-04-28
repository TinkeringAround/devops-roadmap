#!/bin/bash
# This file is to demonstrate bash scripting

CHALLENGE="90DaysOfDevOps"
TOTAL_DAYS=90

echo "Enter your name:"
read -r name
echo "Hey $name, welcome to the $CHALLENGE!"
echo "How Many Days of the $CHALLENGE CHALLENGE have you completed?"
read -r daysCompleted

if [ "$daysCompleted" -eq $TOTAL_DAYS ]; then
  echo "You have finished, well done"
elif [ "$daysCompleted" -lt $TOTAL_DAYS ]; then
  echo "Keep going you are doing great"
else
  echo "You have entered the wrong amount of days"
fi
