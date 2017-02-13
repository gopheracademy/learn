package v1

import (
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	_ "github.com/gopheracademy/learn/grifts" // load grifts
	"github.com/markbates/grift/grift"
	"github.com/markbates/hmax"
	"github.com/pkg/errors"
)

var hmac = hmax.New("X-Hub-Signature", []byte(envy.Get("GITHUB_WEBHOOK_SECRET", "some-secret")))

func GitHubWebhook(c buffalo.Context) error {
	req := c.Request()
	xhs := req.Header.Get(hmac.Header)
	xhs = strings.TrimPrefix(xhs, "sha1=")
	req.Header.Set(hmac.Header, xhs)
	b, err := hmac.VerifyRequest(req)
	if err != nil {
		return err
	}
	if !b {
		return errors.Errorf("could not verify signature with %s", c.Request().Header.Get(hmac.Header))
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
