package models

import (
	"time"
)

type OnlineUser struct {
	DisplayName string    `json:"display_name"`
	Coordinate  Position  `json:"coordinate"`
	Timestamp   time.Time `json:"timestamp"`
}
