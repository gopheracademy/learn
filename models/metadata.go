package models

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type MetaData map[string]string

func (m MetaData) String() string {
	bb := &bytes.Buffer{}
	for k, v := range m {
		bb.WriteString(fmt.Sprintf("%s:%s\n", k, v))
	}
	return bb.String()
}

// Scan implements the Scanner interface.
func (s *MetaData) Scan(value interface{}) error {
	return json.Unmarshal([]byte(value.(string)), s)
}

// Value implements the driver Valuer interface.
func (s MetaData) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return string(b), err
}
