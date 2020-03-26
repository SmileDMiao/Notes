## Postgresql并发控制
>MVCC:Multiversion Currency Control

### 事物隔离级别MVCC
---
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

## Postgresql显示锁定
#### 表级锁
两种锁模式之间真正的区别是它们有着不同的冲突锁集合，两个事务在同一时刻不能在同一个表上持有相互冲突的锁。不过，一个事务决不会和自身冲突。
```sql
BEGIN WORK;
LOCK TABLE films IN SHARE MODE;
SELECT id FROM films
    WHERE name = 'Star Wars: Episode I - The Phantom Menace';
-- 如果记录没有被返回就做 ROLLBACK
INSERT INTO films_user_comments VALUES
    (_id_, 'GREAT! I was waiting for it for so long!');
COMMIT WORK;
```
1. ACCESS SHARE
2. ROW SHARE
3. ROW EXCLUSIVE
4. SHARE UPDATE EXCLUSIVE
5. SHARE
6. SHARE ROW EXCLUSIVE
7. EXCLUSIVE
8. ACCESS EXCLUSIVE

冲突的锁模式
![IMAGE](resources/7268ABD1F570F53C7985BE5ADD048EFE.jpg =1416x316)

#### 行级锁
一个事务可能会在相同的行上保持冲突的锁，甚至是在不同的子事务中。但是除此之外，两个事务永远不可能在相同的行上持有冲突的锁。行级锁不影响数据查询，它们只阻塞对同一行的写入者和加锁者。
1. FOR UPDATE
2. FOR NO KEY UPDATE
3. FOR SHARE
4. FOR KEY SHARE

冲突的行级锁
![IMAGE](resources/038EA298B4BB12DF75CEAF8CD306FA0F.jpg =747x200)

### drop truncate delete
1. drop table, truncate table , 最简单直接
2. DELETE ，版本被保留。所以需要delete+vacuum 。