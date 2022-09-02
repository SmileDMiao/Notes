## Mysql命令
---
```sql
-- 查看建表语句
show create table table_name
-- 数据库的函数，存储过程查询表
mysql.proc
-- 查看函数存储过程创建代码
show create procedure proc_name;
show create function func_name;
-- 查看表的索引
show index from table_name

-- 描述一个表
DESCRIBE TABLE NAME

-- 创建数据库
-- character:字符集.collate:排序规则(utf8_bin区分大小写,utf_general_ci不区分大小写)
create schema database_name default character set utf8 collate utf8_bin

# 导入sql文件
mysql -u name -p db < .sql

# 强制使用索引
select * from table_name FORCE index(PRIMARY) where xx;

# create table
CREATE TABLE IF NOT EXISTS users (
    id bigint UNSIGNED AUTO_INCREMENT,
    username text COMMENT '用户登录名' ,
    password text COMMENT '用户登录密码',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
```


## EXplain
---
我们经常为表添加索引以方便更快的搜索，但是我经常有疑问，我怎么知道这条sql执行了到底有没有命中索引呢？总不能通过执行时间吧，这个虽然是最直观的，但我希望从数据上有个直观的感受。

**mysql数据库我所知道的一种方式就是通过慢查询日志**
满查询日志的结果会显示这条sql搜索到多少条结果，总共扫描了多少条数据。
```sql
-- mysql.conf
slow_query_log：是否开启慢查询，0或者OFF为关闭，1或者ON为开启，默认值为OFF，即为关闭
slow_query_log_file：指定慢查询日志存放路径
long_query_time：大于等于此时间记录慢查询日志，精度可达微秒级别，默认为10s。当设置为0时表示记录所谓查询记录
-- 日志:Query_time: 5.007305 Lock_time: 0.000112 Rows_sent: 5 Rows_examined: 10
Query_time: 查询花费的总时间
Lock_time: 等待锁的时间
Rows_sent: 实际获取的数据行数
Rows_examined: 实际扫描的数据行数
```

1. select_type: 显示本行是简单或复杂的查询
2. type: 数据访问, 读取操作类型
3. possible_keys: 揭示哪些索引可以利于提高搜索
4. key: mysql采用哪个索引来优化搜索
5. key_len: mysql在索引里使用的字节数
6. ref: 显示了之前的表在key列记录的索引中查找值所用的列或常量
7. rows: 为了找到所需的行而需要读取的行数, 估算值, 不精确。通过把所有rows列值相乘, 可粗略估算整个查询会检查的行数
8. Extra: 额外信息，如using index、filesort等

SELECT_TYPE:
![IMAGE](resources/2252DEB026760A6BD170C337BBA4A1FB.jpg =1132x359)

TYPE:
type显示的是访问类型，是较为重要的一个指标，结果值从好到坏依次是：
system > const > eq_ref > ref > fulltext > ref_or_null > index_merge > unique_subquery > index_subquery > range > index > ALL
ref: All rows with matching index values are read from this table for each combination of rows from the previous tables.