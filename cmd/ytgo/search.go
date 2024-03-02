package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"regexp"
	"slices"
)

func fetch(u string) (string, error) {
	res, err := http.Get(u)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func search(query string) (string, error) {
	params := url.Values{"search_query": []string{query}}.Encode()
	return fetch(YtURL + "results?" + params)
}

func getVideos(query string) ([]VID, error) {
	res, err := search(query)
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`(?m)watch\?v=(\S{11})`)
	matches := re.FindAllStringSubmatch(res, -1)
	var v VID
	var vids []VID
	for _, match := range matches {
		v = VID(match[1])
		if !slices.Contains(vids, v) {
			vids = append(vids, v)
		}
	}
	return vids, nil
}

func NthVideo(query string, n int) (VID, error) {
	vids, err := getVideos(query)
	if err != nil {
		return "", err
	}
	if n <= 0 || n > len(vids) {
		return "", errors.New("no video found")
	}
	return vids[n-1], nil
}

func fetchVideoInfo(id VID) (Video, error) {
	j := "%(.{id,title,channel,duration_string,original_url})#j"
	out, err := exec.Command("yt-dlp", "-O", j, YtURL+"watch?v="+string(id)).Output()
	if err != nil {
		return Video{}, err
	}

	var v Video
	err = json.Unmarshal(out, &v)
	if err != nil {
		return Video{}, err
	}

	return v, nil
}
