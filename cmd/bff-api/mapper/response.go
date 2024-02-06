package mapper

import (
	mapservice "github.com/ZecretBone/ips-bff/apps/map"
	mapResponse "github.com/ZecretBone/ips-bff/apps/map/models/response"
	mapv1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/map/v1"
)

func ToGetFloorListResponse(raw *mapv1.GetFloorListResponse) mapResponse.GetMapFloorListResponse {
	result := mapResponse.GetMapFloorListResponse{
		Floors:  make([]mapservice.Map, len(raw.Floors)),
		IsAdmin: false,
	}

	for i, v := range raw.Floors {
		result.Floors[i] = mapservice.Map{
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

func ToGetMapURLResponse(raw *mapv1.GetMapURLResponse) mapResponse.GetMapURLResponse {
	result := mapResponse.GetMapURLResponse{
		Detail: mapservice.Map{
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
