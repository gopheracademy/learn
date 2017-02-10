package actions

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/buffalo/render/resolvers"
	"github.com/gopheracademy/learn/models"
	"github.com/leekchan/accounting"
	"github.com/pkg/errors"
)

var r *render.Engine
var money = accounting.Accounting{Symbol: "$", Precision: 2}

func init() {
	r = render.New(render.Options{
		HTMLLayout:     "application.html",
		CacheTemplates: ENV == "production",
		FileResolverFunc: func() resolvers.FileResolver {
			return &resolvers.RiceBox{
				Box: rice.MustFindBox("../templates"),
			}
		},
		Helpers: map[string]interface{}{
			"currency": currencyHelper,
		},
	})
}

func TrainingAssets(c buffalo.Context) error {
	asset := c.Param("asset")
	if strings.ToLower(filepath.Base(asset)) == "module.md" {
		return c.Error(404, errors.Errorf("could not find %s", asset))
	}
	f, err := os.Open(filepath.Join(models.ModulesPath, asset))
	if err != nil {
		return c.Error(404, errors.Wrapf(err, "could not find %s", asset))
	}
	defer f.Close()
	_, err = io.Copy(c.Response(), f)
	return err
}

func currencyHelper(price int) string {
	return money.FormatMoney(price / 100)
}

func assetsPath() http.FileSystem {
	box := rice.MustFindBox("../public/assets")
	return box.HTTPBox()
}
