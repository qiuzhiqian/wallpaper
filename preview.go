package main

import (
	"path/filepath"
	"wallpaper-toolbox/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Preview struct {
	obj  fyne.CanvasObject
	foot *Foot
	data *DataEngine
}

func NewPreview() *Preview {
	dir, err := utils.GetImageDir()
	if err != nil {
		panic("err")
	}
	preview := &Preview{
		data: NewDataEngine(dir, []string{".png", ".jpg", ".jpeg"}),
	}

	preview.initListWidget()

	return preview
}

func (p *Preview) Init() {
	go p.EventHandle()
	go p.data.Run()
}

func (p *Preview) initListWidget() {
	// 用作容器，用来刷新图片预览。如果不用布局包裹起来，好像无法实时刷新。
	content := container.NewMax()

	list := widget.NewList(
		func() int {
			return p.data.Size()
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template Object")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			dataitem, err := p.data.at(id)
			if err != nil {
				return
			}
			name := filepath.Base(dataitem)
			item.(*widget.Label).SetText(name)
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		dataitem, err := p.data.at(id)
		if err != nil {
			return
		}

		logo := canvas.NewImageFromFile(dataitem)
		logo.FillMode = canvas.ImageFillContain

		content.Objects = []fyne.CanvasObject{logo}
		content.Refresh()
	}
	list.OnUnselected = func(id widget.ListItemID) {
		dataitem, err := p.data.at(id)
		if err != nil {
			return
		}

		logo := canvas.NewImageFromFile(dataitem)
		logo.FillMode = canvas.ImageFillContain

		content.Objects = []fyne.CanvasObject{logo}
		content.Refresh()
	}

	sp := container.NewHSplit(list, content)
	sp.Offset = 0.2

	p.foot = NewFoot()
	p.foot.SetCount(p.data.Size())

	p.obj = container.NewBorder(nil, p.foot.obj, nil, nil, sp)
}

func (p *Preview) EventHandle() {
	ch := make(chan DataEvent, 100)
	p.data.Register(ch)

	for ev := range ch {
		switch ev.EventType {
		case ADD:
			//do add
			p.obj.(*container.Split).Leading.Refresh()
			p.foot.SetCount(p.data.Size())
		case REMOVE:
			// do remove
			p.obj.(*container.Split).Leading.Refresh()
			p.foot.SetCount(p.data.Size())
		}
	}
}

func (p *Preview) SetShowName(name string) {
	p.foot.SetName(name)
}
