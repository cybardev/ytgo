package main

import "encoding/json"

type YTRES string

func (r YTRES) Parse() []Video {
	var j interface{}
	json.Unmarshal([]byte(r), &j)
	res := j.(map[string]interface{})["contents"].(map[string]interface{})["twoColumnSearchResultsRenderer"].(map[string]interface{})["primaryContents"].(map[string]interface{})["sectionListRenderer"].(map[string]interface{})["contents"].([]interface{})[0].(map[string]interface{})["itemSectionRenderer"].(map[string]interface{})["contents"].([]interface{})
	return getVideoList(&res)
}

func getVideoList(d *[]interface{}) []Video {
	var vids []Video
	for _, i := range *d {
		v, ok := getVideoFromEntry(&i)
		if !ok {
			continue
		}
		vids = append(vids, v)
	}
	return vids
}

func getVideoFromEntry(i *interface{}) (Video, bool) {
	x := (*i).(map[string]interface{})["videoRenderer"]
	if x == nil {
		return Video{}, false // when radioRenderer, shelfRenderer, reelShelfRenderer, etc.
	}
	y := x.(map[string]interface{})
	return Video{
		Id:       VID(y["videoId"].(string)),
		Title:    y["title"].(map[string]interface{})["runs"].([]interface{})[0].(map[string]interface{})["text"].(string),
		Channel:  y["ownerText"].(map[string]interface{})["runs"].([]interface{})[0].(map[string]interface{})["text"].(string),
		Duration: y["lengthText"].(map[string]interface{})["simpleText"].(string),
	}, true
}
