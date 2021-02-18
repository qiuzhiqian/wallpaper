package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Foot struct {
	obj        fyne.CanvasObject
	count      int
	name       string
	updateTime time.Time

	countLabel  *widget.Label
	nameLabel   *widget.Label
	updateLabel *widget.Label

	ch []chan bool
}

func NewFoot() *Foot {
	foot := &Foot{
		count:      0,
		name:       "",
		updateTime: time.Now(),
		ch:         make([]chan bool, 0),
	}
	foot.countLabel = widget.NewLabel(fmt.Sprintf("count: %d", foot.count))
	foot.nameLabel = widget.NewLabel(foot.name)
	foot.updateLabel = widget.NewLabel("update: " + foot.updateTime.Format("2006-01-02 15:04:05"))
	infoWidget := container.NewGridWithRows(1, foot.countLabel, foot.nameLabel, foot.updateLabel)

	btnSet := widget.NewButton("Set", func() {
		for _, ch := range foot.ch {
			ch <- true
		}
	})
	foot.obj = container.NewBorder(nil, nil, nil, btnSet, infoWidget)
	return foot
}

func (f *Foot) RegisterClickedEvent(ch chan bool) {
	f.ch = append(f.ch, ch)
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

func (f *Foot) SetUpdateTime(t time.Time) {
	if t.Unix() != f.updateTime.Unix() {
		f.updateTime = t
		f.updateLabel.SetText("update: " + f.updateTime.Format("2006-01-02 15:04:05"))
	}
}
