package main

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
type Properties struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Group       int    `json:"group"`
	MediaURL    string `json:"media_url,omitempty"`
	MarkerColor string `json:"marker-color"`
}
type Entity struct {
	Type       string     `json:"type"`
	ID         int64      `json:"id"`
	Geometry   Geometry   `json:"geometry"`
	Properties Properties `json:"properties,omitempty"`
}

type NormalisedEntity struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Level     uint8   `json:"level"`
	Date      string  `json:"date"` // probably string is better
	Title     string  `json:"title"`
	MediaURL  string  `json:"media_url,omitempty"`
}

type MapObject struct {
	Casualties  []NormalisedEntity `json:"casualties"`
	Shellings   []NormalisedEntity `json:"shelling"`
	InfraDamage []NormalisedEntity `json:"infra_damage"`
}

type Entities = []Entity
