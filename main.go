package main

import (
	"fmt"
	"time"
	"wallpaper/autorun"
	"wallpaper/utils"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Title string
	Mgr   Manager   `toml:"manager"`
	Wh    Wallhaven `toml:"wallhaven"`
}

type Manager struct {
	Period int
	Mode   string
}

type Wallhaven struct {
	Config []WallhavenConfig `toml:"config"`
}

type WallhavenConfig struct {
	Page       int
	Categories string
	Tag        string
}

var cmdch chan int
var cfg Config

var check *widget.Check

const (
	version string = "1.1.0"
)

func getVersion() string {
	return version
}

func libInit(configpath string) {
	cmdch = make(chan int, 0)
	fmt.Println("core lib version:", version)
	_, err := toml.DecodeFile(configpath, &cfg)
	if err != nil {
		fmt.Println("err:", err)
	}
}

func running() {
	for {
		Handler(cfg)
		select {
		case num := <-cmdch:
			fmt.Println("cmdch:", num)
		case <-time.After(time.Second * time.Duration(cfg.Mgr.Period)):
			fmt.Println("do next")
		}
	}
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
	libInit(utils.GetCurrentDirectory() + "/config.toml")
	go running()
	loadUI()
}
