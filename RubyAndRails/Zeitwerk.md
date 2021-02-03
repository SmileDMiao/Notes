Zeitwerk是Rails6中使用的新代码加载引擎

## classic mode
依赖 `autoload_paths` `const_missing callback` 来加载文件

工作流程:
在代码执行过程中, 每次Ruby发现一个未知常量引用时, 就会触发一个const_ missing回调。Rails重写了Ruby的默认const_missing回调, 这通常只会引发NameError, 相反，Rails里它尝试加载与正在查找的常量相关联的文件。这时候自动加载路径发挥作用, Rails遍历autoloadpaths列表, 查找引用常量的 "snake-case" 版本, 如果存在, 则加载该文件。

```ruby
# app/models/flight_model.rb
class FlightModel
end

# app/models/bell_x1/flight_model.rb
module BellX1
  class FlightModel < FlightModel
  end
end
 
# app/models/bell_x1/aircraft.rb
module BellX1
  class Aircraft
    def initialize
      @flight_model = FlightModel.new
    end
  end
end
```

上面代码本是想要创建一个`BellX1:：FlightModel`, 嵌套有BellX1。但是, 如果加载了默认的`FlightModel`, 而`BellX1`没有加载, 那么解释器就能够解析顶级`FlightModel`，因此`BellX1::FlightModel`不会触发自动加载。

## zeitwerk
Zeitwerk依赖 `autoload_paths` `Module#autoload` 来加载文件

工作流程:
当项目启动时, Rails将调用 `Zeitwerk#setup`, 这个方法负责为所有已知的 `autoload_paths` 设置自动加载。

```ruby
autoload :Comment, Rails.root.join('app/models/comment.rb')
autoload :Post, Rails.root.join('app/models/post.rb')
```
以`Post`为例:
1. 在 `Symbol Table` 检查是否有 `:Post` 的引用
2. 如果没有找到, 那么会检查是否为 `:Post` 设置了自动加载
3. 如果没有自动加载那么触发 `const_missing`
4. 如果有设置自动加载, 那么加载对应的文件