package models

import "time"

type DeathLink struct {
	Link
	DeathID uint `json:"-"`
}
type BirthLink struct {
	Link
	BirthID uint `json:"-"`
}
type EventLink struct {
	Link
	EventID uint `json:"-"`
}

type Link struct {
	ID        uint      `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Title     string    `json:"title"`
	Link      string    `json:"link"`
}

func (DeathLink) TableName() string {
	return "death_links"
}

func (BirthLink) TableName() string {
	return "birth_links"
}

func (EventLink) TableName() string {
	return "event_links"
}
