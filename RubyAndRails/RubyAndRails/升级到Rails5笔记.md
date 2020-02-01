## 更新Gem
其实[guides](http://guides.ruby-china.org/upgrading_ruby_on_rails.html)已经很详细了，但是我在升级过程中还是遇到了好多的问题，这里记录我的问题

- Gemfile更改新版本
- bundle update
- rails app:update
- 根据新版本的变化更改一些文件，可以参考框架的更新日志

## Model 和 Job 的类多了一层继承
```ruby
ApplicationRecord < ActiveRecord::Base
end

ApplicationJob < ActiveJob::Base
end
```

## rack-contrib
之前由于需要安装了这个gem:rack-contrib，但是我rails s死活不成功， 由于rails5使用了rack2.0，该gem已经无法使用，相应的使用了该gem的路由配置也要删除，只好寻找替代品。

## config.active_record.raise_in_transactional_callbacks
**该配置已经弃用，这东西也导致我rails s失败**

## 我创建一个用户的时候，执行@user.save的时候 提示undefined method split for nil
```ruby
validates :full_name, uniqueness: { case_sensitive: false }
```
通过排除大法发现是这一段验证的问题，之前rails4这么写的确没有问题，但是升级后的rails5就有问题，guides文档这么写的，看到ruby-china中这么写好像也没啥问题， 反正我这里这么写报错了，我只有暂且删除这种写法了。

## 当我更新到rails5的时候，发现migration的写法有点变化
```ruby
# 这里的写法就有点奇怪了，你继承就继承，后面的[5.1]是什么意思？
class CreateUsers < ActiveRecord::Migration[5.1]
end
# 后来在源码里看到这么一段
def self.[](version)
   Compatibility.find(version)
end

# 其实就是根据传入的version返回对应的constant
Module.const_get(name)
```

## 自定义UUID遇到的问题
之前在公司的一个项目上使用过uuid作为主键，后来在一个自己写的玩具上也试了一下，在rails4版本的时候使用的做法是打开ActiveRecord:Base,添加一个全局的回调
```ruby
  before_create :set_uuid
  def set_uuid
    self.id = UUID.generate(:compact)
  end
```
但是在rails5的版本时，schema_migrations这张表生成的记录也变成了uuid，导致migration运行验证过不去，schema表在rails4时不受ActiveRecord回调影响，rails5里有了影响。
跳过ActiveRecord::SchemaMigration
```ruby
def set_uuid
  self.id = UUID.generate(:compact) unless self.class.name == 'ActiveRecord::SchemaMigration'
end
```

## redirect_to :back 弃用
```ruby
redirect_back(fallback_location:  request.headers['HTTP_REFERER'])
```

## manifest.js
```ruby
# app/assets/config/manifest.js
//= link application.css
//= link application.js
```

## Rails5.1 and 5.2 cache_store
```ruby
# 5.1
# Gemfile
gem 'redis-rails'
config.cache_store = :redis_store, {url: 'redis://localhost:6379/1', expires_in: 120.minutes, dirver: :hiredis}
# 5.2
config.cache_store = :redis_cache_store, {url: 'redis://localhost:6379/1', expires_in: 120.minutes, dirver: :hiredis}
```