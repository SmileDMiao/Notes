## Mina
在rails项目中添加
```ruby
gem 'mina'
```
执行 bundle install 安装gem，运行 mina init 生成mina部署的配置文件

### [Mina配置文件](https://github.com/SmileDMiao/personal-practice/blob/master/config/deploy.rb)
改配置文件可以使用3个命令：mina setup; mina deploy; mina restart
第一次必须依次运行 *mina setup* *mina deploy*


### Mina配置文件的一些说明
---
```ruby
# 远程服务器用户
set :user, 'miao'
# 远程服务器地址
set :domain, '139.224.133.155'
# 部署的目录，没有的话需要手动去新建
set :deploy_to, '/home/miao/practice/personal-practice'
# 项目地址
set :repository, 'https://github.com/SmileDMiao/personal-practice.git'
# 项目分支
set :branch, 'master'
#如果设置为：pretty，则使用缩进来优化输出。否则，只需调用命令即可
# term_mode :default :pretty,但是在mac上会出现问题，没有任何输出
set :term_mode, nil

# shared_paths: 共享目录，有些文件是不希望别人看到的，或是无需上传的，比如database.yml的配置，不可能把生产环境的配置
# 上传，所以这里设置其为共享目录，代码下载之后，利用软连接的方式指到分享的目录，而不使用远程代码库的对应文件
set :shared_paths, ['config/database.yml', 'log', 'tmp/pids']

# 利用rvm设置版本
task :environment do
  invoke :'rvm:use[2.3.0]'
end

# mina setup
# 在服务器上创建shared_paths相关内容，并赋予权限
task :setup => :environment do
  # queue! %[mkdir -p "#{deploy_to}/#{shared_path}/log"]
  # queue! %[chmod g+rx,u+rwx "#{deploy_to}/#{shared_path}/log"]ids"]
end
```


Nginx配置
```shell
upstream app {
    # Path to Unicorn SOCK file, as defined previously
    server unix:/tmp/unicorn_blog.sock fail_timeout=0;
}
server {
    listen 80;
    server_name localhost;

    # Application root, as defined previously
    # 由于mina会默认生成current文件夹，确保改目录位置的正确
    root /home/miao/practice/personal-practice/current/public;

    try_files $uri/index.html $uri @app;

    location @app {
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header Host $http_host;
        proxy_redirect off;
        proxy_pass http://app;
    }

    error_page 500 502 503 504 /500.html;
    client_max_body_size 4G;
    keepalive_timeout 10;
}
```


Rails production.rb配置
```ruby
# nginx 管理静态资源
config.action_dispatch.x_sendfile_header = 'X-Accel-Redirect'
```

---
### Mac上deploy无输出
mina有个term_mode的设置，默认pretty，是用来优化输出的，但是mac上需要设置其为nil才可 mina deploy
```shell
set :term_mode, nil
```

---
### Mina deploy 时出现 locale 问题:Perl: warning: Setting locale failed
```shell
sudo vim /etc/environment 加上 LC_ALL=C
```

---
### Mina console 无法输入中文
本机的console可以输入中文，但是mina console在远程打开console无法输入中文。
我这里的问题是和ssh的配置有关
本地的ssh配置
/etc/ssh/ssh_config里面有 SendEnv LANG LC_*
而远程ssh服务端是没有这个配置的，本地注释掉该配置，问题解决

---
### sidekiq 无法启动
sidekiq 默认配置 config/sidekiq.yml
但是mina部署的时候sidekiq老是无法启动，提示找不到pid，设置到tmp／pids下，日志到log下，setup 时建了对应到shared文件夹，也建立了软连接，但还是无法启动
```ruby
# config／deploy.rb
set :sidekiq_pid, "#{deploy_to}/shared/tmp/pids/sidekiq.pid"
```

---
### 命令
+ 使用`mina -T` 可以看到很多方便的帮助命令
+ 如果是引入的gem中有静态资源需要编译，发布时需要强制`mina deploy force_assets=1`， 不然如果没有引入新的静态资源的话，而如果有添加gem中的静态资源，将跳过预编译，但实际这些是需要编译的

---
### Mina whenever-task
[mina-whenever.rb](https://gist.github.com/SmileDMiao/b91b1ca6b1c5a7c90f03260e10ac0be4)

---
### Mina run rake task
```ruby
mina 'rake[namespace:task]'
```
由于没有使用最新版本的mina，但是whenever plugin 需要mina版本为新版mina,这里我们可以参考whenever plugin 插件的写法，写一个自己的task，引入deploy.rb(**require_relative**)

---
## 部署多个环境
如果同一套代码，需要部署到测试环境，stag环境，生产环境,可以建立不同的deploy文件，部署时可以通过 -f 来指定配置文件
在配置文件中，可以通过 **set :ails_env** 这个变量来区分不同的环境,这个要看一些周边插件的设置，最起码 puma,unicorn都是通过这个来区分环境配置的。

---
## 热部署
我自己的部署方案是mina,unicorn,nginx，在真实的项目中，一般都有热部署的需求，总不能因为发布服务就停了吧。mina的部署脚本中的重启unicorn部分是可以热重启的，那么是怎么做到的呢？
[相关代码部分](https://github.com/scarfacedeb/mina-unicorn/blob/master/lib/mina/unicorn/utility.rb)见这里：
其中涉及到热重启的核心部分就是：
```shell
kill -s USR2 pid
```
此时旧的主进程将会把自己的进程文件改名为.oldbin，然后启动新的进程服务。新旧进程共同处理请求。
这里mina的unicorn插件的处理方式是这样的：
先是kill发送usr2信号到老的进程，然后启动新的服务。之后等待2秒（默认，可以修改）如果老的服务还在运行，那么就直接发送quit信号到老的服务。

---
## 内存占用过高
[puma_worker_killer](https://github.com/schneems/puma_worker_killer)
[unicorn_worker_killer](https://github.com/kzk/unicorn-worker-killer)