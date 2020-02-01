## 单记录Upsert
+ Mysql语法: INSERT...ON DUPLICATE KEY UPDATE
+ Postgresql: INSERT ...ON CONFLICT(KEY) DO UPDATE
```sql
-- Mysql
INSERT INTO bulk_upserts (name, email, city, created_at, updated_at) VALUES ('miao', 'aa@qq.com', 'a', '2016-01-01', '2016-01-01')
ON DUPLICATE KEY UPDATE
email = VALUES(email);
-- Postgresql
INSERT INTO bulk_upserts (name, email, city, created_at, updated_at) VALUES ('miao', 'aa@qq.com', 'a', '2016-01-01', '2016-01-01')
ON CONFLICT(name) DO  UPDATE
SET email = EXCLUDED.email;
```

## 批量Upsert
```sql
-- mysql
INSERT INTO bulk_upserts (name, email, city, created_at, updated_at) VALUES ( 'miao1', 'bbcccc@qq.com', 'a', '2016-01-01', '2016-01-01'),
 ('miao2', 'bbcccc@qq.com', 'a', '2016-01-01', '2016-01-01'),
 ('miao2', 'bbccccdd@qq.com', 'a', '2016-01-01', '2016-01-01')
ON DUPLICATE KEY UPDATE
email = VALUES(email);
-- postgresql
-- In PG,  Ensure that no rows proposed for insertion within the same command have duplicate constrained values.
INSERT INTO bulk_upserts (name, email, city, created_at, updated_at) VALUES ( 'miao1', 'bbcccc@qq.com', 'a', '2016-01-01', '2016-01-01'),
 ('miao2', 'bbcccc@qq.com', 'a', '2016-01-01', '2016-01-01'),
 ('miao2', 'bbccccdd@qq.com', 'a', '2016-01-01', '2016-01-01')
ON CONFLICT(name) DO  UPDATE
SET email = EXCLUDED.email;
```

## Ruby Gem activerecord-import
[activerecord-import](https://github.com/zdennis/activerecord-import)这个gem可以批量的进行upsert的操作，其原理就是使用的是 **insert into on duplicate key**,
我们在控制台中可以看到其输出的sql，下面是其批量upsert的语法，这里是pg数据库的写法，在mysql中的写法是不一样的。
```ruby
gem 'activerecord-import'
books = Book.all
books.each do |book|
  book.name = "updated #{book.name}"
end
Book.import books.to_a, :on_duplicate_key_update => [:name]
```

## Ruby Gem upsert and Tips
和upsert类似或相关的Gem包

+ [upsert](https://github.com/seamusabshere/upsert)
+ [activerecord-import](https://github.com/zdennis/activerecord-import)
+ [activerecord-bulkwrite](https://github.com/coin8086/activerecord-bulkwrite)

1. 这个gem的确实现了upsert的功能，也提供了简便使用的语法，但实现方式并不是使用的上面提到的语法，而是会在数据库层创建存储过程，在存储过程里循环判断插入更新。
2. 在使用上的确很方便但是如果类似的需求多了，存储过程的维护又是一个问题。
3. activerecord-import,activerecord-bulkwrite:
这些gem没有使用数据库层面的存储过程，需要数据库支持upsert的操作，高版本的pg和mysql都有支持，在使用的时候由于使用rails的model，在rails中一般created_at，updated_at是默认必填的，而在upsert时，这两个字段是必填的，哪怕数据库中已经存在了同样的数据。


参考[Bulk Upsert for MySQL & PostgreSQL](https://ruby-china.org/topics/32424)