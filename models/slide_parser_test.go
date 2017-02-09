package models

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SlideParser(t *testing.T) {
	r := require.New(t)

	f, err := os.Open("./slide_parser_test.md")
	r.NoError(err)
	p := NewParser(f)
	err = p.Parse()
	r.NoError(err)

	m := p.Module
	r.Equal("Concurrency", m.Title)
	r.Len(m.Slides, 10)
	r.Equal(m.Slides[0].Title, "Goroutines")
}
