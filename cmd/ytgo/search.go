package main

import (
	"errors"
	"io"
	"net/http"
	"net/url"
)

func GetVideoFromSearch(query string, n int) (Video, error) {
	vids, err := GetSearchResults(query)
	if err != nil {
		return Video{}, err
	}
	if n <= 0 || n > len(vids) {
		return Video{}, errors.New("no video found")
	}
	return vids[n-1], nil
}

func GetSearchResults(query string) ([]Video, error) {
	params := url.Values{"search_query": []string{query}}.Encode()
	r, err := GetRequest(YtURL + "results?" + params)
	if err != nil {
		return nil, err
	}
	return YTRES(r).Parse()
}

func GetRequest(u string) (string, error) {
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
