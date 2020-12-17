## PostgreSQL interactive terminal
---
```shell
# 生成当前用户的同名数据库
createdb
```
```sql
-- 我本地执行sql不显示执行时间，这个是个开关
\timing

-- 查看当前时区信息
show timezone

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
-- 给用户添加权限
CREATE USER postgres SUPERUSER;
ALTER USER postgres CREATEDB;
ALTER DATABASE postgres OWNER TO postgres;
ALTER ROLE themadeknight WITH CREATEDB;

-- 添加删除列
ALTER TABLE old_metrics ADD used character varying;
ALTER TABLE table_name drop column column_name;

-- 设置改变 NOT NULL
ALTER TABLE table_nae ALTER column drop not null;
ALTER TABLE table_nae ALTER column set not null;

-- 添加删除外建约束
ALTER TABLE OPTIONS DROP CONSTRAINT constranint_name;
ALTER TABLE OPTIONS ADD FOREIGN KEY(topic_id) REFERENCES TOPICS(id) ON UPDATE CASCADE ON DELETE CASCADE;

-- 改变类型
ALTER TABLE users ALTER COLUMN username TYPE citext;
ALTER TABLE users ALTER COLUMN username type citext USING username::bigint;
ALTER TABLE topics ALTER COLUMN switch DROP DEFAULT;
ALTER TABLE topics ALTER switch TYPE bool USING CASE WHEN switch=0 THEN FALSE ELSE TRUE END;
ALTER TABLE topics ALTER COLUMN switch SET DEFAULT true;

-- sleep
SELECT * from pg_sleep(5);

-- 清空表
TRUNCATE projects;

-- 删除表
drop table table_name
-- cascade:delete dependent objects like foreign keys views...
drop table table_name cascade

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
CREATE INDEX ... WHERE: 只为过滤后的数据加索引

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
    
-- 查看所有的参数设置
show all;
-- bitmap scan on or off   
set enable_bitmapscan=off;

-- DROP TRUNCATE DELETE
-- equal rebuild table
drop table, truncate table
-- versions are keeped so need vacuum
DELETE
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

### SQL函数
---
+ gen_random_uuid: PostgreSql自带的类型，函数支持需要安装模块pgcrypto，这个模块提供了好多关于加密的函数
+ generate_series: 生成多条记录
+ random: random value in the range 0.0 <= x < 1.0
+ substring: 切割字符串
+ string || string: 串接字符串
+ clock_timestamp: 当天时间戳

#### POSIX正则表达式
---
1
| 操作符 | 描述 | 例子 |
| :-----| :---- | :----: |
| ~ | 表示查询关键字左边的字段匹配右边表达式的记录,大小写敏感 | 'thomas' ~ '.*thomas.*' |
| ~* | 表示查询关键字左边的字段匹配右边表达式的记录,并且不区分大小写 | 'thomas' ~* '.*Thomas.*' |
| !~ | 表示查询关键字左边的字段匹配右边表达式的记录,大小写敏感 | 'thomas' !~ '.*Thomas.*' |
| !~* | 表示查询关键字左边的字段不匹配右边表达式的记录,并且不区分大小写 | 'thomas' !~* '.*Thomas.*' |

2
| 分隔符 | 描述 |
| :-----| :---- | 
| ? | 一个匹配 0 或者更多个原子的序列. ab?c matches only ac or abc |
| + | 一个匹配 1 或者更多个原子的序列. ab+c matches abc, abbc, abbbc, and so on, but not ac |
| \| |	 abc\|def matches abc or def |

3
| 例子 | 描述 |
| :-----| :---- |
| [hc]?at | matches "at", "hat", and "cat" |
| [hc]*at | matches "at", "hat", "cat", "hhat", "chat", "hcat", "cchchat" |
| [hc]+at | matches "hat", "cat", "hhat", "chat", "hcat", "cchchat", and so on, but not "at" |
| cat|dog | matches "cat" or "dog" |

4
| 分隔符 | 描述 |
| :-----| :---- | 
| ^ |	匹配字符串中的起始位置 |
| . |	匹配任何单个字符. a.c matches "abc", etc., but [a.c] matches only "a", ".", or "c". |
| [] | 括号表达式。匹配方括号中包含的单个字符. [abc] matches "a", "b", or "c". [abcx-z] matches "a", "b", "c", "x", "y", or "z", 等同 [a-cx-z] |
|[^] | 匹配不包含在方括号内的单个字符. [^abc] 匹配除 "a", "b", or "c" 之外的任何字符. 文字字符和范围也可以混合使用 |
| $ | 匹配字符串的结束位置或字符串结束换行符之前的位置 |
| () | 定义标记的子表达式 |
| \n | 匹配第n个标记的子表达式匹配的内容，其中n是从1到9的数字 |
| * | 与前面的元素匹配零次或多次. ab*c matches "ac", "abc", "abbbc". [xyz]* matches "", "x", "y", "z", "zx", "zyx", "xyzzy". (ab)* matches "", "ab", "abab", "ababab" |
| {m,n} | 与前面的元素匹配至少m次，但不超过n次. a{3,5} matches only "aaa", "aaaa", and "aaaaa". |

5
| 例子 | 描述 |
| :-----| :---- | 
| .at | 匹配以 at 结尾, including "hat", "cat", and "bat". |
| [hc]at | 匹配 "hat" and "cat". |
| [^b]at | 匹配 任何匹配 .at 除了 "bat". |
| [^hc]at | 匹配 任何匹配 .at 除了 "hat" and "cat". |
| ^[hc]at | 匹配 "hat" and "cat", 但只有是开头或者新行 |
| [hc]at$ | 匹配 "hat" and "cat", 但只有结尾或新行 |
| \\[.\\] | 匹配被[] 包含的字符串. "[a]" and "[b]". |
| s.* | 匹配零个或多个字符跟在后面. "s" and "saw" and "seed". |

```sql
-- User.where("name ~* ?", 'ruby')
select * from articles where title ~* 'ruby';
```

#### LIKE操作符
---
| 操作符 | SQL | 
| :-----| :---- | 
| ~~ | LIKE | 
| ~~*: | ILIKE |
| !~~ | NOT LIKE |
| !~~* | NOT ILIKE |