package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/yuin/goldmark"
)

// MarkdownHandler contains code & configuration
type MarkdownHandler struct {
	Filename string
	Markdown string
	HTML     template.HTML // as converted from Markdown
	Template *template.Template
}

// NewMarkdownHandler returns obj
func NewMarkdownHandler(filename string) *MarkdownHandler {
	t, err := template.New("md").Parse(string(Client))
	if err != nil {
		log.Fatal(err)
	}
	return &MarkdownHandler{
		Filename: filename,
		Template: t,
	}
}

// ServeHTTP implements the Handler interface to respond to request by converting markdown & rendering html
func (s *MarkdownHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET /")
	s.Refresh()
	s.Template.Execute(w, s)
}

// Refresh updates the file Input Markdown & Output HTML stored in the struct
func (s *MarkdownHandler) Refresh() {
	markdown, err := ioutil.ReadFile(s.Filename)
	if err != nil {
		log.Fatal(err)
	}
	s.Markdown = string(markdown)
	s.Markdown = strings.ReplaceAll(s.Markdown, "\r\n", "\n")

	var html bytes.Buffer
	if err := goldmark.Convert([]byte(s.Markdown), &html); err != nil {
		log.Println("ERROR: Unable to parse Markdown")
	}
	s.HTML = template.HTML(html.String())
}
