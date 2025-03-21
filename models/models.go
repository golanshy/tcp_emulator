package models

import "time"

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type DataPoint struct {
	Id        string     `json:"id"`
	Type      int        `json:"type"`
	Location  LatLng     `json:"location"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
