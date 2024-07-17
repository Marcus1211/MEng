# Implementation
This folder contains all programs and scripts to simulate the distributed k-core decomposition algorithm.
## Prerequisite
- Golang version: 1.22
- Steps in the `preparation` folder need to be completed before conducting this experiment.
## What's in this repository

### - distributed-k-core.go
This go program simulates the distributed k-core decomposition algorithm. Details of implementation can be found in the report.
This go program only takes `json` formatted graph data as input. Only one graph data can be passed as input.
### - run.sh
This shell script automatically runs the `distributed-k-core.go` with each data graph as input.
It loops through all processed graph data. For each graph data, this script executes the simulation 20 times in order to obtain sufficient results.

## How to use this repository
1. Navigate to the same directory where the `run.sh` is. Execute following command
    ```shell
    ./run.sh
    ```
2. The results of each simulation will be saved into different log files for future analysis.



