package models

import "time"

type Event struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Organization string    `json:"organization"`
	Date         time.Time `json:"date"`
	Price        float64   `json:"price"`
	Rating       string    `json:"rating"`
	ImageURL     string    `json:"image_url"`
	CreatedAt    time.Time `json:"created_at"`
	Location     string    `json:"location"`
}

type Spot struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	EventID int    `json:"event_id"`
}

type Data struct {
	Events []Event `json:"events"`
	Spots  []Spot  `json:"spots"`
}

type Reservation struct {
	Spots []string `json:"spots"`
}
