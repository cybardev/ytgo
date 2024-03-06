package main

import (
	"fmt"
	"regexp"
	"sync"
	"testing"
)

var vs []VID = []VID{
	"dQw4w9WgXcQ",
	"kVqUuYKH77o",
	"0Uhh62MUEic",
	"XtK50cbCAdk",
	"KqRl5OAFYCQ",
}

var qs []string = []string{
	"rickroll",
	"gurenge cover",
	"hotaru maiko fujita",
	"one last kiss utada",
	"face my fears utada",
}

func TestGetRequest(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(len(vs))

	for _, v := range vs {
		go func(v VID, t *testing.T, wg *sync.WaitGroup) {
			defer wg.Done()
			r, err := GetRequest(v.URL())
			if err != nil {
				t.Error(err)
			}
			if r == "" {
				t.Error("Empty response")
			}
		}(v, t, &wg)
	}

	wg.Wait()
}

func TestGetSearchResults(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(len(qs))

	for _, q := range qs {
		go func(s string, t *testing.T, wg *sync.WaitGroup) {
			defer wg.Done()
			v, err := GetSearchResults(s)
			if err != nil {
				t.Error(err)
			}
			for _, v := range *v {
				testGottenVideo(&v, t)
			}
		}(q, t, &wg)
	}

	wg.Wait()
}

func TestGetVideoFromSearch(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(len(qs))

	for i, q := range qs {
		go func(s string, t *testing.T, wg *sync.WaitGroup) {
			defer wg.Done()
			v, err := GetVideoFromSearch(s, i+1)
			if err != nil {
				t.Error(err)
			}
			testGottenVideo(v, t)
		}(q, t, &wg)
	}

	wg.Wait()
}

func TestGetVideoFromURL(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(len(vs))

	for _, u := range vs {
		go func(u string, t *testing.T, wg *sync.WaitGroup) {
			defer wg.Done()
			v, err := GetVideoFromURL(u)
			if err != nil {
				t.Error(err)
			}
			testGottenVideo(v, t)
		}(u.URL(), t, &wg)
	}

	wg.Wait()
}

func testGottenVideo(v *Video, t *testing.T) {
	re := regexp.MustCompile(`^.{11}$`)
	if re.MatchString(string(v.Id)) == false {
		t.Error("Id does not match pattern:", v.Id)
	}
	if v.Title == "" {
		t.Error("Title is empty")
	}
	if v.Channel == "" {
		t.Error("Channel is empty")
	}
	re = regexp.MustCompile(`^(\d+:)?\d+:\d+$`)
	if re.MatchString(v.Duration) == false {
		t.Error("Duration does not match pattern:", v.Duration)
	}
	if v.Desc() != fmt.Sprintf("(%s) [%s]", v.Channel, v.Duration) {
		t.Error("Description not in expected form:", v.Desc())
	}
	if v.String() != fmt.Sprintf("\x1b[36m%s \x1b[32m(%s) \x1b[31m[%s]\x1b[00m", v.Title, v.Channel, v.Duration) {
		t.Error("String is not in expected form:", v.String())
	}
	re = regexp.MustCompile(`^https://www\.youtube\.com/watch\?v=.{11}$`)
	if re.MatchString(v.Id.URL()) == false {
		t.Error("URL does not match pattern:", v.Id.URL())
	}
}
