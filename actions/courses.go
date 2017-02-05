package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gopheracademy/learn/models"
	"github.com/markbates/pop"
)

// CoursesIndex default implementation.
func CoursesIndex(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	courses := &models.Courses{}
	err := tx.All(courses)
	if err != nil {
		return err
	}
	c.Set("courses", courses)
	return c.Render(200, r.HTML("courses/index.html"))
}

// CoursesShow default implementation.
func CoursesShow(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	course := &models.Course{}
	err := tx.Find(course, c.Param("course_id"))
	if err != nil {
		return err
	}
	c.Set("course", course)
	return c.Render(200, r.HTML("courses/show.html"))
}
