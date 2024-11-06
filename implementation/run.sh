#!/bin/bash

# Function to display help
show_help() {
    echo "Usage: $0 [options]"
    echo
    echo "Options:"
    echo "  -enableLog   Specify how many runs with log enabled"
    echo "  -disableLog  Specify how many runs with log disabled "
    echo "  -h, --help           Show this help message and exit."
    echo
    echo "Example:"
    echo "  $0 -enableLog 10 -disableLog 10"
}

# Default values
enableLog=""
disableLog=""

# Parse command-line options
while [[ $# -gt 0 ]]; do
    case $1 in
        -enableLog)
            enableLog="$2"
            shift 2
            ;;
        -disableLog)
            disableLog="$2"
            shift 2
            ;;
        -h|--help)
            show_help
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            show_help
            exit 1
            ;;
    esac
done

# Check if required options are provided
if [[ -z $enableLog || -z $disableLog ]]; then
    echo "Error: Both name and age must be provided."
    show_help
    exit 1
fi

# Folder containing the files
FOLDER_PATH="./../preparation/data_graph"

# Check if the folder exists
if [ ! -d "$FOLDER_PATH" ]; then
  echo "Folder $FOLDER_PATH does not exist."
  exit 1
fi

echo "Note: This program runs against all test data graphs"
echo "      Based on the input, the program will run multiple times against all data graphs to obtain surfficient data"
echo "      To check the progress of each run, you can check the log output txt file"
echo "      For large graphs contains over 200000 vertices, it may take hours to run the program"
echo "      It is not recommened to run the entire experiment in the local laptop"

# Loop through each file in the folder
for FILE in "$FOLDER_PATH"/*.json; do
  # Check if it is a file
  if [ -f "$FILE" ]; then
    total_run=$((enableLog + disableLog))
    echo "$total_run"
    for ((i=1; i<=$total_run; i++)); do
      SOURCE_FILE=$(echo $FILE | cut -d'/' -f 5)
      echo $SOURCE_FILE
      echo "run"
      if [ "$i" -gt $enableLog ]; then
      	start_time=$(date +%s)
      	go run distributed-k-core.go $FILE disableLog > "${SOURCE_FILE}_output_run_${i}_logDisabled.txt" 2>&1
      	end_time=$(date +%s)
      	elapsed_time=$(($end_time - $start_time))
      	echo "The ${SOURCE_FILE}_output_run_${i}.txt took $elapsed_time seconds to complete." >> all_time.log
      else
	go run distributed-k-core.go $FILE enableLog > "${SOURCE_FILE}_output_run_${i}_logEnabled.txt" 2>&1
      fi
    done
  fi
done

echo "All files have been processed."

### Process data###
echo "Starting generating diagrams"
cd python_script
./generate_diagrams.sh -enableLog $enableLog -disableLog $disableLog
echo "All diagrams have been generated"
