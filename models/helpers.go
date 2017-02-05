package models

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/gopheracademy/material/content"
)

type FullCourse struct {
	Course     *content.Course
	ModuleList []*FullModule
	Instructor *content.Instructor
}

type FullModule struct {
	Module     *content.Module
	LessonList []*content.Lesson
}

func GetFullCourse(id int) (*FullCourse, error) {
	fmt.Printf("using %s\n", BaseURL)
	var fc FullCourse
	c, err := GetCourse(id)
	if err != nil {
		fmt.Println("Get Course Error", err)
		return &fc, err
	}
	fc = FullCourse{
		Course:     &c,
		ModuleList: make([]*FullModule, len(c.Modules)),
	}
	for i, id := range ModulesIDsForCourse(c) {
		fm, err := GetFullModule(id)
		if err != nil {
			fmt.Println("GFM Error", err)
			return &fc, err
		}
		fc.ModuleList[i] = fm
	}
	if c.Instructor != "" {
		id, err := getID(c.Instructor)
		if err == nil {
			inst, err := GetInstructor(id)
			if err == nil {
				fc.Instructor = &inst
			}
		}
	}
	return &fc, nil
}
func GetFullModule(id int) (*FullModule, error) {
	var fm FullModule
	c, err := GetModule(id)
	if err != nil {
		return &fm, err
	}
	fm = FullModule{
		Module:     &c,
		LessonList: make([]*content.Lesson, len(c.Lessons)),
	}

	for i, id := range LessonsIDsForModule(c) {
		lesson, err := GetLesson(id)
		if err != nil {
			return &fm, err
		}
		fm.LessonList[i] = &lesson
	}
	return &fm, nil
}
func ModulesIDsForCourse(c content.Course) []int {
	var modules []int
	for _, s := range c.Modules {
		i, err := getID(s)
		if err == nil {
			modules = append(modules, i)
		}
	}
	return modules
}

func LessonsIDsForModule(m content.Module) []int {
	var lessons []int
	for _, s := range m.Lessons {
		i, err := getID(s)
		if err == nil {
			lessons = append(lessons, i)
		}
	}
	return lessons
}

func getID(s string) (int, error) {
	//?type=Module&id=4
	u, err := url.Parse(s)
	if err != nil {
		return 0, err
	}
	vals := u.Query()
	ii, ok := vals["id"]
	if !ok {
		return 0, err
	}
	return strconv.Atoi(ii[0])
}
