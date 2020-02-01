## Rails中的缓存
[参考文档](http://guides.ruby-china.org/caching_with_rails.html)

>可以用Rails.cache.class 查看使用了哪种缓存，如果没有设置的话，Rails.cache.class是ActiveSupport::Cache::FileStore，默认使用文件缓存，位置在/tmp/cache*

## 设置使用哪种缓存方式
```ruby
# 使用文件缓存,
config.cache_store = :file_store, Rails.root.join('tmp')
# 使用内存缓存
config.cache_store = ：memory_store
# 使用memcached服务器缓存数据
namespace:缓存命名空间，expires_in:换粗 过期时间,compress:缓存过大时是否压缩,pool_size:dalli connection pool
config.cache_store = [:mem_cache_store, '127.0.0.1', { :namespace => NAME_OF_RAILS_APP, :expires_in => 1.day, :compress => true }]
```

* 使用文件缓存：在tmp目录下生成缓存文件
```ruby
Rails.cache.class
ActiveSupport::Cache::FileStore
Rails.cache.write("miao", "malzahar")
Rails.cache.read("miao")；Rails.cache.fetch("miao")
$file_store = ActiveSupport::Cache::FileStore.new(Rails.root.join('tmp/cache'))
$file_store.write(user.github_repositories_cache_key, items, expires_in: 15.days)
cache_key = "github-repos:#{github}:1"
items = $file_store.read(cache_key)
```

* 使用内存缓存（默认有大小限制，时间限制）
```ruby
Rails.cache.write("miao","malzahar")
 => true
Rails.cache
 => <#ActiveSupport::Cache::MemoryStore entries=1, size=252, options={}>
```

* 使用memcached缓存服务器：gemfile中添加dalli;安装memcached
```ruby
gem 'dalli'
apt-get install memcached
Rails.cache.class
 => ActiveSupport::Cache::MemCacheStore
Rails.cache.write("miao","malzahar")
Dalli::Server#connect 127.0.0.1:11211
 => 72057594037927936
Rails.cache.read("miao")
 => "malzahar"
```

## 客户端的缓存
[文档](http://api.rubyonrails.org/classes/ActionController/ConditionalGet.html#method-i-stale-3F)

[Etag](https://ruby-china.org/topics/24996)

stale?和fresh_when
在http的request和response的header中会有一个etag或者last_modified之类的标签，后台在接收到请求之后会对比etag，如果一致那么浏览器则直接使用cache的内容渲染页面.
测试结果：
action中代码还是会全部执行，不过内容没有变化的时候，不会重新渲染页面，由客户端从缓存中加载。
若你有特定的响应处理，请使用stale?方法；若你没有特定的响应处理，例如你不需要使用respond_to或调用render方法，请使用fresh_when。

## 片段缓存(套娃缓存)
这里两篇文章写的很好
[说说 Rails 的套娃缓存机制](https://ruby-china.org/topics/21488)
[Redis 实现 Cache 系统实践](https://ruby-china.org/topics/27939)

```ruby
# actionview/lib/action_view/helpers/cache_helper
def cache(name = {}, options = {}, &block)
  # 这里先判断有没有打开换粗的设置
  if controller.respond_to?(:perform_caching) && controller.perform_caching
    # 拿到支持的参数
    # skip_digest的作用就是用于跳过下面的那个缓存key计算方式，如果你想在cache的block之外的改动不会刷新缓存的话。
    name_options = options.slice(:skip_digest, :virtual_path)
    safe_concat(fragment_for(cache_fragment_name(name, name_options), options, &block))
  else
    # 没有打开缓存直接渲染
    yield
  end

  nil
end

  # 这个方法的作用在于计算缓存key
def fragment_name_with_digest(name, virtual_path)
  virtual_path ||= @virtual_path
  if virtual_path
    name = controller.url_for(name).split("://").last if name.is_a?(Hash)
    digest = Digestor.digest name: virtual_path, finder: lookup_context, dependencies: view_cache_dependencies
    [ name, digest ]
  else
    name
  end
end

# 判断缓存是否命中
def fragment_for(name = {}, options = nil, &block)
  if content = read_fragment_for(name, options)
    @cache_hit = true
    content
  else
    @cache_hit = false
    write_fragment_for(name, options, &block)
  end
end

# 读取缓存
def read_fragment_for(name, options)
  controller.read_fragment(name, options)
end

# 写入缓存
def write_fragment_for(name, options)
  # pos = 0
  pos = output_buffer.length
  yield
  output_safe = output_buffer.html_safe?
  # fragment就是需要缓存的内容
  fragment = output_buffer.slice!(pos..-1)
  if output_safe
    # output_buffer变为''
    self.output_buffer = output_buffer.class.new(output_buffer)
  end
  controller.write_fragment(name, fragment, options)
end
```

```ruby
# 读取缓存部分代码
def read_fragment(key, options = nil)
    # 先看是否配置了开启缓存
    return unless cache_configured?
    # 计算换粗key
    key = fragment_cache_key(key)
    instrument_fragment_cache :read_fragment, key do
      # 读取缓存
      result = cache_store.read(key, options)
      # 安全输出html
      result.respond_to?(:html_safe) ? result.html_safe : result
  end
end
  # 写入缓存
def write_fragment(key, content, options = nil)
  return content unless cache_configured?

  key = fragment_cache_key(key)
  instrument_fragment_cache :write_fragment, key do
    content = content.to_str
    cache_store.write(key, content, options)
  end
  content
end

  # 真正执行缓存读取写入的地方在block中
  # 这里使用ActiveSupport::Notifications发布缓存读取写入的事件，可能是日志订阅了这个事件
def instrument_fragment_cache(name, key) # :nodoc:
  payload = instrument_payload(key)
  ActiveSupport::Notifications.instrument("#{name}.#{instrument_name}", payload) { yield }
end
```

## 注意点：
1. 避免直接对nil做缓存，添加标示
2. 在使用套娃缓存的时候注意touch问题

## 总结
在使用片段缓存的时候，基本就是计算缓存key，然后在缓存中读取key值，如果读取到，那么就返回读取到内容，如果没有读取到，那么就缓存block中的缓存内容。