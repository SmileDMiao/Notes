## Erb ruby_array js中使用问题
erb文件中：
```ruby
user_ids = current_user.follower_ids + current_user.following_ids
users = User.where(:id => user_ids).collect{|i|i.full_name}.to_s
```
js中使用：
```javascript
<script>
     data = <%= users %>;
</script>
```
但是实际页面加载出来是这样第：
```js
data = [&quot;胥建锋&quot;];
```
解决方法：

```ruby
user_ids = current_user.follower_ids + current_user.following_ids
users = User.where(:id => user_ids).collect{|i|i.full_name}.to_s.html_safe
```

---
## 发送邮件时指定queue
```ruby
Notifier.welcome(User.first.id).deliver_later(queue: "low")
```

---
## 如何在一个action中或者视图中知道当前页面所在的layout名称是什么？
```ruby
# 在controller中 
self.send(:_layout)
# 在视图中
controller.send(:_layout)
```

---
## Rails 4 paperclip 上传文件类型验证
问题：在rails4中，paperclip在上传文件的时候默认是必须要验证文件类型的，在上传图片的时候，验证的content_type很好写，
但是在上传docx，xls，pdf之类的文件的时候总是失败。
验证方法:
```ruby
validates_attachment_content_type :xxx, content_type: { content_type: ["image/jpeg", "image/gif", "image/png"] }
#paperclip 提供了一种跳过验证的方法:
do_not_validate_attachment_file_type :xxx
```

如果需要验证的话，可以直接写入文件类型的content_type，可以在controller中打个断点，在params方法中可以看到文件类型对应的content_type，
也可以先跳过验证，看保存到数据库中的content_type是什么，之后在写到类型验证方法中。

## Rails generate skip test file
```ruby
config.generator do |g|
  g.test_framework nil
end
```

---
## 静态资源无法加载
之后发现静态资源没有加载，预编译也没用，发现production.rb中
```ruby
config.serve_static_files = ENV['RAILS_SERVE_STATIC_FILES'].present?
```
由于没有环境变量，只要设置成true就好了

---
## Cancancan Strong Parameters
controller中的strong parameters参数方法的名字必须是model名 + 下划线 + params


---
## 序列化xml
nokogiri
```ruby
require 'nokogiri'
source = '<some><nested><xml>value</xml></nested></some>'
doc = Nokogiri::XML source
puts doc.to_xml

Hash.from_xml(source)
# or we can use activesupport supply method Hash.from_xml, but i meet a problem, a post request with a xml body, but i can not use Hash.from_xml to parse the xml, but the below way can parse it correctly.

# xml转hash
def xml_to_hash(xml)
  dom = Nokogiri::XML(xml)
  def xml_dom(dom)
    hash = dom.element_children.each_with_object(Hash.new) do |e, h|
      h[e.name.to_sym] = e.element_children.empty? ? e.content : xml_dom(e)
    end
    return hash
  end
  hash = xml_dom(dom)
  return hash
end
```

---
## Rails skip callback
```ruby
Model.skip_callback(:save, :before, :calculate_average)
Model.set_callback(:save, :before, :calculate_average) 
```

---
## Rails migration comment
model中对应字段的注释women可以使用数据库层的注释：
gem 'migration_comments'
rails5 框架自带
对应sql:
```sql
create table test( 
    id int not null default 0 comment '用户id' ) ;
alter table test 
change column id id int not null default 0 comment '测试表id';
show full columns from test;
create table test1 ( 
    field_name int comment '字段的注释' 
)comment='表的注释';
alter table test1 modify column field_name int comment '修改后的字段注释'; 
show  create  table  test1;
```

---
## 在Rake任务中调用其他的Rake任务
```ruby
# invoke方法，由rake task提供，可以阅读api看看其他方法
task :init_system => :environment do
  Rake::Task["db:seed:system_base"].invoke
  Rake::Task["db:seed:system_suport"].invoke
  Rake::Task["irm:initdata"].invoke
end
```


## Rake Task Arguments
```ruby
task :add, [:num1, :num] do |t, args|
  puts args[:num1].to_i + args[:num].to_i
end
rake add[1,2]

task :add do
  ARGV.each { |a| task a.to_sym do ; end }
  puts ARGV[1].to_i + ARGV[2].to_i
end
rake add 1 2

task :add do
  puts ENV['NUM1'].to_i + ENV['NUM2'].to_i
end
rake add NUM1=1 NUM2=2
```


## Make an interactive Rake task?
```ruby
task :action do
  STDOUT.puts "I'm acting!"
end

task :check do
  STDOUT.puts "Are you sure? (y/n)"
  input = STDIN.gets.strip
  if input == 'y'
    Rake::Task["action"].reenable
    Rake::Task["action"].invoke
  else
    STDOUT.puts "So sorry for the confusion"
  end
end
```


## 在Rake任务中执行数据库操作
```ruby
# 只要在mingration中使用的方法前面加上ActiveRecord::Base.connection.
# 疑问1:ActiveRecord::ConnectionAdapters::SchemaStatements
# 疑问2:ActiveRecord::Tasks
ActiveRecord::Base.connection.
create_table :look_user do |t|
  t.string    :name,      :limit => 20
  t.string    :category,  :limit => 20
  t.timestamps null: false
end
```

## Rake Task 方法重复的问题
项目中有两个task，task1定义了方法hello, task2中也定义了方法hello，我在task1中调用方法hello的时候，实际上调用的却是task2中的hello方法，后面我定义了uniq method name解决了这个问题。还有就是我在task里面调用find_in_batches这个方法的时候，却报错了，报错报到另一个task里面去了，原来find_in_batches里面有调用logger这个方法，但是另外的一个task里面已经定义了这个方法。
原因: 在运行Rake任务的时候，不管你运行哪一个任务Rake都会加载所有的task，这也是不区分命名空间的，那么先加载rails环境，再加载task代码，所以后面加载的会覆盖前面加载的，所以会出现上面遇到的这么个情况。