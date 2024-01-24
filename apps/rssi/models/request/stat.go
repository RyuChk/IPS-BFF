package request

import (
	"time"

	"github.com/ZecretBone/ips-bff/apps/constants"
	"github.com/ZecretBone/ips-bff/apps/rssi/models"
)

type StatCollectionRequest struct {
	Signals             []models.RSSI                  `json:"signals"`
	Position            models.Position                `json:"position"`
	Duration            int                            `json:"duration"` //capture duration in second
	PollingRate         int                            `json:"polling_rate"`
	StatCollectionStage constants.DataCollectionStages `json:"stat_collection_stage"`
	StartedAt           time.Time                      `json:"started_at"`
	EndedAt             time.Time                      `json:"ended_at"`
	CreatedAt           time.Time                      `json:"created_at"`
}
