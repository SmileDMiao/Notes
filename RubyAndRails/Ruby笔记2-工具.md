## Gem本地服务器
---
#### 内置的gem server
```ruby
ruby -v
rvm gemset list
gem list
gem server
=>Server started at http://0.0.0.0:8808
```

#### geminabox
```ruby
gem install geminabox
mkdir data
vi config.ru
  require "rubygems"
  require "geminabox"

  Geminabox.data = "./data"
  run Geminabox::Server
rackup
=>Host:  http://localhost:9292
# push gem to the server
gem inabox secretgem-0.0.1.gem
```
如果不是共享的服务，放在本地：
```ruby
source "file://gem_sources_folder" do
  gem 'secretgem'
end
```