package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/yuin/goldmark"
)

// MarkdownHandler contains code & configuration
type MarkdownHandler struct {
	Filename string
	Markdown string
	HTML     template.HTML // as converted from Markdown
	ModTime  time.Time
	Interval int
	Template *template.Template
	Verbose  bool
}

// NewMarkdownHandler returns obj
func NewMarkdownHandler(filename string, interval int, verbose bool) *MarkdownHandler {
	t, err := template.New("md").Parse(string(Client))
	if err != nil {
		log.Fatal(err)
	}
	return &MarkdownHandler{
		Filename: filename,
		Interval: interval,
		Template: t,
		Verbose:  verbose,
	}
}

// ServeHTTP implements the Handler interface to respond to request by converting markdown & rendering html
func (s *MarkdownHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	s.Log("GET /" + path)
	if strings.Compare(path, "md") == 0 {
		if s.IsModified() {
			s.Refresh()
		}
		fmt.Fprint(w, s.HTML)
	} else {
		s.Refresh()
		s.Template.Execute(w, s) // Execute against MarkdownHandler struct members
	}
}

// IsModified checks whether ModTime has changed
func (s *MarkdownHandler) IsModified() bool {
	stat, err := os.Stat(s.Filename)
	if err != nil {
		log.Fatal(err)
	}
	prev := s.ModTime
	curr := stat.ModTime()
	if prev.Equal(curr) {
		return false
	} else {
		s.ModTime = curr
		return true
	}
}

// Refresh updates the input Markdown and converts to output HTML stored in the struct
func (s *MarkdownHandler) Refresh() {
	s.Log("Refresh Markdown!")
	markdown, err := ioutil.ReadFile(s.Filename)
	if err != nil {
		log.Fatal(err)
	}
	s.Markdown = string(markdown)
	s.Markdown = strings.ReplaceAll(s.Markdown, "\r\n", "\n") // Windows
	s.Markdown = strings.ReplaceAll(s.Markdown, "\r", "\n")   // Mac

	var html bytes.Buffer
	if err := goldmark.Convert([]byte(s.Markdown), &html); err != nil {
		s.Log("ERROR: Unable to parse Markdown")
	}
	s.HTML = template.HTML(html.String())
}

func (s *MarkdownHandler) Log(message string) {
	if s.Verbose {
		log.Print(message)
	}
}
