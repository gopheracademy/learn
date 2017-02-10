package models

import (
	"bytes"
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

var ModulesPath = envy.Get("MODULES_PATH", filepath.Join(envy.Get("GOPATH", ""), "src", "github.com", "gopheracademy", "training"))
var PublicModulesPath = envy.Get("MODULES_PATH", filepath.Join(envy.Get("GOPATH", ""), "src", "github.com", "gopheracademy", "code"))

type Module struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Title     string    `json:"title" db:"title"`
	Path      string    `json:"path" db:"path"`
	Slug      string    `json:"slug" db:"slug"`
	MetaData  MetaData  `json:"metadata" db:"metadata"`
	Slides    Slides    `json:"slides" db:"slides"`
}

// String is not required by pop and may be deleted
func (m Module) String() string {
	b, _ := json.Marshal(m)
	return string(b)
}

func (m Module) Remarkize() string {
	bb := &bytes.Buffer{}

	for i, s := range m.Slides {
		if i != 0 {
			bb.WriteString("\n---\n")
		}
		bb.WriteString(s.MetaData.String())
		bb.WriteString("\n")
		bb.WriteString(s.Content)
		if s.Notes != "" {
			bb.WriteString("\n???\n")
			bb.WriteString(s.Notes)
		}
	}

	return bb.String()
}

func (m Module) Length() time.Duration {
	if t, ok := m.MetaData["module-time"]; ok {
		d, err := time.ParseDuration(t + "m")
		if err != nil {
			return time.Duration(0)
		}
		return d
	}
	return time.Duration(0)
}

// Modules is not required by pop and may be deleted
type Modules []Module

// String is not required by pop and may be deleted
func (m Modules) String() string {
	b, _ := json.Marshal(m)
	return string(b)
}

func (mm Modules) Length() time.Duration {
	l := time.Duration(0)
	for _, m := range mm {
		l += m.Length()
	}
	return l
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

func FindModuleFiles() ([]string, error) {
	files := []string{}
	err := filepath.Walk(ModulesPath, func(path string, info os.FileInfo, err error) error {
		if info != nil && info.Name() == "module.md" {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func RebuildModules() error {
	return DB.Transaction(func(tx *pop.Connection) error {
		ids := []interface{}{}
		files, err := FindModuleFiles()
		if err != nil {
			return err
		}
		for _, path := range files {
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
		// get rid of any modules in the database that no longer exist in the repo
		if len(ids) > 0 {
			err = tx.RawQuery("delete from modules where id not in (?)", ids...).Exec()
			if err != nil {
				return err
			}
			return tx.RawQuery("delete from course_modules where module_id not in (?)", ids...).Exec()
		}
		return nil
	})
}

func buildModule(tx *pop.Connection, path string, in io.Reader) (*Module, error) {
	root := filepath.Dir(path)
	fmt.Printf("Found a module at %s\n", root)

	slug := filepath.Base(root)

	m := &Module{Slug: slug, Path: path}
	b, err := tx.Where("slug = ?", slug).Exists(m)
	if err != nil {
		return m, errors.WithStack(err)
	}
	if b {
		err = tx.Where("slug = ?", slug).First(m)
		if err != nil {
			return m, errors.WithStack(err)
		}
	}

	sp := NewParser(in)
	err = sp.Parse(m)
	if err != nil {
		return m, errors.WithStack(err)
	}

	// throw away the first slide
	m.Slides = m.Slides[1:]

	verrs, err := tx.ValidateAndSave(m)
	if verrs.HasAny() {
		return m, verrs
	}

	return m, errors.WithStack(err)
}
