package main

import (
	"fmt"
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func VideoFromMenu(query string) (Video, error) {
	vids, err := getVIDs(query)
	if err != nil {
		return Video{}, err
	}
	n := min(10, len(vids))
	v := vids[:n]
	m := make(VideoMap, n)
	return menuUI(&v, &m)
}

func menuUI(vids *[]VID, m *VideoMap) (Video, error) {
	var wg sync.WaitGroup
	wg.Add(len(*vids))

	fmt.Println("Loading video list. Please wait...")
	for _, vid := range *vids {
		go mapVideos(vid, m, &wg)
	}

	wg.Wait()

	tview.Styles.PrimitiveBackgroundColor = tcell.ColorDefault
	app := tview.NewApplication()
	return getVideoFromList(app, m, vids)
}

func mapVideos(id VID, m *VideoMap, wg *sync.WaitGroup) error {
	defer (*wg).Done()
	v, err := getVideoInfo(id)
	if err != nil {
		return err
	}
	(*m)[id] = v
	return nil
}

func getVideoFromList(app *tview.Application, m *VideoMap, vids *[]VID) (Video, error) {
	var selected Video

	l := tview.NewList()
	l = l.
		AddItem("Quit", "Press to exit", 'q', func() {
			selected = Video{}
			app.Stop()
		})

	for i, vid := range *vids {
		v := (*m)[vid]
		l.AddItem(v.Title, v.Desc(), rune(i+48), func() {
			selected = v
			app.Stop()
		})
	}

	err := app.SetRoot(l, true).EnableMouse(true).Run()
	if err != nil {
		return Video{}, err
	}
	return selected, nil
}
