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
	slides := Slides{}
	md, err := ioutil.ReadAll(sp)
	if err != nil {
		return errors.WithStack(err)
	}

	var parsedFirstSlide bool
	lines := bytes.Split(md, []byte("\n"))
	h1 := []byte("# ")
	sep := []byte("---")

	for i := 0; i < len(lines); i++ {
		s := Slide{MetaData: MetaData{}}
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
				k := bytes.TrimSpace(d[0])
				v := bytes.TrimSpace(d[1])
				s.MetaData[string(k)] = string(v)
			}
			i++
			line = lines[i]
		}

		// it's an h1 so it's a new slide
		t := string(bytes.TrimPrefix(line, h1))
		s.Title = t
		if !parsedFirstSlide {
			sp.Module.Title = t
			sp.Module.MetaData = s.MetaData
			parsedFirstSlide = true
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
		slides = append(slides, s)
	}
	sp.Module.Slides = slides
	return nil
}

func NewParser(r io.Reader) SlideParser {
	return SlideParser{
		Reader: r,
		Module: Module{
			Slides:   Slides{},
			MetaData: MetaData{},
		},
	}
}
