package models

type QueryDate struct {
	Month int `json:"month" query:"month"`
	Day   int `json:"day" query:"day"`
}
