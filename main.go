package main

import (
	"fmt"
	"wallpaper/autorun"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

var cmdch chan int

var check *widget.Check

const (
	version string = "1.2.0"
)

func getVersion() string {
	return version
}

func running() {
	m := NewManager()
	go m.DownloadHandle()
	go m.SettingHandle()
}

func next() {
	cmdch <- 0
}

func loadUI() {
	appObj := app.New()

	check = widget.NewCheck("Auto run", func(checked bool) {
		fmt.Println("check:", checked)
		res := autorun.Enable(checked)
		if !res {
			check.SetChecked(!checked)
		}
	})

	w := appObj.NewWindow("Wallpaper")
	w.SetContent(fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		widget.NewLabel("version:"+getVersion()),
		fyne.NewContainerWithLayout(
			layout.NewGridLayout(2),
			check,
			widget.NewButton("Next", func() {
				next()
			}),
		),
		widget.NewButton("Quit", func() {
			appObj.Quit()
		}),
	))

	w.Resize(fyne.Size{Width: 300, Height: 100})

	w.ShowAndRun()
}

func main() {
	running()
	loadUI()
}
