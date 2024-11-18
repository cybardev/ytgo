package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/chzyer/readline"
)

const VERSION string = "v3.1.3"

const (
	C_RED   string = "\x1b[31m"
	C_GREEN string = "\x1b[32m"
	C_CYAN  string = "\x1b[36m"
	C_RESET string = "\x1b[00m"
)

func main() {
	var (
		// command-line args
		d, i, m, p, u, ver bool
		n                  int
		// declare vars
		err   error
		query string
		v     *Video
		rl    *readline.Instance
	)

	// parse CLI args
	flag.BoolVar(&ver, "v", false, "Display version")
	flag.BoolVar(&d, "d", false, "Display URL only")
	flag.BoolVar(&i, "i", false, "Interactive selection")
	flag.BoolVar(&m, "m", false, "Play music only")
	flag.BoolVar(&p, "p", false, "Prompt mode")
	flag.BoolVar(&u, "u", false, "Play from URL")
	flag.IntVar(&n, "n", 1, "Play nth media")
	flag.Parse()

	// display version
	if ver {
		fmt.Println(VERSION)
		return
	}

	// get search query
	if p {
		// create line reader for search
		rl, err = readline.New(fmt.Sprintf("%sSearch:%s ", C_CYAN, C_RESET))
		if err != nil {
			log.Fatalln(err)
		}
		defer rl.Close()

		query = getQuery(rl)
	} else {
		query = strings.Join(flag.Args(), " ")
	}

	for {
		for query == "" {
			if !p {
				flag.Usage()
				fmt.Println()
				log.Fatalln("no query provided")
			}
			fmt.Println("No search query provided.")
			query = getQuery(rl)
		}

		// play media from YT or display URL
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
		} else if d {
			fmt.Println(v.Id.URL())
		} else {
			err = v.Play(m)
		}
		if err != nil {
			log.Fatalln(err)
		}

		if !p {
			break
		}
		query = getQuery(rl)
	}
}

func getQuery(r *readline.Instance) string {
	query, err := r.Readline()
	if err != nil {
		os.Exit(0) // exit on EOF/SIGINT
	}
	return query
}
