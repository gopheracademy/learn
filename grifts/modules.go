package grifts

import (
	"os"
	"os/exec"

	"github.com/gopheracademy/learn/models"
	. "github.com/markbates/grift/grift"
)

var _ = Add("pull:modules", func(c *Context) error {
	err := os.Chdir(models.ModulesPath)
	if err != nil {
		return err
	}
	cmd := exec.Command("git", "pull", "origin", "master")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
})

var _ = Add("build:modules", func(c *Context) error {
	return models.RebuildModules()
})
