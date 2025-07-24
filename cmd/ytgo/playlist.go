package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ergochat/readline"
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

	// create line reader for search
	rl, err := readline.NewFromConfig(&readline.Config{
		Prompt:  fmt.Sprintf("%sSearch:%s ", C_CYAN, C_RESET),
		VimMode: true,
	})
	if err != nil {
		return err
	}
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
			continue
		}
		v, err := GetVideoFromMenu(query)
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
	playlist, err := os.ReadFile(f)
	if err != nil {
		return err
	}
	lines := strings.Split(string(playlist), "\n")
	for i := 0; i < len(lines)-1; i++ {
		id := lines[i]
		if len(id) == 11 {
			v := Video{Id: VID(id)}
			v.Play(m)
		} else {
			log.Printf("%s[WARN]%s Skipped invalid Video ID: %s\n", C_YELLOW, C_RESET, id)
		}
		playlist, err := os.ReadFile(f)
		if err != nil {
			return err
		}
		lines = strings.Split(string(playlist), "\n")
	}
	return nil
}
