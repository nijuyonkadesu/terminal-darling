package main

import (
	"io"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func stdinCapture(textView *tview.TextView) {
	w := tview.ANSIWriter(textView)
	if _, err := io.Copy(w, os.Stdin); err != nil {
		panic(err)
	}
}

func main() {
	app := tview.NewApplication().EnableMouse(true)
	killButton := tview.NewButton("nuke!").SetSelectedFunc(func() {
		app.Stop()
	})
	textView := tview.NewTextView().SetDynamicColors(true).SetChangedFunc(func() { app.Draw() }) // bad practice? on the only purpose of textView.SetChangedFunc is to call app.Draw() in a separate goroutine? why?
	// TODO: grid is not suitable for having multiple premitives that require focus... only one can be focused ig... or only non interactable elements
	// try to do the same setup using flexbox
	grid := tview.NewGrid().SetRows(0, 10).SetColumns(0)

	killButton.SetBorder(true)

	go stdinCapture(textView)
	textView.SetBorder(true).SetTitle("Hey minna!!")
	textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEsc:
			textView.Clear()
		}

		return event
	})

	grid.AddItem(textView, 0, 0, 1, 3, 30, 0, false)
	grid.AddItem(killButton, 1, 1, 1, 1, 10, 0, false)

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}

// Check type heirarchy: https://pkg.go.dev/github.com/rivo/tview#hdr-Type_Hierarchy
// Continue from here: https://github.com/rivo/tview/wiki/Grid
// TODO: do fun stuffs like rendering image on screen like this: https://github.com/rivo/tview/blob/master/demos/image/main.go
