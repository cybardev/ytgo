package main

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

func GetNthVideo(query string, n int) (Video, error) {
	vids, err := getVideos(query)
	if err != nil {
		return Video{}, err
	}
	if n <= 0 || n > len(vids) {
		return Video{}, errors.New("no video found")
	}
	return vids[n-1], nil
}

func getVideos(query string) ([]Video, error) {
	res, err := getSearchResults(query)
	if err != nil {
		return nil, err
	}
	return res.Parse()
}

func getSearchResults(query string) (YTRES, error) {
	params := url.Values{"search_query": []string{query}}.Encode()
	s, err := getRequest(YtURL + "results?" + params)
	if err != nil {
		return YTRES(""), err
	}
	re := regexp.MustCompile(`var ytInitialData = ({.*?});`)
	return YTRES(re.FindStringSubmatch(s)[1]), nil
}

func getRequest(u string) (string, error) {
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
