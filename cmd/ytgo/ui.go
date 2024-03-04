package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func GetVideoFromMenu(query string) (Video, error) {
	vids, err := GetVideos(query)
	if err != nil {
		return Video{}, err
	}
	tview.Styles.PrimitiveBackgroundColor = tcell.ColorDefault
	app := tview.NewApplication()
	return getVideoFromList(app, &vids)
}

func getVideoFromList(app *tview.Application, vids *[]Video) (Video, error) {
	var selected Video

	l := tview.NewList()
	l = l.
		AddItem("Quit", "Press to exit", 'q', func() {
			selected = Video{}
			app.Stop()
		})

	for i, v := range *vids {
		l.AddItem(v.Title, v.Desc(), getShortcut(i), func() {
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

func getShortcut(n int) rune {
	if n > 36 {
		return 0
	}
	if n > 9 {
		return rune(n + 87)
	}
	return rune(n + 48)
}
