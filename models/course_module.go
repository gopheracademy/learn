package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type CourseModule struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CourseID  uuid.UUID `json:"course_id" db:"course_id"`
	ModuleID  uuid.UUID `json:"module_id" db:"module_id"`
	Position  int       `json:"position" db:"position"`
}

// String is not required by pop and may be deleted
func (c CourseModule) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// CourseModules is not required by pop and may be deleted
type CourseModules []CourseModule

// String is not required by pop and may be deleted
func (c CourseModules) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// Validate gets run everytime you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (c *CourseModule) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: c.Position, Name: "Position"},
	), nil

}

// ValidateSave gets run everytime you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (c *CourseModule) ValidateSave(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run everytime you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (c *CourseModule) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
