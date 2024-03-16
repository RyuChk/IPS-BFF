package models

type Building struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	OriginLat   float64 `json:"origin_lat"`
	OriginLong  float64 `json:"origin_long"`
	FloorList   []Floor `json:"floor_list,omitempty"`
}

type Floor struct {
	Name        string  `json:"name"`        //Map Name
	Description string  `json:"description"` //Floor Description
	Floor       int     `json:"floor"`       //Floor number
	Symbol      string  `json:"symbol"`
	Building    string  `json:"building"` //Name of the building
	IsAdmin     bool    `json:"is_admin"` //Only admin can view
	OriginLat   float64 `json:"origin_lat"`
	OriginLong  float64 `json:"origin_long"`
	MapUrl      string  `json:"map_url"`
}

type FloorDetail struct {
	Key      string `json:"key,omitempty"` //Key
	Info     Floor  `json:"info"`
	RoomList []Room `json:"room"`
}

type Room struct {
	RoomID      string  `json:"room_id"`     // Room ID
	Name        string  `json:"name"`        // Name of room
	Description string  `json:"description"` // Description
	Latitude    float64 `json:"latitude"`    // lat location
	Longitude   float64 `json:"longitude"`   // long location
}
