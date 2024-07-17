#!/bin/bash

# Folder containing the files
FOLDER_PATH="small_test_graph"

# Check if the folder exists
if [ ! -d "$FOLDER_PATH" ]; then
  echo "Folder $FOLDER_PATH does not exist."
  exit 1
fi

# Loop through each file in the folder
for FILE in "$FOLDER_PATH"/*.json; do
  # Check if it is a file
  if [ -f "$FILE" ]; then
      echo "run"
      SOURCE_FILE=$(echo $FILE | cut -d'/' -f 2 | cut -d'.' -f 1)
      go run ../implementation/distributed-k-core.go $FILE > "test_small_${SOURCE_FILE}_output.txt"

      go run bz-origin.go $FILE > "test_small_${SOURCE_FILE}_bz_output.txt"
  fi
done

echo "All files have been processed."

