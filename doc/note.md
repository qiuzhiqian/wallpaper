# wallpaper开发笔记

一般的桌面都有一个自动切换壁纸的功能，但是这个功能有个很致命德缺点：只能切换本地指定目录下德壁纸。我一直想找一个能自动切换网络上面的壁纸应用，然而并没有发现合适的，所以打算自己写一个。

既然是要自己写一个网络壁纸应用，需要实现的功能也就是两个：
1. 获取网络壁纸
2. 设置壁纸

## 获取网络壁纸
既然是网络壁纸，获取方式无非是两种：
1. api
2. 爬虫

api方式自然是最高效和简单德，但是支持api方式的图片网站都很少，更别说壁纸网站了。
爬虫方式虽然麻烦一些，而且效率低一些，但是通用性却更高。

从多方面考虑后，前期德版本仅支持api方式，因为更加简单。

### api图片网站

API | Description | Auth | HTTPS | CORS |
|---|---|---|---|---|
| [Flickr](https://www.flickr.com/services/api/) | Flickr Services | `OAuth` | Yes | Unknown |
| [Getty Images](http://developers.gettyimages.com/en/) | Build applications using the world's most powerful imagery | `OAuth` | Yes | Unknown |
| [Gfycat](https://developers.gfycat.com/api/) | Jiffier GIFs | `OAuth` | Yes | Unknown |
| [Giphy](https://developers.giphy.com/docs/) | Get all your gifs | `apiKey` | Yes | Unknown |
| [Gyazo](https://gyazo.com/api/docs) | Upload images | `apiKey` | Yes | Unknown |
| [Imgur](https://apidocs.imgur.com/) | Images | `OAuth` | Yes | Unknown |
| [Lorem Picsum](https://picsum.photos/) | Images from Unsplash | No | Yes | Unknown |
| [Pexels](https://www.pexels.com/api/) | Free Stock Photos and Videos | `apiKey` | Yes | Yes |
| [Pixabay](https://pixabay.com/sk/service/about/api/) | Photography | `apiKey` | Yes | Unknown |
| [Pixhost](https://pixhost.org/api/index.html) | Upload images, photos, galleries | No | Yes | Unknown |
| [PlaceKitten](https://placekitten.com/) | Resizable kitten placeholder images | No | Yes | Unknown |
| [ScreenShotLayer](https://screenshotlayer.com) | URL 2 Image | No | Yes | Unknown |
| [Unsplash](https://unsplash.com/developers) | Photography | `OAuth` | Yes | Unknown |
| [Wallhaven](https://wallhaven.cc/help/api) | Wallpapers | `apiKey` | Yes | Unknown |

[资料原网址](https://github.com/public-apis/public-apis)

综合考虑后，最终选择使用Wallhaven网站，原因是该网站本身就是一个壁纸网站，所以图片更加适合壁纸德风格，而且壁纸数量庞大，api使用也比较简单，国内访问速度也不慢。

### wallhaven支持的api

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