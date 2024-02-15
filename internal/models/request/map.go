package request

type GetMapFloorListRequest struct {
	Building string `json:"building"`
	Role     string `json:"role"`
}
