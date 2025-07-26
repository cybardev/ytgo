package main

import (
	"errors"
	"io"
	"net/http"
	"net/url"
)

func GetVideoFromURL(url string) (*Video, error) {
	res, err := GetRequest(url)
	if err != nil {
		return nil, err
	}
	return VideoRes(res).Parse()
}

func GetVideoFromSearch(query string, n int) (*Video, error) {
	vs, err := GetSearchResults(query)
	if err != nil {
		return nil, err
	}
	if n <= 0 || n > len(vs) {
		return nil, errors.New("no video found")
	}
	return &(vs)[n-1], nil
}

func GetSearchResults(query string) ([]Video, error) {
	params := url.Values{"search_query": []string{query}}.Encode()
	res, err := GetRequest(YtURL + "results?" + params)
	if err != nil {
		return nil, err
	}
	vs, err := SearchRes(res).Parse()
	return vs, err
}

func GetRequest(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", errors.New("HTTP response status not OK")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
