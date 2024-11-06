# Implementation
This folder contains all scripts to process the data from the experiment.
Each script will generate diagrams used in the report
## Prerequisite
- python 3
- Required python libraries
-- matplotlib
-- re
-- numpy
-- scipy
-- collections
-- datetime
If there are any libraries missing, you can install it via `Pip` or `homebrew`
## What's in this folder
- barChart_message_auto.py
    This script generates a bar Chart diagram of all messages passed during k-core decomposition for each test graph
- barChart_time_auto.py
    This script generates a bar Chart diagram of time comsumed during k-core decomposition for each test graph
- core_dist_auto.py
    This script generates the number of nodes distribution of different core numbers for each graph
- message_dist_auto.py
    This script generates the number of passed message over time for each graph during distributed k-core decomposition
- node_dist_auto.py
    This script generates the number of active nodes over time for each graph during distributed k-core decomposition
### - generate_diagrams.sh
This shell script automatically runs all python scripts to generate all diagrams.
There is no need to manually execute individual python script.

## How to use this folder
1. Navigate to the same directory where the `generate_diagrams.sh` is. Execute following command
    ```shell
    ./generate_diagrams.sh -enableLog 10 -disableLog 10
    ```
   The option value of `enableLog` and `disableLog` should be the same as value used for running `run.sh`
2. The results of each python script will be saved as `svg` image files.
