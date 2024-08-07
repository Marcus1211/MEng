# Test
This folder contains all programs and scripts to test the correctness of the distributed k-core decomposition algorithm.
The testing is using black-box testing methodology.
The k-core of each node in the test graph will be calculated by Bienstock–Zuckerberg algorithm (BZ algorithm) to check the correctness.
We also added synthetic test graphs to test the edge cases of different graphs and manually calculate the k-core of each node as reference.
## Prerequisite
- Golang version: 1.22
- Steps in the `preparation` folder need to be completed before conducting this experiment.
## What's in this repository

### - bz-origin.go
This go program implements the BZ algorithm to calculate the k-core of each node.
The program takes the `json` formatted graph data from the `preparation` folder as input and saves the core number of each node into a text file.
### - small_test_graph
This folder contains all the synthetic graphs to test the edge cases of the different graphs
Cases:
- Circle graph: All nodes in this graph form a circle.
- Linear graph: All nodes in this graph form a linear line.
- Missing edge: Some of the edges are missing in this graph. e.g node A knows neighbour B but Node B does not know neighbour A. This can be used for testing directed graphs.
- Missing nodes: Some of the nodes are missing in this graph. e.g Node A knows neighbour B but Node B is missing in this graph data.
- Small graph: This is the normal graph that can be used for regular testing/troubleshoot purposes.

### - small_test_graph_result
All k-core of small test graphs are saved into text files stored in this folder. This is used to check the correctness of BZ algorithm and Distributed k-core decomposition algorithm
### - test-small-graph.sh
This script automatically runs the `bz-origin.go` program and `distributed-k-core.go` program against all small test graphs. The result of each run is saved into a text file for future analysis. Execute following command to run the programs

### - test-snap-graph.sh
This script automatically runs the `bz-origin.go` program against all downloaded graphs in the `preparation` folder. The result for each graph is saved into a text file for future analysis.

## How to use this repository
1. Navigate to the same directory where the `test-small-graph.sh` is. Execute following command
    ```shell
    ./test-small-graph.sh
    ```
2. Once `preparation` steps are completed, execute following command
    ```shell
    ./test-snap-graph.sh
    ```
2. The results of each simulation will be saved into different log files for future analysis.
