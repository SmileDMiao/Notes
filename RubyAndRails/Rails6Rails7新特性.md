# Rails7
## Active Record Encryption
---
```ruby
# generate random key setc
rails db:encryption:init

class Article < ApplicationRecord
  # 默认非确定性加密(对同一内容加密两次结果不一致)
  encrypts :title
end

class Article < ApplicationRecord
  # 确定性加密
  encrypts :title, deterministic: true
end
```

## Trace Query Origins With Marginalia-Style Tagging
---
https://github.com/basecamp/marginalia
https://docs.gitlab.com/ee/development/database_query_comments.html

## Asynchronous Query Loading
---
```ruby
@genres = Genre.all.load_async
@videos = Video.order(published_at: :desc).load_async
```

# Rails6
[Rails 6.1: Horizontal Sharding, Multi-DB Improvements, Strict Loading, Destroy Associations in Background, Error Objects, and more!](https://weblog.rubyonrails.org/2020/12/9/Rails-6-1-0-release/)
## 多数据库支持
---
```ruby
// yml config
production:
  primary:
    database: my_primary_database
    username: root
    password: <%= ENV['ROOT_PASSWORD'] %>
    adapter: mysql2
  primary_replica:
    database: my_primary_database
    username: root_readonly
    password: <%= ENV['ROOT_READONLY_PASSWORD'] %>
    adapter: mysql2
    replica: true
  animals:
    database: my_animals_database
    username: animals_root
    password: <%= ENV['ANIMALS_ROOT_PASSWORD'] %>
    adapter: mysql2
    migrations_paths: db/animals_migrate
  animals_replica:
    database: my_animals_database
    username: animals_readonly
    password: <%= ENV['ANIMALS_READONLY_PASSWORD'] %>
    adapter: mysql2
    // let Rails know which one is a replica and which one is the writer
    replica: true
    
// command
`rails g migration CreateCats name:string --database animals`
`rails db:migrate:status:animals`
`rails db:create:animals`
`rails db:migrate:primary`

// model
class AnimalsRecord < ApplicationRecord
  self.abstract_class = true
  // write db read db configuration
  connects_to database: { writing: :animals, reading: :animals_replica }
end

// configuration
config.active_record.writing_role = :default
config.active_record.reading_role = :readonly

// automatic connection switching
config.active_record.database_selector = { delay: 2.seconds }
config.active_record.database_resolver = ActiveRecord::Middleware::DatabaseSelector::Resolver
config.active_record.database_resolver_context = ActiveRecord::Middleware::DatabaseSelector::Resolver::Session
```

## 并行测试
---
```ruby
class ActiveSupport::TestCase
  parallelize_setup do |worker|
    # setup databases
  end
   parallelize_teardown do |worker|
    # cleanup database
  end

  # Run tests in parallel with specified workers
  parallelize(workers: :number_of_processors)

  # Setup all fixtures in test/fixtures/*.yml for all tests in alphabetical order.
  fixtures :all

  # Add more helper methods to be used by all tests here...
end

`PARALLEL_WORKERS=10 bin/rails test`
```

## Destory Association Background
---
```ruby
class Team < ApplicationRecord
  has_many :players, dependent: :destroy_async
end
```

## Disallowed Deprecation Support
---
```ruby
// 使用已经弃用的api时报错
ActiveSupport::Deprecation.disallowed_warnings = [
  "calling bad_method is deprecated",
  :worse_method,
  /(horrible|unsafe)_method/
]

ActiveSupport::Deprecation.disallowed_warnings = :all

ActiveSupport::Deprecation.disallowed_behavior = ->(message, callstack, deprecation_horizon, gem_name) {
  # custom deprecation handling ...
}
```