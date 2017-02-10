package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gobuffalo/envy"
	"github.com/gopheracademy/learn/models"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var rootPath = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "gopheracademy", "learn", "teach", "cmd")
var modules []*models.Module
var moduleMap = map[string]*models.Module{}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "teach",
	Short: "Teach a course locally",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	PreRunE: func(cmd *cobra.Command, args []string) error {
		files, err := models.FindModuleFiles()
		if err != nil {
			return err
		}
		return setupModules(files)
	},
	Run: func(cmd *cobra.Command, args []string) {
		mux := mux.NewRouter()
		a := http.StripPrefix("/assets/", http.FileServer(http.Dir(filepath.Join(rootPath, "assets"))))
		mux.PathPrefix("/assets/").Handler(a)
		ta := http.StripPrefix("/training-assets/", http.FileServer(http.Dir(models.ModulesPath)))
		mux.PathPrefix("/training-assets/").Handler(ta)
		mux.HandleFunc("/", IndexHandler)
		mux.HandleFunc("/module/{name}", ModuleHandler)
		log.Fatal(http.ListenAndServe(envy.Get("PORT", ":8000"), mux))
	},
}

func setupModules(files []string) error {
	for _, path := range files {
		root := filepath.Dir(path)
		slug := filepath.Base(root)
		in, err := os.Open(path)
		if err != nil {
			return errors.WithStack(err)
		}

		m := &models.Module{Slug: slug, Path: path}
		sp := models.NewParser(in)
		err = sp.Parse(m)
		if err != nil {
			return errors.WithStack(err)
		}
		modules = append(modules, m)
		moduleMap[m.Slug] = m
	}
	return nil
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
