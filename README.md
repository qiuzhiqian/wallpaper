# wallpaper

一个定时自动切换壁纸的应用。使用golang实现，壁纸来源从wallhaven获取，然后随机选择一张设置为当前壁纸。由于设置壁纸的方法，目前只知道windows的方法，linux桌面(gnome和kde)没有比较好的实现思路，所以该应用理论上只支持windows。当然壁纸下载部分是通用的。

## 特性
- 从wallhaven获取壁纸列表。
- 下载参数固定为categories=anime，后续可做灵活配置。
- 随机选择第一页中的一张。
- 30分钟自动更换，这个时间后面会设置为可调整的。
- 之前下载的壁纸并不会自动清除，所以长时间运行后，壁纸占用磁盘空间会增大，可手动删除。后面会添加定时清理功能。

## 使用
修改start.vbs文件中的路径
双击start.vbs