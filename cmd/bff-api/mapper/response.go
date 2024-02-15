package mapper

import (
	mapv1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/map/v1"
	"github.com/ZecretBone/ips-bff/internal/models"
	"github.com/ZecretBone/ips-bff/internal/models/response"
)

func ToGetFloorListResponse(raw *mapv1.GetFloorListResponse) response.GetMapFloorListResponse {
	result := response.GetMapFloorListResponse{
		Floors:  make([]models.Map, len(raw.Floors)),
		IsAdmin: false,
	}

	for i, v := range raw.Floors {
		result.Floors[i] = models.Map{
			Name:        v.Name,
			Description: v.Description,
			Building:    v.Building,
			Number:      int(v.Number),
			Symbol:      v.Symbol,
			IsAdmin:     v.IsAdmin,
		}
	}

	return result
}

func ToGetMapURLResponse(raw *mapv1.GetMapURLResponse) response.GetMapURLResponse {
	result := response.GetMapURLResponse{
		Detail: models.Map{
			Name:        raw.Detail.Name,
			Description: raw.Detail.Description,
			Building:    raw.Detail.Building,
			Number:      int(raw.Detail.Number),
			Symbol:      raw.Detail.Symbol,
			IsAdmin:     raw.Detail.IsAdmin,
		},
		URL:       raw.Url,
		UpdatedAt: raw.UpdatedAt.AsTime(),
	}

	return result
}

func MapFetchResponseToOnlineUserStruct(resp *mapv1.FetchOnlineUserResponse) []models.OnlineUser {
	result := make([]models.OnlineUser, len(resp.OnlineUsers))
	for i, v := range resp.OnlineUsers {
		result[i] = models.OnlineUser{
			DisplayName: v.DisplayName,
			Coordinate: models.Position{
				X: float64(v.Coordinate.X),
				Y: float64(v.Coordinate.Y),
				Z: float64(v.Coordinate.Z),
			},
			Timestamp: v.Timestamp.AsTime(),
		}
	}

	return result
}
