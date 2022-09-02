## brew 依赖
---
```shell
# 删除的时候删除依赖
brew tap beeftornado/rmtree
# 显示依赖
brew deps --tree --installed

brew uninstall
brew autoremove
```

## Mac Imagemagick: Unable to revert mtime: /Library/Fonts
---
Edited /usr/local/Cellar/imagemagick/7.0.10-7/etc/ImageMagick-7/type.xml and replaced type-ghostscript with type-apple
```shell
# 合并两张照片
montage -geometry 100% left.jpeg right.jpeg merge.jpeg
```

## mac dock栏全屏不会自动隐藏
---
```shell
defaults write com.apple.dock autohide-delay -int 0
defaults write com.apple.dock autohide-time-modifier -float 1.0
killall Dock
```

## Mac上pip install遇到权限问题
---
mac的sip机制导致的。取消sip机制：重启电脑command+r进入恢复模式，左上角菜单里找到实用工具 -> 终端
输入csrutil disable回车重启Mac即可。
```shell
# 查看状态
csrutil status
```

## MacBook Pro 16inch 2019 外接显示器的情况下CPU爆满-kernel_task
---
1. 强制禁用独立显卡(老子加多少钱才有的独立显卡结果不能用?) `sudo pmset -a GPUSwitch 0`   `pmset -g` (0: 使用核显卡, 1: 独立显卡, 2: 自动切换)
2. 然后在电源 ( Energy Saver )配置里反选自动切换显卡 Automatic graphics switching
3. 取消选中"显示器具有单独的空间", 在 调度中心 (Mission Control)里反选 显示器具有单独的空间
4. 显示器设置选择和mac一样的色彩方案

## 快捷键
---
1. `command + shift + 5` 快速录屏
2. `command + shift + 4` 快速截取屏幕
3. `command + shift + 3` 快速截屏