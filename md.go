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
	Markdown server:

	--help		Show this help
	--port		Port to serve from
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
			{{.Output}}
		</article>
	</body>
</html>
`

// MarkdownServer contains code & configuration
type MarkdownServer struct {
	Port     string
	Filename string
	Input    string
	Output   template.HTML
}

func main() {
	s := MarkdownServer{}
	help := flag.Bool("help", false, "show help")

	flag.StringVar(&s.Port, "port", "8080", "Server port")
	flag.Parse()

	args := flag.Args()
	if *help {
		fmt.Println(helpText)
		os.Exit(0)
	}

	if len(args) == 0 {
		log.Fatal("Error: Please provide a filename")
	}
	if len(args) > 1 {
		log.Fatal("Error: Please provide a single filename to serve")
	}
	s.Filename = args[0]
	s.Serve()
}

// Serve starts the http server
func (s *MarkdownServer) Serve() {
	h := http.NewServeMux()
	h.HandleFunc("/", s.Render)
	log.Printf("Serving (%s) at http://localhost:%s", s.Filename, s.Port)
	err := http.ListenAndServe(":"+s.Port, h)
	log.Fatal(err)
}

// Render handles the request by converting markdown & rendering html
func (s *MarkdownServer) Render(w http.ResponseWriter, r *http.Request) {
	log.Printf("Server: refreshing")
	s.Refresh()

	output := blackfriday.MarkdownCommon([]byte(s.Input))
	s.Output = template.HTML(output)

	tmpl, _ := template.New("test").Parse(html)
	tmpl.Execute(w, s)
	//fmt.Fprintf(w, s.output)
}

// Refresh updates the file stored in the struct
func (s *MarkdownServer) Refresh() {
	input, err := ioutil.ReadFile(s.Filename)
	if err != nil {
		log.Fatal(err)
	}
	s.Input = string(input)
}
