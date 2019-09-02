package models

import "time"

type Birth struct {
	Item
	Links []BirthLink `json:"links" gorm:"foreignkey:BirthID"`
}

func (Birth) TableName() string {
	return "births"
}

type Death struct {
	Item
	Links []DeathLink `json:"links" gorm:"foreignkey:DeathID"`
}

func (Death) TableName() string {
	return "deaths"
}

type Event struct {
	Item
	Links []EventLink `json:"links" gorm:"foreignkey:EventID"`
}

func (Event) TableName() string {
	return "events"
}

type Item struct {
	ID        uint      `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Year      string    `json:"year"`
	HTML      string    `json:"html"`
	Text      string    `json:"text"`
	DateID    uint      `json:"-"`
}
