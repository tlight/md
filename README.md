# md

## About

Minimal configuration markdown server for local rendering.

## Features

* Fast start with zero or minimal CLI config
* Live reload
* Airplane mode (no external assets)

## Installation

via go get:
```
$ go get github.com/tlight/md
```

via [gobinaries.com](https://gobinaries.com) (if you don't have Go installed):
```
curl -sf https://gobinaries.com/tlight/md | sh
```

## Usage

```sh
Usage: md FILE.md
       md -p 3000 -n 5 FILE.md

    -p, --port           Port to serve from (default 8080)
    -n, --interval       Set update interval in seconds (default 1)
    -h, --help           Output usage information
    -v, --verbose        Enable verbose log output
        --version        Show application version
```

## Examples
```sh
$ md README.md
Starting Markdown Server for 'README.md' at http://localhost:8080

$ md -p 3000 -n 5 -v README.md
Starting Markdown Server for 'README.md' at http://localhost:3000
2020/06/02 12:19:10 GET /
2020/06/02 12:19:10 Refresh Markdown!
2020/06/02 12:19:15 GET /md
```
# License

 MIT
