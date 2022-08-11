package models

import (
	"time"

	"github.com/lib/pq"
)

type mpaa string

type Film struct {
	FilmID          int            `json:"film_id" gorm:"primary_key"`
	Title           string         `json:"title"`
	Description     string         `json:"description"`
	ReleaseYear     int            `json:"release_year"`
	LanguageID      uint8          `json:"language_id"`
	RentalDuration  uint8          `json:"rental_duration"`
	RentalRate      float32        `json:"rental_rate"`
	Length          uint8          `json:"length"`
	ReplacementCost float32        `json:"replacement_cost"`
	Rating          mpaa           `json:"rating"`
	LastUpdate      time.Time      `json:"last_update"`
	SpecialFeatures pq.StringArray `json:"special_features"`
	Fulltext        string         `json:"fulltext" gorm:"tsvector"`
}
