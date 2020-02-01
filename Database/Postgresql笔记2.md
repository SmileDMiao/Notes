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


## FILLFACTOR
---
PostgreSQL每个表和索引的数据都是由很多个固定尺寸的页面存储  ,PostgreSQL中数据操作永远是Append操作,具体含义如下:
+ insert 时向页中添加一条数据
+ update 将历史数据标记为无效,然后向页中添加新数据
+ delete 将历史数据标记为无效
因为这个特性,所以需要定期对数据库vacuum,否则会导致数据库膨胀,建议打开autovacuum.
ctid: (0,59)表示数据存放位置为第0个页面的第59行
```sql
select ctid , * from tablename;
```
对一个表批量的插入数据，查询ctid，如果fillfactor设置为100，假设ctid这样分布: (0,99)(0,100)(1,1)(1,2)，就是指数据的存储将每页存满之后再往下一页存。那么如果将fillfactor设置为60，ctid会这样分布: (0,59)(0, 60)(1,1)(1,2)，存储数据的时候不会将每页都存满在往下一页存。那么在更新和删除的时候是怎样的，如果fillfactor设置为100，假设当前表存储的最新ctid是(22,100)，在更新(0,16)这条记录的时候会往当前表最新ctid后面添加那这里就是(23,1)，如果设置的fillfactor是60，那么在更新和删除的时候不会往最新的ctid后面添加数据，而是会往那40%没有填充数据的地方写入数据。

+ autovacuum非常重要,必须要打开并设置合适的参数
+ fillfactor会降低insert的性能,但是update和delete性能将有提升

## VACUUM AND AUTOVACUUM
---
VACUUM回收死行占据的存储空间，那些已经DELETE的行或者被UPDATE过后过时的行并没有从它们所属的表中物理删除
```sql
-- vacuum不能在事物里执行
vacuum(
full: 完全清理，会锁表也会占用更多的空间
verbose: 打印清理报告
analyze: 更新用于优化器的统计信息，以决定执行查询的最有效方法。
) tablename;

-- 查看每个表的vacuum时间
select relname,last_vacuum, last_autovacuum, last_analyze, last_autoanalyze from pg_stat_user_tables;
```

AUTOVACUUM:
```shell
# 配置在postgresql.conf中
# AUTOVACUUM PARAMETERS
autovacuum：是否启动系统自动清理功能，默认值为on。
autovacuum_max_workers：设置系统自动清理工作进程的最大数量
autovacuum_naptime：设置两次系统自动清理操作之间的间隔时间。
log_autovacuum_min_duration = -1: -1禁用对自动清理动作的记录，0记录所有的自动清理动作

autovacuum_vacuum_threshold 和 autovacuum_analyze_threshold ：设置当表上被更新的元组数的阈值超过这些阈值时分别需要执行vacuum和analyze。
autovacuum_vacuum_scale_factor设置表大小的缩放系数。
autovacuum_freeze_max_age：设置需要强制对数据库进行清理的XID上限值。
autovacuum_vacuum_scale_factor = 0.2
autovacuum_analyze_scale_factor = 0.1
```

```sql
-- 触发条件: pg_stat_all_tables.n_dead_tup >= threshold + pg_class.reltuples * scale_factor
SELECT nspname AS schemaname, 
         relname as tablename, 
         reltuples
    FROM pg_class C 
         LEFT JOIN pg_namespace N ON (N.oid = C.relnamespace)
   WHERE nspname NOT IN ('pg_catalog', 'information_schema') AND relkind='r'
ORDER BY reltuples DESC;

select relname, 
       n_dead_tup, 
       n_live_tup 
  from pg_stat_all_tables 
 where schemaname !~ 'pg_';
```