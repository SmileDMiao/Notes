## PG导入导出csv
---
[用 PostgreSQL 的 COPY 导入导出 CSV](https://ruby-china.org/topics/32293)
```sql
COPY products
TO '/path/to/output.csv'
WITH CSV;

COPY products (name, price)
TO '/path/to/output.csv'
WITH CSV HEADERS;

COPY (
  SELECT name, category_name
  FROM products
  LEFT JOIN categories ON categories.id = products.category_id
)
TO '/path/to/output.csv'
WITH CSV;

COPY products
FROM '/path/to/input.csv'
WITH CSV (HEADER);

-- PG::InsufficientPrivilege: ERROR: must be superuser or a member of the pg_write_server_files role to COPY to a file
-- 有时候会遇到这种权限问题
GRANT pg_read_server_files TO themadeknight;
```
## 多进程COPY
---
[timescaledb-parallel-copy](https://github.com/timescale/timescaledb-parallel-copy)
```shell
timescaledb-parallel-copy --db-name postgres --table test --file ./ii.csv --workers 8 --reporting-period 30s -connection "host=localhost user=postgres password=helloworld sslmode=disable" -truncate -batch-size 100000
```

## Postgresql树形结构插件ltree
---
[文档](http://www.postgres.cn/docs/9.4/ltree.html)
```sql
CREATE EXTENSION ltree;
CREATE TABLE test (path ltree);
INSERT INTO test VALUES ('Top');
INSERT INTO test VALUES ('Top.Science');
INSERT INTO test VALUES ('Top.Science.Astronomy');
INSERT INTO test VALUES ('Top.Science.Astronomy.Astrophysics');
INSERT INTO test VALUES ('Top.Science.Astronomy.Cosmology');
INSERT INTO test VALUES ('Top.Hobbies');
INSERT INTO test VALUES ('Top.Hobbies.Amateurs_Astronomy');
INSERT INTO test VALUES ('Top.Collections');
INSERT INTO test VALUES ('Top.Collections.Pictures');
INSERT INTO test VALUES ('Top.Collections.Pictures.Astronomy');

CREATE INDEX path_gist_idx ON test USING gist(path);
CREATE INDEX path_idx ON test USING btree(path);
```

## PG数据库导入导出
---
```shell
# 导出 pg_dump
pg_dump -h xx -U xx  -v -O db > xxx.dump
# 导入 pg_restore
pg_restore --host $host_name --username $user_name --verbose --no-owner --dbname $db_name xx.dump

# 导出
pg_dump -U USERNAME DBNAME > dbexport.pgsql
# 导入
psql -U USERNAME DBNAME < dbexport.pgsql
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

## UNLOGGED TABLE
---
不会写入XLOG的表

```sql
-- 创建unlogged table
CREATE UNLOGGED TABLE name ();
-- 改为unlogged table
ALTER TABLE name SET UNLOGGED;

-- list unlogged table in current db
SELECT relname FROM pg_class WHERE relpersistence = 'u';
```

## XLOG AND WAL
---
WAL(Write-Ahead Logging)就是写在前面的日志,是事物和数据库故障的一个保护。任何试图修改数据库数据的操作都会写一份日志到磁盘。这个日志在PG中叫XLOG。所有的日志都会写在$PGDATA/pg_wal目录下面。

## PARALLER QUERY
---
```sql
-- setup
create table test(id int4 primary key, create_time timestamp without time zone default clock_timestamp(), name character varying(32));
insert into test(id,name) select n,n*random()*10000  from generate_series(1,10000000) n; 

-- enable 
show max_parallel_workers_per_gather;
set max_parallel_workers_per_gather to 4;

-- explain
explain analyze select count(*)  from test_big1 where id <1000000;
```