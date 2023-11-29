package main

import (
	"image/color"
	

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

)


type FileDialog struct {

}




func main() {

	a := app.New()
	win := a.NewWindow("Ref Video Sorter")
	win.Resize(fyne.NewSize(800, 300))

	title := canvas.NewText("Refereeing Video Sorter", color.White)
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24


	

	

	

	
	hBox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), layout.NewSpacer())
	vBox := container.New(layout.NewVBoxLayout(), title, hBox, widget.NewSeparator())

	win.SetContent(vBox)
	win.ShowAndRun()
}