package models

import "time"

type DataPoint struct {
	DeviceId     string        `json:"device_id"`
	Type         int           `json:"type"`
	Location     LatLng        `json:"location"`
	Measurements []Measurement `json:"measurements"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Measurement struct {
	Index    int     `json:"index"`
	SourceID string  `json:"source_id"`
	Type     string  `json:"type"`
	Value    float64 `json:"value"`
}
