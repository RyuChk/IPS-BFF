package models

import "time"

type Map struct {
	Name        string    `json:"name"`        //Map Name
	Description string    `json:"description"` //Floor Description
	Number      int       `json:"number"`      //Floor number
	Symbol      string    `json:"symbol"`
	Building    string    `json:"building"`   //Name of the building
	IsAdmin     bool      `json:"is_admin"`   //Only admin can view
	CreatedAt   time.Time `json:"created_at"` //Create time
	UpdatedAt   time.Time `json:"updated_at"` //Update time
}

type MapImageURL struct {
	Key       string    `json:"key"`        //<Building>-<Symbol>
	MapDetail Map       `json:"-"`          //Map detail
	URL       string    `json:"url"`        //Map image URL
	CreatedAt time.Time `json:"created_at"` //Create time
	UpdatedAt time.Time `json:"updated_at"` //Update time
}
