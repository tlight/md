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

// MarkdownServer contains code & configuration
type MarkdownServer struct {
	Port     string
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

	s := MarkdownServer{*port, filename, "", ""}
	s.Serve()
}

func usage(note string) {
	if len(note) > 0 {
		fmt.Println("Error: " + note)
	}
	fmt.Println(helpText)
	os.Exit(0)
}

// Serve starts the http server
func (s *MarkdownServer) Serve() {
	h := http.NewServeMux()
	h.HandleFunc("/", s.Render)
	log.Printf("Starting Markdown Server for '%s' at http://localhost:%s", s.Filename, s.Port)
	err := http.ListenAndServe(":"+s.Port, h)
	log.Fatal(err)
}

// Render handles the request by converting markdown & rendering html
func (s *MarkdownServer) Render(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET /")
	s.Refresh()
	tmpl, _ := template.New("test").Parse(html)
	tmpl.Execute(w, s)
}

// Refresh updates the file Input Markdown & Output HTML stored in the struct
func (s *MarkdownServer) Refresh() {
	markdown, err := ioutil.ReadFile(s.Filename)
	if err != nil {
		log.Fatal(err)
	}
	s.Markdown = string(markdown)

	html := blackfriday.MarkdownCommon([]byte(s.Markdown))
	s.HTML = template.HTML(html)
}
