package request

import (
	"time"

	"github.com/ZecretBone/ips-bff/internal/constants"
	"github.com/ZecretBone/ips-bff/internal/models"
)

type StatCollectionRequest struct {
	Signals             []models.RSSI                  `json:"signals"`
	Position            models.Position                `json:"position"`
	Duration            int                            `json:"duration"` //capture duration in second
	PollingRate         int                            `json:"polling_rate"`
	StatCollectionStage constants.DataCollectionStages `json:"stat_collection_stage"`
	Direction           constants.Direction            `json:"selected_direction"`
	StartedAt           time.Time                      `json:"started_at"`
	EndedAt             time.Time                      `json:"ended_at"`
	CreatedAt           time.Time                      `json:"created_at"`
}
