package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func GetVideoFromMenu(vs []Video) (*Video, error) {
	tview.Styles.PrimitiveBackgroundColor = tcell.ColorDefault
	app := tview.NewApplication()
	return getVideoFromList(app, vs)
}

func getVideoFromList(app *tview.Application, vs []Video) (*Video, error) {
	var selected *Video

	list := tview.NewList()
	list.SetSelectedTextColor(tcell.ColorBlack)
	list.SetSelectedBackgroundColor(tcell.ColorWhite)
	list = list.
		AddItem("Quit", "Press to exit", 'Q', func() {
			selected = nil
			app.Stop()
		}).
		AddItem("Back", "Press to go back", 'B', func() {
			selected = BACK_FLAG
			app.Stop()
		})

	for i, v := range vs {
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
