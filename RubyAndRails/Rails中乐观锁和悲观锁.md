## 悲观锁
悲观锁在使用时采用比较保守的策略，当锁住的时候，其他人无法修改锁住的记录，在事务开始之前获取写的权限，事务结束之后释放，在资源竞争比较激烈的时候适用。
rails 有两种方式
```ruby
# 这两种方式在Mysql和Postgresql中都可以使用
account = Account.find(1)
Account.transaction do
    account.lock!
    account.balance -= 100
    account.save!
end
# 和下面是等价的
account.with_lock do
    account.balance -= 100
    account.save!
end
```
给数据加锁形成的sql就是类似这样的语句
```sql
select * from where xxx for update
```

在PG中也可以不手动锁，可以设置事物的隔离级别(:serializable)
对应的数据库操作就是在事务开始的时候设置事务的隔离级别
```ruby
Comment.transaction isolation: :serializable do  
  comment =Comment.find('8f5dea901ead013587ed4ccc6afe409e')
  comment.body = 'update in transaction'
  comment.save!
end
```

```sql
BEGIN / BEGIN TRANSACTION

COMMIT / END TRANSACTION

ROLLBACK;
```

## 乐观锁
在事务提交之前，大家都可以提交自己的数据，但是在提交的时候发现数据有改变，则拒绝提交
需要给表加上lock_version字段
由于Mysql和PG的MVCC实现的不同，Mysql的事务隔离级别并不能实现乐观锁的功能，所以Mysql需要手动管理lock_version
```ruby
# 这里的方式就是添加lock_version的方式，Mysql和PG都适用
add_cloumn :comments, :lock_version, :integer, default: 0

ActiveRecord::Base.transaction do
  comment = Comment.find('01193f601ead013587ed4ccc6afe409e')
  puts comment.lock_version
  sleep 30
  comment.body = 'update in options'
  comment.save!
  puts comment.lock_version
end

ActiveRecord::Base.transaction do
  comment = Comment.find('01193f601ead013587ed4ccc6afe409e')
  puts comment.lock_version
  comment.body = 'should show at last'
  comment.save
  puts comment.lock_version
end

# p2会报异常Raises a ActiveRecord::StaleObjectError
p1 = Product.find(1)
p2 = Product.find(1)

p1.name = "Michael"
p1.save

p2.name = "should fail"
p2.save
```

在Postgresql中可以设置事物隔离级别来实现(PG的MVCC实现可以代替乐观锁 :repeatable_read,但是Mysql的事物隔离级别和PG不同，Mysql的不可以实现乐观锁的功能)
嵌套事务不适用
```ruby
# postgresql中的乐观锁
Comment.transaction isolation: :repeatable_read do
  comment =Comment.find('8f5dea901ead013587ed4ccc6afe409e')
  comment.body = 'uuuuuuuppppppdddddddaaaaatttttee'
  comment.save!
end
```
对于乐观锁，还需要注意如果是前端操作频繁，那么还需要把 lock_version 写入到 form 表单中，否则起不到锁的作用

## 嵌套事物死锁
```ruby
ActiveRecord::Base.transaction do
  a = User.find(1)
  b = User.find(2)
  
  a.with_lock do
    a.increment! money: 100
    b.with_lock do
      b.decrement! money: 100
    end
  end
end
```
![IMAGE](resources/F3038D069FE5A59F0BED1B395B3ED6E8.jpg =679x397)

资源排序一定程度上避免死锁
![IMAGE](resources/579E22309832D6764486AD33AB7907E6.jpg =657x377)

## 参考
https://ruby-china.org/topics/28963
https://ruby-china.org/topics/19499