package actions

import (
	"fmt"
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/gopheracademy/learn/models"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/gplus"
	"github.com/markbates/goth/providers/linkedin"
	"github.com/markbates/goth/providers/twitter"
	"github.com/markbates/pop"
	"github.com/markbates/pop/nulls"
)

func init() {
	gothic.Store = App().SessionStore

	goth.UseProviders(
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/github/callback")),
		linkedin.New(os.Getenv("LINKEDIN_KEY"), os.Getenv("LINKEDIN_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/linkedin/callback")),
		twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/twitter/callback")),
		facebook.New(os.Getenv("FACEBOOK_KEY"), os.Getenv("FACEBOOK_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/facebook/callback")),
		gplus.New(os.Getenv("GPLUS_KEY"), os.Getenv("GPLUS_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/gplus/callback")),
	)
}

func AuthCallback(c buffalo.Context) error {
	guser, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(401, err)
	}

	tx := c.Value("tx").(*pop.Connection)
	u := &models.User{}
	q := tx.Where("provider = ? and provider_id = ?", guser.Provider, guser.UserID)
	b, err := q.Exists(u)
	if err != nil {
		return err
	}
	if b {
		err := q.First(u)
		if err != nil {
			return err
		}
		return login(c, u)
	}
	// the user doesn't exist, so create it
	u.Name = guser.Name
	u.Email = nulls.NewString(guser.Email)
	u.Provider = guser.Provider
	u.ProviderID = guser.UserID
	u.AvatarUrl = nulls.NewString(guser.AvatarURL)

	verrs, err := tx.ValidateAndCreate(u)
	if err != nil {
		return err
	}
	if verrs.HasAny() {
		for _, e := range verrs.Errors {
			c.Flash().Set("danger", e)
		}
		return c.Render(422, r.HTML("index.html"))
	}

	return login(c, u)
}

// AuthLogout default implementation.
func AuthLogout(c buffalo.Context) error {
	c.Session().Delete("current_user_id")
	err := c.Session().Save()
	if err != nil {
		return err
	}
	c.Flash().Add("success", "You have been successfully logged out!")
	return c.Redirect(302, "/")
}

func login(c buffalo.Context, u *models.User) error {
	c.Session().Set("current_user_id", u.ID)
	err := c.Session().Save()
	if err != nil {
		return err
	}

	c.Flash().Add("success", "You have been successfully logged in!")
	return c.Redirect(302, "/")
}

func authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if cu := c.Value("current_user"); cu == nil {
			return c.Redirect(302, "/")
		}
		return next(c)
	}
}

func setCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			tx := c.Value("tx").(*pop.Connection)
			u := &models.User{}
			err := tx.Find(u, uid)
			if err == nil {
				c.Set("current_user_id", u.ID)
				c.Set("current_user", u)
			}
		}
		return next(c)
	}
}
