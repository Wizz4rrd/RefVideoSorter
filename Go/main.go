package main

import (
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)


func main() {

	a := app.New()
	win := a.NewWindow("Ref Video Sorter")
	win.Resize(fyne.NewSize(800, 300))

	title := canvas.NewText("Refereeing Video Sorter", color.Black)
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24



	btnSelectDir := widget.NewButton("Select directory", func() {
		showFolderDialog(win, func(selectedPath string, err error) {
			if err != nil {
				log.Println("Error opening folder:", err)
				return
			}
			if selectedPath != "" {
				fmt.Println("Selected folder:", selectedPath) // DEV TO DO: return & use for label selected dir (also add label)
			} else {
				log.Println("Folder selection cancelled by user")
			}
		})
	})




	hBox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), btnSelectDir, layout.NewSpacer())
	vBox := container.New(layout.NewVBoxLayout(), title, hBox, widget.NewSeparator())

	win.SetContent(vBox)
	win.ShowAndRun()
}


func showFolderDialog(win fyne.Window, callback func(string, error)) {
	dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
		if err != nil {
			callback("", err)
			return
		}
		if uri == nil {
			// User cancelled folder selection
			callback("", nil)
			return
		}
		callback(uri.Path(), nil)
	}, win)
}