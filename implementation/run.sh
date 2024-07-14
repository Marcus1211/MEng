#!/bin/bash

# Folder containing the files
FOLDER_PATH="./../preparation/data_graph"

# Check if the folder exists
if [ ! -d "$FOLDER_PATH" ]; then
  echo "Folder $FOLDER_PATH does not exist."
  exit 1
fi

# Loop through each file in the folder
for FILE in "$FOLDER_PATH"/*.json; do
  # Check if it is a file
  if [ -f "$FILE" ]; then
    # Run the Go program 20 times for each file
    for ((i=1; i<=20; i++)); do
      SOURCE_FILE=$(echo $FILE | cut -d'/' -f 5)
      echo $SOURCE_FILE
      echo "run"
      go run distributed-k-core.go $FILE > "${SOURCE_FILE}_output_run_${i}.txt"
    done
  fi
done

echo "All files have been processed."

