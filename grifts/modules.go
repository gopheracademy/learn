package grifts

import (
	"os"
	"os/exec"
	"path/filepath"

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

var _ = Desc("sync:modules", "Synchronized the public and private repositories, filtering content that doesn't belong.")
var _ = Add("sync:modules", func(c *Context) error {
	// TODO: clone the repo if it hasn't been already
	err := os.Chdir(models.ModulesPath)
	if err != nil {
		return err
	}
	cmd := exec.Command("git", "pull", "origin", "master")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	// rsync -rtv --exclude 'directory' source_folder/ destination_folder/
	// must be run from the learn directory - which is fine because that's where grifts live anyway
	// (otherwise the git ignore will fail)
	err = os.Chdir(models.PublicModulesPath)
	if err != nil {
		return err
	}
	cmd = exec.Command("rsync", "-rtv", "--delete", "--exclude", filepath.Join("training", ".git"), models.ModulesPath, models.PublicModulesPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	err = os.Chdir(models.PublicModulesPath)
	if err != nil {
		return err
	}
	cmd = exec.Command("git", "add", "--all")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("git", "commit", "-m", "synchronization")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("git", "push", "-u", "origin", "master")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
})
