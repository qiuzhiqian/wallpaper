#
# Regular cron jobs for the wallpaper-toolbox package
#
0 4	* * *	root	[ -x /usr/bin/wallpaper-toolbox_maintenance ] && /usr/bin/wallpaper-toolbox_maintenance
