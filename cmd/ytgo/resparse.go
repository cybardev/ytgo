package main

import (
	"encoding/json"
)

type YTRES string

func (r YTRES) Parse() ([]Video, error) {
	var j interface{}
	err := json.Unmarshal([]byte(r), &j)
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

func getVideoFromEntry(i *interface{}) (Video, bool) {
	j := (*i).(map[string]interface{})["videoRenderer"]
	if j == nil {
		return Video{}, false // when radioRenderer, shelfRenderer, reelShelfRenderer, etc.
	}
	k := j.(map[string]interface{})
	return Video{
		Id:       VID(k["videoId"].(string)),
		Title:    k["title"].(map[string]interface{})["runs"].([]interface{})[0].(map[string]interface{})["text"].(string),
		Channel:  k["ownerText"].(map[string]interface{})["runs"].([]interface{})[0].(map[string]interface{})["text"].(string),
		Duration: k["lengthText"].(map[string]interface{})["simpleText"].(string),
	}, true
}
