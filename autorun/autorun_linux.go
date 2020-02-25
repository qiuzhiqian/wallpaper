package autorun

import (
	"os"
)

func genService(fileName, serviceName string) error {
	content := `[Unit]
	Description=Wallpaper
	
	[Service]
	` + fileName + `
	
	[Install]
	WantedBy=default.target`

	_, err := os.Stat(serviceName)
	if err != nil {
		//文件异常
		var fd *os.File
		fd, err = os.OpenFile(serviceName, 0, 0644)
		if err != nil {
			return err
		}

		defer fd.Close()

		_, err = fd.WriteString(content)
		if err != nil {
			return err
		}
	}
	return nil
}

func Enable(yes bool) bool {
	return true
}

func IsEnable() bool {
	return true
}
