package main

import (
	"C"
	"fmt"
	"time"
	"wallpaper/utils"

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

const (
	version string = "1.1.0"
)

//export getVersion
func getVersion() string {
	return version
}

//export running
func running(configpath string) {
	cmdch = make(chan int, 0)
	fmt.Println("core lib version:", version)
	var cfg Config
	_, err := toml.DecodeFile(configpath, &cfg)
	if err != nil {
		fmt.Println("err:", err)
	}

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

//export next
func next() {
	cmdch <- 0
}

func main() {
	running(utils.GetCurrentDirectory() + "/config.toml")
}
