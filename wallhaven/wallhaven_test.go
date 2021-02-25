package wallhaven

import (
	"os"
	"path/filepath"
	"testing"
	"time"
	"wallpaper-toolbox/utils"
)

func TestWallPaperFlash(tt *testing.T) {
	wp := Param{
		Page:       1,
		Categories: "anime",
	}

	jsondata, err := Searching(wp)
	if err != nil {
		tt.Log("err:", err)
	}

	for _, item := range jsondata.Data {
		t, err := time.Parse("2006-01-02 15:04:05", item.CreatedAt)
		if err != nil {
			tt.Log("err:", err)
			continue
		}

		filePath := utils.GetCurrentDirectory() + "/" + "image/" + t.Format("20060102")
		fileName := filepath.Base(item.Path)

		var ok bool = false
		ok, err = utils.PathExists(filePath)
		if ok == true {
			utils.SaveFile(item.Path, filePath+"/"+fileName)
		} else if err == nil {
			os.MkdirAll(filePath, os.ModeDir|0644)
		} else {
			continue
		}

		ok, err = utils.PathExists(filePath + "/" + fileName)
		if ok == false || (ok == true && err != nil) {
			utils.SaveFile(item.Path, filePath+"/"+fileName)
		}
	}
}

func TestToString(t *testing.T) {
	p := Param{
		Page:       1,
		Categories: "anime",
		Tag:        "anime",
		Resolutions: []Resolution{
			Resolution{
				Width:  1920,
				Height: 1080,
			},
			Resolution{
				Width:  2160,
				Height: 1440,
			},
		},
	}

	t.Log("url:", p)
}
