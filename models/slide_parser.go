package models

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/pkg/errors"
)

type SlideParser struct {
	io.Reader
	Module Module
}

func (sp *SlideParser) Parse() error {
	md, err := ioutil.ReadAll(sp)
	if err != nil {
		return errors.WithStack(err)
	}

	var foundTitle bool
	lines := bytes.Split(md, []byte("\n"))
	h1 := []byte("# ")
	sep := []byte("---")

	for i := 0; i < len(lines); i++ {
		s := Slide{MetaData: map[string]string{}}
		line := lines[i]
		// found a new slide:
		for {
			if bytes.HasPrefix(line, h1) {
				// stop reading metadata and break
				break
			}
			// read the metadata:
			d := bytes.Split(line, []byte(":"))
			if len(d) >= 2 {
				s.MetaData[string(d[0])] = string(d[1])
			}
			i++
			line = lines[i]
		}

		// it's an h1 so it's a new slide
		t := string(bytes.TrimPrefix(line, h1))
		s.Title = t
		if !foundTitle {
			sp.Module.Title = t
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
			// found an ---, back up and have a go at the next slide
			if bytes.HasPrefix(line, sep) {
				i--
				break
			}
			bb.Write(line)
			bb.WriteRune('\n')
		}
		s.Content = bb.String()
		sp.Module.Slides = append(sp.Module.Slides, s)
	}
	return nil
}

func NewParser(r io.Reader) SlideParser {
	return SlideParser{
		Reader: r,
		Module: Module{
			Slides: Slides{},
		},
	}
}
