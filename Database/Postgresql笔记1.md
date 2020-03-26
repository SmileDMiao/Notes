## PostgreSQL interactive terminal
---
```shell
# 生成当前用户的同名数据库
createdb
```
```sql
-- 我本地执行sql不显示执行时间，这个是个开关
\timing

-- 查看所有数据库
psql -l
\l
SELECT datname FROM pg_database;

-- 查看安装的扩展
\dx

-- 查看索引(索引大小)
\di || \di+

-- 切换数据库
\c dbname

-- 用户
\du

-- 查看当前库所有表
\dt
\dt+ tablename(optional)查看额外的信息(比如大小)

-- 描述数据表 索引类型
\d tablename

-- 连接数据库
psql -d DBname -h Hostname -p Port -U Username

-- 查看数据文件夹
show data_directory
```

## PostgreSQL SQL
---
```sql
-- 添加列
ALTER TABLE old_metrics ADD used character varying;

-- 清空表
TRUNCATE projects;

-- 创建用户
create user 'username' with password 'password' createdb; ALTER USER 'username' WITH SUPERUSER;

-- 创建数据库
CREATE DATABASE name WITH OWNER = username
                          TEMPLATE = template
                          ENCODING = encoding
                          LC_COLLATE = lc_collate
                          LC_CTYPE = lc_ctype
                          TABLESPACE = tablespace_name
                          CONNETION LIMIT = connlimit
                          
-- ENCODING: 数据库的字符集(CHARACTER SET)，需要和指定的lc_ctype和lc_collate必须兼容。
-- 查看字符集支持的lc_ctype和lc_collate
select pg_encoding_to_char(collencoding) as encoding, collname, collcollate, collctype from pg_collation;
-- LC_TYPE: 字符分类
-- LC_COLLATE: 字符串排序顺序

-- 创建索引
-- 这种方式创建索引不适合线上操作会造成锁表的情况
CREATE INDEX INDEX_NAME ON TABLE_NAME USING GIN(COLUMN_NAME, PATTERN)
-- 支持在线创建索引，不堵塞其他会话
CREATE INDEX CONCURRENTLY

-- 查看索引大小
select pg_size_pretty(pg_relation_size('index_name'));
-- 通过pg_stat_user_indexes.idx_scan可检查利用索引进行扫描的次数
select idx_scan from pg_stat_user_indexes where indexrelname = 'ind_t_id';

-- 位置参数:一个位置参数引用被用来指示一个由 SQL 语句外部提供的值
-- $number
-- $$:实际代码的开始，当遇到下一个 $$ 的时候， 为代码的结束
CREATE FUNCTION add_em(integer, integer) RETURNS integer AS $$
    SELECT $1 + $2;
$$ LANGUAGE SQL;
SELECT add_em(1, 2) AS answer;

-- 这里的两个冒号为PostgreSQL-风格的类型转换
select (random()*100)::int from generate_series(1,10);

-- 查看已定义的操作符类
SELECT am.amname AS index_method,
       opc.opcname AS opclass_name,
       opc.opcintype::regtype AS indexed_type,
       opc.opcdefault AS is_default
    FROM pg_am am, pg_opclass opc
    WHERE opc.opcmethod = am.oid
    ORDER BY index_method, opclass_name;
```

### 创建表Copy Structure from another table
```sql
-- 根据现有表创建新表
-- including all
create table tbl_inherits_partition (like tbl_inherits_parent including constraints including indexes including defaults);

-- create table as select: define a new table from the results of a query
create table tbl_inherits_partition as select * from tbl_inherits_parent;

-- create table ... as table ... with {data|no data}: 创建一个和原表结构相同的新表，保留或不保留数据，但是不会继承原表的约束，索引等。
create table tbl_inherits_partition as table tbl_inherits_parent with data;

-- select * into new_table from table: 将结果集保存在新表中，但是只能执行一次。
select * into tbl_inherits_partition from tbl_inherits_parent ;
```

## 用到的函数和操作符
---
+ gen_random_uuid: PostgreSql自带的类型，函数支持需要安装模块pgcrypto，这个模块提供了好多关于加密的函数
+ generate_series: 生成多条记录
+ random: random value in the range 0.0 <= x < 1.0
+ substring: 切割字符串
+ string || string: 串接字符串
+ clock_timestamp: 当天时间戳

#### POSIX正则表达式
+ ~: 表示查询关键字左边的字段匹配右边表达式的记录
+ ~*: 表示查询关键字左边的字段匹配右边表达式的记录，并且不区分大小写
+ !~: 表示查询关键字左边的字段不匹配右边表达式的记录
+ !~*: 表示查询关键字左边的字段不匹配右边表达式的记录，并且不区分大小写
```ruby
User.where("name ~* ?", 'ruby')
```

#### 操作符
+ ~~: like
+ ~~*: ilike
+ !~~: NOT LIKE
+ !~~*: NOT ILIKE

## 生成测试数据
---
```sql
-- 准备测试数据
create table test( 
  id int8,  
  c1 int8 default 0,  
  c2 int8 default 0,  
  c3 int8 default 0,  
  c4 float8 default 0,  
  c5 text default 'hello world postgresql',  
  ts timestamp default clock_timestamp()  
)
insert into articles select id, substring(md5(random()::text),1,8), md5(random()::text), 0, null, null, 0, '{}', '{}', clock_timestamp(), clock_timestamp() from generate_series(1, 10) id;

-- 可以通过这种方式生成一个随机字符串 
substring(md5(random()::text),1,8)

-- 生成随机的中文
create or replace function gen_hanzi(int) returns text as $$  
declare  
  res text;  
begin  
  if $1 >=1 then  
    select string_agg(chr(19968+(random()*20901)::int), '') into res from generate_series(1,$1);  
    return res;  
  end if;  
  return null;  
end;  
$$ language plpgsql strict;

select gen_hanzi(10) from generate_series(1,10);

-- 随机数组
create or replace function gen_rand_arr(int,int) returns int[] as $$    
  select array_agg((random()*$1)::int) from generate_series(1,$2);    
$$ language sql strict;
select gen_rand_arr(100,50)
```

## 窗口函数
---
窗口函数提供在与当前查询行相关的行集合上执行计算的能力，必须使用窗口函数的语法调用这些窗口函数，一个OVER子句是必需的。
```sql
CREATE TABLE products (
  id varchar(10),
  name text,
  price numeric,
  uid varchar(14),
  type varchar(100)
);
-- 按照价格排序
-- row_number(): 在其分区重点当前行号，从1计
select type, name, price, row_number() over(order by price asc) as idx from products;

-- 在类别内按照价格排序
select type, name, price, row_number() over(PARTITION by type order by price asc)  as idx from products;

-- rank(): 在有间隔的当前排行，与它的第一个相同行的row_number相同(1134)
-- 在上面的基础上，如果一个类别内价格相同那么idx一样
select type, name, price, rank() over(PARTITION by type order by price) from products;

-- dense_rank(): 没有间隔的当前行排名；这个函数计数对等组。(1123)
select type, name, price, dense_rank() over(PARTITION by type order by price) from products;

-- 窗口函数+聚合函数配合使用示例
select 
     id,type,name,price,
     sum(price) over w1 类别金额合计,
     (sum(price) over (order by type))/sum(price) over() 类别总额占所有品类商品百分比,
     round(price/(sum(price) over w2),3) 子除类别百分比,
     rank() over w3 排名,
     sum(price) over() 金额总计
from 
     products 
WINDOW 
     w1 as (partition by type),
     w2 as (partition by type rows between unbounded preceding and unbounded following),
     w3 as (partition by type order by price desc)
 ORDER BY 
     type,price asc; 
```

## 递归查询
---
```sql
create table area(
  id int,
  name varchar(32),
  parent_id int
);
insert into area values (1, '中国'   ,0);
insert into area values (2, '江苏省'   ,1);
insert into area values (3, '湖南省'   ,1);
insert into area values (4, '盐城市'   ,2);
insert into area values (5, '苏州市'   ,2);
insert into area values (6, '长沙市'   ,3);
insert into area values (7, '射阳县' ,4);
insert into area values (8, '建湖县' ,4);

-- RECURSIVE: 递归查询关键字
-- T: 申明的虚拟表，每次递归一层后都会将本层数据写入T中
-- T()里面可以定义参数，个数要与后面的查询参数一致，select * 就不写了
-- union all: 用于合并查询结果
-- area.parent_id = T.id: 取虚拟表的ID和实体表parent_id连，这个条件决定了当前递归查询的查询方式(向上查询还是向下查询)
-- 查询江苏省下面的所有结果(向下查找)
WITH RECURSIVE T AS (
       SELECT * FROM area WHERE id = 2
     union ALL 
       SELECT area.* FROM area, T WHERE area.parent_id = T.id 
     ) 
SELECT * FROM T ORDER BY id;
-- 向上查找
WITH RECURSIVE T AS (
       SELECT * FROM area WHERE id = 7
     union ALL 
       SELECT area.* FROM area, T WHERE area.id = T.parent_id
     ) 
SELECT * FROM T ORDER BY id;
```