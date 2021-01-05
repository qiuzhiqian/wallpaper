package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Version string
	//Title   string
	Setting Setting   `json:"setting"`
	Wh      Wallhaven `json:"wallhaven"`
}

type Setting struct {
	Period int
}

type Wallhaven struct {
	Period int
	Param  WallhavenParam `json:"param"`
}

type WallhavenParam struct {
	Page       []int
	Categories string
	Tag        string
}

var defaultCfg Config = Config{
	Version: "1.0.0",
	Setting: Setting{
		Period: 30,
	},
	Wh: Wallhaven{
		Period: 10,
		Param: WallhavenParam{
			Page:       []int{1, 2, 3},
			Categories: "anime",
			Tag:        "anime",
		},
	},
}

func LoadConfig(name string) (*Config, error) {
	_, err := os.Stat(name)
	if err != nil && os.IsNotExist(err) {
		defaultCfg.Save(name)
		return &defaultCfg, nil
	}

	var cfg Config

	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (c Config) Save(name string) error {
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer f.Close()
	if err != nil {
		return err
	}

	data, err := json.Marshal(&c)
	if err != nil {
		return err
	}

	startIndex := 0
	for {
		l, err := f.Write(data[startIndex:])
		if err != nil && startIndex+l < len(data) {
			startIndex = startIndex + l
		}
		break
	}

	f.Sync()

	return nil
}
