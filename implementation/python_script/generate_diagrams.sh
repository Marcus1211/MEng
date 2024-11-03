#!/bin/bash

# Execute Python scripts to generate report diagrams

echo "Running barChart_message_auto.py"
python3 barChart_message_auto.py

echo "Running barChart_time_auto.py"
python3 barChart_time_auto.py

echo "Running core_dist_auto.py"
python3 core_dist_auto.py

echo "Running message_dist_auto.py"
python3 message_dist_auto.py

echo "Running node_dist_auto.py"
python3 node_dist_auto.py

echo "All scripts executed successfully!"
