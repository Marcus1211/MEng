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
      SOURCE_FILE=$(echo $FILE | cut -d'/' -f 5)
      echo "run"
      go run bz-origin.go $FILE > "${SOURCE_FILE}_output_bz.txt"
  fi
done

echo "All files have been processed."

