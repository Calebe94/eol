package main

import (
	"time"
)

type Cycle struct {
	Cycle        interface{} `json:"cycle"`
	ReleaseDate  string      `json:"releaseDate"`
	EOL          interface{} `json:"eol"`
	Latest       string      `json:"latest"`
	Link         string      `json:"link"`
	LTS          interface{} `json:"lts"`
	Support      interface{} `json:"support"`
	Discontinued interface{} `json:"discontinued"`
}

type CLI struct {
	format  string
	output  string
	watch   time.Duration
	noColor bool
	filter  string
}
