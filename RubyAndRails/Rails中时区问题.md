## 问题描述
>在开发一个接口的过程中遇到一个问题，用grape提供一个接口，接口基本没有啥逻辑，只是将一个表的数据展示出来，只要将时间format一下就好了，心里想着懒得写entity了，就打算在数据库层做这个时间转换的事。（使用postgresql数据库）
当时代码是这么写的
```ruby
orders = current_user.intl_flights.select(:departure_time, "to_char(departure_time, 'YYYY-MM-DD HH24:MI:SS') AS departure_time_text")
```
但是最后展示的时间却与实际的不相符。发现拿出的时间比期望值少8小时。那估计就是时区的问题的，北京时间刚好是标准时间+8小时

## 可能引发的原因
rails中的配置
```ruby
config.time_zone = 'Beijing'
config.active_record.default_timezone = :local
```
数据库时区:PRC
操作系统时区:2017年12月20日 星期三 20时28分35秒 CST

## 测试
rails不做任何设置
```ruby
Rails.application.config.time_zone
UTC
Rails.application.config.active_record.default_timezone
nil
created_at = Food.last.created_at
DB时间                         => 2017-12-20 12:07:48.782666
created_at                    => Wed, 20 Dec 2017 12:07:48 UTC +00:00
created_at.localtime          => 2017-12-20 20:07:48 +0800
created_at.utc                => 2017-12-20 12:07:48 UTC
```
可以看出数据库时间是UTC时间，created_at也是UTC时间，created_at.localtime是北京时间

只设置time_zone
```ruby
DB时间                         => 2017-12-20 12:12:10.988018
created_at                    => Wed, 20 Dec 2017 20:12:10 CST +08:00
created_at.localtime          => 2017-12-20 20:12:10 +0800
created_at.utc                => 2017-12-20 12:12:10 UTC
可见config.time_zone = 'Beijing'配置的作用，是在ActiveRecord中取时间的时候，将UTC时间转换成Local时间，也就是通过created_at等方法获取到的将直接是Local时间。而存储在DB中的时间仍然是UTC时间。
```

都设置
```ruby
DB时间                        2017-12-20 20:16:44.927086
created_at                    => Wed, 20 Dec 2017 20:16:44 CST +08:00
created_at.localtime          => 2017-12-20 20:16:44 +0800
created_at.utc                => 2017-12-20 12:16:44 UTC
```
可见config.active_record.default_timezone = :local配置的作用，是在ActiveRecord中往数据库存放数据时，将按Local时间进行存储，通过添加这两项配置，就可以实现数据库存放时间以及通过created_at等方法取到的时间均为Local时间。

rails中默认创建的created_at的类型是timestamp，是不带时区的，如果我将其类型改为timestamptz，那么在创建的时候则会默认转化为数据库时区设置的时区的时间格式。这个时候在db中取出的时间和rails中取出的时间就是一样的当地时间。rails在获取时间的时候，在第一个测试的情况下，如果获取到的时间不带时区那么会按照rails设置的时区转化，如果获取的时间带时区那么如果时区和rails一样就不转化了。

## 在rails中创建带时区的时间：
```ruby
class CreateFoo < ActiveRecord::Migration
  def up
    create_table :foos do |t|
      t.column :created_at, "timestamp with time zone"
      t.column :updated_at, "timestamp with time zone"
    end
  end
end
```

## 数据库层面的转化：
需要指定对应的时区,不然则是默认转化为utc时间。
```sql
SELECT to_char(created_at at time zone 'utc+8', 'YYYY-MM-DD HH24:MI:SS')  FROM users WHERE id = 1
```

## rails多时区项目情况处理
rake time:zones:all 可以列出所有的时区
若要支持多时区，你需要保存每个人对应的时区，然后根据具体的时区做对应的时间转化。
```ruby
around_action :user_time_zone, if: :current_user

def user_time_zone(&block)
  Time.use_zone(current_user.time_zone, &block)
end
```