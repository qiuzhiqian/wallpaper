package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"sync"
	"time"
	"wallpaper/background"
	"wallpaper/wallhaven"

	"gitee.com/qiuzhiqian/downloader"
)

var ErrNeedSkip = errors.New("file exist need skip")

type Manager struct {
	cfg        *Config
	mux        sync.Mutex
	wallpapers []string
}

func NewManager() *Manager {
	mgr := &Manager{
		wallpapers: make([]string, 0),
	}

	configDir, err := getConfigDir()
	if err != nil {
		panic(err)
	}

	cfg, err := LoadConfig(filepath.Join(configDir, "config.json"))
	if err != nil {
		panic(err)
	}

	mgr.cfg = cfg

	fmt.Println("cfg:", cfg)

	imageDir, err := getImageDir()
	if err != nil {
		panic(err)
	}
	localList := GetLocalFile(imageDir, []string{".png", ".jpg", ".jpeg"})
	if len(localList) > 0 {
		mgr.wallpapers = append(mgr.wallpapers, localList...)
	}

	return mgr
}

func (m *Manager) DownloadHandle() {
	for {
		var DataList []wallhaven.ImgInfo
		for _, page := range m.cfg.Wh.Param.Page {
			fmt.Println("download page =", page)
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

		for _, url := range DataList {
			fmt.Println("download url =", url)
			file := saveHaven(url)
			if file == "" {
				return
			}
			fmt.Println("download success.")

			m.mux.Lock()
			m.wallpapers = append(m.wallpapers)
			m.mux.Unlock()
		}

		time.Sleep(30 * time.Minute)
	}
}

func (m *Manager) SettingHandle() {
	for {
		if len(m.wallpapers) == 0 {
			time.Sleep(5 * time.Second)
			continue
		}

		m.mux.Lock()
		rand.Seed(time.Now().Unix())
		index := rand.Intn(len(m.wallpapers))
		fmt.Println("index:", index)

		err := background.SetBg(m.wallpapers[index])
		if err != nil {
			fmt.Println("err:", err)
		}
		fmt.Println("set background success.")
		m.mux.Unlock()

		time.Sleep(2 * time.Minute)
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
	dwl := downloader.NewDownloader()
	dwl.SetProgressHandler(func(_ string, _, _ int) {})
	dwl.SetRenameHandler(func(filePath string) (string, error) {
		_, err = os.Stat(filePath)
		if err == nil {
			fmt.Println("find file")
			return filePath, ErrNeedSkip
		}

		fmt.Println("can't find file")
		return filePath, nil
	})
	saveName, err := dwl.Download(item.Path, fileDir)
	if err != nil && err != ErrNeedSkip {
		fmt.Println("err:", err)
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

func getConfigDir() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(u.HomeDir, ".config", "wallhaven"), nil
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
