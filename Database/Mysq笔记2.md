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

-- 创建数据库
-- character:字符集.collate:排序规则(utf8_bin区分大小写,utf_general_ci不区分大小写)
create schema database_name default character set utf8 collate utf8_bin

# 导入sql文件
mysql -u name -p db < .sql
```