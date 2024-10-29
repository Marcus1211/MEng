#!/bin/bash

# Folder containing the files
FOLDER_PATH="./../preparation/data_graph"

# Check if the folder exists
if [ ! -d "$FOLDER_PATH" ]; then
  echo "Folder $FOLDER_PATH does not exist."
  exit 1
fi

echo "Note: This program runs against all test data graphs"
echo "      This program runs each test data graph 20 times"
echo "      To check the progress of each run, you can check the log output txt file"
echo "      For large graphs contains over 200000 vertices, it may take hours to run the program"
echo "      It is not recommened to run the entire experiment in the local laptop"

# Loop through each file in the folder
for FILE in "$FOLDER_PATH"/*.json; do
  # Check if it is a file
  if [ -f "$FILE" ]; then
    # Run the Go program 20 times for each file
    for ((i=1; i<=20; i++)); do
      SOURCE_FILE=$(echo $FILE | cut -d'/' -f 5)
      echo $SOURCE_FILE
      echo "run"
      if [ "$i" -gt 10 ]; then
      	start_time=$(date +%s)
      	go run distributed-k-core.go $FILE disableLog > "${SOURCE_FILE}_output_run_${i}.txt" 2>&1
      	end_time=$(date +%s)
      	elapsed_time=$(($end_time - $start_time))
      	echo "The ${SOURCE_FILE}_output_run_${i}.txt took $elapsed_time seconds to complete." >> all_time.log
      else
	go run distributed-k-core.go $FILE enableLog > "${SOURCE_FILE}_output_run_${i}.txt" 2>&1
      fi
    done
  fi
done

echo "All files have been processed."

