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

		// Seed Beginning Go
		c := &models.Course{
			Title:       "Beginning Go",
			Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			Price:       5000,
		}
		verrs, err := tx.ValidateAndCreate(c)
		if verrs.HasAny() {
			return verrs
		}
		if err != nil {
			return err
		}
		modules := models.Modules{}
		err = tx.Where("slug in (?)", "errmgmt", "concurrency").All(&modules)
		if err != nil {
			return err
		}
		fmt.Printf("### modules -> %+v\n", modules)
		for i, m := range modules {
			err = tx.Create(&models.CourseModule{CourseID: c.ID, ModuleID: m.ID, Position: i})
			if err != nil {
				return err
			}
		}

		// Seed Distributed Systems
		c = &models.Course{
			Title:       "Distributed Systems",
			Description: "Distributed computing has a lot of challenges, including requirements like gossip or consensus protocols.  Additionally, how do you monitor and debug this service?  This course will walk you through the building blocks needed and best practices to tie them together.",
			Price:       5000,
		}
		verrs, err = tx.ValidateAndCreate(c)
		if verrs.HasAny() {
			return verrs
		}
		if err != nil {
			return err
		}
		modules = models.Modules{}
		err = tx.Where("slug in (?)", "distributed-systems", "concepts", "protocols", "grpc", "libraries", "existing-solutions").All(&modules)
		if err != nil {
			return err
		}
		fmt.Printf("### modules -> %+v\n", modules)
		for i, m := range modules {
			err = tx.Create(&models.CourseModule{CourseID: c.ID, ModuleID: m.ID, Position: i})
			if err != nil {
				return err
			}
		}

		return nil
	})
})
