## ActiveSupport Concern
---
为什么我们需要？
类包含模块后可以获得实例方法和类方法。
**concern模块出现之前**
```ruby
# 包含并扩展
module A
  def self.included(base)
    base.extend ClassMethods
  end
  
  module ClassMethods
    def hello
      puts 'in class method'
    end
  end
end
```

**存在的问题**
链式包含的问题
当module B include A的时候base是module B, module A中的class_method_a则成为了Module B的类方法.所以C没有这个方法
```ruby
module A
  def self.included(base)
    base.extend ClassMethods
  end
  
  def instance_method_a
    'ok'
  end
  
  module ClassMethods
    def class_method_a; 'ok'; end
  end
end

module B
  include A
  def self.included(base)
    base.extend ClassMethods
  end
  
    def instance_method_b
    'ok'
  end
  
  module ClassMethods
    def class_method_b; 'ok'; end
  end
end

class C
  include B
end

C.class_method_a => 'undefined method'
C.class_method_b => ok
C.new.instance_method_a => ok
C.new.instance_method_b => ok
```

**concern模块**

```ruby
module A
  extend ActiveSupport::Concern
  
  def instance_method_a
    'ok'
  end
  
  module ClassMethods
    def class_method_a; 'ok'; end
  end
end

class B
  include A
end
```

当模块扩展concern的时候为扩展它的类定义一个类实例变量:@_dependencies,初始为空数组.
append_futures方法: 内核hook方法(include默认空的, append_futures:检查被包含的模块是否在祖先链上,如果不在则将该模块加入祖先链)
一个concern不会包含另外一个concern,如果一个concern试图包含另外一个concern,只是链接到一个依赖中.
首先检查concern是否已经在该包的祖先链中(这种情况在链式包含会发生[base < self]),如果没有出现则进入包含并扩展部分
```ruby
module Concern
  class MultipleIncludedBlocks < StandardError #:nodoc:
    def initialize
      super "Cannot define multiple 'included' blocks for a Concern"
    end
  end

  def self.extended(base) #:nodoc:
    base.instance_variable_set(:@_dependencies, [])
  end

  def append_features(base)
    if base.instance_variable_defined?(:@_dependencies)
      base.instance_variable_get(:@_dependencies) << self
      false
    else
      return false if base < self
      @_dependencies.each { |dep| base.include(dep) }
      super
      base.extend const_get(:ClassMethods) if const_defined?(:ClassMethods)
      base.class_eval(&@_included_block) if instance_variable_defined?(:@_included_block)
    end
  end

  def included(base = nil, &block)
    if base.nil?
      if instance_variable_defined?(:@_included_block)
        if @_included_block.source_location != block.source_location
          raise MultipleIncludedBlocks
        end
      else
        @_included_block = block
      end
    else
      super
    end
  end

  def class_methods(&class_methods_module_definition)
    mod = const_defined?(:ClassMethods, false) ?
      const_get(:ClassMethods) :
      const_set(:ClassMethods, Module.new)

    mod.module_eval(&class_methods_module_definition)
  end
end
```

