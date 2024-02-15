package response

import (
	"time"

	"github.com/ZecretBone/ips-bff/internal/models"
)

type GetMapFloorListResponse struct {
	Floors  []models.Map `json:"floors"`
	IsAdmin bool         `json:"is_admin"`
}

type GetMapURLResponse struct {
	Detail    models.Map `json:"detail"`
	URL       string     `json:"url"`
	UpdatedAt time.Time  `json:"updated_at"`
}
