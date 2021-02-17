package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"wallpaper-toolbox/utils"

	"github.com/fsnotify/fsnotify"
)

type EventType int

const (
	ADD EventType = iota
	REMOVE
	UNKNOWN
)

type DataEvent struct {
	EventType EventType
	Name      string
}

type DataEngine struct {
	dir       string
	mux       sync.Mutex
	data      []string
	notifyChs []chan DataEvent
	tags      []string
}

func NewDataEngine(dir string, tags []string) *DataEngine {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	} else if err != nil {
		return nil
	}

	if tags == nil || len(tags) == 0 {
		return nil
	}

	localList := utils.GetLocalFile(dir, tags)

	return &DataEngine{
		dir:       dir,
		data:      localList,
		notifyChs: make([]chan DataEvent, 0),
		tags:      tags,
	}
}

func (d *DataEngine) reset() {
	d.mux.Lock()
	defer d.mux.Unlock()

	if d.data == nil || len(d.data) > 0 {
		d.data = make([]string, 0)
	}
}

func (d *DataEngine) Size() int {
	return len(d.data)
}

func (d *DataEngine) add(datas ...string) {
	if d.data == nil {
		return
	}

	d.mux.Lock()
	defer d.mux.Unlock()
	d.data = append(d.data, datas...)
}

func (d *DataEngine) removeByIndex(index int) error {
	if d.data == nil {
		return fmt.Errorf("data is nil")
	}

	if index < 0 || index >= len(d.data) {
		return fmt.Errorf("index is invalid")
	}

	d.mux.Lock()
	defer d.mux.Unlock()
	d.data = append(d.data[:index], d.data[index+1:]...)
	return nil
}

func (d *DataEngine) remove(item string) error {
	return d.removeByIndex(d.findIndex(item))
}

func (d *DataEngine) at(index int) (string, error) {
	if d.data == nil {
		return "", fmt.Errorf("data is nil")
	}

	if index < 0 || index >= len(d.data) {
		return "", fmt.Errorf("index is invalid")
	}

	d.mux.Lock()
	item := d.data[index]
	d.mux.Unlock()

	return item, nil
}

func (d *DataEngine) findIndex(item string) int {
	iret := -1
	if d.data == nil || len(d.data) == 0 {
		return iret
	}

	for i, v := range d.data {
		if v == item {
			iret = i
			break
		}
	}

	return iret
}

func (d *DataEngine) Register(ch chan DataEvent) {
	d.mux.Lock()
	d.notifyChs = append(d.notifyChs, ch)
	d.mux.Unlock()
}

func (d *DataEngine) Run() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	err = watcher.Add(d.dir)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			switch {
			case event.Op&fsnotify.Create == fsnotify.Create:
				//log.Println("Create file:", event.Name)
				if utils.IsFileMatch(event.Name, d.tags) {
					d.add(event.Name)
					for _, ch := range d.notifyChs {
						log.Println("new file:", event.Name)
						ch <- DataEvent{EventType: ADD, Name: event.Name}
					}
				}
			case event.Op&fsnotify.Remove == fsnotify.Remove || event.Op&fsnotify.Rename == fsnotify.Rename:
				//log.Println("rename file:", event.Name)
				if utils.IsFileMatch(event.Name, d.tags) {
					d.remove(event.Name)
					for _, ch := range d.notifyChs {
						log.Println("remove file:", event.Name)
						ch <- DataEvent{EventType: REMOVE, Name: event.Name}
					}
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}
