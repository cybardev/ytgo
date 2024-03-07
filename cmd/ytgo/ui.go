package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func GetVideoFromMenu(query string) (*Video, error) {
	vs, err := GetSearchResults(query)
	if err != nil {
		return nil, err
	}
	tview.Styles.PrimitiveBackgroundColor = tcell.ColorDefault
	app := tview.NewApplication()
	return getVideoFromList(app, vs)
}

func getVideoFromList(app *tview.Application, vs *[]Video) (*Video, error) {
	var selected *Video

	list := tview.NewList()
	list = list.
		AddItem("Quit", "Press to exit", 'q', func() {
			selected = nil
			app.Stop()
		})

	for i, v := range *vs {
		list.AddItem(v.Title, v.Desc(), getShortcut(i), func() {
			selected = &v
			app.Stop()
		})
	}

	err := app.SetRoot(list, true).EnableMouse(true).Run()
	if err != nil {
		return nil, err
	}
	return selected, nil
}

func getShortcut(n int) rune {
	if n > 36 || n < 0 {
		return 0
	}
	if n > 9 {
		return rune(n + 87)
	}
	return rune(n + 48)
}
