package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"tcp_emulator/utils"
	"time"
)

const timeInterval = 5
const numberOfSources = 10
const numberOfInstancesOfEachSource = 500

func main() {
	log.Println("Starting TCP Hello World emulator")

	points, err := utils.ReadData()
	if err != nil {
		log.Fatal(err)
	}

	indexMap := utils.CreateIndexMap(numberOfSources, numberOfInstancesOfEachSource, points)
	idMap := utils.CreateIdsMap(numberOfSources, numberOfInstancesOfEachSource)

	var conn net.Conn
	conn, err = connect()

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
				if strings.Contains(err.Error(), "broken pipe") {
					conn, err = connect()
					if err != nil {
						log.Fatalf("Error connecting to server: %v", err)
						return
					}

					// Resend the message post re-connection
					_, err = conn.Write(byteArray)
					if err != nil {
						fmt.Printf("conn.Write success: %s", string(byteArray))
					}
				}
				continue
			}

			fmt.Printf("conn.Write success: %s", string(byteArray))

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
		fmt.Println("Attempting connection to server on port 8085")

		conn, err = net.Dial("tcp", "localhost:8085")
		if err != nil {
			log.Printf("Error connecting to server: %v", err)
		}
		if conn != nil {
			break
		}
		time.Sleep(time.Duration(timeInterval) * time.Second)
	}

	return conn, err
}
