package main

var cmdch chan int

const (
	version string = "2.1.0"
)

func getVersion() string {
	return version
}

func next() {
	cmdch <- 0
}

func main() {
	m := NewManager()
	c := NewCenter()
	c.init(m)
	m.setCenter(c)
	c.SyncData()

	go m.DownloadHandle()
	go m.SettingHandle()

	c.ShowAndRun()
}
