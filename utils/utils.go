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

	initialMap := make([]string, 0)
	initialMap = append(initialMap, "ccf32c20-8475-4c71-9e63-80b15a506713")
	initialMap = append(initialMap, "3b32105c-8252-4f3b-938a-04c828edc6e1")
	initialMap = append(initialMap, "8e655452-90a3-4964-9d4e-633a115bef8d")
	initialMap = append(initialMap, "7ba4700d-2df4-4a37-bd58-747cee104ff1")
	initialMap = append(initialMap, "96337bad-ad44-4461-aa5c-023cb0d080ae")
	initialMap = append(initialMap, "a7634b7e-1b51-4bb2-b909-9c8d65ef1d90")
	initialMap = append(initialMap, "534ba91b-1c81-4508-a7e8-5e9944f87d24")
	initialMap = append(initialMap, "2e812d48-2903-4c38-8893-383aeaccc457")
	initialMap = append(initialMap, "5eda4ad7-0169-47af-bdd9-e30bcf263d14")
	initialMap = append(initialMap, "a8ef2449-cce2-4140-8cc5-baeb73aade1e")
	initialMap = append(initialMap, "03bc3576-2fec-4b54-aa81-f6a15264f979")
	initialMap = append(initialMap, "0149c5eb-210f-4af9-bf3f-d7cbde27a115")
	initialMap = append(initialMap, "feddc523-25ca-4ded-84ae-18c65cf10f63")
	initialMap = append(initialMap, "73c620c2-01c1-440d-9f37-13e34c4a3c34")
	initialMap = append(initialMap, "ad01409d-9ef7-4b73-9212-854c0ff0363e")
	initialMap = append(initialMap, "8ba69b9d-d600-40bd-8396-2e235b05238a")
	initialMap = append(initialMap, "5ad43b2d-5b5d-4150-bd3f-e463dbf4a13d")
	initialMap = append(initialMap, "491d6a70-8948-4da8-b21e-a136f8958b54")
	initialMap = append(initialMap, "4da07d03-0c96-4d1b-9aa4-7a33077d5d58")
	initialMap = append(initialMap, "4fcc8959-3370-427a-af64-ea5205b5c988")

	idMap := map[int]string{}
	for i := 0; i < numberOfSources*numberOfInstancesOfEachSource; i++ {
		if i < 20 {
			idMap[i] = initialMap[i]
		} else {
			id, _ := uuid.NewRandom()
			idMap[i] = id.String()
		}
	}
	return idMap
}

func CreateDataToSend(points *[]models.LatLng, dataPointId string, index int, t int) (*[]byte, error) {
	point := (*points)[index]

	now := time.Now()
	dataPoint := &models.DataPoint{
		DeviceId:  dataPointId,
		Type:      t,
		Location:  point,
		CreatedAt: now,
		UpdatedAt: now,
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
