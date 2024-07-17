# Preparation
This folder contains all programs and scripts to download and process test graph data.
Each graph data after the process can be directly used by the distributed k-core decomposition program.
## Prerequisite
- Golang version: 1.22

## What's in this repository

### - data_graph
All test data graphs will be downloaded to this folder.
This folder also contains the go program `txt-go-json.go`
This go program converts text based data graph files into json format.
### - download.sh
This shell scripts automatically downloads all selected test graph data from
https://snap.stanford.edu/data/
If the test graph data is compressed, this script decompresses the graph data.
The script also removes any unwanted lines in the graph data files and makes it ready to be converted in to json format
### - convert-txt-to-json.sh
This script automatically converts any text based graph data files into json format.
## How to use this repository
1. Make sure `data_graph` folder only contains the go program `txt-to-json.go`
2. Navigate to the same directory where the `download.sh` is. Execute following command
    ```shell
    ./download.sh
    ```
3. Once download completes, execute following command
    ```shell
    ./convert-txt-to-json.sh
    ```
4. All graph data files should be converted into `json` format and ready to be used by the simulation go program in the `implementation` folder



