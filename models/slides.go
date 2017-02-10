package models

import (
	"database/sql/driver"
	"encoding/json"
)

type Slide struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Notes    string   `json:"notes"`
	MetaData MetaData `json:"metadata"`
}

type Slides []Slide

// Scan implements the Scanner interface.
func (s *Slides) Scan(value interface{}) error {
	return json.Unmarshal([]byte(value.(string)), s)
}

// Value implements the driver Valuer interface.
func (s Slides) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return string(b), err
}
