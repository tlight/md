# md

## About

Zero configuration minimal markdown server for local rendering.

## Features

* Fast start with zero or minimal CLI config
* Live reload
* Airplane mode (no external assets)

## Installation

via go get:
```
$ go get github.com/tlight/md
```

via gobinaries (if you don't have Go installed):
```
curl -sf https://gobinaries.com/tlight/md | sh
```

## Usage

```sh
$ md README.md
Starting Markdown Server for 'README.md' at http://localhost:8080

$ md README.md --port 3000
Starting Markdown Server for 'README.md' at http://localhost:3000
```

# License

 MIT
