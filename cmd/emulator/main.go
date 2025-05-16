package main

import (
	"fmt"
	"log"
	"net"
	"tcp_emulator/utils"
	"time"
)

const timeInterval = 1
const numberOfSources = 2
const numberOfInstancesOfEachSource = 5

func main() {
	log.Println("Starting TCP emulator")

	points, err := utils.ReadData()
	if err != nil {
		log.Fatal(err)
	}

	indexMap := utils.CreateIndexMap(numberOfSources, numberOfInstancesOfEachSource, points)
	idMap := utils.CreateIdsMap(numberOfSources, numberOfInstancesOfEachSource)

	var conn net.Conn

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	fmt.Println("Connected to server on port 8085")

	for {
		for i := 0; i < numberOfSources*numberOfInstancesOfEachSource; i++ {

			if indexMap[i] >= len(*points) {
				indexMap[i] = 0
			}
			data, err := utils.CreateDataToSend(points, idMap[i], indexMap[i], i%numberOfSources+1)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("-----------------------------------\n")
			conn, _ = connect()
			if conn != nil {

				fmt.Printf("Sending message: %s\n", string(*data))
				_, err = conn.Write(*data) // Send the message to the server
				if err != nil {
					fmt.Printf("conn.Write error: %s", string(*data))
				}
				fmt.Printf("Message successfully sent\n")
				err := conn.Close()
				if err != nil {
					fmt.Printf("conn.Close error: %s", err.Error())
					continue
				}
				fmt.Printf("Connection closed\n")
			}

			indexMap[i] += timeInterval
			time.Sleep(10 * time.Millisecond) //Optional delay. Prevents flooding the connection.
		}
		time.Sleep(timeInterval * time.Second)
	}
}

func connect() (net.Conn, error) {

	var conn net.Conn
	var err error

	for {
		conn, err = net.Dial("tcp", "localhost:8085")
		if err != nil {
			log.Printf("Error connecting to server: %v", err)
			break
		}
		if conn != nil {
			fmt.Println("Connected to server on port 8085")
			break
		}
		time.Sleep(time.Duration(timeInterval) * time.Second)
	}

	return conn, err
}
