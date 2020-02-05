package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
