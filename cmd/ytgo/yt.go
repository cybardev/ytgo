package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ergochat/readline"
)

const VERSION string = "v3.1.4"

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
		home, _ := os.UserHomeDir()

		// create line reader for search
		rl, err = readline.NewFromConfig(&readline.Config{
			Prompt:            fmt.Sprintf("%sSearch:%s ", C_CYAN, C_RESET),
			HistoryFile:       fmt.Sprintf("%s/%s", home, ".ytgo_history"),
			HistoryLimit:      48,
			HistorySearchFold: true,
			VimMode:           true,
		})
		rl.SetVimMode(true)
		if err != nil {
			log.Fatalln(err)
		}
		defer rl.Close()

		goto prompt
	}
	query = strings.Join(flag.Args(), " ")

entrypoint:
	if query == "" {
		if p {
			fmt.Println("No search query provided.")
			goto prompt
		}
		flag.Usage()
		fmt.Println()
		log.Fatalln("no query provided")
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

prompt:
	if p {
		query, err = rl.ReadLine()
		if err != nil {
			return // exit on EOF/SIGINT
		}
		goto entrypoint
	}
}
