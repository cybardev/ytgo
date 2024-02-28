package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	// specify available flags
	var query string
	var n int
	var f bool
	var m bool
	var u bool

	// parse CLI args
	flag.IntVar(&n, "n", 1, "Play nth media")
	flag.BoolVar(&f, "f", false, "Play from URL")
	flag.BoolVar(&m, "m", false, "Play music only")
	flag.BoolVar(&u, "u", false, "Display URL only")
	flag.Parse()

	query = strings.Join(flag.Args(), " ")
	if query == "" {
		flag.Usage()
		fmt.Println()
		log.Fatalln("No query provided")
	}

	// play media from YT or display URL
	var v VID
	if f {
		v = VIDfromURL(query)
	} else {
		v = nthVideo(query, n)
	}
	if u {
		fmt.Println(v.url())
	} else {
		v.play(m)
	}
}
