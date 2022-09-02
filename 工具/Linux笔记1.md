## CURL使用
---
+ -b: 设置cookie
+ -i/I: 显示header/只显示header
+ -X: 请求方式
+ -A: 设置user-agent
curl post json数据
```shell
curl -H "Content-Type: application/json" -X POST  --data '{"data":"1"}'  http://127.0.0.1/

# -w: time_total: 总时间按秒计; time_connect: 连接时间,从开始到建立TCP连接完成所用时间,包括前边DNS解析时间;
# time_starttransfer: 开始传输时间。在发出请求之后，Web 服务器返回数据的第一个字节所用的时间
# time_namelookup: DNS解析时间
curl -w "time_connect: %{time_connect} time_starttransfer: %{time_starttransfer} time_nslookup:%{time_namelookup} time_total: %{time_total}" "https://api.weixin.qq.com"
```

&nbsp;

## 上传下载文件
---
### SCP
```shell
#从服务器下载文件
scp root@171.19.1.15:/home/abc.rb /home/miao/
#本地文件上传到服务器
scp /home/miao/news.rb root@171.19.1.15:/home/miao/
#复制本地的目录到服务器
scp -r /home/miao/music/ root@171.19.1.15:/root/others/
#复制服务器目录到本地
scp -r root@171.19.1.15:/root/others/ /home/miao/music/
```
### SFTP上传下载文件
scp命令用起来感觉有点麻烦,sftp用起来就很方便了.
```shell
sftp user@remoteaddress
put xx :上传
get xx :下载
lls :查看本地目录下的文件
mput *:上传文件夹下所有文件
help:Available Commands
```
### SFTP VS SCP
1. 速度: SCP is much faster than SFTP
2. 功能： SFTP support more functions than SCP
3. 安全: Both protocols run on SSH, provide the same security features.

&nbsp;

## Nohup后台运行任务
---
**先说一下linux重定向**：
0、1和2分别表示标准输入、标准输出和标准错误信息输出，可以用来指定需要重定向的标准输入或输出。
/dev/null 黑洞文件
```shell
# 后台运行shadowsocks，不打印日志
nohup sslocal -c /home/knight/software/shadowsocks/shadowsocks_wly.json >/dev/null 2>&1 &
```

&nbsp;

## Time 命令
---
用来检测命令的运行时间
```shell
#简单用法
time ruby test.rb
outputs:
real 10.03
user 0.02
sys 0.00
```

&nbsp;

## Centos7 Firewall 简单命令
---
*在centos上安了shadowsocks之后，本机连接补上报errno113，centos7中使用firewall，端口被这东西毙掉了*
```shell
# 添加端口
firewall-cmd --zone=public --add-port=6022/tcp --permanent
firewall-cmd --zone=public --add-port=6022/udp --permanent
# 删除一个端口
firewall-cmd --zone=public --remove-port=6022/tcp --permanent
firewall-cmd --zone=public --remove-port=6022/udp --permanent
# 重启
firewall-cmd --reload
firewall-cmd --complete-reload
```

&nbsp;

## Ubuntu防火墙简单使用
---
ubuntu默认使用的防火墙是 *ufw*，默认情况下是关闭的,这里只记录简单命令的使用：
```shell
# 开启
sudo ufw enable
# 关闭
sudo ufw disable
# 状态
sudo ufw status
# 允许访问80端口
sudo ufw allow 80
# 允许80端口的tcp协议许可
sudo ufw allow 80/tcp
# 删除
sudo ufw delete allow 80
# 允许ssh
sudo ufw allow ssh
# 日志开
sudo ufw logging on
# 日志关
sudo ufw logging off
```

&nbsp;

## Linux 添加用户
---
```shell
#添加用户
adduser(好)
useradd(不好)建出来的用户属于三无产品
#添加用户到sudo组
sudo usermod -aG sudo <username>
```

&nbsp;

## 查看某个命令所在目录位置
---
```shell
which -a command_name :-a是列出所有的，不加-a显示找到的第一个
```

&nbsp;

## 使用dd制作U盘启动盘
---
```shell
sudo dd if=xxx.iso of=/dev/sdb
```

&nbsp;

## Ubuntu dpkg apt-get的一些命令
---
dpkg -l |  grep vim

iU 表示软件包未安装成功

ii 表示安装成功

rc 表示软件包已经被卸载,但配置文件仍在(删除／卸载 dpkg --purge xxxx)

```shell
# 卸载,删除配置文档
apt-get remove --purge xxxx
# 搜索
apt-cache search xxxx
# 修复依赖
apt-get install -f
# 删除包与配置
apt-get clean/autoclean/remove/autoremove
apt-get --purge remove
```

&nbsp;

##  Samba文件分享
---
```shell
# 安装
apt-get insall samba
# 配置
[share]
   path = /home/miao/shared
   browseable = yes
   writable = yes
   comment = smb share
   public = yes (无需密码，否则是你创建的用户和密码)

# 添加samba账户(必须是一个系统已经存在的用户)
sudo smbpasswd -a smbuser
# 启动
sudo service smbd restart
```

&nbsp;

## Tar解压文件
---
+ -c: 建立压缩档案
+ -x：解压
+ -t：查看内容
+ -r：向压缩归档文件末尾追加文件
+ -u：更新原压缩包中的文件
+ -z：有gzip属性的
+ -j：有bz2属性的
+ -Z：有compress属性的
+ -v：显示所有过程
+ -O：将文件解开到标准输出
+ -f: 使用档案名字，切记，这个参数是最后一个参数，后面只能接档案名。
```shell
tar zxvf xxx
tar czvf xxx.zip xxx
```

&nbsp;

## 关于压缩文件
---
1. zip -(1-9: default6): 数字越大压缩率越高
2. 压缩率: 7z > zip > gzip
3. 压缩速度: gzip > 7z > zip

&nbsp;

## 查看进程
---
每一列的含义:
+ USER: 进程的所属用户
+ PID: 进程的进程ID号
+ %CPU: 进程占用的 CPU资源 百分比
+ %MEM: 进程占用的 物理内存 百分比
+ VSZ: 进程使用掉的虚拟内存量 (Kbytes)
+ RSS: 进程占用的固定的内存量 (Kbytes)
+ TTY: 与进程相关联的终端（tty),?代表无关,tty1-tty6是本机上面的登入者程序,pts/0表示为由网络连接进主机的程序
+ STAT: 进程的状态，具体见2.1列出来的部分
+ START: 进程开始创建的时间
+ TIME: 进程使用的总cpu时间
+ COMMAND: 进程对应的实际程序

命令行参数
+ -A: Display information about other users' processes, including those without controlling terminals.
+ -a: Display information about other users' processes as well as your own.  This will skip any processes which do not have a controlling terminal, unless the -x option is also specified.
+ -X: When displaying processes matched by other options, skip any processes which do not have a controlling terminal.
+ -x: When displaying processes matched by other options, include processes which do not have a controlling terminal.  This is the opposite of the -X option.  If both -X and -x are specified in the same command, then ps will use the one which was specified last.
+ -U: Display the processes belonging to the specified real user IDs.
+ -u: Display the processes belonging to the specified usernames.
```shell
ps -ef
ps aux
ps aux | grep ruby
```

&nbsp;

## Linux Process Signal
---
+ O: 进程正在处理器运行,这个状态从来没有见过.
+ S: 休眠状态（sleeping）
+ R: 等待运行（runable）R Running or runnable (on run queue) 进程处于运行或就绪状态
+ I: 空闲状态（idle）
+ Z: 僵尸状态（zombie）
+ T: 跟踪状态（Traced）
+ B: 进程正在等待更多的内存页
+ D: 不可中断的深度睡眠，一般由IO引起，同步IO在做读或写操作时，cpu不能做其它事情，只能等待，这时进程处于这种状态，如果程序采用异步IO，这种状态应该就很少见到了
其中就绪状态表示进程已经分配到除CPU以外的资源，等CPU调度它时就可以马上执行了。运行状态就是正在运行了，获得包括CPU在内的所有资源。等待状态表示因等待某个事件而没有被执行，这时候不耗CPU时间，而这个时间有可能是等待IO、申请不到足够的缓冲区或者在等待信号。
&nbsp;

## 文本搜索
---
### grep
+ -b: 显示行数
+ -v: 显示不包含匹配文本的所有行
+ -i: 忽略字符大小写的差别
+ -C n: 查看结果上下的n行
```shell
# 在文件中查找ruby
grep ruby file_name
# 在文件夹中查找ruby
grep ruby -r directoruy_name
# 查找文件中包含hello和ruby的
grep hello file_name | grep ruby
# 查找包含ruby但不包含mruby
grep ruby file_name | grep -v mruby
# 文件头5G
head -c 5GB file | grep "xx"
# 文件尾5G
tail -c 5GB | grep "xx"
```
### ack(默认递归搜索)
+ -c: count出现次数
+ -g: Print searchable files where the relative path + filename matches PATTERN.
+ --group: 默认
+ --nogroup: 不分组
+ -i: 忽略字符大小写的差别
+ --[no]ignore-dir=DIRNAME, --[no]ignore-directory=DIRNAME
+ -v: 显示不包含匹配文本的所有行
+ -w: Turn on "words mode"

```shell
ack -g foo = ack -f | ack foo
```
### ag(默认会忽略.gitignore中的文件)
+ -g PATTERN: 搜索匹配pattern的文件名
+ -o(--only-matching): 只打印匹配字串，不是一行
+ -l: 只打印匹配了的文件名，不打印行
+ -i: 大小写不敏感
+ -s: 大小写敏感
+ -S: 自动大小写敏感（有大写就敏感，全小写则不敏感，默认）


&nbsp;

## df
---
df命令用于显示目前在Linux系统上的文件系统的磁盘使用情况统计
```shell
df
# -h: "Human-readable" output
df -h
# only show locally-mounted filesystems
df -l
# -g: Use 1073741824-byte (1-Gbyte) blocks rather than the default
df -g
```

&nbsp;

## DU: display disk usage statistics
---
查看当前指定文件或目录(会递归显示子目录)占用磁盘空间大小
+ -s: 显示总和
+ -h: 以K，M，G为单位，提高信息的可读性

&nbsp;

## 创建一个指定大小的文件
---
```shell
mkfile 24m outputfile.out
dd if=/dev/zero of=output.dat  bs=24M  count=1
dd if=/dev/zero of=output.dat  bs=1M  count=24
```

&nbsp;

## 查看当前使用的shell
---
```shell
 echo $0;
```

&nbsp;

## Tmux
---
default prefix: `ctrl + b`
Session -> Window -> Panel

### 窗口(Window)操作
```shell
<prefix> c  创建新窗口
<prefix> w  列出所有窗口
<prefix> n  后一个窗口
<prefix> p  前一个窗口
<prefix> f  查找窗口
<prefix> ,  重命名当前窗口
<prefix> &  关闭当前窗口
```

### 窗格(Panel)操作
```
<prefix> "  垂直分割
<prefix> %  水平分割
<prefix> o  交换窗格
<prefix> x  关闭窗格
<prefix> <space> 切换布局
<prefix> q 显示每个窗格是第几个，当数字出现的时候按数字几就选中第几个窗格
<prefix> { 与上一个窗格交换位置
<prefix> } 与下一个窗格交换位置
<prefix> u 切换窗格最大化
<prefix> n 切换窗格最小化
<prefix> 0-9 切换到指定编号的窗%口
```

### 会话(Session)相关
```
<prefix> :new<回车>  启动新会话
<prefix> s           列出所有会话
<prefix> $           重命名当前会话
```

### .tmux.conf
```shell
# Set default shell to zsh
set-option -g default-shell /bin/zsh

set-window-option -g mouse on

# 设置状态栏显示内容和内容颜色。这里配置从左边开始显示，使用绿色显示session名称，黄色显示窗口号，蓝色显示窗口分割号
set -g status-left "#[fg=colour52]#S #[fg=yellow]#I #[fg=cyan]#P"
# 窗口信息居中显示
set -g status-justify centre
# 监视窗口信息，如有内容变动，进行提示
setw -g monitor-activity on
set -g visual-activity on
# 支持鼠标选择窗口，调节窗口大小
set -g mouse on
set -s escape-time 1
# 修改默认的窗口分割快捷键，使用更直观的符号
bind | split-window -h
bind - split-window -v
```
### tmux path_helper
*tmux nvm is not compatible with the npm config "prefix" option*
```shell
# .zshrc
if [ -f /etc/profile ]; then
    PATH=""
    source /etc/profile
fi
```