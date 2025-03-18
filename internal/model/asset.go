package model


import "time"


type Asset struct {
	ID          int64
	Category    string
	ImgUrl      string
	Name        string
	Price       float64
	Description *string
	City        string
	Address     string
	MapsUrl     *string
	StartDate   time.Time
	EndDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
