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