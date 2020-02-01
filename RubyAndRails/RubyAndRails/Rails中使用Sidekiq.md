##  关于Sidekiq
---
sidekiq在GitHub上的wiki内容真的很全面，概括了方方面面，我这里使用的不是Pro或企业版的Sidkeiq，这里使用了一些第三方Sidekiq插件，主要内容还是Github上文档上写的，这里我只是自己实践了一下记录了个人认为比较重要的东西。

## Sidekiq配置
---
安装好sidekiq之后，直接执行sidekiq就可以运行了，但是为了更好的运行，需要
```ruby
:concurrency: 5
:pidfile: tmp/pids/sidekiq.pid
:logfile: log/sidekiq.log
:queues:
  - [notifications, 100]
  - [default, 3]
```

1. concurrency:并发，开启5个线程
2. pidfile/logfile进程文件和日志文件
3. queues:队列
4. notifications, 100:队列名称和权重
5. 没有权重则按照写的顺序执行,权重一样则随机执行
6. verbose: Print more verbose output
7. timeout: shutdown timeout
8. max_retries: 1 最大重试次数

```ruby
# config/initializers/sideiq.rb
Sidekiq.configure_server do |config|
  config.redis = { :namespace => 'personal_practice', url: 'redis://localhost:6379/6' }

  config.on(:startup) do
    Sidekiq.schedule = YAML.load_file(File.expand_path('../../sidekiq-scheduler.yml', __FILE__))
    SidekiqScheduler::Scheduler.instance.reload_schedule!
  end

  Sidekiq::Status.configure_server_middleware config, expiration: 30.minutes
  Sidekiq::Status.configure_client_middleware config, expiration: 30.minutes
end

Sidekiq.configure_client do |config|
  config.redis = { :namespace => 'personal_practice', url: 'redis://localhost:6379/6' }

  Sidekiq::Status.configure_client_middleware config, expiration: 30.minutes
end
```

## 基本使用
---
```ruby
include Sidekiq::Worker
# processed asynchronously
HardWorker.perform_async('bob', 5)

# processed in the future
HardWorker.perform_in(5.minutes, 'bob', 5)
MyWorker.perform_at(3.hours.from_now, 'mike', 1)
```

## Best Practice
---
1. 参数小而简洁
2. Job业务要保持幂等，事务。
3. 拥抱并发

## 处理Retry之后依然失败的Job
---
```ruby
Sidekiq.configure_server do |config|
  config.death_handlers << ->(job, ex) do
    puts "Uh oh, #{job['class']} #{job["jid"]} just died with error #{ex.message}."
  end
end
```

## Delayed-Extension
---
```ruby
# 这是一个配置上的开关，可以在发邮件或者ActiveRecord执行某些方法时，利用Sidekiq将这些操作变成异步（很像Delayed_job）
Sidekiq::Extensions.enable_delay!
UserMailer.delay.welcome_email(@user.id)
```


## Error Handle
---
这里可以看到具体的报错信息，但是只有直接的报错信息，没有显示异常栈，有记录详细的关于job的信息包括job_id,args,run_time,class_name之类的信息，对于troubleshooting还是很有帮助的。
```ruby
config.error_handlers << Proc.new do |ex,ctx_hash|
    puts ex, ctx_hash
end
ex: "undefined method `size' for nil:NilClass"
ctx_hash: 
{:context=>"Job raised exception", :job=>{"class"=>"SchedulerFirstWorker", "args"=>[], "retry"=>false, "queue"=>"scheduler", "backtrace"=>true, "jid"=>"f6bbb7c29edc4273095aeba0", "created_at"=>1533623376.256413, "enqueued_at"=>1533623376.2639098}, :jobstr=>"{\"class\":\"SchedulerFirstWorker\",\"args\":[],\"retry\":false,\"queue\":\"scheduler\",\"backtrace\":true,\"jid\":\"f6bbb7c29edc4273095aeba0\",\"created_at\":1533623376.256413,\"enqueued_at\":1533623376.2639098}"}
```

## Sidekiq Test
---
1. Sidekiq::Testing.fake!（不用redis）
2. Sidekiq::Testing.inline!（同步执行）
3. Sidekiq::Testing.disable!（禁用所有异步任务）


## Sidekiq Plugin
---
```ruby
# 相似任务打包处理
gem 'sidekiq-grouping'
# sidekiq 定时任务
gem 'sidekiq-scheduler'
# 记录Job的状态
gem 'sidekiq-status'
# Ensure uniqueness of your Sidekiq jobs
gem 'sidekiq-unique-jobs'
```

### Sidekiq Scheduler
+ Under normal conditions, cron and at jobs are pushed once regardless of the number of sidekiq-scheduler running instances, assumming that time deltas between hosts is less than 24 hours.
+ every, interval and in jobs will be pushed once per host.

### Sidekiq Grouping Retry Not Work
>Sidekiq group在服务端循环将redis里的数据写成sidekiq worker格式, worker_class传入string, sdiekiq将去获取部分默认的sidekiq options
```ruby
module Sidekiq
  module Grouping
    class Batch

      def flush
        chunk = pluck
        return unless chunk

        chunk.each_slice(chunk_size) do |subchunk|
          Sidekiq::Client.push(
            # change @worker_class from string to Class
            'class' => @worker_class.constantize,
            'queue' => @queue,
            'args' => [true, subchunk]
          )
        end
        set_current_time_as_last
      end
    end
  end
end
```