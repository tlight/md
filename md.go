package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var helpText = `
Usage: md FILE.md
       md --port 3000 README.md

	--port          Port to serve from
	--help	        Show this help screen
`

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
		usage("Please limit to a single file")
	}
	filename := args[0]

	// Serve MarkdownHandler
	log.Printf("Starting Markdown Server for '%s' at http://localhost:%s", filename, *port)
	http.Handle("/", NewMarkdownHandler(filename))
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
