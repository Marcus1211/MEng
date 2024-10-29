package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

type sendMsg struct {
	ID         string
	coreNumber int
}

func node(id string, selfChan chan sendMsg, neighbourChan []chan sendMsg, heartbeat chan bool, selfTerminationChan chan bool, wg *sync.WaitGroup, enableLog string) {
	//fmt.Println("Node ", id, " is starting")
	defer wg.Done()
	coreNumber := len(neighbourChan)
	storedNeighbourK := map[string]int{}
	active := true
	//first send
	send(id, coreNumber, neighbourChan, enableLog)

	for {
		// node receives message from its own channel
		select {
		case receivedMsg := <-selfChan:
			if receivedMsg.coreNumber < storedNeighbourK[receivedMsg.ID] || storedNeighbourK[receivedMsg.ID] == 0 {
				//node stores neighbours core number
				storedNeighbourK[receivedMsg.ID] = receivedMsg.coreNumber
				if len(storedNeighbourK) >= coreNumber {
					//calculate k-core
					active = true
					if active {
						//log.Println("Node ", id, " status is active ", time.Now().Format("2006-01-02 15:04:05"))
					}
					//k-core calculation
					new_core := updateCore(coreNumber, storedNeighbourK)
					if new_core < coreNumber {
						coreNumber = new_core
						heartbeat <- true
						send(id, coreNumber, neighbourChan, enableLog)
						active = false
					}
				}
			}
		case <-selfTerminationChan:
			active = false
			log.Println("Node ", id, " has ", len(neighbourChan), " neighbours and final core number of ", coreNumber)
			return
			//default:
		}
	}
}

func updateCore(coreNumber int, storedNeighbourK map[string]int) (k int) {
	for {
		count := 0
		for _, v := range storedNeighbourK {
			if v >= coreNumber {
				count++
			}
		}
		if count >= coreNumber {
			return coreNumber
		} else {
			coreNumber--
		}
	}
}

func send(id string, coreNumber int, neighbourChan []chan sendMsg, enableLog string) {
	msg := sendMsg{id, coreNumber}
	for _, c := range neighbourChan {
		//send messages to all neighbour nodes
		c <- msg
	}
	if enableLog == "enableLog" {
		log.Println("Node ", id, " sent ", len(neighbourChan), " messages ")
	}

}

func watchdog(heartbeat chan bool, terminationChan map[string]chan bool) {
	done := false
	for !done {
		select {
		case <-heartbeat:
			done = false
		//termination happens when no heartbeat receives in the past 300 seconds
		case <-time.After(300 * time.Second):
			fmt.Println("No heartbeat has received in the past 300 seconds, start terminating all nodes")
			for _, c := range terminationChan {
				c <- true
			}
			done = true
			//default:
		}
	}
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
	filename := os.Args[1]
	fileContent, err := os.Open(filename)
	if err != nil {
		fmt.Println("Can't open file")
		return
	}

	enableLog := os.Args[2]

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	byteResult, _ := io.ReadAll(fileContent)
	var res map[string][]int
	json.Unmarshal([]byte(byteResult), &res)

	res = dataCleanse(res)
	res = dataCleanse2(res)

	allNodeChan := make(map[string]chan sendMsg)
	terminationChan := make(map[string]chan bool)
	//neighbourChan := make(map[string][]chan sendMsg)

	for k, v := range res {
		//create synchronized channel for each node
		bufferSize := 0
		for v := range v {
			bufferSize = bufferSize + len(res[strconv.Itoa(v)])
		}
		bufferSize = bufferSize + len(v)
		allNodeChan[k] = make(chan sendMsg, bufferSize)
		terminationChan[k] = make(chan bool)
	}

	neighbourChan := make(map[string][]chan sendMsg)

	for k, v := range res {
		for _, c := range v {
			neighbourChan[k] = append(neighbourChan[k], allNodeChan[strconv.Itoa(c)])
		}
	}

	heartbeat := make(chan bool, len(res))

	go watchdog(heartbeat, terminationChan)
	var wg sync.WaitGroup
	wg.Add(len(res))
	for k, _ := range res {
		go node(k, allNodeChan[k], neighbourChan[k], heartbeat, terminationChan[k], &wg, enableLog)
	}
	wg.Wait()
}
