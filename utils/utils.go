package utils

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"os"
	"tcp_emulator/models"
	"time"
)

func CreateIndexMap(numberOfSources int, numberOfInstancesOfEachSource int, points *[]models.LatLng) map[int]int {
	indexMap := map[int]int{}
	for i := 0; i < numberOfSources*numberOfInstancesOfEachSource; i++ {
		// Generate a random index between 1 and n
		// Seed the random number generator with the current time
		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		randomIndex := random.Intn(len(*points))
		indexMap[i] = randomIndex
	}
	return indexMap
}

func CreateIdsMap(numberOfSources int, numberOfInstancesOfEachSource int) map[int]string {
	idMap := map[int]string{}
	for i := 0; i < numberOfSources*numberOfInstancesOfEachSource; i++ {
		id, _ := uuid.NewRandom()
		idMap[i] = id.String()
	}
	return idMap
}

func CreateDataToSend(points *[]models.LatLng, dataPointId string, index int, t int) (*[]byte, error) {
	point := (*points)[index]

	now := time.Now()
	dataPoint := &models.DataPoint{
		Id:        dataPointId,
		Type:      t,
		Location:  point,
		CreatedAt: &now,
		UpdatedAt: nil,
	}
	jsonValue, err := json.Marshal(dataPoint)
	if err != nil {
		return nil, err
	}

	return &jsonValue, nil
}

func ReadData() (*[]models.LatLng, error) {
	// Specify the path to your JSON file
	filePath := "points.json"

	// Read the entire file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
		return nil, err
	}

	// Create a slice of LatLng structs to unmarshal the JSON array into
	var points []models.LatLng

	// Unmarshal the JSON data into the slice
	err = json.Unmarshal(data, &points)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
		return nil, err
	}

	// Iterate through the list of lat,lng points
	fmt.Println("List of LatLng points:")
	for i, point := range points {
		fmt.Printf("[%d] Lat: %f, Lng: %f\n", i+1, point.Lat, point.Lng)
	}
	return &points, nil
}
