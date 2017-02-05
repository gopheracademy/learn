package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type Course struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Price       int       `json:"price" db:"price"`
}

// String is not required by pop and may be deleted
func (c Course) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// Courses is not required by pop and may be deleted
type Courses []Course

// String is not required by pop and may be deleted
func (c Courses) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// Validate gets run everytime you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (c *Course) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.Title, Name: "Title"},
		&validators.StringIsPresent{Field: c.Description, Name: "Description"},
		&validators.IntIsPresent{Field: c.Price, Name: "Price"},
	), nil

}

// ValidateSave gets run everytime you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (c *Course) ValidateSave(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run everytime you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (c *Course) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
