package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
	"wallpaper/background"
	"wallpaper/utils"
	"wallpaper/wallhaven"
)

func Handler(c Config) {
	var DataList []wallhaven.ImgInfo
	for _, item := range c.Wh.Config {
		wp := wallhaven.Param{
			Page:       item.Page,
			Categories: item.Categories,
			Tag:        item.Tag,
		}
		var jsondata wallhaven.SearchList
		err := wallhaven.Searching(wp, &jsondata)
		if err != nil {
			fmt.Println("err:", err)
		}

		if len(jsondata.Data) == 0 {
			continue
		}

		DataList = append(DataList, jsondata.Data...)
		fmt.Println("data len:", len(DataList))
	}

	if len(DataList) == 0 {
		return
	}

	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(DataList))
	fmt.Println("index:", index)

	file := saveHaven(DataList[index])
	if file == "" {
		return
	}
	fmt.Println("download success.")
	background.SetBg(file)
	fmt.Println("set background success.")
}

func saveHaven(item wallhaven.ImgInfo) string {
	filePath := utils.GetCurrentDirectory() + "/" + "image"
	fileName := filepath.Base(item.Path)

	var ok bool = false
	var err error
	ok, err = utils.PathExists(filePath)
	if ok == false && err == nil {
		os.MkdirAll(filePath, os.ModeDir|0644)
	} else if ok == false && err != nil {
		return ""
	}

	savename := filePath + "/" + fileName
	ok, err = utils.PathExists(savename)
	if ok == false || (ok == true && err != nil) {
		tempname := savename + ".wptemp"
		ok, err = utils.PathExists(tempname)
		if ok == true && err == nil {
			os.Remove(tempname)
		}

		err = utils.SaveFile(item.Path, tempname)
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
