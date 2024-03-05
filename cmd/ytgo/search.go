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
		return &Video{}, err
	}
	return VRES(res).Parse()
}

func GetVideoFromSearch(query string, n int) (*Video, error) {
	vs, err := GetSearchResults(query)
	if err != nil {
		return &Video{}, err
	}
	if n <= 0 || n > len(*vs) {
		return &Video{}, errors.New("no video found")
	}
	return &(*vs)[n-1], nil
}

func GetSearchResults(query string) (*[]Video, error) {
	params := url.Values{"search_query": []string{query}}.Encode()
	r, err := GetRequest(YtURL + "results?" + params)
	if err != nil {
		return nil, err
	}
	vs, err := YTRES(r).Parse()
	return vs, err
}

func GetRequest(url string) (string, error) {
	res, err := http.Get(url)
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
