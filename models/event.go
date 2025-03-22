package models

import "time"

// Event struct holds event information.
type Event struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Address     string    `json:"address"`
	ImageURL    string    `json:"image_url"`
	Description string    `json:"description"`
}
