package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
		return tx.RawQuery("delete from modules where id not in (?)", ids...).Exec()
	})
}

func buildModule(tx *pop.Connection, path string, in io.Reader) (Module, error) {
	root := filepath.Dir(path)
	fmt.Printf("Found a module at %s\n", root)
	slug := filepath.Base(root)
	m := Module{Slug: slug, Path: path}
	b, err := tx.Where("slug = ?", slug).Exists(&m)
	if err != nil {
		return m, errors.WithStack(err)
	}
	if b {
		err = tx.Where("slug = ?", slug).First(&m)
		if err != nil {
			return m, errors.WithStack(err)
		}
	}

	md, err := ioutil.ReadAll(in)
	if err != nil {
		return m, errors.WithStack(err)
	}

	var foundTitle bool
	lines := bytes.Split(md, []byte("\n"))
	h1 := []byte("# ")

	m.Slides = Slides{}

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if bytes.HasPrefix(line, h1) {
			// it's an h1 so it's a new slide
			t := string(bytes.TrimPrefix(line, h1))
			s := Slide{Title: t}
			if !foundTitle {
				m.Title = t
				foundTitle = true
			}
			bb := &bytes.Buffer{}
			for {
				// keep reading the contents of this slide
				i++
				if i == len(lines) {
					break
				}
				line := lines[i]
				// found an h1, back up and have a go at the next slide
				if bytes.HasPrefix(line, h1) {
					i--
					break
				}
				bb.Write(line)
				bb.WriteRune('\n')
			}
			s.Content = bb.String()
			m.Slides = append(m.Slides, s)
		}
	}

	verrs, err := tx.ValidateAndSave(&m)
	if verrs.HasAny() {
		return m, verrs
	}

	return m, errors.WithStack(err)
}
