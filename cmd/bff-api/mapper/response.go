package mapper

import (
	mapv1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/map/v1"
	userv1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/user/v1"
	"github.com/ZecretBone/ips-bff/internal/models"
	"github.com/ZecretBone/ips-bff/internal/models/response"
)

func ToGetCoordinateResponse(raw *userv1.GetCoordinateResponse) response.GetSingleCoordinateResponse {
	result := response.GetSingleCoordinateResponse{
		X:        float64(raw.Position.X),
		Y:        float64(raw.Position.Y),
		Z:        float64(raw.Position.Z),
		Label:    raw.Label,
		Building: raw.Building,
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

func ToBuildingModelList(body *mapv1.GetBuildingListResponse) []models.Building {
	result := make([]models.Building, len(body.Buildings))

	for i, v := range body.Buildings {
		result[i] = models.Building{
			Name:        v.Name,
			Description: v.Description,
			OriginLat:   v.OriginLat,
			OriginLong:  v.OriginLong,
		}
	}

	return result
}

func ToBuildingDetailModel(body *mapv1.GetBuildingInfoResponse) models.Building {
	result := models.Building{
		Name:        body.Name,
		Description: body.Description,
		FloorList:   make([]models.Floor, len(body.Floors)),
		OriginLat:   body.OriginLat,
		OriginLong:  body.OriginLong,
	}

	for i, v := range body.Floors {
		result.FloorList[i] = models.Floor{
			Name:     v.Name,
			Building: v.Building,
			Floor:    int(v.Floor),
			Symbol:   v.Symbol,
			IsAdmin:  v.IsAdmin,
		}
	}

	return result
}

func ToFloorDetailModel(body *mapv1.GetFloorInfoResponse) models.FloorDetail {
	result := models.FloorDetail{
		Info: models.Floor{
			Name:     body.Info.Name,
			Building: body.Info.Building,
			Floor:    int(body.Info.Floor),
			Symbol:   body.Info.Symbol,
			IsAdmin:  body.Info.IsAdmin,
		},
		RoomList: make([]models.Room, len(body.Rooms)),
	}

	for i, v := range body.Rooms {
		result.RoomList[i] = models.Room{
			RoomID:      v.RoomId,
			Name:        v.Name,
			Description: v.Description,
			Latitude:    v.Latitude,
			Longitude:   v.Longitude,
		}
	}

	return result
}
