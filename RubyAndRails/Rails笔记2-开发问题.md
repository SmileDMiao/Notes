## Rails config_for YAML.load_file
---
Convenience for loading (config/foo.yml) for the current Rails env.
```ruby
# config/exception_notification.yml:
production:
  url: http://127.0.0.1:8080
  namespace: my_app_production
development:
  url: http://localhost:3001
  namespace: my_app_development
# config/production.rb
Rails.application.configure do
  config.middleware.use ExceptionNotifier, config_for(:exception_notification)
end
# YAML load file
YAML.load_file(Rails.root.join('config', 'cleanup.yml'))
```

## Rails change default host
---
```ruby
# bin/rails
if ARGV.first == 's' || ARGV.first == 'server'
  require 'rails/commands/server'
  module Rails
    class Server
      def default_options
        super.merge(Host:  '0.0.0.0', Port: 3000)
      end
    end
  end
end
```

## Rails Temporarily Change Locale
---
```ruby
class UserMailer < ActionMailer::Base
  default from: 'noreply@translation.io'

  def invitation(user)
    I18n.with_locale(user.locale) do
      mail({
        :subject => _("You have been invited"),
        :to      => user.email
      })
    end
  end
end
```


## Gem install 指定source
---
```ruby
source 'source-host' do
  gem 'gem-name'
end

gem 'gem-name' source 'source-host'

gem install gem-name --source source-host
```

## ActiveModel::Dirty track changes
---
http://api.rubyonrails.org/classes/ActiveModel/Dirty.html
```ruby
2.3.0 :001 > a =  Article.find("635462606f5701349e75308d99265ad4")
2.3.0 :003 > a.liked_user_ids_will_change!
=> ["a", "b"]
2.3.0 :005 >  a.liked_user_ids << 'finances'
=> ["a", "b", "finances"]
2.3.0 :009 > a.liked_user_ids_changed?
=> true
2.3.0 :010 > a.changes
=> {"liked_user_ids"=>[["a", "b"], ["a", "b", "finances"]]}
```


## 查看一个model的关联model
---
```ruby
# it's a arry
Model_name.reflect_on_all_associations
# several methods
Model_name.reflect_on_all_associations.first[:macro, :klass, :table_name]
```


## Rails assets precompile every js css
---
rails可以为单一的文件进行预编译，rails默认的做法是将所有assets打包，每个页面夹在一样的，但是这样存在一个问题，静态资源过多时，就明显有些浪费了。
毕竟不是所有的js,css每个页面都需要。rails手动添加编译方式：
```ruby
Rails.application.config.assets.precompile += %w( wechat.scss, wechat.js )
```
但是如果需要手动添加编译的很多这样就很麻烦了，今天在项目中看到一段代码：
```ruby
Rails.application.config.assets.precompile << Proc.new do |path|
  if path =~ /\.(css|js)\z/
    full_path = Rails.application.assets.resolve(path).to_s
    app_assets_path = Rails.root.join('app', 'assets').to_s
    if full_path.starts_with? app_assets_path
      puts "including asset: " + full_path
      true
    else
      puts "excluding asset: " + full_path
      false
    end
  else
    false
  end
end
```
一开始对这里的Proc有些不理解,发现在console中查看 **Rails.application.config.assets.precompile* 中默认会有一个proc，而这个proc就是rails默认编译的静态资源，其默认值为:
```ruby
[ Proc.new { |filename, path| path =~ /app\/assets/ && !%w(.js .css).include?(File.extname(filename)) },
/application.(css|js)$/ ]
```
rails本身是可以对application.js和application.css编译的，这里就是运用rails本身的方式对每一个静态资源进行编译。


## has many 模型关系 返回JSON
---
```ruby

# 这样返回的数据格式很完美，层级结构清晰，但是有个问题,实际上返回这些数据只需要查询两次sql,但是这里的方式确是有所少数据查询多少次
# 如果使用jbuilder之类的工具可以只查询两次,然后遍历，形成自己想要的数据结构

# 控制top level of json {user: {id: 1}} and {id: 1}
ActiveRecord::Base.include_root_in_json = false

orders = orders.where(organization_id: current_user.organization_id) if current_user.organization_id
orders = orders.as_json(:include => { :air_ticket_details => {:include => :air_order_insurances}}, :only => [:order_no, :detail])
```

## JSON转化为object
---
```ruby
user_json = User.take.to_json
user = JSON.parse(user_json, object_class: User)
```

## index_by
---
Convert an enumerable to a hash keying it by the block return value.
```ruby
people.index_by(&:login)
=> { "nextangle" => <Person ...>, "chade-" => <Person ...>, ...}
people.index_by { |person| "#{person.first_name} #{person.last_name}" }
=> { "Chade- Fowlersburg-e" => <Person ...>, "David Heinemeier Hansson" => <Person ...>, ...}
```


## 找到某个model字段对应的翻译
---
只要遵循rails风格的写法, 可以用下面的方法来获取某一字段的翻译
```ruby
Model.human_attribute_name(attr)
```

## Rails log tag: uuid
---
在rails的配置文件production.rb中配置
```ruby
# 那么每次的请求后端访问的记录都会由同一个id起头，这样就可以很容易在日志中找到同一个请求的所有日志了
config.log_tags = [:uuid]
config.log_tags = [:request_id]
```

## Rails ActiveRecord autosave association
---
遇到一个一对多模型关系同时保存的问题，发现子项更改后保存失败。后面才知道是autosave的作用。
我在update!方法里没找到头绪, 其实**auto_save: true**是为模型添加了回调方法实现association model 保存的
```ruby
def add_autosave_association_callbacks(reflection)
  save_method = :"autosave_associated_records_for_#{reflection.name}"

  if reflection.collection?
    before_save :before_save_collection_association
    after_save :after_save_collection_association

    define_non_cyclic_method(save_method) { save_collection_association(reflection) }
    # Doesn't use after_save as that would save associations added in after_create/after_update twice
    after_create save_method
    after_update save_method
  elsif reflection.has_one?
    define_method(save_method) { save_has_one_association(reflection) } unless method_defined?(save_method)
    # Configures two callbacks instead of a single after_save so that
    # the model may rely on their execution order relative to its
    # own callbacks.
    #
    # For example, given that after_creates run before after_saves, if
    # we configured instead an after_save there would be no way to fire
    # a custom after_create callback after the child association gets
    # created.
    after_create save_method
    after_update save_method
  else
    define_non_cyclic_method(save_method) { throw(:abort) if save_belongs_to_association(reflection) == false }
    before_save save_method
  end
  
  define_autosave_validation_callbacks(reflection)
end
```

## Rails5 ar_internal_metadata
---
> 阻止Rake指令意外清空线上数据库

简而言之就是第一次migration会把当前环境写入DB, 后面运行Rake Task的时候会检查环境变量, 如果包含production则会抛出异常