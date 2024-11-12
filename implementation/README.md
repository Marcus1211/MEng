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
### - python_script
This folder contains all scripts to automatically generate diagrams for this experiment. Follow the README instructions inside the folder after previous steps are completed.

## How to use this repository
1. Navigate to the same directory where the `run.sh` is. Execute following command
    ```shell
    ./run.sh -enableLog 10 -disableLog 10
    ```
   The helper of the `run.sh` has more information of how to run it

   ```shell
   ./run.sh -h
   Usage: ./run.sh [options]

   Options:
     -enableLog   Specify how many runs with log enabled
     -disableLog  Specify how many runs with log disabled
     -h, --help           Show this help message and exit.

   Example:
     ./run.sh -enableLog 10 -disableLog 10
   ```
   The logs output will affect the overall runtime of the algorithm, that's why to have separate runs with log enabled and disabled
   The runs with `-enableLog` is to track the overall runtime
   The runs with `-disableLog` is to track the overall status of the system, e.g. what's the status of each node
   It's recommend to have 10 runs of each, the experiment in the report is based on 10 runs of each

2. The results of each simulation will be saved into different log files for future analysis.

3. The `run.sh` will also triggers the `generate_diagrams.sh` script inside the python_script folder to genreate all diagrams. In case you need to run `generate_diagrams.sh` manually
    ```shell
    ./generate_diagrams.sh -enableLog 10 -disableLog 10
    ```
   The numbers of `enabledLog` and `disableLog` options should match with the input for `run.sh`


## Note
By default, the distributed-k-core.go should use all computing cores available. The program prints out the number of cores used.
In case of different local environment setups, Go may only use single computing core. 
Following code can be added to the main function to specify how many cores to use.
```go
runtime.GOMAXPROCS(runtime.NumCPU())
```
