package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/abema/go-mp4"
)

var lblPath *canvas.Text

func main() {
	a := app.New()
	win := a.NewWindow("Ref Video Sorter")
	win.Resize(fyne.NewSize(1200, 600))

	title := canvas.NewText("Refereeing Video Sorter", color.Black)
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	lblPath = canvas.NewText("No selected directory", color.Black)
	lblPath.Alignment = fyne.TextAlignLeading
	lblPath.TextSize = 16
	lblPath.TextStyle = fyne.TextStyle{Bold: true}

	btnSelectDir := widget.NewButton("Select directory", func() {
		showFolderDialog(win, func(selectedPath string, err error) {
			if err != nil {
				log.Println("Error opening folder:", err)
				return
			}
			if selectedPath != "" {
				updateLabel(selectedPath)
				videoFiles, err := scanForVideos(selectedPath)
				if err != nil {
					log.Println("Error scanning for videos:", err)
					return
				}
				displayVideoFiles(win, videoFiles)
			} else {
				log.Println("Folder selection cancelled by user")
			}
		})
	})

	content := container.NewVBox(
		title,
		lblPath,
		btnSelectDir,
	)

	win.SetContent(content)
	win.ShowAndRun()
}

// showFolderDialog opens a folder selection dialog and returns the selected path.
func showFolderDialog(win fyne.Window, callback func(string, error)) {
	dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
		if err != nil {
			callback("", err)
			return
		}
		if uri == nil {
			callback("", nil)
			return
		}
		callback(uri.Path(), nil)
	}, win)
}

// updateLabel updates the label with the selected directory path.
func updateLabel(path string) {
	lblPath.Text = "Selected directory: " + path
	lblPath.TextStyle = fyne.TextStyle{Bold: false}
	lblPath.Refresh()
}

// getVideoLength retrieves the length of the video file in HH:MM:SS format.
func getVideoLength(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "00:00", err
	}
	defer file.Close()

	info, err := mp4.Probe(file)
	if err != nil {
		return "00:00", err
	}

	durationInSeconds := info.Duration / 1000
	if durationInSeconds <= 0 {
		return "00:00", fmt.Errorf("invalid duration")
	}

	hours := int(durationInSeconds) / 3600
	minutes := (int(durationInSeconds) % 3600) / 60
	seconds := int(durationInSeconds) % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds), nil
}

// scanForVideos scans the specified directory for video files and returns their paths.
func scanForVideos(dir string) ([]string, error) {
	var videoFiles []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && (filepath.Ext(path) == ".mp4" || filepath.Ext(path) == ".mkv" || filepath.Ext(path) == ".avi") {
			videoFiles = append(videoFiles, path)
		}
		return nil
	})
	return videoFiles, err
}

// displayVideoFiles displays the list of video files in a table format.
func displayVideoFiles(win fyne.Window, files []string) {
	var videoData [][]string

	for _, file := range files {
		info, err := os.Stat(file)
		if err != nil {
			continue
		}
		date := info.ModTime().Format("2006-01-02")
		length, err := getVideoLength(file)
		if err != nil {
			length = "00:00"
		}
		relativePath := truncatePath(file)

		videoData = append(videoData, []string{relativePath, date, length})
	}

	videoTable := widget.NewTable(
		func() (int, int) {
			return len(videoData) + 1, 3
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			if tci.Row == 0 {
				if tci.Col == 0 {
					co.(*widget.Label).SetText("File Path")
				} else if tci.Col == 1 {
					co.(*widget.Label).SetText("Date")
				} else if tci.Col == 2 {
					co.(*widget.Label).SetText("Length")
				}
				co.(*widget.Label).TextStyle = fyne.TextStyle{Bold: true}
			} else if tci.Row-1 < len(videoData) {
				co.(*widget.Label).SetText(videoData[tci.Row-1][tci.Col])
			}
		},
	)

	totalWidth := float32(1200.0)
	videoTable.SetColumnWidth(0, totalWidth*0.75)
	videoTable.SetColumnWidth(1, totalWidth*0.125)
	videoTable.SetColumnWidth(2, totalWidth*0.125)

	scrollContainer := container.NewScroll(videoTable)
	scrollContainer.SetMinSize(fyne.NewSize(0, 400))

	restartButton := widget.NewButton("Restart", func() {
		resetApplication(win)
	})

	quitButton := widget.NewButton("Quit", func() {
		win.Close()
	})

	titleContainer := container.NewHBox(
		widget.NewLabel("Video Files"),
		restartButton,
		quitButton,
	)

	content := container.NewVBox(
		titleContainer,
		scrollContainer,
	)

	win.SetContent(content)
}

// resetApplication resets the application state to its initial state.
func resetApplication(win fyne.Window) {
	title := canvas.NewText("Refereeing Video Sorter", color.Black)
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	lblPath = canvas.NewText("No selected directory", color.Black)
	lblPath.Alignment = fyne.TextAlignLeading
	lblPath.TextSize = 16
	lblPath.TextStyle = fyne.TextStyle{Bold: true}

	btnSelectDir := widget.NewButton("Select directory", func() {
		showFolderDialog(win, func(selectedPath string, err error) {
			if err != nil {
				log.Println("Error opening folder:", err)
				return
			}
			if selectedPath != "" {
				updateLabel(selectedPath)
				videoFiles, err := scanForVideos(selectedPath)
				if err != nil {
					log.Println("Error scanning for videos:", err)
					return
				}
				displayVideoFiles(win, videoFiles)
			} else {
				log.Println("Folder selection cancelled by user")
			}
		})
	})

	content := container.NewVBox(
		title,
		lblPath,
		btnSelectDir,
	)

	win.SetContent(content)
}

// truncatePath truncates the full file path to a relative path.
func truncatePath(fullPath string) string {
	baseDir := filepath.Dir(fullPath)
	relativePath, _ := filepath.Rel(baseDir, fullPath)
	return relativePath
}