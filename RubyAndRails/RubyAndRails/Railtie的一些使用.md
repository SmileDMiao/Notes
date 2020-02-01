## Railtie的一些使用
> [Railtie](https://api.rubyonrails.org/classes/Rails/Railtie.html) is the core of the Rails Framework and provides several hooks to extend Rails and/or modify the initialization process.


### 添加配置属性
在config的配置文件中，可以随意添加参数
```ruby
# config/application.rb
# 这里就是新定义一个require_login 的属性
config.require_login = true
```

### Railtie
利用railtie对config进行扩展
写在 initializers中
```ruby
# config/initializers/
class MyRailtie < Rails::Railtie
  initializer "my_railtie.configure_rails_initialization" do
    # some initialization behavior
  end
end
```

写在 lib 文件夹中
```ruby
# lib/fwk/railtie.rb
# 之后可以有两种方式加载；
# 1是在application.rb文件中 手动 require
# 2是可以在 gemfile 中 手动 require
module fwk
  class Railtie < Rails::Railtie
    config.before_configuration do
      # 初始化自定义配置
    config.fwk = Fwk::Config.instance
  end
end
```

**description**:
上面Fwk::Config受一个单例类，可以在initializer方法中为添加的属性赋值，
之后就可以通过 *Rails.application.config.fwk.xxx* 来获取自定义的相关配置

### RakeTask And Generator
```ruby
# lib/fkw/railtie.rb
# 在lib目录下按照标准写会替换rails的原生generator,这里可以通过 *send(:include)之类的方式打开类来扩展rails原生的generator
generator do
  # 可以扩展generator
end

rake_tasks do
    load 'path/to/my_railtie.tasks'
end
```