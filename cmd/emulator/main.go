package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"tcp_emulator/config"
	"tcp_emulator/utils"
	"time"
)

func main() {
	log.Println("Starting TCP emulator")
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", cfg)
	}

	points, err := utils.ReadData()
	if err != nil {
		log.Fatal(err)
	}

	indexMap := utils.CreateIndexMap(cfg.Emulator.NumberOfSources, cfg.Emulator.NumberOfInstancesOfEachSource, points)
	idMap := utils.CreateIdsMap(cfg.Emulator.NumberOfSources, cfg.Emulator.NumberOfInstancesOfEachSource)

	var conn net.Conn

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	fmt.Println("Connected to server on port 8085")

	for {
		for i := 0; i < cfg.Emulator.NumberOfSources*cfg.Emulator.NumberOfInstancesOfEachSource; i++ {

			if indexMap[i] >= len(*points) {
				indexMap[i] = 0
			}
			data, err := utils.CreateDataToSend(points, idMap[i], indexMap[i], i%cfg.Emulator.NumberOfSources+1)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("-----------------------------------\n")
			conn, _ = connect(cfg)
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

			indexMap[i] += cfg.Emulator.TimeInterval
			time.Sleep(10 * time.Millisecond) //Optional delay. Prevents flooding the connection.
		}
		time.Sleep(time.Duration(cfg.Emulator.TimeInterval) * time.Second)
	}
}

func connect(cfg *config.Config) (net.Conn, error) {

	var conn net.Conn
	var err error

	for {
		conn, err = net.Dial(cfg.TCP.Protocol, fmt.Sprintf("%s:%s", cfg.TCP.Address, cfg.TCP.Port))
		if err != nil {
			log.Printf("Error connecting to server: %v", err)
			break
		}
		if conn != nil {
			fmt.Println("Connected to server on port 8085")
			break
		}
		time.Sleep(time.Duration(cfg.Emulator.TimeInterval) * time.Second)
	}

	return conn, err
}
