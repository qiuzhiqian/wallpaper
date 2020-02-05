package wallhaven

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type ThumbsStruct struct {
	Large    string
	Original string
	Small    string
}

type ImgInfo struct {
	Id         string
	Url        string
	ShortUrl   string `json:"short_url"`
	Views      int
	Favorites  int
	Source     string
	Purity     string
	Category   string
	DimensionX int `json:"dimension_x"`
	DimensionY int `json:"dimension_y"`
	Resolution string
	Ratio      string
	FileSize   int    `json:"file_size"`
	FileType   string `json:"file_type"`
	CreatedAt  string `json:"created_at"`
	Colors     []string
	Path       string
	Thumbs     ThumbsStruct
}

type MetaStruct struct {
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	PerPage     int `json:"per_page"`
	Total       int
}

type SearchList struct {
	Data []ImgInfo
	Meta MetaStruct
}

func Searching(page int, v interface{}) error {
	resp, err := http.Get("https://wallhaven.cc/api/v1/search" + "?" + "categories=anime&" + "page=" + strconv.FormatInt(int64(page), 10))
	if err != nil {
		// handle error
		return err
	}
	defer resp.Body.Close()

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err:", err)
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		fmt.Println("err:", err)
		return err
	}
	return nil
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

func GetCurrentDirectory() string {
	//返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	//将\替换成/
	return strings.Replace(dir, "\\", "/", -1)
}
