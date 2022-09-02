>简单点说, rack 是 Ruby web 应用的简单的模块化的接口. 它封装 HTTP 请求与响应, 并提供大量的实用工具.
它需要一个响应 call 方法的对象, 接受 env. 返回三元素的数组: 分别是 status code, header, body. 其中 status code 大于等于 100, 小于 600. header 是一个 hash, body 是一个响应 each 方法的数组.

## 数行代码的Rack示例
---
```ruby
# Rack::Directory.new('./')就是一个相应call方法的那个对象
require 'rack'
Rack::Handler::Thin.run Rack::Directory.new('./'), :Port => 9292

# 自己写最简单的例子
require 'rack'
class HelloWorld
  def call(env)
    [200, {"Content-Type" => "text/html"}, ["Hello Rack!"]]
  end
end
Rack::Handler::Thin.run HelloWorld.new, :Port => 9292
```

## Rack Middleware Stack
---
```ruby
# config.ru
# 将 body 标签的内容转换为全大写.
class ToUpper
  def initialize(app)
    @app = app
  end
  def call(env)
    status, head, body = @app.call(env)
    upcased_body = body.map{|chunk| chunk.upcase }
    [status, head, upcased_body]
  end
end
# 将 body 内容置于标签, 设置字体颜色为红色, 并指明返回的内容为 text/html.
class WrapWithRedP
  def initialize(app)
    @app = app
  end
  def call(env)
    status, head, body = @app.call(env)
    red_body = body.map{|chunk| "<p style='color:red;'>#{chunk}</p>" }
    head['Content-type'] = 'text/html'
    [status, head, red_body]
  end
end
# 将 body 内容放置到 HTML 文档中.
class WrapWithHtml
  def initialize(app)
    @app = app
  end
  def call(env)
    status, head, body = @app.call(env)
    wrap_html = <<-EOF
       <!DOCTYPE html>
       <html>
         <head>
         <title>hello</title>
         <body>
         #{body[0]}
         </body>
       </html>
    EOF
    [status, head, [wrap_html]]
  end
end
# 起始点, 只返回一行字符的 rack app.
class Hello
  def initialize
    super
  end
  def call(env)
    [200, {'Content-Type' => 'text/plain'}, ["hello, this is a test."]]
  end
end
use WrapWithHtml
use WrapWithRedP
use ToUpper
run Hello.new
```
直接运行rackup就可以运行上述代码
use 与 run 本质上没有太大的差别, 只是 run 是最先调用的. 它们生成一个 statck, 本质上是先调用 Hello.new#call,由下向上依次执行.
use ToUpper; run Hello.new本质上是完成如下调用 **ToUpper.new(Hello.new.call(env)).call(env)**


## Rails On Rack
---
Rails.application 是 Rails 应用的主 Rack 应用对象,rails server 负责创建 Rack::Server 对象和启动 Web 服务器。
```ruby
Rails::Server.new.tap do |server|
  require APP_PATH
  Dir.chdir(Rails.application.root)
  server.start
end
```

Rails::Server 继承自 Rack::Server，像下面这样调用 Rack::Server#start 方法：

```ruby
class Server < ::Rack::Server
  def start
    ...
    super
  end
end
```
中间件只加载一次，不会监视变化。若想让改动生效，必须重启服务器。

配置middleware
+ config.middleware.use(new_middleware, args)：在中间件栈的末尾添加一个中间件。
+ config.middleware.insert_before(existing_middleware, new_middleware, args)：在中间件栈里指定现有中间件的前面添加一个中间件。
+ config.middleware.insert_after(existing_middleware, new_middleware, args)：在中间件栈里指定现有中间件的后面添加一个中间件。
+ 可以使用 config.middleware.swap 替换中间件栈里的现有中间件：

## 参考
---
https://ruby-china.org/topics/21517
https://ruby-china.org/topics/24166