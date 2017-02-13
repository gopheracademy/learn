package models

import (
	"fmt"
	"path/filepath"

	"os"

	"github.com/gobuffalo/envy"
	"github.com/markbates/pop"
)

// DB is a connection to your database to be used
// throughout your application.
var DB *pop.Connection

func init() {
	pop.AddLookupPaths(filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "gopheracademy", "learn"))
	var err error
	env := envy.Get("GO_ENV", "development")
	DB, err = pop.Connect(env)
	if err != nil {
		fmt.Println(err)
	}
	pop.Debug = env == "development"
}
