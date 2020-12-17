## 解决Centos中gnome桌面环境扩展插件安装

```shell
yum install gnome-shell-browser-plugin
```

&nbsp;

## Ubuntu无法安装gnome extensions
```shell
sudo apt-get install chrome-gnome-shell
```

&nbsp;

## 识别NTFS格式的硬盘或者U盘：
```shell
sudo yum install ntfs-3g
```

&nbsp;

## Centos安装vmvare：提示环境包不全，丢失：gcc,kernel
需要的环境包
```shell
yum install gcc gcc-c++
yum install kernel-devel
```

&nbsp;

## 在命令行中快速的移动光标
`ctrl + b` 向后移动
`ctrl + f` 向前移动
`ctrl + d` 删除当前字符
`ctrl + l` 清屏
`ctrl + a` 移到行首
`ctrl + e` 移到行尾
`meta + f` 向前移动一个单词
`meta + b` 向后移动一个单词

&nbsp;

## 解压zip文件遇到的问题
```shell
# 解压中文乱码
unzip -O CP936 xxx.zip

# 解压大文件报错
brew install p7zip
7za x big_file.zip

# 破解有密码的zip压缩文件
brew install fcrackzip
# -c 指定字符集，字符集 格式只能为 -c 'aA1!:'
# a 表示小写字母[a-z]
# A 表示大写字母[A-Z]
# 1 表示阿拉伯数字[0-9]
# ! 感叹号表示特殊字符[!:$%&/()=?{[]}+*~#]
# : 表示包含冒号之后的字符（不能为二进制的空字符），例如  a1:$%  表示 字符集包含小写字母、数字、$字符和%百分号
fcrackzip  -b -c 'aA1!:?><.,' -l 1-10   -u rubymine2020.2.zip
# PASSWORD FOUND!!!!: pw == qsb
# -b 表示使用暴利破解的方式
# -c 'aA1' 表示使用大小写字母和数字混合破解的方式
# -l 1-10 表示需要破解的密码长度为1到10位
# -u 表示只显示破解出来的密码，其他错误的密码不显示出
```

&nbsp;

## win10卸载ubuntu子系统方法
**warning：千万不要用windows来修改ubuntu系统的一些重要的文件，比如：/etc/sudoers**
```shell
lxrun /uninstall /full
```

&nbsp;

## SSH简单的安全措施
**说明:**之前在阿里云上安装了postgresqls，它会默认生成一个名为postgres的用户，我自己demo的数据库，密码就没有很复杂，有天我收到了阿里云异地登录的提醒，发现是这个用户。更多的我们搜索**SSH 最佳安全实践**
1. Ubuntu开启防火墙
2. 允许ssh:ufw allow ssh
3. 限制root用户登录，限制其他一些用户登录，只允许某一个用户登录.
4. 重启sshd: service ssh restart

ssh使用pem证书：
```shell
# 生成证书: -t: 类型, -f: 文件
ssh-keygen -t rsa -f my.pem
ssh -i xxx.pem name@ip
# 如果提示pem文件权限过高则: chmod 600 coco.pem
ssh-add -K xxx.pem
ssh name@ip
```

&nbsp;

## ICMP 协议
icmp是tcp/ip协议的一个子协议，面向无连接。
我们在ping一个服务器的时候，使用的就是icmp协议，如果服务器返回应答消息，则认为该服务器可达。
有时候ping不通一个服务器，可能是因为imcp协议被禁止了。
在ubuntu中，想要禁止只需要 将 /proc/sys/net/ipv4/icmp_echo_ignore_all 里的内容改成1 就可以，反之开启。

&nbsp;

## 改变命令行默认EDITOR
git commit 提交的时候默认的message填写界面是git；
gem open gemname 默认的编辑器也是vim
这里的方式是改成sublime
```shell
ln -s /Applications/Sublime\ Text.app/Contents/SharedSupport/bin/subl /usr/local/bin/subl
export PATH=/bin:/sbin:/usr/bin:/usr/local/sbin:/usr/local/bin:$PATH
export EDITOR='subl'
```

&nbsp;

## Telnet测试邮件服务器
```shell
telnet mail_address port
#HELO表示向服务器打招呼，后面内容不限
HELO smtp.163.com
250 OK
#告诉服务器你要登录
auth login
服务器返回334 VXN1cm5hbWU6 #这一串字符串表示“Username：”这是base64码
输入账户11111111@qq.com对应的base64码
输入密码对应的base64码
服务器返回235 Authentication successful  #表明身份认证成功可以发邮件了
quit # 退出
```

&nbsp;

## Mac上pip install遇到权限问题
mac的sip机制导致的。取消sip机制：重启电脑command+r进入恢复模式，左上角菜单里找到实用工具 -> 终端
输入csrutil disable回车重启Mac即可。
```shell
# 查看状态
csrutil status
```

&nbsp;

## 命令行里按Enter打印^M而不是新的一行
```shell
# 一般是由于stty terminal line setting问题
# 所有特殊字符均使用默认值
stty sane
```