## KDE命令行更换壁纸

更改下面的代码 替换/PATH/TO/IMAGE.png 这一部分即可

```bash
dbus-send --session --dest=org.kde.plasmashell --type=method_call /PlasmaShell org.kde.PlasmaShell.evaluateScript 'string:
var Desktops = desktops();
for (i=0;i<Desktops.length;i++) {
        d = Desktops[i];
        d.wallpaperPlugin = "org.kde.image";
        d.currentConfigGroup = Array("Wallpaper",
                                    "org.kde.image",
                                    "General");
        d.writeConfig("Image", "file:///PATH/TO/IMAGE.png");
}'
```

Replace /PATH/TO/IMAGE.png with appropriate path to wallpaper.

[原博客](http://ivo-wang.github.io/2018/02/27/kde-wallpaper-command-set/)

## Gnome更换锁屏和背景壁纸

在设置桌面及锁屏背景的时候，注意Picture标签下只显示~/Pictures文件夹下的图片。如果您想使用不在该文件夹下的图片，请使用下列命令：

### 对于桌面背景：

```bash
$ gsettings set org.gnome.desktop.background picture-uri 'file:///path/to/my/picture.jpg'
```

### 对于锁屏背景

```bash
$ gsettings set org.gnome.desktop.screensaver picture-uri 'file:///path/to/my/picture.jpg'
```

[原文地址](https://wiki.archlinux.org/index.php/GNOME)