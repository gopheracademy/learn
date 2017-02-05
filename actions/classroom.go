package actions

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/gopheracademy/learn/models"
	"github.com/gopheracademy/material/content"
)

// ClassroomShow default implementation.
func ClassroomShow(c buffalo.Context) error {
	var mid, lid int
	var err error
	mid, err = c.ParamInt("mid")
	if err != nil {
		mid = 0
	}
	c.LogField("module", mid)
	lid, err = c.ParamInt("lid")
	if err != nil {
		lid = 0
	}

	c.LogField("lesson", lid)
	c.Set("mid", mid)
	c.Set("lid", lid)
	id := 2
	course, err := models.GetFullCourse(id)
	if err != nil {
		return c.Error(500, err)
	}
	var lsn content.Lesson
	var mdl models.FullModule
	if mid != 0 {
		for _, mod := range course.ModuleList {
			if mod.Module.ID == mid {
				mdl = *mod

				if lid != 0 {
					for _, les := range mod.LessonList {
						if les.ID == lid {
							lsn = *les

						}
					}
				}
			}
		}
	}

	var hasVideo bool
	var hasContent bool
	var hasLesson bool
	var videoCode string
	// if we have a specified lesson, and it contains a video
	// then display the video tab
	if lid != 0 {
		if lsn.VideoCode != "" {
			hasVideo = true
			videoCode = lsn.VideoCode
		}
	} else {
		// no lesson specified, maybe the start of a module?
		if mid != 0 {
			if mdl.Module.VideoCode != "" {
				hasVideo = true
				videoCode = mdl.Module.VideoCode
			}
		} else {
			// no module specified, maybe the start of a course
			if course.Course.VideoCode != "" {
				hasVideo = true
				videoCode = course.Course.VideoCode
			}
		}
	}
	if len(videoCode) > 0 {
		c.Set("videoCode", videoCode)
	}

	if lid != 0 {
		hasLesson = true
	}
	pl, nl := prevNextLinks(course, mid, lid)
	c.Set("mdl", mdl)
	c.LogField("mdl", mdl)
	c.Set("lsn", lsn)
	c.LogField("lsn", lsn)
	c.Set("hasVideo", hasVideo)
	c.Set("hasContent", hasContent)
	c.Set("hasLesson", hasLesson)
	c.Set("prevLink", pl)
	c.Set("nextLink", nl)
	c.Set("course", course)
	return c.Render(200, r.HTML("classroom/show.html"))
}

func prevNextLinks(c *models.FullCourse, mid, lid int) (nextLink, prevLink string) {

	// TODO(BJK) - what a clusterfuck.  Sleep more, clean up.
	// if module is 0, we're at the beginning of the course
	if mid == 0 {
		return fmtLink(c.ModuleList[0].Module.ID, 0), ""
	}
	var previousModule, nextModule, previousLesson, nextLesson int
	for i, m := range c.ModuleList {
		if m.Module.ID == mid {
			// this is the current module
			if lid == 0 {
				fmt.Println("no current lesson")
				// No lesson yet, so next lesson is first lesson of module
				nextLesson = m.LessonList[0].ID
				if i == 0 {
					// first module, no previous
					fmt.Println("no previous module")
					previousModule = 0
					if i < len(c.ModuleList)-1 {
						nextModule = mid
						nextLesson = m.LessonList[0].ID
					} else {

						fmt.Println("no subsequent modules exist")
						nextModule = 0
					}
				} else {
					// not first module, there is a previous

					fmt.Println("not first module")
					previousModule = c.ModuleList[i-1].Module.ID
					if i < len(m.LessonList)-1 {

						fmt.Println("subsequent lessons exist")
						nextModule = mid
						nextLesson = m.LessonList[0].ID
					} else {

						fmt.Println("no subsequent lessons exist")
						nextModule = c.ModuleList[i+1].Module.ID
					}
				}

			} else {
				// we have a current lesson
				for j, les := range m.LessonList {
					if les.ID == lid {
						// we have a current lesson
						if j < len(m.LessonList)-1 {
							fmt.Println(" subsequent lessons exist")
							nextLesson = m.LessonList[j+1].ID
							nextModule = mid
							if i > 1 {
								previousModule = c.ModuleList[i-1].Module.ID
							}
							if j > 1 {
								previousLesson = m.LessonList[j-1].ID
							}
						} else {
							fmt.Println("no subsequent lessons exist")
							if i < len(c.ModuleList)-1 {
								fmt.Println("subsequent modules exist")
								nextModule = c.ModuleList[i+1].Module.ID
							} else {
								fmt.Println("no subsequent modules exist")
								nextModule = 0
							}
						}
						if j > 1 {
							previousLesson = m.LessonList[j-1].ID
							previousModule = mid
						}
					}

				}
			}
		}
	}
	fmt.Printf("MID: %d, LID: %d, PrevMod %d, NextMod %d, PrevLes %d, NextLess %d\n", mid, lid, previousModule, nextModule, previousLesson, nextLesson)
	return fmtLink(nextModule, nextLesson), fmtLink(previousModule, previousLesson)

}

func fmtLink(mid, lid int) string {
	if lid != 0 {
		return fmt.Sprintf("/classroom?mid=%d&lid=%d", mid, lid)
	}

	return fmt.Sprintf("/classroom?mid=%d", mid)
}
