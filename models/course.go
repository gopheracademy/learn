package models

import (
	"encoding/json"
	"fmt"
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
	Status      string    `json:"status" db:"status"`
	Purchased   bool      `json:"-" db:"-"`
}

// String is not required by pop and may be deleted
func (c Course) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

func (c Course) URL() string {
	return fmt.Sprintf("/courses/%s", c.ID)
}

func (c *Course) MarkAsPurchased(tx *pop.Connection, u *User) error {
	if c.Price == 0 {
		c.Purchased = true
		return nil
	}
	b, err := tx.Where("course_id = ? and user_id = ?", c.ID, u.ID).Exists("purchases")
	c.Purchased = b
	return err
}

// Courses is not required by pop and may be deleted
type Courses []Course

// String is not required by pop and may be deleted
func (c Courses) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

func (cc Courses) MarkPurchases(tx *pop.Connection, u *User) error {
	for i, c := range cc {
		err := c.MarkAsPurchased(tx, u)
		if err != nil {
			return err
		}
		cc[i] = c
	}
	return nil
}

// Validate gets run everytime you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (c *Course) Validate(tx *pop.Connection) (*validate.Errors, error) {
	validStates := []string{"public", "private"}
	return validate.Validate(
		&validators.StringIsPresent{Field: c.Title, Name: "Title"},
		&validators.StringIsPresent{Field: c.Description, Name: "Description"},
		&validators.FuncValidator{
			Field:   "Status",
			Message: fmt.Sprintf("%%s '%s' is not valid: %+v", c.Status, validStates),
			Fn: func() bool {
				for _, s := range validStates {
					if s == c.Status {
						return true
					}
				}
				return false
			},
		},
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
