package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Foot struct {
	obj   fyne.CanvasObject
	count int
	name  string

	countLabel *widget.Label
	nameLabel  *widget.Label
}

func NewFoot() *Foot {
	foot := &Foot{
		count: 0,
		name:  "",
	}
	foot.countLabel = widget.NewLabel(fmt.Sprintf("count: %d", foot.count))
	foot.nameLabel = widget.NewLabel(foot.name)
	infoWidget := container.NewGridWithRows(1, foot.countLabel, foot.nameLabel)

	btnSet := widget.NewButton("Set", func() {})
	foot.obj = container.NewBorder(nil, nil, nil, btnSet, infoWidget)
	return foot
}

func (f *Foot) GetCount() int {
	return f.count
}

func (f *Foot) SetCount(cnt int) {
	if cnt != f.count {
		f.count = cnt
		f.countLabel.SetText(fmt.Sprintf("count: %d", f.count))
	}
}

func (f *Foot) GetName() string {
	return f.name
}

func (f *Foot) SetName(name string) {
	if name != f.name {
		f.name = name
		f.nameLabel.SetText(f.name)
	}
}
