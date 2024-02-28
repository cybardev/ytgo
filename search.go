package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"slices"
)

func fetch(u string) string {
	res, err := http.Get(u)
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

func search(query string) string {
	params := url.Values{"search_query": []string{query}}.Encode()
	return fetch("https://www.youtube.com/results?" + params)
}

func getVideos(query string) []VID {
	res := search(query)
	re := regexp.MustCompile(`(?m)watch\?v=(\S{11})`)
	matches := re.FindAllStringSubmatch(res, -1)
	var vids []VID
	for _, match := range matches {
		if !slices.Contains(vids, VID(match[1])) {
			vids = append(vids, VID(match[1]))
		}
	}
	return vids
}

func nthVideo(query string, n int) VID {
	vids := getVideos(query)
	if 0 >= n || n > len(vids) {
		log.Fatalln("No video found")
	}
	return vids[n-1]
}
