## 如何DEBUG
---
[Gem:pry](https://github.com/pry/pry/wiki/Source-browsing)

```ruby
# 查看方法的源码位置
DRb.method(:thread).class => Method
DRb.method(:thread).source_location

# 查看include了多少模块
Array.included_modules

# 查看当前执行的调用栈
caller

# 查看继承链
ActiveRecord::Base.ancestors

# 方法或者proc接收多少个参数
"".method(:upcase).arity
String.instance_method(:upcase).arity

# Gem:pry中
# 查看方法位置
$ Post.create
$ ActiveRecord::Base#save!
# 查看源码位置
show-source object

# 查看某个类的子类(直接的子类[定义在active_support中])
Array.subclasses

# 查找方法
String.methods.grep(/length/)
```


## Ruby里的尾递归
---
斐波纳切数列大家都知道
1 1 2 3 5 8 11 19
给你一个n，求出第n个数是什么

+ 最基本的递归方式
> 但是这种方式在比较大的时候速度十分缓慢，因为递归调用栈太多，内存消耗极大。

```ruby
def fib(n)
  return 1 if n == 1 || n == 2
  return fib(n - 1) + fib(n - 2)
end
```

+ 循环方式
> 这种方式比较容易理解， 算出最后一个然后调换位置

```ruby
def fib(n)
  first = second = third = 1
  while n > 2
    third = first + second
    first = second
    second = third
    n = n - 1
  end
  return third
end
```

+ 尾递归方式
> 尾递归就是函数返回值还是这个函数本身，每次调用都已经将结果计算出来，不用再返回之前的调用栈.

```ruby
def fib(n, a = 1, b = 1)
  return a if n == 1
  return fib(n - 1, b, a + b)
end
```
但是ruby默认是不支持尾递归优化的，ruby vm提供了打开尾递归优化的方式：
```ruby
RubyVM::InstructionSequence.compile_option = {
  tailcall_optimization: true,
  trace_instruction: false
}
```

+ 矩阵方式
```ruby
require 'matrix'
def fib(n)
  (Matrix[[1,1],[1,0]] ** n)[1,0]
end
```


## Ruby线程安全
---
Mutex:互斥锁
>线程要访问一段代码前，先获得一把锁，其他的线程就不能在锁被释放前访问这段代码，只能等锁释放了，其他线程才能进来。Rails之所以是线程安全的原因就在于使用了这种机制，每个请求(action)，最终就是会进入这样一个互斥锁的控制。

线程安全的数据类型:Queue
```ruby
require 'thread'

queue = Queue.new

producer = Thread.new do
  5.times do |i|
    sleep rand(i) # simulate expense
    queue << i
    puts "#{i} produced"
  end
end

consumer = Thread.new do
  5.times do |i|
    value = queue.pop
    sleep rand(i/2) # simulate expense
    puts "consumed #{value}"
  end
end

consumer.join
```

条件变量:ConditonVariable
>增强互斥锁Mutex，可以在关键部分暂停，直到资源可用。

```ruby
require 'thread'

mutex = Mutex.new
resource = ConditionVariable.new

a = Thread.new {
  mutex.synchronize {
    # Thread 'a' now needs the resource
    resource.wait(mutex)
    # 'a' can now have the resource
  }
}

b = Thread.new {
  mutex.synchronize {
    # Thread 'b' has finished using the resource
    resource.signal
  }
}
```


## alias vs alias_method
---
```ruby
# 用法区别
class User
  def full_name
    puts "Johnnie Walker"
  end
  alias name full_name
  alias :name, :full_name
end
User.new.name #=>Johnnie Walker

# alias是关键字, alias将self视为读取源代码时self的值。相反，alias_方法将self视为运行时确定的值。
class User
  def full_name
    puts "Johnnie Walker"
  end

  def self.add_rename
    alias_method :name, :full_name
  end
end

class Developer < User
  def full_name
    puts "Geeky geek"
  end
  add_rename
end
Developer.new.name #=> 'Gekky geek'

class User
  def full_name
    puts "Johnnie Walker"
  end

  def self.add_rename
    alias :name :full_name
  end
end

class Developer < User
  def full_name
    puts "Geeky geek"
  end
  add_rename
end
Developer.new.name #=> 'Johnnie Walker'
```


## Block Proc 和 Lambda 的区别
---
```ruby
# lambda
# 创建方式等同
lambda {|x| x + 1}
->(x) {x + 1}

# proc
Proc.new{|x| x + 1}
```
1. Proc 和 Lambda 都是对象，而 Block 不是
2. 参数列表中最多只能有一个 Block，但是可以有多个 Proc 或 Lambda
3. Lambda 对参数的检查很严格，而 Proc 则比较宽松
4. Proc 和 Lambda 中return关键字的行为是不同的, lambda中，return表示从这个lambda中返回，proc中表示从定义这个proc的作用域返回。
5. 用lambda创建的proc称为lambda,其他方式创建的称为proc,可以通过Proc#lambda?方式判断

```ruby
def f1
  yield
end

# 注意&p 不是参数，&p 类似于一种声明，当方法后面有 block 时，会将 block 捕捉起来存放到变量 p 中，如果方法后面没有 block，那么&p 什么也不干
def f2(&p)
  p.call
end

def f3(p)
  p.call
end

f1 { puts "f1" }
f2 { puts "f2" }
f3(proc{ puts "f3"})
f3(Proc.new{puts "f3"})
f3(lambda{puts "f3"})
```

&操作符: 这是一个proc对象我想当成代码块使用
```ruby
# Symbol#to_proc
[1,2,3].map(&:to_i)
```

## Ruby 中的回调
---
### Method相关
+ respond_to_missing?: 当尝试查看一个missing的方法时执行。
+ method_missing: 当调用一个不存在的方法时执行。
+ method_added: 当定义一个方法时执行
+ method_removed: 当一个方法被移除时执行
+ singleton_method_added: 当添加一个singleton方法时执行。
+ singleton_method_removed: 当一个singleton方法被移除时执行。
+ method_undefined: 当一个方法被undefined的时候执行，undef和remove的区别在于当你在子类中你可以undef掉从父类继承而来的方法，而remove则不可以删除定义在父类中的方法。在父类中无论是使用undef和remove子类都无法继续使用这些方法。
+ singleton_method_undefined: 当一个singleton方法被undefined时刻执行。

### Class/Object相关
+ inherited: 当被继承时执行。
+ initialize_copy: 当调用initialize_clone和initialize_dup时执行。
+ initialize_dup: 当调用dup，返回dup结果前时执行。
+ initialize_clone: 当调用clone，进行frozen判断和返回clone结果之前执行。

### Modules相关
+ included: 当include时执行。
+ append_features: 当include时执行。
+ prepended: 2.0新增prepend的时候执行。
+ prepend_features：跟上面的append_feature可以一同理解。
+ extend_object: 当被extend时执行。
+ extended: 当被extend时执行。
const_missing: 当有常量丢失时执行。

## Difference Between freezing string or not? 
---
```ruby
 insns = RubyVM::InstructionSequence.compile '"hello world".freeze'
 insns = RubyVM::InstructionSequence.compile '"hello world"'
 
 puts insns.disasm
```
1. freeze: opt_str_freeze
2. normal: putstring

## refine
---
```ruby
// 可以在想要打补丁的地方 using 这个模块，这样就可以只在相应打地方使用该扩展，而不会影响到其他地方
module MonkeyPatch
  refine Hash do
    def deep_merge
    end
  end
end 

using MonkeyPatch
```

## to_s 和 inspect 和 to_str之间的区别
---
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
---
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
---
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
---
```ruby
# default: false
# true: all threads will abort (the process will exit(0)) if an exception is raised in any thread
Thread.abort_on_exception
```


## dup和clone的区别
---
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
---
看源码有时候会遇到一些不知道什么意思的全局变量或者常量，比如`$:, STDIN`，官方文档有个列表可以参照下。
[Pre-defined variables](https://ruby-doc.org/core-2.1.1/doc/globals_rdoc.html)


## 并行匹配赋值
---
```
a,b = 1,2 #=> a=1,b=2
a,b,c = 1,2 #=> a=1,b=2,c=nil
a,b = [1,2,3] #=> 效果和上面是一样的
a,b = b,a #=> 交换a,b
```


## Ruby interface for data serialization in YAML format.
---
>store a objec instance
```ruby
user = User.first
user_fake = YAML.dump(user)
user_real = YAML.load(user_fake)
```


## Ruby中require load autoload的区别
---
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

## Ruby中执行shell外部命令
---
``用来执行shell外部命令
```ruby
puts `date`
2016年 10月 14日 星期五 11:14:40 CST
puts %x(date)
```


## 方法参数里面的*
---
```ruby
# a 是一个常规的参数，b 将会得到第一个参数后面的所有参数并把他们放入一个数组中， **c 将会得到传人方法的任何格式为 key: value 的参数。
def my_method(a, *b, **c)
end
my_method(1, 2, 3, 4, a: 1, b: 2)
# => [1, [2, 3, 4], {:a => 1, :b => 2}]
```