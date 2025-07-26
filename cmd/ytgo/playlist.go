package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Playlist string

func (p Playlist) Create() error {
	f := string(p)
	playlist, err := os.Create(f)
	if err != nil {
		return err
	}
	defer playlist.Close()
	w := bufio.NewWriter(playlist)

	rl := GetReadline()

	// keep adding until user quits
	for {
		query, err := rl.ReadLine()
		if err != nil {
			info, err := playlist.Stat()
			if err != nil {
				return err
			}
			if info.Size() == 0 {
				os.Remove(f)
			}
			return nil // exit on EOF/SIGINT
		}
		if query == "" {
			continue
		}
		vs, err := GetSearchResults(query)
		if err != nil {
			return err
		}
		v, err := GetVideoFromMenu(vs)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(w, v.Id)
		if err != nil {
			return err
		}
		w.Flush()
	}
}

func (p Playlist) Play(m bool) error {
	f := string(p)
	prev := BACK_FLAG
	for {
		vs, err := getPlaylistVideos(f)
		if err != nil {
			return err
		}
		v, err := GetVideoFromMenu(vs)
		if err != nil {
			return err
		}
		switch v {
		case nil:
			return nil
		case BACK_FLAG:
			if prev == BACK_FLAG {
				return nil
			}
			v = prev
		}
		v.Play(m)
		prev = v
	}
}

func getPlaylistVideos(f string) (*[]Video, error) {
	playlist, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(playlist), "\n")
	var vs []Video
	for i := 0; i < len(lines)-1; i++ {
		id := lines[i]
		if len(id) == 11 {
			v, err := GetVideoFromURL(VID(id).URL())
			if err != nil {
				return nil, err
			}
			vs = append(vs, *v)
		} else {
			log.Printf("%s[WARN]%s Skipped invalid Video ID on line %d: %s\n", C_YELLOW, C_RESET, i+1, id)
		}
	}
	return &vs, nil
}
