package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
	"wallpaper-toolbox/background"
	"wallpaper-toolbox/utils"
	"wallpaper-toolbox/wallhaven"

	"gitee.com/qiuzhiqian/downloader"
)

var ErrNeedSkip = errors.New("file exist need skip")

type Manager struct {
	cfg    *Config
	center *Center
	mux    sync.Mutex
	nextCh chan bool
}

func NewManager() *Manager {
	mgr := &Manager{
		nextCh: make(chan bool),
	}

	configDir, err := utils.GetConfigDir()
	if err != nil {
		panic(err)
	}

	cfg, err := LoadConfig(filepath.Join(configDir, "config.json"))
	if err != nil {
		panic(err)
	}

	mgr.cfg = cfg

	return mgr
}

func (m *Manager) setCenter(c *Center) {
	m.center = c
}

func (m *Manager) changeDownloadState(text string) {
	m.center.changeDownloadState(text)
}

func (m *Manager) changeWallpaperState(text string) {
	m.center.changeWallpaperState(text)
}

func (m *Manager) DownloadHandle() {
	for {
		var DataList []wallhaven.ImgInfo
		for _, page := range m.cfg.Wh.Param.Page {
			m.changeDownloadState(fmt.Sprintf("parse page = %d", page))
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

		for i, url := range DataList {
			//continue
			m.changeDownloadState(fmt.Sprintf("downloading: %d/%d", i, len(DataList)))
			file, err := saveHaven(url)
			if file == "" {
				return
			}

			if err != nil && err != ErrNeedSkip {
				// 获取了一个错误，并且该错误不能被忽略
				return
			} else if err != nil {
				// 获取了一个错误，并且该错误需要被忽略(通常时发现了相同的文件名，所以需要忽略)
				continue
			} else {
				// 说明成功下载了一个新文件
				//m.changeWallpaperState(fmt.Sprintf("wallpaper count: %d.", len(m.wallpapers)))
			}
		}

		for i := m.cfg.Wh.Period; i > 0; i-- {
			m.changeDownloadState(fmt.Sprintf("download end %d, Sleeping %d minute.", len(DataList), i))
			time.Sleep(time.Minute)
		}
	}
}

func (m *Manager) SettingHandle() {
	//m.changeWallpaperState(fmt.Sprintf("wallpaper count: %d.", len(m.wallpapers)))
	t := time.NewTimer(time.Duration(m.cfg.Setting.Period) * time.Minute)
	for {
		if m.center.DataSize() == 0 {
			time.Sleep(5 * time.Second)
			continue
		} else {
			//先执行一遍
			m.switchRandom()
			break
		}
	}
	setCh := m.center.GetEventCh()
	for {
		select {
		case <-t.C:
			m.switchRandom()
		case <-m.nextCh:
			m.switchRandom()
		case name := <-setCh:
			m.switchByName(name)
		}
	}
}

func (m *Manager) switchRandom() {
	m.mux.Lock()
	rand.Seed(time.Now().Unix())
	index := rand.Intn(m.center.DataSize())

	name, err := m.center.view.data.at(index)
	if err != nil {
		return
	}
	err = background.SetBg(name)
	if err != nil {
		fmt.Println("err:", err)
	}

	m.center.SetShowName(filepath.Base(name))
	m.center.SetUpdateTime(time.Now())
	m.mux.Unlock()
}

func (m *Manager) switchByName(name string) {
	err := background.SetBg(name)
	if err != nil {
		fmt.Println("err:", err)
	}

	m.center.SetShowName(filepath.Base(name))
	m.center.SetUpdateTime(time.Now())
}

func (m *Manager) Next() {
	m.nextCh <- true
}

func saveHaven(item wallhaven.ImgInfo) (string, error) {
	fileDir, err := utils.GetImageDir()
	if err != nil {
		return "", err
	}

	_, err = os.Stat(fileDir)
	if os.IsNotExist(err) {
		os.MkdirAll(fileDir, os.ModeDir|0755)
	} else if err != nil {
		return "", err
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
	return dwl.Download(item.Path, fileDir)
}
