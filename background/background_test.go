package background

import "testing"

func TestSetBg(t *testing.T) {
	err := SetBg(`E:\code\code_go\wallhaven\image\20200203\wallhaven-2e591y.jpg`)
	if err != nil {
		t.Log("err:", err)
	}
}
