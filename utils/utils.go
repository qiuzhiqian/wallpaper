package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"
)

func GetCurrentDirectory() string {
	//返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	//将\替换成/
	return strings.Replace(dir, "\\", "/", -1)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func SaveFile(url, filename string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(filename)
		panic(err)
	}

	defer f.Close()

	_, err = io.Copy(f, res.Body)
	if err != nil {
		fmt.Println("err:", err)
		return err
	}
	return nil
}

func GetImageDir() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(u.HomeDir, ".wallpaper", "wallpaper-toolbox"), nil
}

func GetConfigDir() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(u.HomeDir, ".config", "wallpaper-toolbox"), nil
}

func GetLocalFile(root string, filter []string) []string {
	localList := make([]string, 0)
	filepath.Walk(root, func(pathname string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Mode().IsRegular() {
			if IsFileMatch(pathname, filter) {
				localList = append(localList, pathname)
			}
		}
		return nil
	})
	return localList
}

func IsFileMatch(name string, filter []string) bool {
	var match bool = false
	ext := path.Ext(name)

	for _, item := range filter {
		if ext == item {
			match = true
			break
		}
	}

	return match
}
