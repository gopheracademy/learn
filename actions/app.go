package actions

import (
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gopheracademy/learn/models"
	"github.com/markbates/going/defaults"

	"github.com/markbates/goth/gothic"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = defaults.String(os.Getenv("GO_ENV"), "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.Automatic(buffalo.Options{
			Env:         ENV,
			SessionName: "_learn_session",
		})

		app.Use(middleware.PopTransaction(models.DB))
		app.Use(setCurrentUser)
		app.Use(trackLastURL)
		app.Use(setStripeKeys)

		app.GET("/", HomeHandler)

		app.ServeFiles("/assets", assetsPath())
		auth := app.Group("/auth")
		auth.Middleware.Replace(trackLastURL, func(next buffalo.Handler) buffalo.Handler {
			return func(c buffalo.Context) error {
				return next(c)
			}
		})
		auth.GET("/{provider}", buffalo.WrapHandlerFunc(gothic.BeginAuthHandler))
		auth.GET("/{provider}/callback", AuthCallback)
		app.DELETE("/logout", AuthLogout)
		app.GET("/courses", CoursesIndex)
		app.GET("/courses/{course_id}", CoursesShow)
		app.POST("/courses/{course_id}/purchases", authorize(PurchasesCreate))
	}

	return app
}

func trackLastURL(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		req := c.Request()
		if req.Method == "GET" {
			c.Session().Set("last_url", req.URL.Path)
			err := c.Session().Save()
			if err != nil {
				return err
			}
		}
		return next(c)
	}
}
