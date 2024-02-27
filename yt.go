package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"slices"
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
	flag.BoolVar(&f, "f", false, "Play from URL")
	flag.BoolVar(&u, "u", false, "Display URL only")
	flag.BoolVar(&m, "m", false, "Play music only")
	flag.IntVar(&n, "n", 1, "Play nth media")
	flag.Parse()

	query = strings.Join(flag.Args(), " ")
	if query == "" {
		flag.Usage()
		fmt.Println()
		log.Fatalln("No query provided")
	}

	// play media from YT or display URL
	if u {
		fmt.Println(nthVideo(query, n))
	} else {
		if f {
			play(query, m)
		} else {
			play(nthVideo(query, n), m)
		}
	}
}

func play(url string, m bool) {
	bestaudio, novideo := "", ""
	if m {
		bestaudio, novideo = "--ytdl-format=bestaudio", "--no-video"
	}
	cmd := exec.Command("mpv", bestaudio, novideo, url)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func nthVideo(query string, n int) string {
	vids := getVideos(query)
	if 0 >= n || n > len(vids) {
		log.Fatalln("No video found")
	}
	return videoURL(vids[n-1])
}

func videoURL(id string) string {
	return "https://www.youtube.com/watch?v=" + id
}

func getVideos(query string) []string {
	res := search(query)
	re := regexp.MustCompile(`(?m)watch\?v=(\S{11})`)
	matches := re.FindAllStringSubmatch(res, -1)
	var vids []string
	for _, match := range matches {
		if !slices.Contains(vids, match[1]) {
			vids = append(vids, match[1])
		}
	}
	return vids
}

func search(query string) string {
	params := url.Values{"search_query": []string{query}}.Encode()
	return fetch("https://www.youtube.com/results?" + params)
}

func fetch(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	// We Read the response body on the line below.
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Return the body as string
	msg := string(body)
	return msg
}
