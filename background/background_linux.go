package background

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/godbus/dbus/v5"
)

type DESKTOP int32

const (
	KDE DESKTOP = iota
	DDE
	GNOME
	UNKNOWN
)

func SetBg(file string) error {
	switch getCurrentDesktop() {
	case KDE:
		return setBgForKDE(file)
	case DDE:
		return setBgForDDE(file)
	case GNOME:
		return setBgForGNOME(file)
	}
	return nil
}

func setBgForKDE(file string) error {
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

func setBgForDDE(file string) error {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}

	prop, err := conn.Object("com.deepin.daemon.Display", "/com/deepin/daemon/Display").GetProperty("com.deepin.daemon.Display.Monitors")
	if err != nil {
		return err
	}

	monitors := prop.Value().([]dbus.ObjectPath)
	if len(monitors) == 0 {
		return fmt.Errorf("monitor is empty")
	}

	m, err := conn.Object("com.deepin.daemon.Display", monitors[0]).GetProperty("com.deepin.daemon.Display.Monitor.Name")
	if err != nil {
		return err
	}

	name := m.Value().(string)
	urlFile, err := filepath.Abs(file)
	if err != nil {
		return err
	}
	call := conn.Object("com.deepin.daemon.Appearance", "/com/deepin/daemon/Appearance").Call("com.deepin.daemon.Appearance.SetMonitorBackground", 0, name, "file://"+urlFile)
	if call.Err != nil {
		return call.Err
	}

	return nil
}

func setBgForGNOME(file string) error {
	return nil
}

func getCurrentDesktop() DESKTOP {
	if os.Getenv("XDG_CURRENT_DESKTOP") == "KDE" || os.Getenv("XDG_SESSION_DESKTOP") == "KDE" {
		return KDE
	} else if os.Getenv("XDG_CURRENT_DESKTOP") == "deepin" || os.Getenv("XDG_SESSION_DESKTOP") == "deepin" {
		return DDE
	}
	return UNKNOWN
}
