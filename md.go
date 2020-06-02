package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const VERSION = "v1.2.0"

var helpText = `
Usage: md FILE.md
       md -p 3000 -n 5 FILE.md

    -p, --port           Port to serve from (default 8080)
    -n, --interval       Set update interval in seconds (default 1)
    -h, --help           Output usage information
    -v, --verbose        Enable verbose log output
        --version        Show application version
`

// entry point & validation
func main() {
	var port, interval int
	var help, verbose, version bool

	flag.IntVar(&port, "port", 8080, "Port to serve from (default 8080)")
	flag.IntVar(&port, "p", 8080, "Port to serve from (default 8080)")
	flag.IntVar(&interval, "interval", 1, "Update interval in seconds (default 1)")
	flag.IntVar(&interval, "n", 1, "Update interval in seconds (default 1)")

	flag.BoolVar(&help, "help", false, "Output usage information")
	flag.BoolVar(&help, "h", false, "Output usage information")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose log output")
	flag.BoolVar(&verbose, "v", false, "Enable verbose log output")
	flag.BoolVar(&version, "version", false, "Show application version")
	flag.Parse()
	args := flag.Args()

	if help {
		usage("")
	} else if version {
		fmt.Printf("md version %s\n", VERSION)
	} else if len(args) == 0 {
		usage("Please provide a file as an argument e.g. README.md")
	} else if len(args) > 1 {
		usage("Please limit to a single file")
	} else {
		filename := args[0]

		// Serve MarkdownHandler
		log.Printf("Starting Markdown Server for '%s' at http://localhost:%d", filename, port)
		http.Handle("/", NewMarkdownHandler(filename, port, verbose))
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		log.Fatal(err)
	}
}

func usage(note string) {
	if len(note) > 0 {
		fmt.Println("Error: " + note)
	}
	fmt.Println(helpText)
	os.Exit(0)
}
