package main

import (
	"io"
	"os"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	grid := tview.NewGrid().SetColumns(-1).SetRows(-1)
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

	grid.AddItem(textView, 0, 0, 3, 3, 50, 0, false)
	grid.AddItem(killButton, 1, 3, 1, 1, 20, 0, true)

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
