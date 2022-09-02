## Nginx 帮助命令
---
nginx -h可以看到一些方便的nginx使用命令
nginx -V :version and configuration options
nginx -t :测试nginx配置文件（由有时可能需要root权限运行）是否正确
nginx -s (stop,quit,reopen,reload)

Mac上的日志默认写的是 logs／error.log logs/access.log  logs/nginx.pid，真实的目录位置在 /usr/local/Cellar/nginx/1.10.3/ (--prefix lcoation)
默认没有logs文件夹，还需要自己建立文件夹


## nginx监听端口无权限问题
---
一般rails服务，nginx都是监听的一个socket文件，但前两天和之前同事过来找我玩，遇到一个问题，在nginx下监听一个服务端口，发现总是不起作用，访问直接报错，查看日志显示：nginx没有访问那个端口的权限。
排除问题测试结果如下：
1. nginx启动方式：systemctl start nginx: 结果无法访问。
2. nginx直接启动: 结果可以访问
3. systemctl start nginx并且关闭selinux: 结果可以访问。

**结论**
看来这应该是selinux的一个安全措施，当以系统服务方式启动nginx的时候，nginx没有访问其他端口的权限，但是直接用nginx自己提供的方式启动却没有这个限制。

**那么如何才能即不关闭selinux,又以systemctl服务的方式启动nginx,同时又可以访问服务呢？**
```shell
 /usr/sbin/setsebool httpd_can_network_connect true
```
具体的解释可以[参见](https://www.nginx.com/blog/nginx-se-linux-changes-upgrading-rhel-6-6/
)：和之前的猜测基本一致。


## alias和root区别
---
root与alias主要区别在于nginx如何解释location后面的uri，这会使两者分别以不同的方式将请求映射到服务器文件上。
alias是一个目录别名的定义（仅能用于location上下文），root则是最上层目录的定义
```shell
# 当请求host/123/abc/logo.png时, 返回服务器上的 /data/www/123/abc/logo.png文件，即/data/www+/123/abc/
location ^~ /123/abc/ {
  root /data/www;
}

# 当请求host/123/abc/logo.png时, 返回服务器上的 /data/www/logo.png文件，即/data/www
location ^~ /123/abc/ {
  alias /data/www;
}
```