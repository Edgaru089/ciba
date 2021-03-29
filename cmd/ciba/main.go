package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	tv := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	search := buildSearchTUI(app, tv)

	tv.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune {
			switch event.Rune() {
			case '/':
				search.SetText("")
				app.SetFocus(search)
				return nil
			case '?':
				app.SetFocus(search)
				return nil
			}
		}
		return event
	})

	box := tview.NewFlex().SetDirection(tview.FlexRow)
	box.AddItem(search, 2, 0, false)
	box.AddItem(tv, 0, 1, true)

	app.SetRoot(box, true)
	app.EnableMouse(true)

	if len(os.Args) > 1 {
		tv.SetText(wrapDetailsRequest(os.Args[1]))
		fmt.Println(tv.GetText(true))
	} else {
		app.SetFocus(search)
	}

	if err := app.Run(); err != nil {
		log.Panic(err)
	}
}
