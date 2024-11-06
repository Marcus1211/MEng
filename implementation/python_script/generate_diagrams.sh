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

# Execute Python scripts to generate report diagrams

echo "Running barChart_message_auto.py"
python3 barChart_message_auto.py $enableLog

echo "Running barChart_time_auto.py"
python3 barChart_time_auto.py

echo "Running core_dist_auto.py"
python3 core_dist_auto.py

echo "Running message_dist_auto.py"
python3 message_dist_auto.py $enableLog

echo "Running node_dist_auto.py"
python3 node_dist_auto.py $enableLog

echo "All scripts executed successfully!"
