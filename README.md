# Distributed k-Core Algorithm
This is the project to simulate distributed k-core decomposition algorithm using Golang.
This repository contains all necessary programs and scripts to reproduce the experiment.
## Prerequisite
- Golang version: 1.22

## Folders in this repository

### - implementation
This folder contains the go program to simulate distributed k-core decomposition algorithm and 
the automation shell script to loop through each test graph data.
### - preparation
This folder contains the shell script to download and process each test graph data.
The processed test graph data will be in `json` format and can be direclty used by the simulation go program.
### - test
The test folder contains the sample test graph data to test the edge cases of the simulation go program.
Each test cases follows the blackbox testing methodology to test if the simulation go program produces correct resutls.
## How to use this repository
1. Follow the instructions in the `preparation` folder to download and process the test graph data.
2. Follow the instructions in the `implmentation` folder to run the simulation go program against test graph data.
3. To test the correctness of the go program, follow the instructions in the `test` folder.

## License

[MIT](https://choosealicense.com/licenses/mit/)
