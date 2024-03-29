# wallpaper

一个定时自动切换壁纸的应用。使用golang实现，壁纸来源从wallhaven获取，然后随机选择一张设置为当前壁纸。

该软件使用纯golang语言实现，前端ui使用基于golang的fyne开发。

## 预览
![window预览1](https://raw.githubusercontent.com/qiuzhiqian/wallpaper/master/doc/img/win10_1.png)

![linux kde预览1](https://raw.githubusercontent.com/qiuzhiqian/wallpaper/master/doc/img/linux_kde_1.png)

## 特性
- 从wallhaven获取壁纸列表。
- 下载参数固定为categories=anime，后续可做灵活配置。
- 随机选择一张壁纸。
- 目前分辨率在配置文件中写死了，需要自己灵活配置
- 30分钟自动更换，这个时间后面会设置为可调整的。
- 之前下载的壁纸并不会自动清除，所以长时间运行后，壁纸占用磁盘空间会增大，可手动删除。后面会添加定时清理功能。
- 图片预览

## 系统支持
- windows
- linux(适配了kde和dde，gnome由于没有现成的环境暂未适配)

## 依赖
- [golang 1.16+](https://golang.google.cn/)
- [downloader](https://gitee.com/qiuzhiqian/downloader)
- [fyne](https://github.com/fyne-io/fyne)
- [fsnotify](https://github.com/fsnotify/fsnotify)

## 编译
linux下面：
```bash
$cd wallpaper
$go build
```
或者直接使用make也行
```bash
$cd wallpaper
$make
```

windows下面：
```
go build -ldflags="-H windowsgui"
```