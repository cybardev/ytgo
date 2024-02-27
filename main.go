package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"slices"
)

func main() {
	// specify available flags
	var query string
	var n int
	var m bool

	// parse CLI args
	flag.BoolVar(&m, "m", false, "Play music only")
	flag.IntVar(&n, "n", 1, "Play nth media")
	flag.Parse()

	query = flag.Arg(0)
	if query == "" {
		log.Println(query)
		log.Fatalln("No query provided")
	}

	// play media from YT
	playNth(query, n, m)
}

func playNth(query string, n int, m bool) {
	vids := getVideos(query)
	if 0 > n || n > len(vids) {
		log.Fatalln("No video found")
	}
	play(videoURL(vids[n-1]), m)
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
