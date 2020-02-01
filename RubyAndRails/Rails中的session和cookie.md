## Rails中的cookie
>cookie本身是http协议的内容

cookie的使用
```ruby
# 加密
cookies.encrypted[:discount] = 45
# 永久的（20年）
cookies.permanent[:login] = "XJ-122"
# 可选参数
# value: 值
# expires: 有效期
# secure: 是否仅传输到https服务器
# httponly: 是否只通过脚本或http访问
# path: 指定cookie的路径，只有对应的路径可以访问
# domain: 用来分享cookie的
cookies[:name] = {
  value: 'a yummy cookie',
  expires: 1.year,
  domain: 'domain.com'
}
```

作用cookie的是ActionDispatch::Cookies,这是一个middleware，可以通过 `rails(rake) middleware` 来查看其在rails middleware中的位置。
其call方法 将cookie信息写到header中
```ruby
def call(env)
  request = ActionDispatch::Request.new env

  status, headers, body = @app.call(env)

  if request.have_cookie_jar?
    cookie_jar = request.cookie_jar
    unless cookie_jar.committed?
      cookie_jar.write(headers)
      if headers[HTTP_HEADER].respond_to?(:join)
        headers[HTTP_HEADER] = headers[HTTP_HEADER].join("\n")
      end
    end
  end

  [status, headers, body]
end
```
真正设置cookie则是由rake的utils.rb中的`add_cookie_to_header(header,key,value)`方法完成

## Rails中的session
默认情况下，在rails中session存储在cookie中，由配置
```ruby
Rails.application.config.session_store :cookie_store, key: '_Malzahar'
```
the server encrypet the session value and store in the cookies as a key _Malzahar_

## Rails CSRF(Cross-site request forgery)
the method `csrf_meta_tags` in layout will generate the below html
```html 
<meta name="csrf-param" content="authenticity_token">
<meta name="csrf-token" content="My+O6o2XcBKYxyH8KOVxsEMJDe3ogBILFshFejnFwFBVaea7uF8WiCINYvK1/y6w5MW3wgg/Re3wKXAFgQFCeg==">
```

### encrypet process
```ruby
 # Sets the token value for the current session.
  AUTHENTICITY_TOKEN_LENGTH = 32
  def form_authenticity_token
    masked_authenticity_token(session)
  end
   # like BREACH.
  def masked_authenticity_token(session)
    one_time_pad = SecureRandom.random_bytes(AUTHENTICITY_TOKEN_LENGTH)
    encrypted_csrf_token = xor_byte_strings(one_time_pad, real_csrf_token(session))
    masked_token = one_time_pad + encrypted_csrf_token
    Base64.strict_encode64(masked_token)
  end
  def real_csrf_token(session)
    session[:_csrf_token] ||= SecureRandom.base64(AUTHENTICITY_TOKEN_LENGTH)
    Base64.strict_decode64(session[:_csrf_token])
  end
  def xor_byte_strings(s1, s2)
    s1.bytes.zip(s2.bytes).map { |(c1,c2)| c1 ^ c2 }.pack('c*')
  end
```

in html store a encrypet token and store unencrypet token in session
form表单在提交时会自动带上加密后的token，用来验证。

### 解密过程
application_controller中的protect_from_forgery with: :exception
为action设置了回调.
主要就是解密加密的token，然后和session中的token做比较.
真正解密比较的代码在 `actionpack-4.2.6/lib/action_controller/metal/request_forgery_protection.rb #valid_authenticity_token?`

## 参考
https://ruby-china.org/topics/35199