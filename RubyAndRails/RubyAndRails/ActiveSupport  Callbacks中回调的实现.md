## AcrtiveSupoort::Callbacks
callback是运行在object的生命周期上，某个事件点上的hook代码。
ActiveSupport::Callbacks提供了callback相关的最基本的功能。
在rails中，activerecord提供的回调，controller中的before_action，activejob中可用的回调都是基于
ActiveSupport::callback实现的,所以我们在此之前先了解ActiveSupport::Callbacks

### AcrtiveSupoort::Callbacks的使用
提供的三个方法:
+ define_callbacks: 定义事件点
+ set_callback: 为事件点安装callback
+ run_callbacks: 运行某个事件点上安装的callback

看下文档中写的demo
```ruby
class Record
  include ActiveSupport::Callbacks
  # 定义事件 save，即运行save方法触发某个回调
  define_callbacks :save

  # save方法
  def save
    # 运行save方法上的安装的回调
    run_callbacks :save do
      puts "save..."
    end
  end
end

class PersonRecord < Record
  # 为事件安装hook,这里就是运行save之前运行saving_message
  set_callback :save, :before, :saving_message
  def saving_message
    puts "saving..."
  end

  # 为事件安装hook,这里就是运行save之后运行一断代码快
  set_callback :save, :after do |object|
    puts "saved..."
  end
end

person = PersonRecord.new
person.save
# 输出
saving...
save...
saved...
```

## define_callbacks
define_callbacks方法会定义一个类变量，然后把一个空的callback链赋值给这个类变量。这个变量的值可以被子类继承，一旦该类被继承，子类也会拥有这个callback链

同时在这个方法里会根据name来创建一个run方法以供调用
```ruby
def define_callbacks(*names)
  options = names.extract_options!

  names.each do |name|
    # 定义类变量
    class_attribute "_#{name}_callbacks", instance_writer: false
    # 赋值Callback链到定义到类变量
    set_callbacks name, CallbackChain.new(name, options)

    module_eval <<-RUBY, __FILE__, __LINE__ + 1
        # 定义方法以供调用
        def _run_#{name}_callbacks(&block)
          __run_callbacks__(_#{name}_callbacks, &block)
        end
    RUBY
  end
end
```

## set_callbacks
deffin_callbacks定义了一个类属性：_save_callbacks，set_callbacks则是将before, after之类的回调方法push到这个类的回调链中.
```ruby
def set_callback(name, *filter_list, &block)
  type, filters, options = normalize_callback_params(filter_list, block)
  # 拿到之前定义到class attribute
  self_chain = get_callbacks name
  mapped = filters.map do |filter|
    Callback.build(self_chain, filter, type, options)
  end

  # target: class, chain: class attribute
  # push callbacks to callbacks chain
  __update_callbacks(name) do |target, chain|
    options[:prepend] ? chain.prepend(*mapped) : chain.append(*mapped)
    target.set_callbacks name, chain
  end
end
```

## run_callbacks

```ruby
def run_callbacks(kind, &block)
  send "_run_#{kind}_callbacks", &block
end

def __run_callbacks__(callbacks, &block)
  if callbacks.empty?
    yield if block_given?
  else
    runner = callbacks.compile
    e = Filters::Environment.new(self, false, nil, block)
    runner.call(e).value
  end
end
    
def call(*args)
  @before.each { |b| b.call(*args) }
  value = @call.call(*args)
  @after.each { |a| a.call(*args) }
  value
end
```

## ActiveRecord中使用ActiveSupport
```ruby
# DefineCallbacks模块实现了define_callbacks:定义事件set_callbacks:安装事件的功能
include DefineCallbacks
# Callbacks模块实现了run_callbacks的功能
include Callbacks
```
define_callbacks.rb
```ruby
define_model_callbacks :initialize, :find, :touch, only: :after
define_model_callbacks :save, :create, :update, :destroy

def define_model_callbacks(*callbacks)
  callbacks.each do |callback|
    # 设置hook
    define_callbacks(callback, options)

    types.each do |type|
      # 安装hook
      send("_define_#{type}_model_callback", self, callback)
    end
  end
end

# 定义方法，设置callbacks，为事件安装回调
def _define_before_model_callback(klass, callback)
  klass.define_singleton_method("before_#{callback}") do |*args, &block|
    set_callback(:"#{callback}", :before, *args, &block)
  end
end
```