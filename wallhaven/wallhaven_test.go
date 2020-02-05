package wallhaven

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestWallPaperFlash(tt *testing.T) {
	var jsondata SearchList
	err := Searching(1, &jsondata)
	if err != nil {
		tt.Log("err:", err)
	}

	for _, item := range jsondata.Data {
		t, err := time.Parse("2006-01-02 15:04:05", item.CreatedAt)
		if err != nil {
			tt.Log("err:", err)
			continue
		}

		filePath := GetCurrentDirectory() + "/" + "image/" + t.Format("20060102")
		fileName := filepath.Base(item.Path)

		var ok bool = false
		ok, err = PathExists(filePath)
		if ok == true {
			SaveFile(item.Path, filePath+"/"+fileName)
		} else if err == nil {
			os.MkdirAll(filePath, os.ModeDir|0644)
		} else {
			continue
		}

		ok, err = PathExists(filePath + "/" + fileName)
		if ok == false || (ok == true && err != nil) {
			SaveFile(item.Path, filePath+"/"+fileName)
		}
	}
}
