package background

import (
	"fmt"

	dbus "github.com/guelfey/go.dbus"
)

func SetBg(file string) error {
	fmt.Println("linux setbg")
	//var ret []string
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}

	param := `string:
	var Desktops = desktops();
	for (i=0;i<Desktops.length;i++) {
					d = Desktops[i];
					d.wallpaperPlugin = "org.kde.image";
					d.currentConfigGroup = Array("Wallpaper","org.kde.image","General");
					d.writeConfig("Image","` + file + `");
	}`

	call := conn.Object("org.kde.plasmashell", "/PlasmaShell").Call("org.kde.PlasmaShell.evaluateScript", 0, param)
	if call.Err != nil {
		return call.Err
	}
	return nil
}
