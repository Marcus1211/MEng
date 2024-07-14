package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    // Open the input file
    fileName := os.Args[1]
    inputFile, err := os.Open(fileName)
    if err != nil {
        fmt.Println("Error opening input file:", err)
        return
    }
    defer inputFile.Close()

    // Create the output file
    outputFile, err := os.Create(fileName + ".json")
    if err != nil {
        fmt.Println("Error creating output file:", err)
        return
    }
    defer outputFile.Close()

    // Initialize a map to store the data
    data := make(map[int][]int)

    // Read the input file line by line
    scanner := bufio.NewScanner(inputFile)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Fields(line)
        if len(parts) != 2 {
            fmt.Println("Invalid input line:", line)
            continue
        }

        // Parse integers
        num1, err := strconv.Atoi(parts[0])
        if err != nil {
            fmt.Println("Error parsing integer:", err)
            continue
        }
        num2, err := strconv.Atoi(parts[1])
        if err != nil {
            fmt.Println("Error parsing integer:", err)
            continue
        }

        // Add to the data map
        data[num1] = append(data[num1], num2)
    }

    // Convert data to JSON
    jsonData, err := json.Marshal(data)
    if err != nil {
        fmt.Println("Error converting to JSON:", err)
        return
    }

    // Write JSON data to the output file
    _, err = outputFile.Write(jsonData)
    if err != nil {
        fmt.Println("Error writing JSON to output file:", err)
        return
    }

    fmt.Println("Output has been written to output.json")
}

