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
	"wallpaper-toolbox/background"
	"wallpaper-toolbox/wallhaven"

	"gitee.com/qiuzhiqian/downloader"
)

var ErrNeedSkip = errors.New("file exist need skip")

type Manager struct {
	cfg        *Config
	mux        sync.Mutex
	wallpapers []string
	nextCh     chan bool
}

func NewManager() *Manager {
	mgr := &Manager{
		wallpapers: make([]string, 0),
		nextCh:     make(chan bool),
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
			param := wallhaven.Param{
				Page:       page,
				Categories: m.cfg.Wh.Param.Categories,
				Tag:        m.cfg.Wh.Param.Tag,
				Resolutions: []wallhaven.Resolution{
					wallhaven.Resolution{
						Width:  m.cfg.ScreenSize.Width,
						Height: m.cfg.ScreenSize.Height,
					},
				},
			}
			jsondata, err := wallhaven.Searching(param)
			if err != nil {
				fmt.Println("err:", err)
				continue
			}

			if len(jsondata.Data) == 0 {
				continue
			}

			DataList = append(DataList, jsondata.Data...)
		}

		if len(DataList) == 0 {
			return
		}

		for _, url := range DataList {
			file := saveHaven(url)
			if file == "" {
				return
			}
			fmt.Println("download success.")

			m.mux.Lock()
			m.wallpapers = append(m.wallpapers, file)
			m.mux.Unlock()
		}

		time.Sleep(time.Duration(m.cfg.Wh.Period) * time.Minute)
	}
}

func (m *Manager) SettingHandle() {
	t := time.NewTimer(time.Duration(m.cfg.Setting.Period) * time.Minute)
	for {
		if len(m.wallpapers) == 0 {
			time.Sleep(5 * time.Second)
			continue
		} else {
			//先执行一遍
			m.switchBg()
			break
		}
	}
	for {
		select {
		case <-t.C:
			m.switchBg()
		case <-m.nextCh:
			m.switchBg()
		}
	}
}

func (m *Manager) switchBg() {
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
}

func (m *Manager) Next() {
	m.nextCh <- true
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
	return filepath.Join(u.HomeDir, ".wallpaper", "wallpaper-toolbox"), nil
}

func getConfigDir() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(u.HomeDir, ".config", "wallpaper-toolbox"), nil
}

func GetLocalFile(root string, filter []string) []string {
	var localList []string
	filepath.Walk(root, func(pathname string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
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
