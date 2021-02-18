package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Center struct {
	m              *Manager
	app            fyne.App
	check          *widget.Check
	window         fyne.Window
	downloadState  *widget.Label
	wallpaperState *widget.Label
	view           *Preview
}

func NewCenter() *Center {
	c := &Center{
		app: app.NewWithID("com.qiuzhiqian.wallpaper-toolbox"),
	}

	c.app.SetIcon(titleIconRs)

	c.window = c.app.NewWindow("Wallpaper")
	c.window.SetMaster()
	return c
}

func (c *Center) init(m *Manager) {
	c.m = m
	c.downloadState = widget.NewLabel("Begining download...")
	c.wallpaperState = widget.NewLabel("wallpaper count: 0")
	c.view = NewPreview()

	topWidget := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("version:"+getVersion()),
			layout.NewSpacer(),
			widget.NewButton("Next", func() {
				m.Next()
			}),
		),
		c.downloadState,
		c.wallpaperState,
	)

	x := container.NewBorder(
		topWidget,
		nil, nil, nil,
		c.view.obj,
	)

	c.window.SetContent(x)

	c.window.Resize(fyne.NewSize(980, 600))

	c.view.Init()
}

func (c *Center) GetEventCh() chan string {
	return c.view.SetCh
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

func (c *Center) GetData(index int) string {
	item, err := c.view.data.at(index)
	if err != nil {
		item = ""
	}
	return item
}

func (c *Center) DataSize() int {
	return c.view.data.Size()
}

func (c *Center) SetShowName(name string) {
	c.view.SetShowName(name)
}

func (c *Center) SetUpdateTime(t time.Time) {
	c.view.foot.SetUpdateTime(t)
}

func (c *Center) ShowAndRun() {
	c.window.ShowAndRun()
}
