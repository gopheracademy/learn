package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/satori/go.uuid"
)

type Purchase struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	CourseID  uuid.UUID `json:"course_id" db:"course_id"`
}

// String is not required by pop and may be deleted
func (p Purchase) String() string {
	b, _ := json.Marshal(p)
	return string(b)
}

// Purchases is not required by pop and may be deleted
type Purchases []Purchase

// String is not required by pop and may be deleted
func (p Purchases) String() string {
	b, _ := json.Marshal(p)
	return string(b)
}

// Validate gets run everytime you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (p *Purchase) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateSave gets run everytime you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (p *Purchase) ValidateSave(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run everytime you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (p *Purchase) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
