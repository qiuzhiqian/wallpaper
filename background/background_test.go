package background

import "testing"

func TestSetBg(t *testing.T) {
	err := SetBg(`E:\code\code_go\wallhaven\image\20200203\wallhaven-2e591y.jpg`)
	if err != nil {
		t.Log("err:", err)
	}
}

func TestSetBgLinux(t *testing.T) {
	err := SetBg("/home/xml/code/code_go/wallpaper/image/wallhaven-39e6pv.jpg")
	if err != nil {
		t.Log("err:", err)
	}
}
