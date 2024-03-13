package mapper

import (
	"github.com/ZecretBone/ips-bff/internal/constants"
	v1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/rssi/v1"
	rssiv1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/shared/rssi/v1"
	userv1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/user/v1"
	"github.com/ZecretBone/ips-bff/internal/models/request"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ToEnumDataCollectionStage = map[constants.DataCollectionStages]rssiv1.StatCollectionStage{
		constants.DataCollectionStageSingle:   rssiv1.StatCollectionStage_STAT_COLLECTION_STAGE_SINGLE,
		constants.DataCollectionStageMultiple: rssiv1.StatCollectionStage_STAT_COLLECTION_STAGE_MULTIPLE,
	}

	ToEnumRecordingDirection = map[constants.Direction]rssiv1.RecordingDirection{
		constants.North: rssiv1.RecordingDirection_RECORDING_DIRECTION_NORTH,
		constants.South: rssiv1.RecordingDirection_RECORDING_DIRECTION_SOUTH,
		constants.East:  rssiv1.RecordingDirection_RECORDING_DIRECTION_EAST,
		constants.West:  rssiv1.RecordingDirection_RECORDING_DIRECTION_WEST,
	}
)

func ToGetCoordinateRequest(req request.GetSingleCoordinateRequest, user string, isAdmin bool) *userv1.GetCoordinateRequest {
	var data userv1.GetCoordinateRequest

	signals := make([]*rssiv1.RSSI, len(req.Signals))

	for i, v := range req.Signals {

		time := make([]*timestamppb.Timestamp, len(v.CreatedAt))
		for j, t := range v.CreatedAt {
			time[j] = timestamppb.New(t)
		}

		signals[i] = &rssiv1.RSSI{
			Ssid:       v.SSID,
			MacAddress: v.MacAddress,
			Strength:   v.Strength,
			CreatedAt:  time,
		}
	}

	data.Building = req.Building
	data.User = user
	data.IsAdmin = isAdmin
	data.Signals = signals

	return &data

}

func ToDataCollectionDataRequest(req request.StatCollectionRequest, device_id, model string) *v1.CollectDataRequest {
	signals := make([]*rssiv1.RSSI, len(req.Signals))

	for i, v := range req.Signals {

		time := make([]*timestamppb.Timestamp, len(v.CreatedAt))
		for j, t := range v.CreatedAt {
			time[j] = timestamppb.New(t)
		}

		signals[i] = &rssiv1.RSSI{
			Ssid:       v.SSID,
			MacAddress: v.MacAddress,
			Strength:   v.Strength,
			CreatedAt:  time,
		}
	}

	data := v1.CollectDataRequest{
		Signals: signals,
		DeviceInfo: &rssiv1.DeviceInfo{
			DeviceId: device_id,
			Model:    model,
		},
		Position: &rssiv1.Position{
			X: float32(req.Position.X),
			Y: float32(req.Position.Y),
			Z: float32(req.Position.Z),
		},
		Duration:    int32(req.Duration),
		PollingRate: int32(req.PollingRate),
		Stage:       ToEnumDataCollectionStage[req.StatCollectionStage],
		Direction:   ToEnumRecordingDirection[req.Direction],
		StartedAt:   timestamppb.New(req.StartedAt),
		CreatedAt:   timestamppb.New(req.CreatedAt),
		EndedAt:     timestamppb.New(req.EndedAt),
	}

	return &data
}
