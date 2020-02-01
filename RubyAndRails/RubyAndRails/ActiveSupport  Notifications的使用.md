## 基本使用
Publishers 在对象状态改变且需要触发事件的时候发布事件。
Subscribers 仅接收它们能响应的事件，并且在每个事件中可以接收到被监控的对象。
ActiveSupport::Notifications 主要核心就是两个方法：instrument 和 subscribe。
你可以把 instrument 理解为发布事件。instrument 会在代码块执行完毕并返回结果之后，发布事件 my.custom.event，同时会自动把相关的一组参数：开始时间、结束时间、每个事件的唯一ID等，放入 payload 对象。

+ instrument 可以发布事件

```ruby
ActiveSupport::Notifications.instrument('render', extra: :information) do
  render plain: 'Foo'
end
```

+ Subscriber 可以订阅事件

```ruby
ActiveSupport::Notifications.subscribe('render') do |name, start, finish, id, payload|
  name    # => String, name of the event (such as 'render' from above)
  start   # => Time, when the instrumented block started execution
  finish  # => Time, when the instrumented block ended execution
  id      # => String, unique ID for this notification
  payload # => Hash, the payload
end
```

+ ActiveSupport::Notifications::Event

```ruby
events = []
ActiveSupport::Notifications.subscribe('render') do |*args|
  events << ActiveSupport::Notifications::Event.new(*args)
end

ActiveSupport::Notifications.instrument('render', extra: :information) do
  render plain: 'Foo'
end

event = events.first
event.name      # => "render"
event.duration  # => 10 (in milliseconds)
event.payload   # => { extra: :information }
```

---
## 管理rails中的callback
[使用 Subscriber 来管理 Model Callbacks 北京 Rubyists 活动分享](https://ruby-china.org/topics/32649)

来自ruby-china的精品贴，主要原理就是当有数据变化时，发布对应的事件，然后监听改事件，做对应的改动，实现和model Callbacks一样的功能，
不过这样可以集中管理callback， 在有类似需求的小型项目中可以采用。

---
## 其他gem
[wisper](https://github.com/krisleech/wisper)

也是实现发布订阅的模式，采用事件回调来解耦，使用更加方便些，功能上差不多。