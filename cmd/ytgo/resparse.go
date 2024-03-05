package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

type VRES string // video response

func (r VRES) Parse() (Video, error) {
	re := regexp.MustCompile(`var ytInitialPlayerResponse = ({.*?});`)
	s := re.FindStringSubmatch(string(r))[1]
	var j interface{}
	err := json.Unmarshal([]byte(s), &j)
	if err != nil {
		return Video{}, err
	}
	k := j.(map[string]interface{})["videoDetails"].(map[string]interface{})
	return getVideoFromDetails(&k)
}

func getVideoFromDetails(d *map[string]interface{}) (Video, error) {
	t, err := strconv.Atoi((*d)["lengthSeconds"].(string))
	if err != nil {
		t = 0
	}
	return Video{
		Id:       VID((*d)["videoId"].(string)),
		Title:    (*d)["title"].(string),
		Channel:  (*d)["author"].(string),
		Duration: fmt.Sprintf("%d:%d", t/60, t%60),
	}, nil
}

type YTRES string // search response

func (r YTRES) Parse() ([]Video, error) {
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

func getVideoList(d *[]interface{}) []Video {
	var vids []Video
	for _, i := range *d {
		v, isVideo := getVideoFromEntry(&i)
		if isVideo {
			vids = append(vids, v)
		}
	}
	return vids
}

func getVideoFromEntry(j *interface{}) (Video, bool) {
	k := (*j).(map[string]interface{})["videoRenderer"]
	if k == nil {
		return Video{}, false // when radioRenderer, shelfRenderer, reelShelfRenderer, etc.
	}
	l := k.(map[string]interface{})
	return Video{
		Id:       VID(l["videoId"].(string)),
		Title:    l["title"].(map[string]interface{})["runs"].([]interface{})[0].(map[string]interface{})["text"].(string),
		Channel:  l["ownerText"].(map[string]interface{})["runs"].([]interface{})[0].(map[string]interface{})["text"].(string),
		Duration: l["lengthText"].(map[string]interface{})["simpleText"].(string),
	}, true
}
