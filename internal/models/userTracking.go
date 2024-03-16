package models

import (
	"time"
)

type OnlineUser struct {
	DisplayName string    `json:"display_name"`
	Coordinate  Position  `json:"coordinate"`
	Timestamp   time.Time `json:"timestamp"`
	Building    string    `json:"building"`
	Floor       int32     `json:"floor"`
	Label       string    `json:"label"`
}
