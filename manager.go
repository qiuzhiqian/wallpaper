package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
	"wallpaper/background"
	"wallpaper/wallhaven"
)

func Handler() {
	var jsondata wallhaven.SearchList
	err := wallhaven.Searching(1, &jsondata)
	if err != nil {
		fmt.Println("err:", err)
	}

	if len(jsondata.Data) == 0 {
		return
	}

	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(jsondata.Data))
	fmt.Println("index:", index)

	file := saveHaven(jsondata.Data[index])

	background.SetBg(file)
}

func saveHaven(item wallhaven.ImgInfo) string {
	filePath := wallhaven.GetCurrentDirectory() + "/" + "image"
	fileName := filepath.Base(item.Path)

	var ok bool = false
	var err error
	ok, err = wallhaven.PathExists(filePath)
	if ok == false && err == nil {
		os.MkdirAll(filePath, os.ModeDir|0644)
	} else if ok == false && err != nil {
		return ""
	}

	ok, err = wallhaven.PathExists(filePath + "/" + fileName)
	if ok == false || (ok == true && err != nil) {
		wallhaven.SaveFile(item.Path, filePath+"/"+fileName)
	}
	return filePath + "/" + fileName
}
