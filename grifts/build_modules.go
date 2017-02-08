package grifts

import (
	"github.com/gopheracademy/learn/models"
	. "github.com/markbates/grift/grift"
)

var _ = Add("build:modules", func(c *Context) error {
	return models.RebuildModules()
})
