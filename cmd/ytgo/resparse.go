package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

type VideoRes string // Video Response

func (r VideoRes) Parse() (*Video, error) {
	re := regexp.MustCompile(`var ytInitialPlayerResponse = ({.*?});`)
	s := re.FindStringSubmatch(string(r))[1]
	var j interface{}
	err := json.Unmarshal([]byte(s), &j)
	if err != nil {
		return &Video{}, err
	}
	k := j.(map[string]interface{})["videoDetails"].(map[string]interface{})
	return getVideoFromDetails(&k)
}

func getVideoFromDetails(j *map[string]interface{}) (*Video, error) {
	t, err := strconv.Atoi((*j)["lengthSeconds"].(string))
	if err != nil {
		t = 0
	}
	return &Video{
		Id:       VID((*j)["videoId"].(string)),
		Title:    (*j)["title"].(string),
		Channel:  (*j)["author"].(string),
		Duration: fmt.Sprintf("%d:%d", t/60, t%60),
	}, nil
}

type SearchRes string // Search Response

func (r SearchRes) Parse() (*[]Video, error) {
	re := regexp.MustCompile(`var ytInitialData = ({.*?});`)
	s := re.FindStringSubmatch(string(r))[1]
	var j interface{}
	err := json.Unmarshal([]byte(s), &j)
	if err != nil {
		return nil, err
	}
	res := j.(map[string]interface{})["contents"].(map[string]interface{})["twoColumnSearchResultsRenderer"].(map[string]interface{})["primaryContents"].(map[string]interface{})["sectionListRenderer"].(map[string]interface{})["contents"].([]interface{})[0].(map[string]interface{})["itemSectionRenderer"].(map[string]interface{})["contents"].([]interface{})
	return getVideoList(&res), nil
}

func getVideoList(j *[]interface{}) *[]Video {
	var vs []Video
	for _, i := range *j {
		v, isVideo := getVideoFromEntry(&i)
		if isVideo {
			vs = append(vs, *v)
		}
	}
	return &vs
}

func getVideoFromEntry(j *interface{}) (*Video, bool) {
	k := (*j).(map[string]interface{})["videoRenderer"]
	if k == nil {
		return &Video{}, false // when radioRenderer, shelfRenderer, reelShelfRenderer, etc.
	}
	l := k.(map[string]interface{})
	return &Video{
		Id:       VID(l["videoId"].(string)),
		Title:    l["title"].(map[string]interface{})["runs"].([]interface{})[0].(map[string]interface{})["text"].(string),
		Channel:  l["ownerText"].(map[string]interface{})["runs"].([]interface{})[0].(map[string]interface{})["text"].(string),
		Duration: l["lengthText"].(map[string]interface{})["simpleText"].(string),
	}, true
}
