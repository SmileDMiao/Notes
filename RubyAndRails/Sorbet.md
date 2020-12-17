## Start
```ruby
# Gemfile
gem 'sorbet', :group => :development
gem 'sorbet-runtime'
```

```shell
srb init

srb tc
```

## Type System
```ruby
# Normal
sig { params(x: SomeType, y: SomeOtherType).returns(MethodReturnType) }
def foo(x, y)
end

# Optional param: y
sig { params(x: String, y: T.nilable(String)).void }
def foo(x:, y: nil)
end

# Rest Params
sig do
  params(
    # Integer describes a single element of args
    args: Integer, # rest positional params
    # Float describes a single value of args
    kwargs: Float  # rest keyword params
  )
  .void
end
def self.main(*args, **kwargs)
  # Positional rest args become an Array in the method body:
  T.reveal_type(args) # => Revealed type: `T::Array[Integer]`

  # Keyword rest args become a Hash in the method body:
  T.reveal_type(kwargs) # => Revealed: type `T::Hash[Symbol, Float]`
end

# Returns & void: Annotating return types

```


## File-Level: Strictness Levelss
```ruby
typed: ignore (Sorbet not read this file)
typed: flase (报告关于语法 常量解析 sig检查到相关错误)
typed: true (常规大类型错误 包括不存在的方法调用)
typed: strict (Sorbet不再隐式地将事物标记为动态类型 在这个级别上 所有方法都必须有sig 所有常量和实例变量都必须有显式注释的类型)
typed: strong (Sorbet不再允许T.untyped作为任何方法调用的中间结果。这实际上意味着Sorbet静态地知道文件中所有调用的类型)
```