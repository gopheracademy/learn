package v1

import (
	"errors"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	_ "github.com/gopheracademy/learn/grifts" // load grifts
	"github.com/markbates/grift/grift"
)

func GitHubWebhook(c buffalo.Context) error {

	gws, err := envy.MustGet("GITHUB_WEBHOOK_SECRET")
	if err != nil {
		return c.Error(500, err)
	}

	_, err = parseHook([]byte(gws), c.Request())
	if err != nil {
		return c.Error(422, errors.New("request was not properly signed"))
	}

	go func(l buffalo.Logger) {
		gc := grift.NewContext("modules")
		err := grift.Run("pull:modules", gc)
		if err != nil {
			l.Error(err)
		}
		err = grift.Run("build:modules", gc)
		if err != nil {
			l.Error(err)
		}
	}(c.Logger())
	return nil
}
