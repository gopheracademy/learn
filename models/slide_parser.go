package models

import (
	"bytes"
	"io"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/net/html"

	"github.com/pkg/errors"
)

type SlideParser struct {
	io.Reader
}

func (sp *SlideParser) ParseCode(line []byte) ([]byte, error) {
	bb := &bytes.Buffer{}
	doc, err := html.Parse(bytes.NewReader(line))

	if err != nil {
		return bb.Bytes(), errors.WithStack(err)
	}

	attrs := map[string]string{
		"start": "1",
		"end":   "-1",
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "code" {
			for _, a := range n.Attr {
				attrs[a.Key] = a.Val
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	if src, ok := attrs["src"]; ok {
		sl, err := strconv.Atoi(attrs["start"])
		if err != nil {
			return bb.Bytes(), errors.WithStack(err)
		}
		el, err := strconv.Atoi(attrs["end"])
		if err != nil {
			return bb.Bytes(), errors.WithStack(err)
		}

		fb, err := ioutil.ReadFile(filepath.Join(ModulesPath, src))
		if err != nil {
			return bb.Bytes(), errors.WithStack(err)
		}

		ext := filepath.Ext(src)
		bb.WriteString("```")
		bb.WriteString(strings.TrimPrefix(ext, "."))
		nl := []byte("\n")
		bb.Write(nl)

		fb = bytes.TrimSpace(fb)
		lines := bytes.Split(fb, nl)
		sl -= 1
		if el == -1 {
			el = len(lines)
		}
		lines = lines[sl:el]
		bb.Write(bytes.Join(lines, nl))

		bb.WriteString("\n```")
	}
	return bb.Bytes(), err
}

func (sp *SlideParser) Parse(m *Module) error {
	slides := Slides{}
	md, err := ioutil.ReadAll(sp)
	if err != nil {
		return errors.WithStack(err)
	}

	var parsedFirstSlide bool
	lines := bytes.Split(md, []byte("\n"))
	h1 := []byte("# ")
	sep := []byte("---")
	code := []byte("<code")
	pres := []byte("???")

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
			m.Title = t
			m.MetaData = s.MetaData
			parsedFirstSlide = true
		}
		bb := &bytes.Buffer{}
		bb.Write(line)
		bb.WriteRune('\n')
		for {
			// keep reading the contents of this slide
			i++
			if i == len(lines) {
				break
			}
			line := lines[i]
			// found a line of code that needs injecting
			if bytes.HasPrefix(line, code) {
				line, err = sp.ParseCode(line)
				if err != nil {
					return err
				}
			}

			// found a ??? for presenter notes
			if bytes.HasPrefix(line, pres) {
				pp := &bytes.Buffer{}
				for {
					i++
					line = lines[i]

					// found an ---, back up and have a go at the next slide
					if bytes.HasPrefix(line, sep) {
						s.Notes = pp.String()
						i--
						break
					}
					pp.Write(line)
					pp.WriteRune('\n')
				}
			}

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
	m.Slides = slides
	return nil
}

func NewParser(r io.Reader) SlideParser {
	return SlideParser{
		Reader: r,
	}
}
