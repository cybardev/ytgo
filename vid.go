package main

import (
	"log"
	"net/url"
	"os"
	"os/exec"
)

type VID string

func (v VID) url() string {
	return "https://www.youtube.com/watch?v=" + string(v)
}

func (v VID) play(m bool) {
	bestaudio, novideo := "", ""
	if m {
		bestaudio, novideo = "--ytdl-format=bestaudio", "--no-video"
	}
	cmd := exec.Command("mpv", bestaudio, novideo, v.url())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func VIDfromURL(s string) VID {
	u, err := url.Parse(s)
	if err != nil {
		log.Fatalln(err)
	}
	return VID(u.Query().Get("v"))
}
