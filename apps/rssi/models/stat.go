package models

import "time"

type RSSI struct {
	SSID           string      `json:"ssid"`             //APs SSID
	MacAddress     string      `json:"mac_address"`      // MAC Address of APs
	Strength       []float64   `json:"strength"`         //RSSI signal strength in Dbm.
	AverageStrenth float32     `json:"average_strength"` //Average of all rssi strength
	CreatedAt      []time.Time `json:"created_at"`       //RSSI signal capture time
}

type Position struct {
	X float64 `json:"x"` // X position in meter
	Y float64 `json:"y"` // Y position in meter
	Z float64 `json:"z"` // Z position in floor
}

type DeviceInfo struct {
	DeviceID string `json:"device_id"` //Device ID
	Models   string `json:"models"`    //Device models eg.Samsung
}

type AccessPoint struct {
	SSID       string `json:"ssid"`        //APs SSID
	MacAddress string `json:"mac_address"` // MAC Address of APs
}
