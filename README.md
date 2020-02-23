# wallpaper

一个定时自动切换壁纸的应用。使用golang实现，壁纸来源从wallhaven获取，然后随机选择一张设置为当前壁纸。

## 预览
![window预览1](https://raw.githubusercontent.com/qiuzhiqian/wallpaper/master/doc/img/image_1.png)

![linux kde预览1](https://raw.githubusercontent.com/qiuzhiqian/wallpaper/master/doc/img/image_2.png)

![设置界面](https://raw.githubusercontent.com/qiuzhiqian/wallpaper/master/doc/img/image_3.png)

## 特性
- 从wallhaven获取壁纸列表。
- 下载参数固定为categories=anime，后续可做灵活配置。
- 随机选择第一页中的一张。
- 30分钟自动更换，这个时间后面会设置为可调整的。
- 之前下载的壁纸并不会自动清除，所以长时间运行后，壁纸占用磁盘空间会增大，可手动删除。后面会添加定时清理功能。

## 系统支持
- windows
- linux(仅限kde)

## 使用

### windows
双击应用程序，将开机启动勾选。

### linux kde
将wallpaper应用程序启动即可，推荐添加到开机中启动。
如果使用systemd管理开机启动项的话，则进行一下操作：  
1. 修改wallpaper.service中的路径为实际wallpaper路径
2. 复制wallpaper.service到/usr/lib/systemd/user中。
3. 执行systemctl --user enable wallpaper.service。这样该应用就添加到了开机启动中