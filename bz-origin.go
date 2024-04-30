package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Graph map[string][]int

func dataCleanse(data map[string][]int) map[string][]int {
	graph := data
	for k, v := range graph {
		for _, n := range v {
			_, ok := graph[strconv.Itoa(n)]
			//fmt.Println("node ", n, " exists? ", ok)
			if ok {
				contain := false
				for _, value := range graph[strconv.Itoa(n)] {
					//fmt.Println("node ", n, " has neighbour ", value)
					if strconv.Itoa(value) == k {
						contain = true
					}
				}
				//fmt.Println("node ", n, " has neighbour ", k, " is ", contain)
				if !contain {
					a, _ := strconv.Atoi(k)
					graph[strconv.Itoa(n)] = append(graph[strconv.Itoa(n)], a)
				}
			}
		}
	}
	return graph
}

func dataCleanse2(data map[string][]int) map[string][]int {
	graph := data
	for k, v := range graph {
		var temp []int
		for _, n := range v {
			_, ok := graph[strconv.Itoa(n)]
			//fmt.Println("node ", n, " exists? ", ok)
			if ok {
				temp = append(temp, n)
			}
		}
		graph[k] = temp
	}
	return graph
}

func main() {
        filename := os.Args[1]
	//filename := "/Users/runzezhao/workspace/MEng_project/twitch/PTBR/musae_PTBR_features.json"
	//filename := "/Users/runzezhao/workspace/MEng_project/lasftm_asia/lastfm_asia_features.json"
	//filename := "/Users/runzezhao/workspace/MEng_project/twitch/DE/musae_DE.json"
	//filename := "/Users/runzezhao/workspace/MEng_project/MEng/git_web_ml/musae_git_features.json"
	fileContent, err := os.Open(filename)
	if err != nil {
		fmt.Println("Can't open file")
		return
	}

	byteResult, _ := io.ReadAll(fileContent)
	var res map[string][]int
	json.Unmarshal([]byte(byteResult), &res)

	res = dataCleanse(res)
	res = dataCleanse2(res)

	coreNumbers := calculateCoreNumbers(res)
	temp := 0
	for _, b := range coreNumbers {
		//fmt.Println("Node ", a, " has core number ", b)
		temp = temp + b
	}
	fmt.Print(temp)

}

func calculateCoreNumbers(graph Graph) map[string]int {
	degrees := make(map[string]int)
	for node := range graph {
		//fmt.Println(node)
		degrees[node] = len(graph[node])
	}
	//fmt.Println(degrees)
	coreNumbers := make(map[string]int)

	k := 0
	for len(degrees) > 0 {
		for node, degree := range degrees {
			if degree < k {
				delete(degrees, node)
				for _, neighbor := range graph[node] {
					if degrees[strconv.Itoa(neighbor)] > 0 {
						degrees[strconv.Itoa(neighbor)]--
					}
				}
			}
		}
		fmt.Println(len(degrees), " nodes have degree ", k)
		for node := range degrees {
			coreNumbers[node] = k
		}
		k++
	}

	return coreNumbers
}
