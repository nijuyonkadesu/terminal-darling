package main

import (
	"io"
	"os"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	grid := tview.NewGrid().SetRows(0, 10).SetColumns(0)
	// how to nest items!!! - use grid / flex
	// box := tview.NewBox()

	killButton := tview.NewButton("nuke!").SetSelectedFunc(func() {
		app.Stop()
	})
	killButton.SetBorder(true)

	textView := tview.NewTextView().SetDynamicColors(true).SetChangedFunc(func() { app.Draw() })
	textView.SetBorder(true).SetTitle("Hey minna!!")
	go func() {
		w := tview.ANSIWriter(textView)
		if _, err := io.Copy(w, os.Stdin); err != nil {
			panic(err)
		}
	}()

	grid.AddItem(textView, 0, 0, 1, 3, 30, 0, false)
	grid.AddItem(killButton, 1, 1, 1, 1, 10, 0, true)

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
