package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
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
		v.Id, err = VIDfromURL(query)
	} else {
		if l {
			v, err = VideoFromMenu(query)
			if err == nil && (v == Video{}) {
				fmt.Println("No video selected.\nExiting...")
				return
			}
		} else {
			v.Id, err = NthVID(query, n)
		}
	}
	if err != nil {
		log.Fatalln(err)
	}
	if u {
		fmt.Println(v.Id.URL())
	} else {
		err = v.Play(m)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
