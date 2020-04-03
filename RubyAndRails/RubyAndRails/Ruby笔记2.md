## Gem本地服务器
#### 内置的gem server
```ruby
ruby -v
rvm gemset list
gem list
gem server
=>Server started at http://0.0.0.0:8808
```

#### geminabox
```ruby
gem install geminabox
mkdir data
vi config.ru
  require "rubygems"
  require "geminabox"

  Geminabox.data = "./data"
  run Geminabox::Server
rackup
=>Host:  http://localhost:9292
# push gem to the server
gem inabox secretgem-0.0.1.gem
```
如果不是共享的服务，放在本地：
```ruby
source "file://gem_sources_folder" do
  gem 'secretgem'
end
```


## Ruby中执行shell外部命令
``用来执行shell外部命令
```ruby
puts `date`
2016年 10月 14日 星期五 11:14:40 CST
puts %x(date)
```


## 方法参数里面的*
```ruby
# a 是一个常规的参数，b 将会得到第一个参数后面的所有参数并把他们放入一个数组中， **c 将会得到传人方法的任何格式为 key: value 的参数。
def my_method(a, *b, **c)
end
my_method(1, 2, 3, 4, a: 1, b: 2)
# => [1, [2, 3, 4], {:a => 1, :b => 2}]
```


## to_s 和 inspect 和 to_str之间的区别
```ruby
class Hooopo
   def to_s
     "to_s"
   end  
   def inspect
     "inspect"
   end  
end  
=> nil
pry(main)> hooopo = Hooopo.new
=> inspect
pry(main)> puts hooopo
to_s
=> nil
pry(main)> print hooopo
to_s=> nil
pry(main)> p hooopo
inspect
=> inspect

class Hooopo
  def to_str
    "to_str"
  end

  def to_s
    "to_s"
  end
end

ruby-1.9.3-p0 :075 > hooopo = Hooopo.new
 => to_s
ruby-1.9.3-p0 :076 > "hello #{hooopo}"
 => "hello to_s"
ruby-1.9.3-p0 :078 > ["hello", hooopo].join(" ")
 => "hello to_str" #这里也说明在join的时候数组里不一定要都是string啊，只要能响应to_str或to_s方法就OK！
ruby-1.9.3-p0 :079 > "hello " + hooopo
 => "hello to_str"
ruby-1.9.3-p0 :080 > File.join("hello", hooopo)
 => "hello/to_str"
ruby-1.9.3-p0 :081 >
```


### === 方法
a === b，的含义可以粗略的描述成表示假设a是一个集合，那么b属于a吗？
```ruby
String === "hi"       # True
Object === "hi"       # True
Integer === "hi"      # False
(1..100) === 3        # True
(1..100) === 200      # False
# 数组类型并没有实现===方法，所以在数组对象上调用===，实际上调用的是Object对象的===方法
[1, 2, 3] === 1      # False
```


## Ruby里面的%Q, %q, %W, %w, %x, %r, %s, %i
```ruby
%i：用于生成一个symbol数组
%i(a b c)
=> [:a, :b, :c]

%s:用于表示symbol，但是不会对其中表达式等内容进行转化
%s(foo)
=> :foo

%r:用于正则表达式
%r(/home/#{foo})
=> "/\\/home\\/Foo/"

%x:执行一段shell脚本并返回标准输出内容
%x(echo foo:#{foo})
=> "foo:Foo\n"

%w:用于表示其中元素被单引号括起的数组，比较奇怪的是\(斜杠空格)会被转化成(空格)，但是其他的内容不会.
%w(a b c\ d \#e #{1}f)
=> ["a", "b", "c d", "\\#e", "\#{1}f"]

%W:用于表示其中元素被双引号括起的数组
%W(#{foo} Bar Bar\ with\ space)
=> ["Foo", "Bar", "Bar with space"]

%Q:用于替代双引号的字符串. 当你需要在字符串里放入很多引号时候, 可以直接用下面方法而不需要在引号前逐个添加反斜杠 (\")
what_frank_said = "hello autodesk"
%Q(Joe said: "Frank said: "#{what_frank_said}"")
=> "Joe said: \"Frank said: \"hello autodesk\"\""

%q:表示的是单引号字符串
%q(Joe said: 'Frank said: '#{what_frank_said} ' ')
```


## Thread.abort_on_exception
```ruby
# default: false
# true: all threads will abort (the process will exit(0)) if an exception is raised in any thread
Thread.abort_on_exception
```


## dup和clone的区别
参考[Ruby | Rails - 浅拷贝 | 深拷贝](https://ruby-china.org/topics/22164)
dup不会有单例方法和冻结状态，clone可以
```ruby
# diff 1
a = Object.new
def a.foo
  :foo
end
p a.foo
b = a.dup
p b.foo

a = Object.new
def a.foo
  :foo
end

# diff 2
b = a.clone
p b.foo
a = Object.new
a.freeze
p a.frozen?
b = a.dup
p b.frozen?
c = a.clone
p c.frozen?
```


## Ruby Pre-defined variables
看源码有时候会遇到一些不知道什么意思的全局变量或者常量，比如`$:, STDIN`，官方文档有个列表可以参照下。
[Pre-defined variables](https://ruby-doc.org/core-2.1.1/doc/globals_rdoc.html)


## 并行匹配赋值
```
a,b = 1,2 #=> a=1,b=2
a,b,c = 1,2 #=> a=1,b=2,c=nil
a,b = [1,2,3] #=> 效果和上面是一样的
a,b = b,a #=> 交换a,b
```


## Ruby interface for data serialization in YAML format.
>store a objec instance
```ruby
user = User.first
user_fake = YAML.dump(user)
user_real = YAML.load(user_fake)
```


## Ruby中require load autoload的区别
require
1. kernel method，可以加载 ruby 文件，也可以加载外部的库。
2. 相比 load ,针对同一个文件，它只加载一次

load
1. 与 require 很类似，但是load会每次都重新加载文件。
2. 大部分情况下，除非你加载的库变动频繁，需要重新加载以获取最新版本，一般建议用 require 来代替 load.

autoload
用法稍稍不同：autoload(const_name, 'file_path'), 其中 const_name 通常是模块名，或者类名。
对于 load 和 require，在 ruby 运行到 require／load 时，会立马加载文件，而 autoload 则只有当你调用 module 或者 class 时才会加载文件。 Autoload will be dead, 不建议使用

1. require并非线程安全
2. 全局变量 类实例变量 类变量都应该是只读的，避免写操作