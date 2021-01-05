package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"sync"
	"time"
	"wallpaper/background"
	"wallpaper/utils"
	"wallpaper/wallhaven"

	"gitee.com/qiuzhiqian/downloader"
)

type Manager struct {
	cfg        *Config
	mux        sync.Mutex
	wallpapers []string
}

func NewManager() *Manager {
	mgr := &Manager{
		wallpapers: make([]string, 0),
	}

	imageDir, err := getImageDir()
	if err != nil {
		panic(err)
	}
	localList := GetLocalFile(imageDir, []string{".png", ".jpg", ".jpeg"})
	if len(localList) > 0 {
		mgr.wallpapers = append(mgr.wallpapers, localList...)
	}

	mgr.cfg, err = LoadConfig(filepath.Join(utils.GetCurrentDirectory(), "config.toml"))
	if err != nil {
		panic(err)
	}
	return mgr
}

func (m *Manager) DownloadHandle() {
	for {
		var DataList []wallhaven.ImgInfo
		for _, page := range m.cfg.Wh.Param.Page {
			jsondata, err := wallhaven.Searching(wallhaven.Param{
				Page:       page,
				Categories: m.cfg.Wh.Param.Categories,
				Tag:        m.cfg.Wh.Param.Tag,
			})
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
	}
}

func (m *Manager) SettingHandle() {
	for {
		rand.Seed(time.Now().Unix())
		index := rand.Intn(len(m.wallpapers))
		fmt.Println("index:", index)

		err := background.SetBg(m.wallpapers[index])
		if err != nil {
			fmt.Println("err:", err)
		}
		fmt.Println("set background success.")
	}
}

func saveHaven(item wallhaven.ImgInfo) string {
	fileDir, err := getImageDir()
	if err != nil {
		return ""
	}

	_, err = os.Stat(fileDir)
	if os.IsNotExist(err) {
		os.MkdirAll(fileDir, os.ModeDir|0755)
	} else if err != nil {
		return ""
	}

	fmt.Println("url:", item.Path)
	saveName, err := downloader.DownloadWithProgress(item.Path, fileDir, func(_, _ int) {})
	if err != nil {
		return ""
	}
	return saveName
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
