package autorun

import (
	"strings"
	"wallpaper/utils"

	"golang.org/x/sys/windows/registry"
)

func Enable(yes bool) bool {
	regLoc := `Software\Microsoft\Windows\CurrentVersion\Run`
	execName := utils.GetCurrentDirectory() + "/wallpaper.exe"
	execName = strings.Replace(execName, "/", "\\", -1)

	regKey, err := registry.OpenKey(registry.CURRENT_USER, regLoc, registry.SET_VALUE|registry.QUERY_VALUE)
	if err != nil {
		return false
	}
	defer regKey.Close()

	var info string
	info, _, err = regKey.GetStringValue("Wallpaper")
	if (err != nil || (err == nil && info != execName)) && yes == true {
		err = regKey.SetStringValue("Wallpaper", execName)
		if err != nil {
			return false
		}
	} else if err == nil && yes == false {
		err = regKey.DeleteValue("Wallpaper")
		if err != nil {
			return false
		}
	}

	return true
}

func IsEnable() bool {
	regLoc := `Software\Microsoft\Windows\CurrentVersion\Run`
	execName := utils.GetCurrentDirectory() + "/wallpaper.exe"
	execName = strings.Replace(execName, "/", "\\", -1)

	regKey, err := registry.OpenKey(registry.CURRENT_USER, regLoc, registry.SET_VALUE|registry.QUERY_VALUE)
	if err != nil {
		return false
	}
	defer regKey.Close()

	var info string
	info, _, err = regKey.GetStringValue("Wallpaper")
	if err == nil && info == execName {
		return true
	}

	return false
}
