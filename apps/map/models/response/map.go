package response

import (
	"time"

	mapservice "github.com/ZecretBone/ips-bff/apps/map"
)

type GetMapFloorListResponse struct {
	Floors  []mapservice.Map `json:"floors"`
	IsAdmin bool             `json:"is_admin"`
}

type GetMapURLResponse struct {
	Detail    mapservice.Map `json:"detail"`
	URL       string         `json:"url"`
	UpdatedAt time.Time      `json:"updated_at"`
}
