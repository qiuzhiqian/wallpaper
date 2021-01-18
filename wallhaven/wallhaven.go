package wallhaven

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Resolution struct {
	Width  int
	Height int
}

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

type Param struct {
	Page        int
	Categories  string
	Tag         string
	Resolutions []Resolution
}

func (p Param) String() string {
	paramStr := ""
	rsStr := Resolutions(p.Resolutions)
	if rsStr != "" {
		paramStr = rsStr
	}
	paramStr = fmt.Sprintf("%s&q=%s", paramStr, p.Tag)
	paramStr = fmt.Sprintf("%s&page=%d", paramStr, p.Page)

	url := ""
	if paramStr != "" {
		url = "https://wallhaven.cc/api/v1/search" + "?" + paramStr
	} else {
		url = "https://wallhaven.cc/api/v1/search"
	}
	return url
}

func Resolutions(rs []Resolution) string {
	paramStr := ""
	for i, r := range rs {
		if i > 0 {
			paramStr = fmt.Sprintf("%s,%dx%d", paramStr, r.Width, r.Height)
		} else {
			paramStr = fmt.Sprintf("resolutions=%dx%d", r.Width, r.Height)
		}
	}
	return paramStr
}

func Searching(p Param) (*SearchList, error) {
	var urlstr string = p.String()
	fmt.Println("url:", urlstr)
	resp, err := http.Get(urlstr)
	if err != nil {
		// handle error
		return nil, err
	}
	defer resp.Body.Close()

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}

	var v SearchList
	err = json.Unmarshal(body, &v)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}

	return &v, nil
}
