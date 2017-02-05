package actions

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/buffalo/render/resolvers"
	"github.com/leekchan/accounting"
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

func currencyHelper(price int) string {
	return money.FormatMoney(price / 100)
}

func assetsPath() http.FileSystem {
	box := rice.MustFindBox("../public/assets")
	return box.HTTPBox()
}
