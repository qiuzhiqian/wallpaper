package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type Center struct {
	app            fyne.App
	check          *widget.Check
	window         fyne.Window
	downloadState  *widget.Label
	wallpaperState *widget.Label
}

func NewCenter() *Center {
	c := &Center{
		app: app.New(),
	}

	c.window = c.app.NewWindow("Wallpaper")
	return c
}

func (c *Center) init(m *Manager) {
	c.check = widget.NewCheck("Auto run", func(checked bool) {
		fmt.Println("do nothing")
	})

	c.downloadState = widget.NewLabel("Begining download...")
	c.wallpaperState = widget.NewLabel("wallpaper count: 0")

	c.window.SetContent(fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		widget.NewLabel("version:"+getVersion()),
		fyne.NewContainerWithLayout(
			layout.NewGridLayout(2),
			c.check,
			widget.NewButton("Next", func() {
				m.Next()
			}),
		),
		c.downloadState,
		c.wallpaperState,
		widget.NewButton("Quit", func() {
			c.app.Quit()
		}),
	))

	c.window.Resize(fyne.Size{Width: 300, Height: 100})
}

func (c *Center) changeDownloadState(text string) {
	if c.downloadState == nil {
		return
	}
	c.downloadState.SetText(text)
}

func (c *Center) changeWallpaperState(text string) {
	if c.wallpaperState == nil {
		return
	}
	c.wallpaperState.SetText(text)
}

func (c *Center) ShowAndRun() {
	c.window.ShowAndRun()
}
