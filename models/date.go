package models

import "time"

type DateResponseBody struct {
	URL  string `json:"url"`
	Date string `json:"date"`
	Data struct {
		Events []Event `json:"Events" gorm:"foreignkey:DateID"`
		Births []Birth `json:"Births" gorm:"foreignkey:DateID"`
		Deaths []Death `json:"Deaths" gorm:"foreignkey:DateID"`
	} `json:"data"`
}

type Date struct {
	ID         uint      `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	URL        string    `json:"url"`
	DateString string    `gorm:"unique;not null" json:"date"`
	Count      struct {
		Events uint `json:"events"`
		Births uint `json:"births"`
		Deaths uint `json:"deaths"`
	} `gorm:"-" json:"counts,omitempty"`
	Events []Event `json:"Events" gorm:"foreignkey:DateID"`
	Births []Birth `json:"Births" gorm:"foreignkey:DateID"`
	Deaths []Death `json:"Deaths" gorm:"foreignkey:DateID"`
}

func (Date) TableName() string {
	return "dates"
}
