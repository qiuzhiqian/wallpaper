package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"time"
	"wallpaper/background"
	"wallpaper/utils"
	"wallpaper/wallhaven"
)

func Handler(c Config) {
	var file string
	switch c.Mgr.Mode {
	case "wallhaven":
		var DataList []wallhaven.ImgInfo
		for _, item := range c.Wh.Config {
			wp := wallhaven.Param{
				Page:       item.Page,
				Categories: item.Categories,
				Tag:        item.Tag,
			}

			jsondata, err := wallhaven.Searching(wp)
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

		file = saveHaven(DataList[index])
		if file == "" {
			return
		}
		fmt.Println("download success.")
	case "offline":
		localList := GetLocalFile(utils.GetCurrentDirectory()+"/"+"image", []string{".png", ".jpg", ".jpeg"})

		if len(localList) == 0 {
			return
		}

		rand.Seed(time.Now().Unix())
		index := rand.Intn(len(localList))
		fmt.Println("index:", index)

		file = localList[index]
	}

	err := background.SetBg(file)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("set background success.")
}

func saveHaven(item wallhaven.ImgInfo) string {
	filePath, err := getImageDir()
	if err != nil {
		return ""
	}
	fileName := filepath.Base(item.Path)

	var ok bool = false
	ok, err = utils.PathExists(filePath)
	if ok == false && err == nil {
		os.MkdirAll(filePath, os.ModeDir|0755)
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

func getImageDir() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(u.HomeDir, ".wallpaper", "wallhaven"), nil
}

func GetLocalFile(root string, filter []string) []string {
	var localList []string
	filepath.Walk(root, func(pathname string, info os.FileInfo, err error) error {
		fmt.Println(pathname)
		if info.Mode().IsRegular() {
			ext := path.Ext(pathname)
			var match bool = false

			for _, item := range filter {
				if ext == item {
					match = true
					break
				}
			}
			if match {
				localList = append(localList, pathname)
			}
		}
		return nil
	})
	return localList
}
