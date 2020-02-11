package main

import (
	"fmt"
	"testing"
	"wallpaper/utils"

	"github.com/BurntSushi/toml"
)

func TestMain(t *testing.T) {
	var cfg Config
	_, err := toml.DecodeFile(utils.GetCurrentDirectory()+"/config.toml", &cfg)
	if err != nil {
		fmt.Println("err:", err)
	}
}

func TestLocal(t *testing.T) {
	localList := GetLocalFile(utils.GetCurrentDirectory()+"/"+"image", []string{".png", ".jpg", ".jpeg"})
	t.Log(localList)
}
