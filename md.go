package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/russross/blackfriday"
)

var helpText = `
Usage: md FILE.md
       md --port 3000 README.md
			 
	--port		Port to serve from
	--help		Show this help screen
`

var html = `
<html>
	<head>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/github-markdown-css/2.10.0/github-markdown.min.css">
		<style>
			.markdown-body {
				box-sizing: border-box;
				min-width: 200px;
				max-width: 980px;
				margin: 0 auto;
				padding: 45px;
			}
			@media (max-width: 767px) {
				.markdown-body {
					padding: 15px;
				}
			}
		</style>
	</head>
	<body>
		<article class="markdown-body">
			{{.HTML}}
		</article>
	</body>
</html>
`

// MarkdownHandler contains code & configuration
type MarkdownHandler struct {
	Filename string
	Markdown string
	HTML     template.HTML
}

// entry point & validation
func main() {
	help := flag.Bool("help", false, "show help")
	port := flag.String("port", "8080", "Server port")
	flag.Parse()
	args := flag.Args()

	if *help {
		usage("")
	} else if len(args) == 0 {
		usage("Please provide a file as an argument e.g. README.md")
	} else if len(args) > 1 {
		usage("Provide limit to single files")
	}
	filename := args[0]

	// Serve MarkdownHandler
	log.Printf("Starting Markdown Server for '%s' at http://localhost:%s", filename, *port)
	http.Handle("/", &MarkdownHandler{filename, "", ""})
	err := http.ListenAndServe(":"+*port, nil)
	log.Fatal(err)
}

func usage(note string) {
	if len(note) > 0 {
		fmt.Println("Error: " + note)
	}
	fmt.Println(helpText)
	os.Exit(0)
}

// ServeHTTP implements the Handler interface to respond to request by converting markdown & rendering html
func (s *MarkdownHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET /")
	s.Refresh()
	t, err := template.New("md").Parse(html)
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, s)
}

// Refresh updates the file Input Markdown & Output HTML stored in the struct
func (s *MarkdownHandler) Refresh() {
	markdown, err := ioutil.ReadFile(s.Filename)
	if err != nil {
		log.Fatal(err)
	}
	s.Markdown = string(markdown)

	html := blackfriday.MarkdownCommon([]byte(s.Markdown))
	s.HTML = template.HTML(html)
}
