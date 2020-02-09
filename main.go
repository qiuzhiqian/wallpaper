package main

import (
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
}

type Wallhaven struct {
	Config []WallhavenConfig `toml:"config"`
}

type WallhavenConfig struct {
	Page       int
	Categories string
	Tag        string
}

func main() {
	fmt.Println("version:", "1.0.0")
	var cfg Config
	_, err := toml.DecodeFile(utils.GetCurrentDirectory()+"/config.toml", &cfg)
	if err != nil {
		fmt.Println("err:", err)
	}

	for {
		Handler(cfg)
		time.Sleep(time.Second * time.Duration(cfg.Mgr.Period))
	}
}
