package main

import (
	"bufio"
	"fmt"
	"hash/crc32"
	"log"
	"os"
	"strings"
)

type videoCache map[uint32][]Video // cache: content hash -> parsed videos

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
	defer rl.Close()

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
			fmt.Println("No search query provided.")
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
	prevVideo := BACK_FLAG
	c := make(videoCache)
	f := string(p)
	for {
		vs, err := getPlaylistVideos(f, c)
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
			if prevVideo == BACK_FLAG {
				return nil
			}
			v = prevVideo
		}
		v.Play(m)
		prevVideo = v
	}
}

func getPlaylistVideos(filename string, cache videoCache) ([]Video, error) {
	pl, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// check if we've already parsed this playlist content
	contentHash := crc32.ChecksumIEEE(pl)
	if cached, exists := cache[contentHash]; exists {
		return cached, nil // use cached parsed videos
	}

	lines := strings.Split(string(pl), "\n")
	var vs []Video
	for i, id := range lines {
		switch len(id) {
		case 0:
			continue // skip empty lines
		case 11:
			v, err := GetVideoFromURL(VID(id).URL())
			if err != nil {
				return nil, err
			}
			vs = append(vs, *v)
		default:
			log.Printf("%s[WARN]%s Skipped invalid Video ID on line %d: %s\n", C_YELLOW, C_RESET, i+1, id)
		}
	}
	cache[contentHash] = vs // cache the result
	return vs, nil
}
