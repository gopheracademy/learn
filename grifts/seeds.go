package grifts

import (
	"fmt"

	"github.com/gopheracademy/learn/models"
	. "github.com/markbates/grift/grift"
	"github.com/markbates/pop"
)

var _ = Desc("seed:all", "Updates the git repo, builds the modules, and seeds the courses")
var _ = Add("seed:all", func(c *Context) error {
	err := Run("pull:modules", c)
	if err != nil {
		return err
	}
	err = Run("build:modules", c)
	if err != nil {
		return err
	}
	err = Run("seed:courses", c)
	if err != nil {
		return err
	}
	return nil
})

var _ = Desc("seed:courses", "Deletes all the courses and purchases in the database and seeds new courses")
var _ = Add("seed:courses", func(c *Context) error {
	return models.DB.Transaction(func(tx *pop.Connection) error {
		for _, x := range []string{"courses", "course_modules", "purchases"} {
			err := tx.RawQuery(fmt.Sprintf("delete from %s", x)).Exec()
			if err != nil {
				return err
			}
		}

		// Seed Distributed Systems
		c := &models.Course{
			Title:       "Distributed Systems",
			Description: "Distributed computing has a lot of challenges, including requirements like gossip or consensus protocols.  Additionally, how do you monitor and debug this service?  This course will walk you through the building blocks needed and best practices to tie them together.",
			Price:       5000,
			Status:      "public",
		}
		verrs, err := tx.ValidateAndCreate(c)
		if verrs.HasAny() {
			return verrs
		}
		if err != nil {
			return err
		}
		for i, slug := range []string{"distributed-systems", "concepts", "protocols", "grpc", "libraries", "existing-solutions"} {
			m := &models.Module{}
			err = tx.Where("slug = ?", slug).First(m)
			if err != nil {
				return err
			}
			err = tx.Create(&models.CourseModule{CourseID: c.ID, ModuleID: m.ID, Position: i})
			if err != nil {
				return err
			}
		}

		return nil
	})
})
