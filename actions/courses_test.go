package actions

func (as *ActionSuite) Test_Courses_Index_NotLoggedIn() {
	pubc, err := publicCourse(as.DB)
	as.NoError(err)

	privc, err := privateCourse(as.DB)
	as.NoError(err)

	res := as.HTML("/courses").Get()

	body := res.Body.String()
	as.Contains(body, pubc.Title)
	as.Contains(body, "must be logged in")
	as.NotContains(body, privc.Title)
}

func (as *ActionSuite) Test_Courses_Index_LoggedIn() {
	pubc, err := publicCourse(as.DB)
	as.NoError(err)

	_, err = as.Login()
	as.NoError(err)

	res := as.HTML("/courses").Get()

	body := res.Body.String()
	as.Contains(body, pubc.Title)
	as.NotContains(body, "must be logged in")
}
