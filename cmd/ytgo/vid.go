package main

import (
	"fmt"
	"os"
	"os/exec"
)

const YtURL = "https://www.youtube.com/"

type VID string

func (v VID) URL() string {
	return YtURL + "watch?v=" + string(v)
}

type Video struct {
	Id       VID    `json:"id"`
	Title    string `json:"title"`
	Channel  string `json:"channel"`
	Duration string `json:"duration_string"`
}

func (v Video) String() string {
	return fmt.Sprintf("%s%s %s(%s) %s[%s]%s", C_CYAN, v.Title, C_GREEN, v.Channel, C_RED, v.Duration, C_RESET)
}

func (v Video) Desc() string {
	return fmt.Sprintf("(%s) [%s]", v.Channel, v.Duration)
}

func (v Video) Play(m bool) error {
	fmt.Println("Playing:", v)
	bestaudio, novideo := "", ""
	if m {
		bestaudio, novideo = "--ytdl-format=bestaudio", "--no-video"
	}
	cmd := exec.Command("mpv", "--script-opts=visualizer-forcewindow=no", bestaudio, novideo, v.Id.URL())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
