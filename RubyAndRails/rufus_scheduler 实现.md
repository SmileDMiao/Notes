## Refus_scheduler基本介绍
**refus_scheduler是一个用于定时任务的gem，和whenever不同的是，这个gem是一个纯ruby实现，没有依赖与操作系统的crontab定时任务。该gem也没有依赖比如redis之类的存储，所以这个工具是没有将任务保存起来的，这样看起来刚好简单一些。之前一直好奇用纯ruby的方式来实现定时任务该怎么做，这里就先看一下这个gem是怎么实现的**


## 基本用法
[refus_scheduler](https://github.com/jmettraux/rufus-scheduler)

Rufus-scheduler（开箱即用）是一个进程内的内存调度程序。 它使用线程。

```ruby
require 'rufus-scheduler'

scheduler = Rufus::Scheduler.new

scheduler.in '3s' do
  puts 'Hello... Rufus'
end

scheduler.join
```


## 处理过程
我是根据gem的用法来一步步看源码的，那么就从第一行代码一步步看下去，看每一行代码到底做了些什么事情。
**初始化的时候，定义了初始化对象的好多属性，比如任务队列，开始时间，最大工作线程数，记录object_id，创建新的线程循环执行任务。并对任务进行了分类，只运行一次的运行之后删除，执行那些需要循环执行的，执行那些超时的任务。**
```ruby
# 初始化对象
scheduler = Rufus::Scheduler.new
# 初始化过程
def initialize(opts={})
      @opts = opts
      # 定时任务开始时间
      @started_at = nil
      # 是否暂停
      @paused = false
      # 初始化任务队列数组
      @jobs = JobArray.new
      # 默认频率
      @frequency = Rufus::Scheduler.parse(opts[:frequency] || 0.300)
      @mutexes = {}
      @work_queue = Queue.new
      # 最大工作线程
      @max_work_threads = opts[:max_work_threads] || MAX_WORK_THREADS
      @stderr = $stderr
      # 记录object_id
      @thread_key = "rufus_scheduler_#{self.object_id}"
      @scheduler_lock =
        if lockfile = opts[:lockfile]
          Rufus::Scheduler::FileLock.new(lockfile)
        else
          opts[:scheduler_lock] || Rufus::Scheduler::NullLock.new
        end
      @trigger_lock = opts[:trigger_lock] || Rufus::Scheduler::NullLock.new
      # If we can't grab the @scheduler_lock, don't run.
      lock || return
      # 开始运行
      start
end
# start 方法
def start
  # 初始化开始时间
  @started_at = EoTime.now
  # 创建一个新的线程来处理任务队列
  @thread =
    Thread.new do
      # 开始死循环任务队列
      while @started_at do
        # 删除那些只运行一次的任务
        unschedule_jobs
        # 执行任务,除非设置了暂停参数
        trigger_jobs unless @paused
        # 处理那些超时的任务
        timeout_jobs
        # 根据初始化的频率，sleep相应的时间
        sleep(@frequency)
      end
    end
  # 设置线程的相关属性，便于清除任务，杀死任务，关掉任务
  @thread[@thread_key] = true
  @thread[:rufus_scheduler] = self
  @thread[:name] = @opts[:thread_name] || "#{@thread_key}_scheduler"
end
```


## 定义任务
```ruby
# 定义不同类型的任务用的都是一个方法，区别在与参数的不同，这里不同的任务类型也会对应到下面提到的不同的类，每一个任务类型有一个对应的类
# 定义只执行一次的任务
def at(time, callable=nil, opts={}, &block)
  do_schedule(:once, time, callable, opts, opts[:job], block)
end
# 定义循环执行的任务
def every(duration, callable=nil, opts={}, &block)
  do_schedule(:every, duration, callable, opts, opts[:job], block)
end

# do_schedule 方法
# 这里先忽略掉一些异常掉判断，和一些默认空掉参数
def do_schedule(job_type, t, callable, opts, return_job_instance, block)
  # 判断任务类型，返回相应的任务类
  job_class =
    case job_type
      when :once
        opts[:_t] ||= Rufus::Scheduler.parse(t, opts)
        opts[:_t].is_a?(Numeric) ? InJob : AtJob
      when :every
        EveryJob
      when :interval
        IntervalJob
      when :cron
        CronJob
    end
  # 初始化一个相应掉任务类对象
  job = job_class.new(self, t, opts, block || callable)
  # 将这个任务对象放入任务队列中
  @jobs.push(job)
end
```


## 任务初始化
所有类型的任务都继承自Job类，在不同类型任务初始化的时候，都在initialize方法中使用了super，那么来看下Job类中任务初始化的时候做了什么事：
```ruby
# scheduler对象 任务执行时间参数 opts选项默认为{} block执行任务代码块 callback默认nil
job = job_class.new(self, t, opts, block || callable)
# class job initialize方法中参数
# 任务处理频率
@original = original
# 任务选项参数
@opts = opts
# 处理代码快
@handler = block
# 真正执行
@callable
# class every job initialize
# 设置任务的频率
@frequency = Rufus::Scheduler.parse_in(@original)
# 设置任务的下次执行时间
set_next_time(nil)
# 下次执行时间的判断， 开始时间大于当前时间则为开始时间，否则为下次执行时间||当前时间(第一次next time为空) + 一个执行频率
n = EoTime.now

@next_time =
if @first_at && (trigger_time == nil || @first_at > n)
  @first_at
else
  (@next_time || n) + @frequency
end
```


## 任务执行
在scheduler初始化的时候会创建一个死循环，那么分别来看关于删除任务，执行任务的代码
```ruby
# 删除任务
@jobs.delete_unscheduled
def delete_unscheduled
@mutex.synchronize {
  # 如果下次执行时间为空，或者设置了任务的截止时间那么将该任务从任务队列中删除
  @array.delete_if { |j| j.next_time.nil? || j.unscheduled_at }
}
end

# 正常执行任务
trigger_jobs
def trigger_jobs
  now = EoTime.now
  @jobs.each(now) do |job|
    job.trigger(now)
  end
end
# 这里job array 类本身写了个each方法覆盖了默认的数组迭代器，可以让其传入一个时间参数
def each(now, &block)
  # 先将所有任务按照下次执行时间排序，先执行的在前面，后执行的在后面
  to_a.sort_by do |job|
    job.next_time || (now + 1)
  end.each do |job|
    # 遍历job， 如果没有下次执行时间或下次执行时间大于当前时间则跳出循环，先排序后遍历
    nt = job.next_time
    break if ( ! nt) || (nt > now)
    # 到这里说明任务是到时候了需要执行
    # block传入job执行任务
    block.call(job)
  end
end
# 执行任务
def trigger(time)
  # 下次执行时间变上次执行时间
  @previous_time = @next_time
  # 设置下次执行时间
  set_next_time(time)
  # 真正执行任务的地方
  do_trigger(time)
end
```


## 总结
这个gem使用存ruby实现，基本思路就是开一个死循环，遍历任务，时间符合就执行，通过对scheduler初始化，设置若干参数来控制scheduler，对任务区分类型，初始化设置执行频率下次执行时间各种参数来方便任务对管理与执行。