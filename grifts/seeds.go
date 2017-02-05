package grifts

import (
	"github.com/gopheracademy/learn/models"
	. "github.com/markbates/grift/grift"
	"github.com/markbates/pop"
)

var _ = Add("seed:all", func(c *Context) error {
	err := Run("seed:courses", c)
	if err != nil {
		return err
	}
	return nil
})

var _ = Add("seed:courses", func(c *Context) error {
	return models.DB.Transaction(func(tx *pop.Connection) error {
		c := &models.Course{
			Title:       "Beginning Go",
			Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			Price:       5000,
		}
		verrs, err := tx.ValidateAndCreate(c)
		if verrs.HasAny() {
			return verrs
		}
		return err
	})
})
