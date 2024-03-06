package main

import (
	"fmt"
	"testing"
)

type ctrlVideo struct {
	Id       VID
	Title    string
	Channel  string
	Duration string
	Desc     string
	String   string
}

func TestVID(t *testing.T) {
	var ctrlStr string = "ABCDEFG_V1"
	var ctrlUrl string = "https://www.youtube.com/watch?v=ABCDEFG_V1"
	v := VID("ABCDEFG_V1")

	if string(v) != ctrlStr {
		t.Errorf("Expected string %s, got %s", ctrlStr, string(v))
	}

	if v.URL() != ctrlUrl {
		t.Errorf("Expected URL %s, got %s", ctrlUrl, v.URL())
	}
}

func TestVideo(t *testing.T) {
	ctrl := ctrlVideo{
		Id:       VID("ABCDEFG_V2"),
		Title:    "Test",
		Channel:  "Tester",
		Duration: "0:12",
		Desc:     "(Tester) [0:12]",
		String:   fmt.Sprintf("%s%s %s(%s) %s[%s]%s", C_CYAN, "Test", C_GREEN, "Tester", C_RED, "0:12", C_RESET),
	}

	v := Video{
		Id:       VID("ABCDEFG_V2"),
		Title:    "Test",
		Channel:  "Tester",
		Duration: "0:12",
	}

	if v.Id != ctrl.Id {
		t.Errorf("Expected Id %s, got %s", ctrl.Id, v.Id)
	}
	if v.Title != ctrl.Title {
		t.Errorf("Expected Title %s, got %s", ctrl.Title, v.Title)
	}
	if v.Channel != ctrl.Channel {
		t.Errorf("Expected Channel %s, got %s", ctrl.Channel, v.Channel)
	}
	if v.Duration != ctrl.Duration {
		t.Errorf("Expected Duration %s, got %s", ctrl.Duration, v.Duration)
	}
	if v.Desc() != ctrl.Desc {
		t.Errorf("Expected Desc %s, got %s", ctrl.Desc, v.Desc())
	}
	if v.String() != ctrl.String {
		t.Errorf("Expected String %s, got %s", ctrl.String, v.String())
	}
	if v.Id.URL() != ctrl.Id.URL() {
		t.Errorf("Expected Id URL %s, got %s", ctrl.Id.URL(), v.Id.URL())
	}
}
