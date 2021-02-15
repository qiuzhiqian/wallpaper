package main

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed misc/icons/wallpaper-toolbox_16px.png
var titleIcon []byte

var titleIconRs = &fyne.StaticResource{
	StaticName:    "wallpaper-toolbox_icon.png",
	StaticContent: titleIcon,
}
