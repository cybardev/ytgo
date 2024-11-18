package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/chzyer/readline"
)

const VERSION string = "v3.1.1"

const (
	C_RED   string = "\x1b[31m"
	C_GREEN string = "\x1b[32m"
	C_CYAN  string = "\x1b[36m"
	C_RESET string = "\x1b[00m"
)

func main() {
	// specify available flags
	var (
		d, i, m, p, u, ver bool
		n                  int
		query              string
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

	// declare necessary vars
	var v *Video
	var err error
	var rl *readline.Instance

	// get search query
	if p {
		// create line reader for search
		rl, err = readline.New(fmt.Sprintf("%sSearch:%s ", C_CYAN, C_RESET))
		if err != nil {
			log.Fatalln(err)
		}
		defer rl.Close()

		goto endloop
	} else {
		query = strings.Join(flag.Args(), " ")
	}

loop:
	if query == "" {
		if p {
			fmt.Println("No search query provided.")
			goto endloop
		} else {
			flag.Usage()
			fmt.Println()
			log.Fatalln("no query provided")
		}
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

endloop:
	if p {
		query, err = rl.Readline()
		if err != nil {
			return // exit on EOF/SIGINT
		}
		goto loop
	}
}
