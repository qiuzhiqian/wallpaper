# KDE命令行更换壁纸

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