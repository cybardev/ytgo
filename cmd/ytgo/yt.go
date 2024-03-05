package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

const (
	C_RED   string = "\x1b[31m"
	C_GREEN string = "\x1b[32m"
	C_CYAN  string = "\x1b[36m"
	C_RESET string = "\x1b[00m"
)

func main() {
	// specify available flags
	var (
		f, l, m, u bool
		n          int
		query      string
	)

	// parse CLI args
	flag.BoolVar(&f, "f", false, "Play from URL")
	flag.BoolVar(&l, "l", false, "Select from list")
	flag.BoolVar(&m, "m", false, "Play music only")
	flag.BoolVar(&u, "u", false, "Display URL only")
	flag.IntVar(&n, "n", 1, "Play nth media")
	flag.Parse()

	query = strings.Join(flag.Args(), " ")
	if query == "" {
		flag.Usage()
		fmt.Println()
		log.Fatalln("no query provided")
	}

	// play media from YT or display URL
	var v Video
	var err error
	if f {
		v, err = GetVideoFromURL(query)
	} else if l {
		v, err = GetVideoFromMenu(query)
		if err == nil && (v == Video{}) {
			return
		}
	} else {
		v, err = GetVideoFromSearch(query, n)
	}
	if err != nil {
		log.Fatalln(err)
	}
	if u {
		fmt.Println(v.Id.URL())
		return
	}
	err = v.Play(m)
	if err != nil {
		log.Fatalln(err)
	}
}
