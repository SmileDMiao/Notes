## Postgresql并发控制
---
>MVCC:Multiversion Currency Control

### 事物隔离级别MVCC
在数据库中，并发的数据库操作会面临脏读（Dirty Read）、不可重复读（Nonrepeatable Read）、幻读（Phantom Read）和串行化异常等问题，为了解决这些问题，在标准的SQL规范中对应定义了四种事务隔离级别：

1. RU(Read uncommitted): 读未提交
2. RC(Read committed): 读已提交
3. RR(Repeatable read): 重复读
4. SERIALIZABLE(Serializable): 串行化

需要注意的是，在PostgreSQL中:
RU隔离级别不允许脏读，实际上和Read committed一样, RR隔离级别不允许幻读.

在各个级别上被禁止出现的现象是：
1. 脏读(Dirty Read): 一个事务读取了另一个并行未提交事务写入的数据。
2. 不可重复读:(Nonrepeatable Read) 一个事务重新读取之前读取过的数据，发现该数据已经被另一个事务（在初始读之后提交）修改。
3. 幻读(Phantom Read): 一个事务重新执行一个返回符合一个搜索条件的行集合的查询， 发现满足条件的行集合因为另一个最近提交的事务而发生了改变。
4. 序列化异常(Serialization Anomaly): 成功提交一组事务的结果与一次运行这些事务的所有可能顺序不一致。

![IMAGE](resources/C01DF053CB0C28A34B9F14B5BA8F2A54.jpg =793x171)

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
```sql
-- 创建表时设置fillfactorc
reate table table_name()with (fillfactor=100);

-- 更新fillfactor
ALTER TABLE table_name SET ( fillfactor = 50);
VACUUM FULL table_name;
```

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