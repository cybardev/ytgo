package ytgo

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"slices"

	"github.com/cybardev/ytgo/internal/pkg/vid"
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
	return fetch(vid.YtURL + "results?" + params)
}

func getVideos(query string) ([]vid.VID, error) {
	res, err := search(query)
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`(?m)watch\?v=(\S{11})`)
	matches := re.FindAllStringSubmatch(res, -1)
	var v vid.VID
	var vids []vid.VID
	for _, match := range matches {
		v = vid.VID(match[1])
		if !slices.Contains(vids, v) {
			vids = append(vids, v)
		}
	}
	return vids, nil
}

func NthVideo(query string, n int) (vid.VID, error) {
	vids, err := getVideos(query)
	if err != nil {
		return "", err
	}
	if n <= 0 || n > len(vids) {
		return "", errors.New("no video found")
	}
	return vids[n-1], nil
}
