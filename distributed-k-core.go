package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"time"
)

type Node struct {
	ID              string
	coreNumber      int
	storedNeighborK map[string]int
	status          string
	selfChan        chan sendMsg
	serverChan      chan string
	terminationChan chan bool
	neighbors       []chan sendMsg
}

type sendMsg struct {
	ID   string
	core int
}

func receive(node *Node) {
	//heartbeatInterval := make(chan bool)
	//go func() {
	//	for {
	//		time.Sleep(10 * time.Second)
	//		heartbeatInterval <- true
	//	}
	//}()

	go func() {
		for {
			select {
			case receivedMsg := <-node.selfChan:
				if receivedMsg.coreNumber != node.storedNeighborK[receivedMsg.ID] {
					node.storedNeighborK[receivedMsg.ID] = receivedMsg.coreNumber
					fmt.Println("Node ", node.ID, " with stored neighbor count", len(node.storedNeighborK), " and core number ", node.coreNumber, " received core number", receivedMsg.coreNumber, " from node ", receivedMsg.ID)
					node.status = "active"

					if len(node.storedNeighborK) >= node.coreNumber {
						fmt.Println("node ", node.ID, " with stored neighbor count", len(node.storedNeighborK), " and core number ", node.coreNumber, "received msg from ", receivedMsg.ID, " with core number ", receivedMsg.coreNumber, " is calling updateCore method")
						updateCore(node)
						sendHeartBeat(node)
						//fmt.Println("Node ", node.ID, " is sending hb because of processing selfchan update message")
					}
				} else {
					lenN := len(node.storedNeighborK)
					fmt.Println("Node ", node.ID, "with ", lenN, "neighbours Received duplicated core number", receivedMsg.coreNumber, " from node ", receivedMsg.ID, " as node has stored ", node.storedNeighborK[receivedMsg.ID])
				}
			//case <-heartbeatInterval:
			//	if node.status == "active" {
			//		sendHeartBeat(node)
			//		fmt.Println("Node ", node.ID, " is sending hb because of interval")
			//		//fmt.Println("Node ", node.ID, " is active, hb sent")
			//	}
			case <-node.terminationChan:
				fmt.Println("node ", node.ID, " has final core number ", node.coreNumber, " and status ", node.status)
				return
			default:
			}
		}
	}()
}

func send(node *Node, txt string) {
	msg := sendMsg{node.ID, node.coreNumber}
	for _, c := range node.neighbors {
		c <- msg
		fmt.Println("Node ", node.ID, " is sending core number", node.coreNumber, " to neighbour", c, " ", txt)
	}
}

func updateCore(node *Node) {
	origin_core := node.coreNumber
	for {
		count := 0
		for _, v := range node.storedNeighborK {
			if v >= node.coreNumber {
				count++
			}
		}
		if count >= node.coreNumber {
			if node.coreNumber != origin_core {
				go send(node, "updating")
			}
			node.status = "deactive"
			//fmt.Println("node ", node.ID, " with core number", node.coreNumber, " core is about to ", node.status)
			return
		} else {
			//fmt.Println("node ", node.ID, " with core number", node.coreNumber, " core is about to reduce 1 to ")
			node.coreNumber--
			//fmt.Println("node ", node.ID, " core is reducing 1 to ", node.coreNumber)
		}
	}
}

func sendHeartBeat(node *Node) {
	//fmt.Println("node ", node.ID, " is sending hearthbeat")
	node.serverChan <- node.status
}

func receiveHeartBeat(serverchan chan string, terminationChannels map[string]chan bool) {
	receiveInterval := make(chan bool)
	receivedHeartBeat := false
	receiveHB := 0
	go func() {
		for {
			time.Sleep(120 * time.Second)
			receiveInterval <- true
		}
	}()

	go func() {
		for {
			select {
			case <-serverchan:
				//fmt.Println(heartBeat)
				receivedHeartBeat = true
				receiveHB += 1
			case <-receiveInterval:
				if receivedHeartBeat {
					fmt.Println("There is heartbeat received in the past 40 seconds")
					receivedHeartBeat = false
				}
				if receiveHB > 0 && !receivedHeartBeat {
					fmt.Println("There's no heartbeat received in the past 40 seconds")
					for _, c := range terminationChannels {
						c <- true
					}
					return
				}
			}
		}
	}()
}

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
	runtime.GOMAXPROCS(10)
	//filename := "/Users/runzezhao/workspace/MEng_project/lasftm_asia/super-simple.json"
	//filename := "/Users/runzezhao/workspace/MEng_project/lasftm_asia/sample.json"
	//filename := "/Users/runzezhao/workspace/MEng_project/lasftm_asia/test-data-clean-sample.json"
	//filename := "/Users/runzezhao/workspace/MEng_project/lasftm_asia/3-core-simple.json"
	//filename := "/Users/runzezhao/workspace/MEng_project/lasftm_asia/lastfm_asia_features.json"
	//filename := "/Users/runzezhao/workspace/MEng_project/twitch/DE/musae_DE.json"
	//filename := "/Users/runzezhao/workspace/MEng_project/twitch/PTBR/musae_PTBR_features.json"
	filename := "/Users/runzezhao/workspace/MEng_project/MEng/git_web_ml/musae_git_features.json"
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

	channels := make(map[string]chan sendMsg, 500)
	terminationChannels := make(map[string]chan bool)

	for k := range res {
		channels[k] = make(chan sendMsg)
		terminationChannels[k] = make(chan bool, 1)
	}

	channelMap := make(map[string][]chan sendMsg)
	for k, v := range res {
		for _, c := range v {
			channelMap[k] = append(channelMap[k], channels[strconv.Itoa(c)])
		}
	}

	serverReceive := make(chan string, 10000)

	var nodes []Node
	for k, v := range res {
		var temp = map[string]int{}
		newNode := Node{
			ID:              k,
			core:            len(v),
			status:          "active",
			storedNeighborK: temp,
			selfChan:        channels[k],
			serverChan:      serverReceive,
			terminationChan: terminationChannels[k],
			neighbors:       channelMap[k],
		}
		nodes = append(nodes, newNode)
	}

	for i := 0; i < len(nodes); i++ {
		go receive(&nodes[i])
	}
	fmt.Println("All node started", len(nodes))

	go receiveHeartBeat(serverReceive, terminationChannels)
	fmt.Println("server started")

	for i := 0; i < len(nodes); i++ {
		go send(&nodes[i], "init")
	}
	fmt.Println("first send completed")

	for true {
	}
}
