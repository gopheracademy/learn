/*
* CODE GENERATED AUTOMATICALLY WITH github.com/bketelsen/ponzigen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package models

import (
	"time"

	"github.com/bketelsen/ponzi"
	"github.com/gopheracademy/material/content"
	"github.com/pkg/errors"
)

var BaseURL string

type CourseListResult struct {
	Data []content.Course `json:"data"`
}
type CourseResultListResult struct {
	Data []content.CourseResult `json:"data"`
}
type InstructorListResult struct {
	Data []content.Instructor `json:"data"`
}
type JobListResult struct {
	Data []content.Job `json:"data"`
}
type LessonListResult struct {
	Data []content.Lesson `json:"data"`
}
type LessonResultListResult struct {
	Data []content.LessonResult `json:"data"`
}
type ModuleListResult struct {
	Data []content.Module `json:"data"`
}
type ModuleResultListResult struct {
	Data []content.ModuleResult `json:"data"`
}
type ResourceListResult struct {
	Data []content.Resource `json:"data"`
}

var courseCache *ponzi.Cache
var courseresultCache *ponzi.Cache
var instructorCache *ponzi.Cache
var jobCache *ponzi.Cache
var lessonCache *ponzi.Cache
var lessonresultCache *ponzi.Cache
var moduleCache *ponzi.Cache
var moduleresultCache *ponzi.Cache
var resourceCache *ponzi.Cache

func initCourseCache() {
	if courseCache == nil {
		courseCache = ponzi.New(BaseURL, 5*time.Minute, 30*time.Second)
	}
}
func initCourseResultCache() {
	if courseresultCache == nil {
		courseresultCache = ponzi.New(BaseURL, 5*time.Minute, 30*time.Second)
	}
}
func initInstructorCache() {
	if instructorCache == nil {
		instructorCache = ponzi.New(BaseURL, 5*time.Minute, 30*time.Second)
	}
}
func initJobCache() {
	if jobCache == nil {
		jobCache = ponzi.New(BaseURL, 5*time.Minute, 30*time.Second)
	}
}
func initLessonCache() {
	if lessonCache == nil {
		lessonCache = ponzi.New(BaseURL, 5*time.Minute, 30*time.Second)
	}
}
func initLessonResultCache() {
	if lessonresultCache == nil {
		lessonresultCache = ponzi.New(BaseURL, 5*time.Minute, 30*time.Second)
	}
}
func initModuleCache() {
	if moduleCache == nil {
		moduleCache = ponzi.New(BaseURL, 5*time.Minute, 30*time.Second)
	}
}
func initModuleResultCache() {
	if moduleresultCache == nil {
		moduleresultCache = ponzi.New(BaseURL, 5*time.Minute, 30*time.Second)
	}
}
func initResourceCache() {
	if resourceCache == nil {
		resourceCache = ponzi.New(BaseURL, 5*time.Minute, 30*time.Second)
	}
}

func GetCourse(id int) (content.Course, error) {
	initCourseCache()
	var sp CourseListResult
	err := courseCache.Get(id, "Course", &sp)
	if err != nil {
		return content.Course{}, err
	}
	if len(sp.Data) == 0 {
		return content.Course{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetCourseResult(id int) (content.CourseResult, error) {
	initCourseResultCache()
	var sp CourseResultListResult
	err := courseresultCache.Get(id, "CourseResult", &sp)
	if err != nil {
		return content.CourseResult{}, err
	}
	if len(sp.Data) == 0 {
		return content.CourseResult{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetInstructor(id int) (content.Instructor, error) {
	initInstructorCache()
	var sp InstructorListResult
	err := instructorCache.Get(id, "Instructor", &sp)
	if err != nil {
		return content.Instructor{}, err
	}
	if len(sp.Data) == 0 {
		return content.Instructor{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetJob(id int) (content.Job, error) {
	initJobCache()
	var sp JobListResult
	err := jobCache.Get(id, "Job", &sp)
	if err != nil {
		return content.Job{}, err
	}
	if len(sp.Data) == 0 {
		return content.Job{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetLesson(id int) (content.Lesson, error) {
	initLessonCache()
	var sp LessonListResult
	err := lessonCache.Get(id, "Lesson", &sp)
	if err != nil {
		return content.Lesson{}, err
	}
	if len(sp.Data) == 0 {
		return content.Lesson{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetLessonResult(id int) (content.LessonResult, error) {
	initLessonResultCache()
	var sp LessonResultListResult
	err := lessonresultCache.Get(id, "LessonResult", &sp)
	if err != nil {
		return content.LessonResult{}, err
	}
	if len(sp.Data) == 0 {
		return content.LessonResult{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetModule(id int) (content.Module, error) {
	initModuleCache()
	var sp ModuleListResult
	err := moduleCache.Get(id, "Module", &sp)
	if err != nil {
		return content.Module{}, err
	}
	if len(sp.Data) == 0 {
		return content.Module{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetModuleResult(id int) (content.ModuleResult, error) {
	initModuleResultCache()
	var sp ModuleResultListResult
	err := moduleresultCache.Get(id, "ModuleResult", &sp)
	if err != nil {
		return content.ModuleResult{}, err
	}
	if len(sp.Data) == 0 {
		return content.ModuleResult{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetResource(id int) (content.Resource, error) {
	initResourceCache()
	var sp ResourceListResult
	err := resourceCache.Get(id, "Resource", &sp)
	if err != nil {
		return content.Resource{}, err
	}
	if len(sp.Data) == 0 {
		return content.Resource{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}

func GetCourseBySlug(slug string) (content.Course, error) {
	initCourseCache()
	var sp CourseListResult
	err := courseCache.GetBySlug(slug, "Course", &sp)
	if err != nil {
		return content.Course{}, err
	}
	if len(sp.Data) == 0 {
		return content.Course{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetCourseResultBySlug(slug string) (content.CourseResult, error) {
	initCourseResultCache()
	var sp CourseResultListResult
	err := courseresultCache.GetBySlug(slug, "CourseResult", &sp)
	if err != nil {
		return content.CourseResult{}, err
	}
	if len(sp.Data) == 0 {
		return content.CourseResult{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetInstructorBySlug(slug string) (content.Instructor, error) {
	initInstructorCache()
	var sp InstructorListResult
	err := instructorCache.GetBySlug(slug, "Instructor", &sp)
	if err != nil {
		return content.Instructor{}, err
	}
	if len(sp.Data) == 0 {
		return content.Instructor{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetJobBySlug(slug string) (content.Job, error) {
	initJobCache()
	var sp JobListResult
	err := jobCache.GetBySlug(slug, "Job", &sp)
	if err != nil {
		return content.Job{}, err
	}
	if len(sp.Data) == 0 {
		return content.Job{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetLessonBySlug(slug string) (content.Lesson, error) {
	initLessonCache()
	var sp LessonListResult
	err := lessonCache.GetBySlug(slug, "Lesson", &sp)
	if err != nil {
		return content.Lesson{}, err
	}
	if len(sp.Data) == 0 {
		return content.Lesson{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetLessonResultBySlug(slug string) (content.LessonResult, error) {
	initLessonResultCache()
	var sp LessonResultListResult
	err := lessonresultCache.GetBySlug(slug, "LessonResult", &sp)
	if err != nil {
		return content.LessonResult{}, err
	}
	if len(sp.Data) == 0 {
		return content.LessonResult{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetModuleBySlug(slug string) (content.Module, error) {
	initModuleCache()
	var sp ModuleListResult
	err := moduleCache.GetBySlug(slug, "Module", &sp)
	if err != nil {
		return content.Module{}, err
	}
	if len(sp.Data) == 0 {
		return content.Module{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetModuleResultBySlug(slug string) (content.ModuleResult, error) {
	initModuleResultCache()
	var sp ModuleResultListResult
	err := moduleresultCache.GetBySlug(slug, "ModuleResult", &sp)
	if err != nil {
		return content.ModuleResult{}, err
	}
	if len(sp.Data) == 0 {
		return content.ModuleResult{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetResourceBySlug(slug string) (content.Resource, error) {
	initResourceCache()
	var sp ResourceListResult
	err := resourceCache.GetBySlug(slug, "Resource", &sp)
	if err != nil {
		return content.Resource{}, err
	}
	if len(sp.Data) == 0 {
		return content.Resource{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}

func GetCourseList() ([]content.Course, error) {
	initCourseCache()
	var sp CourseListResult
	err := courseCache.GetAll("Course", &sp)
	if err != nil {
		return []content.Course{}, err
	}
	if len(sp.Data) == 0 {
		return []content.Course{}, errors.New("Not Found")
	}
	return sp.Data, err

}
func GetCourseResultList() ([]content.CourseResult, error) {
	initCourseResultCache()
	var sp CourseResultListResult
	err := courseresultCache.GetAll("CourseResult", &sp)
	if err != nil {
		return []content.CourseResult{}, err
	}
	if len(sp.Data) == 0 {
		return []content.CourseResult{}, errors.New("Not Found")
	}
	return sp.Data, err

}
func GetInstructorList() ([]content.Instructor, error) {
	initInstructorCache()
	var sp InstructorListResult
	err := instructorCache.GetAll("Instructor", &sp)
	if err != nil {
		return []content.Instructor{}, err
	}
	if len(sp.Data) == 0 {
		return []content.Instructor{}, errors.New("Not Found")
	}
	return sp.Data, err

}
func GetJobList() ([]content.Job, error) {
	initJobCache()
	var sp JobListResult
	err := jobCache.GetAll("Job", &sp)
	if err != nil {
		return []content.Job{}, err
	}
	if len(sp.Data) == 0 {
		return []content.Job{}, errors.New("Not Found")
	}
	return sp.Data, err

}
func GetLessonList() ([]content.Lesson, error) {
	initLessonCache()
	var sp LessonListResult
	err := lessonCache.GetAll("Lesson", &sp)
	if err != nil {
		return []content.Lesson{}, err
	}
	if len(sp.Data) == 0 {
		return []content.Lesson{}, errors.New("Not Found")
	}
	return sp.Data, err

}
func GetLessonResultList() ([]content.LessonResult, error) {
	initLessonResultCache()
	var sp LessonResultListResult
	err := lessonresultCache.GetAll("LessonResult", &sp)
	if err != nil {
		return []content.LessonResult{}, err
	}
	if len(sp.Data) == 0 {
		return []content.LessonResult{}, errors.New("Not Found")
	}
	return sp.Data, err

}
func GetModuleList() ([]content.Module, error) {
	initModuleCache()
	var sp ModuleListResult
	err := moduleCache.GetAll("Module", &sp)
	if err != nil {
		return []content.Module{}, err
	}
	if len(sp.Data) == 0 {
		return []content.Module{}, errors.New("Not Found")
	}
	return sp.Data, err

}
func GetModuleResultList() ([]content.ModuleResult, error) {
	initModuleResultCache()
	var sp ModuleResultListResult
	err := moduleresultCache.GetAll("ModuleResult", &sp)
	if err != nil {
		return []content.ModuleResult{}, err
	}
	if len(sp.Data) == 0 {
		return []content.ModuleResult{}, errors.New("Not Found")
	}
	return sp.Data, err

}
func GetResourceList() ([]content.Resource, error) {
	initResourceCache()
	var sp ResourceListResult
	err := resourceCache.GetAll("Resource", &sp)
	if err != nil {
		return []content.Resource{}, err
	}
	if len(sp.Data) == 0 {
		return []content.Resource{}, errors.New("Not Found")
	}
	return sp.Data, err

}
