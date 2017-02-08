package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gobuffalo/envy"
	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

type Module struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Title     string    `json:"title" db:"title"`
	Path      string    `json:"path" db:"path"`
	Slug      string    `json:"slug" db:"slug"`
	Slides    Slides    `json:"slides" db:"slides"`
}

// String is not required by pop and may be deleted
func (m Module) String() string {
	b, _ := json.Marshal(m)
	return string(b)
}

// Modules is not required by pop and may be deleted
type Modules []Module

// String is not required by pop and may be deleted
func (m Modules) String() string {
	b, _ := json.Marshal(m)
	return string(b)
}

// Validate gets run everytime you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (m *Module) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: m.Title, Name: "Title"},
		&validators.StringIsPresent{Field: m.Path, Name: "Path"},
		&validators.StringIsPresent{Field: m.Slug, Name: "Slug"},
	), nil

}

// ValidateSave gets run everytime you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (m *Module) ValidateSave(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run everytime you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (m *Module) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func RebuildModules() error {
	return DB.Transaction(func(tx *pop.Connection) error {
		ids := []interface{}{}
		root := envy.Get("MODULES_PATH", filepath.Join(envy.Get("GOPATH", ""), "src", "github.com", "gopheracademy", "training"))
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if info != nil && info.Name() == "module.md" {
				f, err := os.Open(path)
				if err != nil {
					return errors.WithStack(err)
				}
				defer f.Close()
				m, err := buildModule(tx, path, f)
				if err != nil {
					return errors.WithStack(err)
				}
				ids = append(ids, m.ID.String())
			}
			return nil
		})
		if err != nil {
			return err
		}
		// get rid of any modules in the database that no longer exist in the repo
		if len(ids) > 0 {
			return tx.RawQuery("delete from modules where id not in (?)", ids...).Exec()
		}
		return nil
	})
}

func buildModule(tx *pop.Connection, path string, in io.Reader) (Module, error) {
	root := filepath.Dir(path)
	fmt.Printf("Found a module at %s\n", root)

	slug := filepath.Base(root)

	sp := NewParser(in)
	sp.Module = Module{Slug: slug, Path: path}
	b, err := tx.Where("slug = ?", slug).Exists(&sp.Module)
	if err != nil {
		return sp.Module, errors.WithStack(err)
	}
	if b {
		err = tx.Where("slug = ?", slug).First(&sp.Module)
		if err != nil {
			return sp.Module, errors.WithStack(err)
		}
	}

	err = sp.Parse()
	if err != nil {
		return sp.Module, errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndSave(&sp.Module)
	if verrs.HasAny() {
		return sp.Module, verrs
	}

	return sp.Module, errors.WithStack(err)
}
