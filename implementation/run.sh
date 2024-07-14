#!/bin/bash

# Folder containing the files
FOLDER_PATH="./../preparation"

# Go program to run
GO_PROGRAM="./distributed-k-core.go"

# Check if the folder exists
if [ ! -d "$FOLDER_PATH" ]; then
  echo "Folder $FOLDER_PATH does not exist."
  exit 1
fi

# Loop through each file in the folder
for FILE in "$FOLDER_PATH"/*; do
  # Check if it is a file
  if [ -f "$FILE" ]; then
    # Run the Go program 20 times for each file
    ABSOLUTE_PATH=$(realpath "$FOLDER_PATH/$FILE")
    for ((i=1; i<=20; i++)); do
      OUTPUT_FILE="./$FILE_output_run_$i.txt"
      $GO_PROGRAM "$ABSOLUTE_PATH" > "$OUTPUT_FILE"
    done
  fi
done

echo "All files have been processed."

