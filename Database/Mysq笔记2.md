## Mysql帮助命令
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
```


## EXplain
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