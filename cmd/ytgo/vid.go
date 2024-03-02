package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
)

const YtURL = "https://www.youtube.com/"

type VideoMap map[VID]Video

type Video struct {
	Id       VID    `json:"id"`
	Title    string `json:"title"`
	Channel  string `json:"channel"`
	Duration string `json:"duration_string"`
	Url      string `json:"original_url"`
}

func (v Video) Desc() string {
	return fmt.Sprintf("(%s) [%s]", v.Channel, v.Duration)
}

type VID string

func (v VID) URL() string {
	return YtURL + "watch?v=" + string(v)
}

func (v VID) Play(m bool) error {
	bestaudio, novideo := "", ""
	if m {
		bestaudio, novideo = "--ytdl-format=bestaudio", "--no-video"
	}
	cmd := exec.Command("mpv", bestaudio, novideo, v.URL())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func VIDfromURL(s string) (VID, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	return VID(u.Query().Get("v")), nil
}
