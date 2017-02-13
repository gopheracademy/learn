package v1

import (
	"fmt"
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
	fmt.Printf("### hmac.Secret -> %+v\n", string(hmac.Secret))
	c.Logger().Infof("### hmac.Secret -> %+v\n", string(hmac.Secret))
	req := c.Request()
	xhs := req.Header.Get(hmac.Header)
	fmt.Printf("### original github header -> %+v\n", xhs)
	c.Logger().Infof("### original github header -> %+v\n", xhs)
	xhs = strings.TrimPrefix(xhs, "sha1=")
	fmt.Printf("### modified github header -> %+v\n", xhs)
	c.Logger().Infof("### modified github header -> %+v\n", xhs)
	req.Header.Set(hmac.Header, xhs)
	b, err := hmac.VerifyRequest(req)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	if !b {
		err = errors.Errorf("could not verify signature with %s", xhs)
		c.Logger().Error(err)
		return c.Error(422, err)
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
