package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/envy"
	"github.com/gopheracademy/learn/models"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var rootPath = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "gopheracademy", "learn", "teach", "cmd")
var modules []*models.Module
var moduleMap = map[string]*models.Module{}
var iptFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "teach",
	Short: "Teach a course locally",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if iptFile != "" {
			b, err := ioutil.ReadFile(iptFile)
			if err != nil {
				return err
			}
			files := strings.Split(string(bytes.TrimSpace(b)), "\n")
			for i, f := range files {
				f = filepath.Join(models.ModulesPath, f, "module.md")
				files[i] = f
			}
			return setupModules(files)
		}
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

		port := envy.Get("PORT", ":8000")
		fmt.Printf("Starting teach on %s\n", port)
		log.Fatal(http.ListenAndServe(port, mux))
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

func init() {
	RootCmd.Flags().StringVarP(&iptFile, "file", "f", "", "a file with a list of modules to use.")
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
