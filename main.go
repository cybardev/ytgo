package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"regexp"
	"slices"
)

func main() {
	play(videoURL(getVideos("gurenge band cover")[0]))
}

func play(url string) {
	cmd := exec.Command("mpv", url)
	cmd.Run()
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

	// Convert the body to type string
	msg := string(body)
	// log.Println(msg)
	return msg
}
