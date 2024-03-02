package main

import (
	"fmt"
	"sync"

	"github.com/rivo/tview"
)

func VIDFromMenu(query string) (VID, error) {
	vids, err := getVideos(query)
	if err != nil {
		return "", err
	}
	m := make(VideoMap, len(vids))
	return menuUI(&vids, &m)
}

func menuUI(vids *[]VID, m *VideoMap) (VID, error) {
	var wg sync.WaitGroup
	wg.Add(len(*vids))

	fmt.Println("Loading video list. Please wait...")
	for _, vid := range *vids {
		go mapVideos(vid, m, &wg)
	}

	wg.Wait()

	var v VID
	app := tview.NewApplication()
	list := videoList(app, &v, m, vids)
	err := app.SetRoot(list, true).EnableMouse(true).Run()
	if err != nil {
		return "", err
	}
	return v, nil
}

func mapVideos(id VID, m *VideoMap, wg *sync.WaitGroup) error {
	defer (*wg).Done()
	v, err := fetchVideoInfo(id)
	if err != nil {
		return err
	}
	(*m)[id] = v
	return nil
}

func videoList(app *tview.Application, selected *VID, m *VideoMap, vids *[]VID) *tview.List {
	l := tview.NewList()
	l = l.
		AddItem("Quit", "Press to exit", 'q', func() {
			*selected = ""
			app.Stop()
		})

	var v Video
	for _, vid := range *vids {
		v = (*m)[vid]
		l.AddItem("\n"+v.Title, v.Desc(), 0, func() {
			*selected = vid
			app.Stop()
		})
	}

	return l
}
