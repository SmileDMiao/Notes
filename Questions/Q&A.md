## Ping得到的IP可以访问百度却不可以访问知乎?
当我们用域名访问知乎的时候，CDN服务器可以根据访问的域名知道你想要的是哪个网站的资源，然后直接给你返回对应的资源。
但是当你用公网ip访问就不一样了，由于一个CDN服务器的公网ip对应多个域名网站，他不知道你想要的是哪个网站的资源，也就是说，当你用 118.89.204.192 去访问知乎的时候，CDN服务器不知道你要访问的是 zhihu.com，还是访问 a.com 或 b.com，所以他也干脆明了点，直接拒绝你的访问。
当客户端用域名访问知乎的时候，DNS会解析成对应的ip去访问CDN服务器，然后CDN服务器可以根据SNI机制获得该ip对应的来源域名，然后返回对应的资源。
SNI机制: 该机制主要是用来解决一个服务器对应多个域名时产生的一些问题，通过这种机制，服务器可以提前知道（还没建立链接）客户端想要访问的网站
为啥百度ip和域名都可以访问呢？这其实很简单，就是百度用的CDN服务器，只对应一个网站域名，百度有钱。

## 内存溢出与内存泄漏
内存泄露就是一块申请了一块内存以后，无法去释放掉这块内存，丢失了这段内存的引用；内存溢出就是申请的内存不够，撑不起我们需要的内存

## 幂等
一个幂等操作的特点是其任意多次执行所产生的影响均与一次执行的影响相同