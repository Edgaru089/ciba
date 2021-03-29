package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func buildSearchTUI(app *tview.Application, root *tview.TextView) *tview.InputField {

	in := tview.NewInputField()
	//in.SetBorder(true)
	/*in.SetAutocompleteFunc(func(currentText string) (entries []string) {
		if len(currentText) == 0 {
			return nil
		}

		words, err := ciba.GetBriefings(currentText, 10)
		if err != nil {
			return nil
		}

		entries = make([]string, len(words))
		for i, w := range words {
			entries[i] = w.Key + ".    " + w.Paraphrase
		}
		return
	})

	in.SetChangedFunc(func(text string) {
		if i := strings.IndexByte(text, '.'); i != -1 {
			in.SetText(text[:i])
		}
	})*/

	in.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			app.SetFocus(root)
			root.SetText(wrapDetailsRequest(in.GetText()))
			root.ScrollToBeginning()

		case tcell.KeyEscape, tcell.KeyTab, tcell.KeyBacktab:
			app.SetFocus(root)
		}
	})

	return in
}
