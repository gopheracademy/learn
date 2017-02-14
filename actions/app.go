package actions

import (
	"fmt"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/envy"
	"github.com/gopheracademy/learn/actions/api/v1"
	"github.com/gopheracademy/learn/models"

	"github.com/markbates/goth/gothic"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {

		host := envy.Get("HOST", fmt.Sprintf("http://127.0.0.1:%s", envy.Get("PORT", "3000")))
		app = buffalo.Automatic(buffalo.Options{
			Env:         ENV,
			SessionName: "_learn_session",
			Host:        host,
		})

		app.Use(middleware.PopTransaction(models.DB))
		app.Use(setCurrentUser)
		app.Use(trackLastURL)
		app.Use(setStripeKeys)
		app.Use(func(next buffalo.Handler) buffalo.Handler {
			return func(c buffalo.Context) error {
				c.Set("year", time.Now().Year())
				return next(c)
			}
		})

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
		app.GET("/training-assets/{asset:.+}", TrainingAssets)

		api := app.Group("/api/v1")
		api.POST("/github", v1.GitHubWebhook)
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
