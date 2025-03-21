package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"tcp_emulator/utils"
	"time"
)

const timeInterval = 5
const numberOfSources = 4
const numberOfInstancesOfEachSource = 2

func main() {
	log.Println("Starting TCP Hello World emulator")

	points, err := utils.ReadData()
	if err != nil {
		log.Fatal(err)
	}

	indexMap := utils.CreateIndexMap(numberOfSources, numberOfInstancesOfEachSource, points)
	idMap := utils.CreateIdsMap(numberOfSources, numberOfInstancesOfEachSource)

	conn, err := net.Dial("tcp", "localhost:8085") // Connect to the server
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	fmt.Println("Connected to server on port 8080")

	for {
		for i := 0; i < numberOfSources*numberOfInstancesOfEachSource; i++ {

			if indexMap[i] >= len(*points) {
				indexMap[i] = 0
			}
			data, err := utils.CreateDataToSend(points, idMap[i], indexMap[i], i%numberOfSources)
			if err != nil {
				log.Fatal(err)
			}
			str := fmt.Sprintf("%s\n", string(*data))
			fmt.Printf("Sending message: %s", str)

			byteArray := append(*data, '\n')
			_, err = conn.Write(byteArray) // Send the message to the server
			if err != nil {
				log.Printf("Error sending message: %v", err)
				continue
			}

			//// Optionally, read the server's response
			response, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Printf("Error reading response: %v", err)
				continue
			}
			fmt.Printf("Server response: %s", response)

			indexMap[i] += timeInterval
			time.Sleep(10 * time.Millisecond) //Optional delay. Prevents flooding the connection.
		}
		time.Sleep(timeInterval * time.Second)
	}
}
