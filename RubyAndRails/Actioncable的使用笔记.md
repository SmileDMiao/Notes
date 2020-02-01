## Actioncable的一些概念理解
+ channel.rb: 用于封装共享逻辑
+ connection.rb: 负责连接的类，完成对连接的授权
+ ../chat_channel.rb: 自己创建的频道类。

创建自己的频道类：
+ class: ChatChannel method: speak
+ rails generate channel chat speak

默认生成的js文件
```js
# 这里的ChatChannel可以理解为这里的websocket使用服务端的ChatChannel这个类来处理
App.chat = App.cable.subscriptions.create "ChatChannel",
  connected: ->
# Called when the subscription is ready for use on the server

  disconnected: ->
# Called when the subscription has been terminated by the server

  received: (data) ->
# 收到服务端消息时的处理
```
默认的自定义频道类
```ruby
class ChatChannel < ApplicationCable::Channel

  def subscribed
    # 用户订阅时候的处理
    # 这里相当于生成一个通道，服务端和客户端的通信基于这个通道，服务端不需要知道这个客户端，服务端向这个通道广播消息，客户端监听这个通道，收到消息做出处理
    stream_from "some_channel"
  end

  def unsubscribed
    # 退订时候的处理
    # Any cleanup needed when channel is unsubscribed
  end

end
```

## 利用Actioncable完成一个聊天的功能
```js
received: (data) ->
  # 收到消息，现实聊天内容
  $('#chats').append data['message']
speak: (message) ->
  # 调用ChatChannel的speak方法，message为内容
  @perform 'speak', message: message

$(document).on 'keypress', '[data-behavior~=chat_speaker]', (event) ->
if event.keyCode is 13 # return = send
  # 调用speak方法，拼装发送的参数内容
  App.chat.speak {'message':event.target.value, 'send_user_id': App.current_user_id, 'receive_user_id': App.receive_user_id}
  event.target.value = ""
  event.preventDefault()
```

```ruby
# ChatChannel中的speak方法
def speak(data)
  # 保存信息
  message = Message.create(data['message'])
  # 由于用户订阅了之前的通道，现在服务端向这个通道发送消息
  ActionCable.server.broadcast "chat_channel#{self.current_user.id}", message: render_message(message)
  ActionCable.server.broadcast "chat_channel#{message.receive_user_id}", message: render_message(message)
end

private
# 这里调用render方法，利用partial,最后生成的是一段html字符串，方便浏览器插入dom
# 由于在聊天内容现实页面也是用类似方式来现实聊天内容的，但是在页面中可以使用current_user之类的方法，但是这里的partial是无法使用该方法的，所以我这里相当于写了两个相同的方法
# 不过第二个方法我是手动传入了current_user这个对象
def render_message(message)
  ApplicationController.renderer.render(partial: 'chats/chat', locals: { message: message, current_user: self.current_user })
end
```

```ruby
# connection中方法，通过cookie来判断设置当前用户
identified_by :current_user
def connect
  self.current_user = find_verified_user
end
protected
def find_verified_user
  User.find_by_auth_token(cookies[:auth_token]) || reject_unauthorized_connection
end
```

## 关于websocket
WebSocket协议是基于TCP的一种新的网络协议。它实现了浏览器与服务器全双工(full-duplex)通信——允许服务器主动发送信息给客户端。
它的原理是这样的，由于它是一个协议，它不用发送跟http同样多的头信息，它比较轻量，速度快。为了建立一个 WebSocket 连接，客户端浏览器首先要向服务器发起一个 HTTP 请求，这个请求和通常的 HTTP 请求不同，包含了一些附加头信息，其中附加头信息”Upgrade: WebSocket”表明这是一个申请协议升级的 HTTP 请求，服务器端解析这些附加的头信息然后产生应答信息返回给客户端，客户端和服务器端的 WebSocket 连接就建立起来了，双方就可以通过这个连接通道自由的传递信息，并且这个连接会持续存在直到客户端或者服务器端的某一方主动的关闭连接。

## 关于adapter
默认rails的actioncable使用redis作为adapter
```ruby
# gem Tubesock:对rack hijack的封装,关于hijack下面说
# 该gem利用redis实现了pub/sub
def chat
  hijack do |tubesock|
    redis_thread = Thread.new do
      Redis.new.subscribe "chat" do |on|
        on.message do |channel, message|
          tubesock.send_data message
        end
      end
    end

    tubesock.onmessage do |m|
      Redis.new.publish "chat", m
    end

    tubesock.onclose do
      redis_thread.kill
    end
  end
end
```

所有的客户端连接的websocket地址，出发chat方法，Redis subscribed的作用就是订阅频道，然而所有的客户端都会连接websocket地址，所以相当于每一个客户端都订阅了这个频道。
Redis publish的作用是向这个频道里广播消息，由于redis实现了pub/sub所以所有的客户端都可以收这个消息。
浏览器A -> server -> redis -> server -> 浏览器ABC

## Actioncable设置
访问设置：
Actioncable默认只使用3000端口，如果想要使用其他端口:
```ruby
config.action_cable.allowed_request_origins = ['http://rubyonrails.com', %r{http://ruby.*}]
config.action_cable.disable_request_forgery_protection = true
```
服务运行：
以rack方式挂载到web程序中：
```ruby
Rails.application.routes.draw do
  mount ActionCable.server => '/cable'
end
```
独立运行：
```ruby
# cable/config.ru
require ::File.expand_path('../../config/environment', __FILE__)
Rails.application.eager_load!
run ActionCable.server
# puma -p 28080 cable/config.ru
```

## 参考
https://ruby-china.org/topics/29927
https://ruby-china.org/topics/30494
http://guides.ruby-china.org/action_cable_overview.html