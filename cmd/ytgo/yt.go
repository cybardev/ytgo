package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ergochat/readline"
)

const VERSION string = "v3.3.1"

const (
	C_RED    string = "\x1b[31m"
	C_GREEN  string = "\x1b[32m"
	C_YELLOW string = "\x1b[33m"
	C_CYAN   string = "\x1b[36m"
	C_RESET  string = "\x1b[00m"
)

func main() {
	var (
		// command-line args
		d, i, m, p, u, ver bool
		n                  int
		f                  string
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
	flag.StringVar(&f, "f", "", "Play from playlist file")
	flag.IntVar(&n, "n", 1, "Play nth media")
	flag.Parse()

	// display version
	if ver {
		fmt.Println(VERSION)
		return
	}

	// playlist mode
	if f != "" {
		_, err := os.Stat(f) // check if file exists
		p := Playlist(f)
		if errors.Is(err, os.ErrNotExist) {
			err = p.Create()
		} else {
			err = p.Play(m)
		}
		if err != nil {
			log.Fatalln(err)
		}
		return // exit program once playlist is handled
	}

	// get search query
	if p {
		rl = GetReadline()
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
		vs, err := GetSearchResults(query)
		if err == nil {
			v, err = GetVideoFromMenu(vs)
			if v == BACK_FLAG {
				goto prompt
			}
		}
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
		fmt.Println("Playing:", v)
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

func GetReadline() *readline.Instance {
	home, _ := os.UserHomeDir()
	rl, err := readline.NewFromConfig(&readline.Config{
		Prompt:            fmt.Sprintf("%sSearch:%s ", C_CYAN, C_RESET),
		HistoryFile:       fmt.Sprintf("%s/%s", home, ".ytgo_history"),
		HistoryLimit:      48,
		HistorySearchFold: true,
		VimMode:           true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer rl.Close()
	return rl
}
