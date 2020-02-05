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
	if file == "" {
		return
	}
	fmt.Println("download success.")
	background.SetBg(file)
	fmt.Println("set background success.")
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

	savename := filePath + "/" + fileName
	ok, err = wallhaven.PathExists(savename)
	if ok == false || (ok == true && err != nil) {
		tempname := savename + ".wptemp"
		ok, err = wallhaven.PathExists(tempname)
		if ok == true && err == nil {
			os.Remove(tempname)
		}

		err = wallhaven.SaveFile(item.Path, tempname)
		if err != nil {
			return ""
		}

		//下载成功，重命名
		err = os.Rename(tempname, savename)
		if err != nil {
			fmt.Println("err:", err)
			return ""
		}
	}
	return savename
}
