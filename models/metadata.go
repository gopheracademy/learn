package models

import (
	"database/sql/driver"
	"encoding/json"
)

type MetaData map[string]string

// Scan implements the Scanner interface.
func (s *MetaData) Scan(value interface{}) error {
	return json.Unmarshal([]byte(value.(string)), s)
}

// Value implements the driver Valuer interface.
func (s MetaData) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return string(b), err
}
