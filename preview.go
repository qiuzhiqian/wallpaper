package main

import (
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Preview struct {
	obj  fyne.CanvasObject
	data []string
}

func NewPreview() *Preview {
	preview := &Preview{
		data: make([]string, 0),
	}

	preview.initListWidget()

	return preview
}

func (p *Preview) InitData(data []string) {
	p.data = make([]string, len(data))
	copy(p.data, data)
}

func (p *Preview) initListWidget() {
	// 用作容器，用来刷新图片预览。如果不用布局包裹起来，好像无法实时刷新。
	content := container.NewMax()

	list := widget.NewList(
		func() int {
			return len(p.data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template Object")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			name := filepath.Base(p.data[id])
			item.(*widget.Label).SetText(name)
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		logo := canvas.NewImageFromFile(p.data[id])
		logo.FillMode = canvas.ImageFillContain

		content.Objects = []fyne.CanvasObject{logo}
		content.Refresh()
	}
	list.OnUnselected = func(id widget.ListItemID) {
		logo := canvas.NewImageFromFile(p.data[id])
		logo.FillMode = canvas.ImageFillContain

		content.Objects = []fyne.CanvasObject{logo}
		content.Refresh()
	}

	sp := container.NewHSplit(list, content)
	sp.Offset = 0.2
	p.obj = sp
}

func (p *Preview) AddDataItem(item string) {
	if p.data == nil {
		p.data = make([]string, 0)
	}
	p.data = append(p.data, item)
}
