package models

import "time"

// Hotel model
type Hotel struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	Email         string     `json:"email,omitempty"`
	Country       string     `json:"country,omitempty"`
	City          string     `json:"city,omitempty"`
	Description   string     `json:"description,omitempty"`
	Location      string     `json:"location"`
	Rating        float64    `json:"rating"`
	Image         *string    `json:"image,omitempty"`
	Photos        []string   `json:"photos,omitempty"`
	CommentsCount int        `json:"comments_count,omitempty"`
	Latitude      *float64   `json:"latitude,omitempty"`
	Longitude     *float64   `json:"longitude,omitempty"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}
