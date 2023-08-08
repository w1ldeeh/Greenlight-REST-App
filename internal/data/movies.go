package data

import (
	"time"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` //`json:"created_at"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`           //`json:"year"`
	Runtime   int32     `json:"runtime,omitempty,string"` //`json:"runtime"`
	Genres    []string  `json:"genres,omitempty"`         //`json:"genres"`
	Version   int32     `json:"version"`
}
