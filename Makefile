
all:build

build:wallpaper-toolbox

wallpaper-toolbox:
	go build -o $@

install:
	mkdir -pv ${DESTDIR}${PREFIX}/bin
	cp -rf wallpaper-toolbox ${DESTDIR}${PREFIX}/bin/

	mkdir -pv ${DESTDIR}${PREFIX}/share/applications/
	cp -rf misc/wallpaper-toolbox.desktop ${DESTDIR}${PREFIX}/share/applications/

	mkdir -pv ${DESTDIR}${PREFIX}/share/icons/hicolor/16x16/apps
	cp -rf misc/icons/wallpaper-toolbox_16px.png ${DESTDIR}${PREFIX}/share/icons/hicolor/16x16/apps/wallpaper-toolbox.png
	mkdir -pv ${DESTDIR}${PREFIX}/share/icons/hicolor/32x32/apps
	cp -rf misc/icons/wallpaper-toolbox_32px.png ${DESTDIR}${PREFIX}/share/icons/hicolor/32x32/apps/wallpaper-toolbox.png
	mkdir -pv ${DESTDIR}${PREFIX}/share/icons/hicolor/64x64/apps
	cp -rf misc/icons/wallpaper-toolbox_64px.png ${DESTDIR}${PREFIX}/share/icons/hicolor/64x64/apps/wallpaper-toolbox.png
	mkdir -pv ${DESTDIR}${PREFIX}/share/icons/hicolor/72x72/apps
	cp -rf misc/icons/wallpaper-toolbox_72px.png ${DESTDIR}${PREFIX}/share/icons/hicolor/72x72/apps/wallpaper-toolbox.png
	mkdir -pv ${DESTDIR}${PREFIX}/share/icons/hicolor/128x128/apps
	cp -rf misc/icons/wallpaper-toolbox_128px.png ${DESTDIR}${PREFIX}/share/icons/hicolor/128x128/apps/wallpaper-toolbox.png

clean:
	rm -rf wallpaper-toolbox