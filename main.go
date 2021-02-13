package main

import (
	"fyne.io/fyne/widget"
)

var cmdch chan int

var check *widget.Check

const (
	version string = "2.0.1"
)

func getVersion() string {
	return version
}

func next() {
	cmdch <- 0
}

func main() {
	m := NewManager()
	c := NewCenter()
	c.init(m)
	m.setCenter(c)

	go m.DownloadHandle()
	go m.SettingHandle()

	c.ShowAndRun()
}
