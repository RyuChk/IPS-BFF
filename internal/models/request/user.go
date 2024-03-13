package request

import (
	"github.com/ZecretBone/ips-bff/internal/models"
)

type GetSingleCoordinateRequest struct {
	Signals  []models.RSSI `json:"signals"`
	Building string        `json:"building"`
}
