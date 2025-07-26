package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type VideoRes string // Video Response

func (r VideoRes) Parse() (*Video, error) {
	re := regexp.MustCompile(`var ytInitialPlayerResponse = ({.*?});`)
	s := re.FindStringSubmatch(string(r))[1]
	var j any
	err := json.Unmarshal([]byte(s), &j)
	if err != nil {
		return nil, err
	}
	k, ok := j.(map[string]any)["videoDetails"].(map[string]any)
	if ok {
		return getVideoFromDetails(&k)
	} else {
		return &Video{}, errors.New("interface type mismatch")
	}
}

func getVideoFromDetails(j *map[string]any) (*Video, error) {
	t, err := strconv.Atoi((*j)["lengthSeconds"].(string))
	if err != nil {
		t = 0
	}
	hh, mm, ss := t/3600, (t%3600)/60, t%60
	var tf string
	if hh > 0 {
		tf = fmt.Sprintf("%d:%02d:%02d", hh, mm, ss)
	} else {
		tf = fmt.Sprintf("%d:%02d", mm, ss)
	}
	return &Video{
		Id:       VID((*j)["videoId"].(string)),
		Title:    (*j)["title"].(string),
		Channel:  (*j)["author"].(string),
		Duration: tf,
	}, nil
}

type SearchRes string // Search Response

func (r SearchRes) Parse() ([]Video, error) {
	re := regexp.MustCompile(`var ytInitialData = ({.*?});`)
	s := re.FindStringSubmatch(string(r))[1]
	var j any
	err := json.Unmarshal([]byte(s), &j)
	if err != nil {
		return nil, err
	}
	res := j.(map[string]any)["contents"].(map[string]any)["twoColumnSearchResultsRenderer"].(map[string]any)["primaryContents"].(map[string]any)["sectionListRenderer"].(map[string]any)["contents"].([]any)[0].(map[string]any)["itemSectionRenderer"].(map[string]any)["contents"].([]any)
	return getVideoList(&res), nil
}

func getVideoList(j *[]any) []Video {
	var vs []Video
	for _, i := range *j {
		v, isVideo := getVideoFromEntry(&i)
		if isVideo {
			vs = append(vs, *v)
		}
	}
	return vs
}

func getVideoFromEntry(j *any) (*Video, bool) {
	k := (*j).(map[string]any)["videoRenderer"]
	if k == nil {
		return nil, false // when radioRenderer, shelfRenderer, reelShelfRenderer, etc.
	}
	l := k.(map[string]any)
	return &Video{
		Id:       VID(l["videoId"].(string)),
		Title:    l["title"].(map[string]any)["runs"].([]any)[0].(map[string]any)["text"].(string),
		Channel:  l["ownerText"].(map[string]any)["runs"].([]any)[0].(map[string]any)["text"].(string),
		Duration: l["lengthText"].(map[string]any)["simpleText"].(string),
	}, true
}
