package response

type GetSingleCoordinateResponse struct {
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Z        float64 `json:"z"`
	Label    string  `json:"label"`
	Building string  `json:"building"`
}
