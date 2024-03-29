## File and Dir
---
```ruby
# 列出文件夹下的文件
Dir[local_tmp_path + '/*.zip']
# 查看文件夹下内容
Dir.entries(folder_path)
# 查看子文件夹的内容
Dir.children("dir")
# 最后一部分的文件名
File.basename(path)
# 最后部分的扩展名
File.extname(path)
# 是否存在
File.file?(file_path)
# 绝对路径
File.absolute_path(path)
# 是否文件夹
File.directory? name
# 创建文件夹
FileUtils.mkdir_p(local_tmp_path)
# 路径转换为绝对路径
File.expand_path(patah)
# 返回文件所在文件名
File.dirname(name)
```


## String
---
```ruby
# chomp
"hello".chomp                #=> "hello"
"hello\n".chomp              #=> "hello"
"hello\r\n".chomp            #=> "hello"
"hello\r".chomp              #=> "hello"
"hello".chomp("llo")         #=> "he"

# slice 截取字符串
slice!(integer) → new_str
slice!(integer, integer) → new_str
slice!(range) → new_str
slice!(regexp) → new_str
slice!(other_str) → new_str

# 分割字符串成数组
"abcdefg".scan(/.{3}/) # => ["abc", "def"]
"abcdef".scan(/.{3}/) # => ["abc", "def"]
```


## Struct
---
简单使用:
```ruby
#struct
Person = Struct.new(:age,:name,:sex)
me = Person.new(24,"spirit","male")
=> #<struct Person age=24, name="spirit", sex="male">
me.age #=> 24
me.name #=> "spirit"
me.sex #=> "male"
me.height #=> NoMethodError

#openstruct
require 'ostruct'
me = OpenStruct.new(age: 24, name: 'Spirit', sex: 'male')
me.height # => nil
me.height = '178'
me.height # => '178'
```


## Array#dig, Hash#dig, OpenStruct#dig
---
ruby2.3.0中的新特性：
这里hash的dig方法一开始没觉得怎么样，但是后面在做API的对接和开发，经常会有对nil继续取key的情况出现，恰好hash#dig可以优雅的解决这个问题
```ruby
# array
results = [[[1, 2, 3]]]
results.dig(0, 0, 0) # => 1
results.dig(0, 1, 2) # => nil

# hash
user = {
  user: {
    address: {
      street1: '123 Main street'
    }
  }
}

user.dig(:user, :address, :street1) # => '123 Main street'
user.dig(:user, :address, :street2) # => nil

# OpenStruct
address = OpenStruct.new('city' => "Anytown NC", 'zip' => 12345)
person = OpenStruct.new('name' => 'John Smith', 'address' => address)
person.dig(:address, 'zip') # => 12345
person.dig(:business_address, 'zip') # => nil
```


## attr_reader, attr_writer, attr_accessor, config_accessor
---
reader：只定义读的方法，writer：只定义写的方法，accessor:定义读和写的方法
config_accessor:
```ruby
require 'active_support/configurable'
include ActiveSupport::Configurable
class User
  include ActiveSupport::Configurable
  config_accessor :allowed_access
end
User.allowed_access
user = User.new
user.allowed_access
config_accessor :allowed_access, instance_reader: false, instance_writer: false
config_accessor :allowed_access, instance_accessor: false
config_accessor :hair_colors do
  [:brown, :black, :blonde, :red]
end
```


## protected private public
---
1. public方法可以被定义它的类和子类访问，并可以被类和子类的实例对象调用；
2. protected方法可以被定义它的类和子类访问，不能被类和子类的实例对象调用，但可以被该类的实例对象(所有)访问；
3. private方法可以被定义它的类和子类访问，不能被类和子类的实例对象调用，且实例对象只能访问自己的private方法。


## Array
---
```ruby
# product方法:返回两个数组元素所有的排列组合
[1,2,3].product([4,5,6])
=> [[1, 4], [1, 5], [1, 6], [2, 4], [2, 5], [2, 6], [3, 4], [3, 5], [3, 6]]
```

```ruby
# group_by
(1..6).group_by { |i| i%3 }   #=> {0=>[3, 6], 1=>[1, 4], 2=>[2, 5]}
```

```ruby
# partion: 若对某元素执行块的结果为真，则把该元素归入第一个数组；若为假则将其归入第二个数组,最后生成并返回一个包含这两个数组的新数组。
# partition {|item| ... }
(1..6).partition { |v| v.even? }
# => [[2, 4, 6], [1, 3, 5]]
```

```ruby
# chunk方法:之前在codewars上看到一个题目，实现一个unique_in_order方法,实现如下功能
unique_in_order('AAAABBBCCDAABBB') == ['A', 'B', 'C', 'D', 'A', 'B']
unique_in_order('ABBCcAD')         == ['A', 'B', 'C', 'c', 'A', 'D']
unique_in_order([1,2,2,3,3])       == [1,2,3]

def unique_in_order(iterable)
  (iterable.is_a?(String) ? iterable.chars : iterable).chunk { |x| x }.map(&:first)
end

'AAAABBBCCDAABBB'.chars.chunk{|x| x }.to_a
[["A", ["A", "A", "A", "A"]], ["B", ["B", "B", "B"]], ["C", ["C", "C"]], ["D", ["D"]], ["A", ["A", "A"]], ["B", ["B", "B", "B"]]]
```

```ruby
# inject: 接受的参数会成为blockh中的第一个参数,下面例子中result就是0
a = [{a: 1, b: 2, c: 3}, {a:3, b:2, c: 1}, {a:4, b: 5, c: 6}]
a.inject(0) {|result, h| result + h[:a]} => 8

# 数组遍历删除
# 这种做法是不对的，出现的结果和预期不一至
# 取而代之我们应该使用 delete_if 这个方法或者 reject!这个方法
a = [1, 2, 3, 4, 5]
a.each do |x|
  next if x < 3
  a.delete x
end
a #=> [1, 2, 4]

a.delete_if do |x|
  x< 3
end

# 删除指定位置的元素：
[1,2,3,4].delete_at(2)

# get the index of first match element
[1,2,3,4].index(2)

# each_slice
a = [1,2,3,4,5,6,7]
a.each_slice(3)
#=> <Enumerator: [1, 2, 3, 4, 5, 6, 7]:each_slice(3)>
a.each_slice(3).to_a
=> [[1, 2, 3], [4, 5, 6], [7]]


#操作单个对象和数组用同样的方式
obj  = 1
obj_arr = [1, 2, 3]
[*obj].each { |s| s }  # => [1]
[*obj_arr].each { |s| s } # => [1, 2, 3]
Array(stuff).each { |s| s } # => [1]
Array(stuff_arr).each { |s| s } # => [1, 2, 3]
```

```ruby
# combination 指定个数的所有排列组合
[1,2,3].combination(2).to_a # => [[1, 2], [1, 3], [2, 3]]
```

```ruby
# min min_by
a = %w(albatross dog horse)
a.min                                     #=> "albatross"
a.min { |a, b| a.length <=> b.length }    #=> "dog"
a.min(2)                                  #=> ["albatross", "dog"]
a.min(2) {|a, b| a.length <=> b.length }  #=> ["dog", "horse"]
a.min_by { |x| x.length }                 #=> "dog"
```

## <=>
---
a <=> b means
if a < b then return -1
if a = b then return  0
if a > b then return  1
if a and b are not comparable then return nil


## Hash
---
```ruby
# slice:返回given keys hash
{a: 1, b: 2, c: 3}.slice(:a, :b)

## with_indifferent_access:activesupport提供的方法，可以让hash无差别的访问
demo_1 = {a: 1}
=> {:a=>1}
demo_2 = demo_1.with_indifferent_access
=> {"a"=>1}
demo_2[:a]
=> 1
demo_2["a"]
=> 1

# 转换所有key为symbol
hash.deep_symbolize_keys
hash.deep_symbolize_keys!
# 转换所有key为string
hash.deep_stringify_keys
hash.deep_stringify_keys!

# assert_valid_keys: Validate all keys in a hash match *valid_keys, raising ArgumentError on a mismatch. Note that keys are treated differently than HashWithIndifferentAccess, meaning that string and symbol keys will not match
```


## Process
---
```ruby
# Detach the process from controlling terminal and run in the background as system daemon
Process.daemon
loop do
  puts "Hello ruby"
end

# kill process
pid = Process.pid
pid = $$
Process.kill(signal, pid)
```


## Ruby Tap
---
tap 是 Object 的 instance_method，传递 self 给一个 block，最后返回 self
```ruby
Person = Struct.new(:name, :age, :address)
person = Person.new
person.name = 'hello'
person.age = 'ruby'
person.address = 'world'

person.tap do |p|
  p.name = 'hello'
  p.age = 'ruby'
  p.address = 'world'
end
```

## Numeric step
---
```ruby
0.step(to: 10, by: 2) {|index| puts index}
```

## ruby保留两位小数 与 浮点数计算 Rational
---
>我们在计算钱的时候，由于直接使用float类型进行计算，是不准确的，比如 **0.1+0.2**, 为了是计算的准确，可以使用 BigDecimal库来计算。
也可以使用Rational有理数来表示，之后讲计算的结果转换为浮点数。
```ruby
i = 1.quo(3).to_f
'%.2f' % i
```

## Integer#digits
---
```ruby
567321.digits # =>[1, 2, 3, 7, 6, 5] 

```

## 进制转换
---
```ruby
# 输出16进制的对应字符串 =>“516"
puts 1000.to_s(16)
# 输出3进制的对应字符串 =>“200112"
puts 500.to_s(3) 

# 输出2进制的对应数字 =>255
puts Integer("0b"+"11111111")
puts "11111111".to_i(2) 

# 输出16进制的对应数字 =>255
puts Integer("0x"+"ff")
puts "ff".to_i(16)

# 输出10进制的对应数字 =>255
puts Integer("0d"+"255")
puts "255".to_i(10)

# 输出8进制的对应数字 =>255
puts Integer("0o"+"377")
puts “377”.to_i(8) 
```

## range
---
```ruby
# ...
1..10  # range 1-10
1...10 # range 1-9
```

## String Encoding ASCII & Unicode
---
```ruby
"a".ord       # 97
"abc".bytes   # [97, 98, 99]
97.chr        # "a"
("a".ord - 32).chr  # "A"
("A".ord + 32).chr  # "a"
"abc".encoding.name # "UTF-8"
"abcΣΣΣ".encode("ASCII", "UTF-8", undef: :replace)  # "abc???"
"abcΣΣΣ".encode("ASCII", "UTF-8", invalid: :replace, undef: :replace, replace: "")  # "abc"
"abcΣΣΣ".encode("ASCII", "UTF-8", fallback: {"Σ" => "E"})  # "abcEEE"
```