package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

const VERSION string = "v3.0.12"

const (
	C_RED   string = "\x1b[31m"
	C_GREEN string = "\x1b[32m"
	C_CYAN  string = "\x1b[36m"
	C_RESET string = "\x1b[00m"
)

func main() {
	// specify available flags
	var (
		d, i, m, u, ver bool
		n               int
		query           string
	)

	// parse CLI args
	flag.BoolVar(&ver, "v", false, "Display version")
	flag.BoolVar(&d, "d", false, "Display URL only")
	flag.BoolVar(&i, "i", false, "Interactive selection")
	flag.BoolVar(&m, "m", false, "Play music only")
	flag.BoolVar(&u, "u", false, "Play from URL")
	flag.IntVar(&n, "n", 1, "Play nth media")
	flag.Parse()

	// display version
	if ver {
		fmt.Println(VERSION)
		return
	}

	// get search query
	query = strings.Join(flag.Args(), " ")
	if query == "" {
		flag.Usage()
		fmt.Println()
		log.Fatalln("no query provided")
	}

	// play media from YT or display URL
	var v *Video
	var err error
	if u {
		v, err = GetVideoFromURL(query)
	} else if i {
		v, err = GetVideoFromMenu(query)
	} else {
		v, err = GetVideoFromSearch(query, n)
	}
	if err != nil {
		log.Fatalln(err)
	} else if v == nil {
	    return
	}
	if d {
		fmt.Println(v.Id.URL())
		return
	}
	err = v.Play(m)
	if err != nil {
		log.Fatalln(err)
	}
}
