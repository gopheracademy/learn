package grifts

import (
	"os"
	"os/exec"

	"github.com/gopheracademy/learn/models"
	. "github.com/markbates/grift/grift"
)

var _ = Desc("pull:modules", "Runs a git pull origin master on the `models.ModulesPath` repo")
var _ = Add("pull:modules", func(c *Context) error {
	// TODO: clone the repo if it hasn't been already
	err := os.Chdir(models.ModulesPath)
	if err != nil {
		return err
	}
	cmd := exec.Command("git", "pull", "origin", "master")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
})

var _ = Desc("build:modules", "Rebuilds the modules in the database from the files at `models.ModulesPath`")
var _ = Add("build:modules", func(c *Context) error {
	return models.RebuildModules()
})
