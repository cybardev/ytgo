package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/cybardev/ytgo/internal/app/ytgo"
)

func main() {
	// specify available flags
	var (
		f, m, u bool
		n       int
		query   string
	)

	// parse CLI args
	flag.BoolVar(&f, "f", false, "Play from URL")
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
	var v ytgo.VID
	var err error
	if f {
		v, err = ytgo.VIDfromURL(query)
	} else {
		v, err = ytgo.NthVideo(query, n)
	}
	if err != nil {
		log.Fatalln(err)
	}
	if u {
		fmt.Println(v.URL())
	} else {
		err = v.Play(m)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
