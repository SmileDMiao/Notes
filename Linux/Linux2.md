## Flock文件锁

---
> 之前在使用whenenver(基于crontab)有遇到一个任务没有运行完成，下个任务就不需要运行这个需求，当时就有用到这个命令: flock
```ruby
# usage script_with_lock 'script_name', lock: 'lock_name'
job_type :script_with_lock, "cd :path && :environment_variable=:environment flock -n /var/lock/:lock.lock bundle exec script/:task :output"
# usage runner_with_lock 'ruby code', lock: 'lock_name'
job_type :runner_with_lock, "cd :path && flock -n /var/lock/:lock.lock script/rails runner -e :environment ':task' :output"
 Ad
```
+ -s, --shared：共享锁
+ -x, -e, --exclusive：独占锁，默认类型
+ -u, --unlock：解锁
+ -n, --nb, --nonblock：非阻塞，若指定的文件正在被其他进程锁定，则立即以失败返回
+ -w, --wait, --timeout seconds：若指定的文件正在被其他进程锁定，则等待指定的秒数；指定为0将被视为非阻塞
+ -o, --close：锁定文件后与执行命令前，关闭用于引用加锁文件的文件描述符
+ -E, --conflict-exit-code number：若指定-n时请求加锁的文件正在被其他进程锁定，或指定-w时等待超时，则以该选项的参数作为返回值
+ -c, --command command：运行无参数的命令


## 查看文件的Access/Modify/Change时间
---
```shell
stat file_name
```


## wc:计算文件有多少行
---
```shell
wc -l file_name
```


## List Open File(lsof):列出当前系统打开文件的工具
---
linux系统中一切皆文件
```shell
# 通过某个进程号显示该进程打开的文件
lsof -p pid
# 列出所有的网络连接
lsof -i
# 列出所有tcp 网络连接信息
lsof -i tcp
# 列出谁在使用某个端口
lsof -i :3306
# 列出某个用户的所有活跃的网络端口
lsof -a -u username -i
# 列出某个程序进程所打开的文件信息
lsof -c postgres
```

## 软连接
```shell
# ln -s 源文件 目标位置
ln -s /usr/local/mysql/bin/mysql /usr/bin
```