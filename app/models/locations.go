package models

type Locations struct {
	BaseModel
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
