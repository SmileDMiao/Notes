## 视图pg_locks
---
提供了数据库服务器上打开事务中保持的锁的信息
| 名字 | 引用 | 描述 |
| :-----: | :----: | :----: |
| locktype |  | 可锁对象的类型 |
| database | pg_database.oid | 锁目标存在的数据库的OID |
| relation | pg_class.oid | 锁目标的关系的OID |
| virtualxid | | 锁目标的事务虚拟ID |
| transactionid |  | 锁目标的事务ID |
| virtualtransaction |  | 保持这个锁或者正在等待这个锁的事务的虚拟ID |
| mode |  | 此进程已持有或者希望持有的锁模式 |
| granted |  | 如果锁已授予则为真，如果锁被等待则为假 |

## 视图pg_stat_activity
---
这非常有用的视图, 可以分析排查当前运行的SQL任务以及一些异常问题。pg_stat_activity 每行展示的是一个process的相关信息，这里的“process”可以理解为一个用户连接。
| 名字 | 类型 | 描述 |
| :-----: | :----: | :----: |
| datid | oid | 数据库的OID |
| datname | name | 数据库的名称 |
| pid | integer | 后端的进程ID |
| usesysid | oid | 登陆后端的用户OID |
| usename | name | 登陆到该后端的用户名 |
| backend_start |  | 当前后台进程开始的时间 |
| xact_start |  | 这个进程的当前事务被启动的时间 |
| query_start |  | 当前活动查询被开始的时间 |
| waiting | boolean | 如果后端当前正等待锁则为true否则为false |
| query | text | 这个后端最近查询的文本, 如果state为active，这个域显示当前正在执行的查询。在所有其他状态下，它显示上一个被执行的查询 |
| state_change | timestamp | 上次状态改变的时间 |
| state | text | 这个后端目前的状态 |
| wait_event | text | 如果后端当前正在等待则显示等待事件名称 |
| wait_event_type | text | 后端正在等待的事件类型 |

#### state状态:
1. active: 后端正在执行一个查询
2. idle: 后端正在等待一个新的客户端命令
3. idle in transaction: 后端在事务中, 但是目前无法执行查询
4. idle in transaction (aborted): 空闲事务(被终止):这个情况类似于空闲事务，除了事务导致错误的一个语句之一
5. fastpath function call: 后端正在执行一个快速路径函数。
6. disabled: 如果后端禁用track_activities，则报告这个状态。

#### wait_event 和 wait_event_type
wait_event(LOCK): 后端正在等待一个重量级的锁。重量级锁，也称为锁管理器锁或简单锁，主要保护SQL可见对象
wait_event_type(relation): 正在等待获取一个relation的锁
wait_event_type(tuple): 正在等待获取一个元组的锁
wait_event_type(transactionid): 正在等待事物完成

## 表级锁
---
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

| 锁类型 | 对应的数据库操作 | 冲突的锁模式 |
| :-----: | :----: | :----: |
| ACCESS SHARE | SELECT(一般来说只读操作都会获取) | 只有ACCESS EXCLUSIVE阻塞SELECT(不包含FOR UPDATE/SHARE语句)。 |
| ROW SHARE | SELECT FOR UPDATE, SELECT FOR SHARE | EXCLUSIVE, ACCESS EXCLUSIVE  |
| ROW EXCLUSIVE | UPDATE,DELETE,INSERT |  SHARE, SHARE ROW EXCLUSIVE, EXCLUSIVE, ACCESS EXCLUSIVE   |
| SHARE UPDATE EXCLUSIVE | VACUUM(WITHOUT FULL), ANALYZE, CREATE INDEX CONCURRENTLY | SHARE UPDATE EXCLUSIVE, SHARE, SHARE ROW EXCLUSIVE, EXCLUSIVE, ACCESS EXCLUSIVE |
| SHARE |  CREATE INDEX(WITHOUT CONCURRENTLY) | ROW EXCLUSIVE, SHARE UPDATE EXCLUSIVE, SHARE ROW EXCLUSIVE, EXCLUSIVE, ACCESS EXCLUSIVE |
| SHARE ROW EXCLUSIVE | 任何Postgresql命令不会自动获得这种锁 | ROW EXCLUSIVE, SHARE UPDATE EXCLUSIVE, SHARE, SHARE ROW EXCLUSIVE, EXCLUSIVE, ACCESS EXCLUSIVE  |
| EXCLUSIVE | 任何Postgresql命令不会自动获得这种锁 | ROW SHARE, ROW EXCLUSIVE, SHARE UPDATE EXCLUSIVE, SHARE, SHARE ROW EXCLUSIVE, EXCLUSIVE |
| ACCESS EXCLUSIVE | DROP TABLE, TRUNCATE, REINDEX, CLUSTER, VACUUM FULL, REFRESH MATERIALIZED VIEW(WITHOUT CONCURRENTLY) | 和所有锁模式冲突 |  

## 冲突的锁模式
![IMAGE](resources/2CF19E939E7701648794631B8EACE61B.jpg =1131x278)

## 事务的锁
看两个场景:
1. 事务1和2更新同一条记录, 第一个事务先执行, 第二个事务等待第一个事务完成, 现在查看锁信息, 发现更新行的锁 `(ROW EXCLUSIVE)` 都是 `granted`, 但是 ROW EXCLUSIVE 本身不会和自己冲突, 那么到底是什么锁了什么?
2. 死锁问题, 事务1更细 AB, 事务2更新 BA, 死锁之后的日志提示的是大概是这个样子 `Process 9220 waits for ShareLock on transaction 7197890; blocked by process 9336.` 问题来了, 这里为什么是等待的是一个 `ShareLock` ? 更细数据获取的明明是 `(ROW EXCLUSIVE)` 且查看锁信息, `ROW EXCLUSIVE` 也都是 `granted`

那么如何解释这两个问题:
关键是数据库的 ShareLock, 这意味着一个事务在等待另一个事务 commit/rollback.
PostgreSQL 事务在开始的时候会在当前 TransactionID 获取一个 ExclusiveLock, 如果两个transaction更新同一条记录, 等待这个事务完成对其他事物会先尝试获取该事务的 ShareLock, 在 commit/abort 时释放ExclusiveLock之前，它将被阻塞。

## 并发更新时候的死锁
对更新的资源做好排序, 并发更新也不会有死锁(删除同理)
```sql
-- 会有死锁
UPDATE topics SET switch = 0 WHERE topicable_id = 3 AND topicable_type = 'Battle';

-- 不会有
UPDATE
	topics t1
SET
	switch = 0
FROM (
	SELECT
		id,
		topicable_type
	FROM
		topics
	WHERE
		topicable_id = 11
		AND topicable_type = 'Match'
	ORDER BY
		id) t
WHERE
	t1.id = t.id
	AND t1.topicable_type = t.topicable_type;
UPDATE topics SET switch = 0 WHERE topicable_id = 3 AND topicable_type = 'Battle';
```


## 行级锁
一个事务可能会在相同的行上保持冲突的锁，甚至是在不同的子事务中。但是除此之外，两个事务永远不可能在相同的行上持有冲突的锁。行级锁不影响数据查询，它们只阻塞对同一行的写入者和加锁者。
跳过锁:
for update skip locked;

冲突的行级锁
![IMAGE](resources/1F1C78C78C0F8F6D7FFE659B77109C07.jpg =600x162)

## 页面级锁 
除了表级别和行级别的锁以外, 页面级别的共享/排他锁也用于控制共享缓冲池中表页面的 读/写。 这些锁在抓取或者更新一行后马上被释放。应用程序员通常不需要关心页级锁。

## 锁和索引
B-tree GiST SP-GiST indexes:
短期的页面级共享/排他锁用于读/写访问。锁在索引行被插入/抓取后立即释放。 这种索 引类型提供了无死锁条件的最高级的并发性。
Hash index: 
Hash 桶级别的共享/排他锁用于读/写访问。锁在整个 Hash 桶处理完成后释放。 Hash 桶 级锁比索引级的锁提供了更好的并发性但是可能产生死锁， 因为锁持有的时间比一次索 引操作时间长。

## SQL
---
```sql
-- 设置锁超时时间
SET lock_timeout TO '2s'

-- 查询耗时较长的查询
SELECT
	CURRENT_TIMESTAMP - query_start AS "runtime",
	usename,
	datname,
	state,
	query
FROM
	pg_stat_activity
WHERE
	now() - query_start > '2 minutes'::interval
ORDER BY
	runtime DESC;
	
-- 查询当前因为lock而被block的SQL
SELECT
	datname,
	usename,
	query
FROM
	pg_stat_activity
WHERE
	waiting;
	
-- 查询由于等待锁而阻塞的查询
SELECT
	*
FROM
	pg_stat_activity
WHERE
	datname = ''
	AND wait_event_type = 'Lock'
	
-- 取消一个正在运行的查询
SELECT pg_cancel_backend(pid)
-- 终止一个正在运行的查询
SELECT pg_terminate_backend(pid);

-- 查询当前运行的blocking语句与blocked语句
SELECT
	pg_locks.mode,
	pg_locks.relation,
	pg_class.relname,
	pg_locks.pid AS blocked_pid,
	blocked_activity.usename AS blocked_user,
	blocking_locks.pid AS blocking_pid,
	blocking_activity.usename AS blocking_user,
	blocked_activity.query AS blocked_statement,
	blocking_activity.query AS current_statement_in_blocking_process
FROM
	pg_locks
	JOIN pg_stat_activity blocked_activity ON blocked_activity.pid = pg_locks.pid
	JOIN pg_locks blocking_locks ON blocking_locks.locktype = pg_locks.locktype
		AND blocking_locks.database IS NOT DISTINCT FROM pg_locks.database
		AND blocking_locks.relation IS NOT DISTINCT FROM pg_locks.relation
		AND blocking_locks.page IS NOT DISTINCT FROM pg_locks.page
		AND blocking_locks.tuple IS NOT DISTINCT FROM pg_locks.tuple
		AND blocking_locks.virtualxid IS NOT DISTINCT FROM pg_locks.virtualxid
		AND blocking_locks.transactionid IS NOT DISTINCT FROM pg_locks.transactionid
		AND blocking_locks.classid IS NOT DISTINCT FROM pg_locks.classid
		AND blocking_locks.objid IS NOT DISTINCT FROM pg_locks.objid
		AND blocking_locks.objsubid IS NOT DISTINCT FROM pg_locks.objsubid
		AND blocking_locks.pid != pg_locks.pid
	JOIN pg_stat_activity blocking_activity ON blocking_activity.pid = blocking_locks.pid
	LEFT JOIN LATERAL (
		SELECT
			relname,
			oid
		FROM
			pg_class
		WHERE
			pg_class.oid = pg_locks.relation) pg_class ON pg_class.oid = pg_locks.relation
WHERE
	NOT pg_locks.granted;
	
-- 查询运行的时间
SELECT
	a.datname,
	l.relation::regclass,
	l.transactionid,
	l.mode,
	l.GRANTED,
	a.usename,
	a.query,
	a.query_start,
	age(now(), a.query_start) AS "age",
	a.pid
FROM
	pg_stat_activity a
	JOIN pg_locks l ON l.pid = a.pid
ORDER BY
	a.query_start;
	
-- 查看当前事务锁等待和持锁信息的SQL
WITH WAIT AS (
	SELECT
		a.mode,
		a.locktype,
		a.database,
		a.relation,
		a.page,
		a.tuple,
		a.classid,
		a.granted,
		a.objid,
		a.objsubid,
		a.pid,
		a.virtualtransaction,
		a.virtualxid,
		a.transactionid,
		a.fastpath,
		b.state,
		b.query,
		b.xact_start,
		b.query_start,
		b.usename,
		b.datname,
		b.client_addr,
		b.client_port,
		b.application_name
	FROM
		pg_locks a,
		pg_stat_activity b
	WHERE
		a.pid = b.pid
		AND NOT a.granted
),
RUN AS (
	SELECT
		a.mode,
		a.locktype,
		a.database,
		a.relation,
		a.page,
		a.tuple,
		a.classid,
		a.granted,
		a.objid,
		a.objsubid,
		a.pid,
		a.virtualtransaction,
		a.virtualxid,
		a.transactionid,
		a.fastpath,
		b.state,
		b.query,
		b.xact_start,
		b.query_start,
		b.usename,
		b.datname,
		b.client_addr,
		b.client_port,
		b.application_name
	FROM
		pg_locks a,
		pg_stat_activity b
	WHERE
		a.pid = b.pid
		AND a.granted
),
OVERLAP AS (
	SELECT
	r.*
FROM
	WAIT w
	JOIN RUN r ON (r.locktype IS NOT DISTINCT FROM w.locktype
			AND r.database IS NOT DISTINCT FROM w.database
			AND r.relation IS NOT DISTINCT FROM w.relation
			AND r.page IS NOT DISTINCT FROM w.page
			AND r.tuple IS NOT DISTINCT FROM w.tuple
			AND r.virtualxid IS NOT DISTINCT FROM w.virtualxid
			AND r.transactionid IS NOT DISTINCT FROM w.transactionid
			AND r.classid IS NOT DISTINCT FROM w.classid
			AND r.objid IS NOT DISTINCT FROM w.objid
			AND r.objsubid IS NOT DISTINCT FROM w.objsubid
			AND r.pid <> w.pid)),
UNIONALL AS (
	SELECT
		r.*
	FROM
		OVERLAP r
	UNION ALL
	SELECT
		w.*
	FROM
		WAIT w) 
select 
	locktype,
	datname,
	relation::regclass,
	page,
	tuple,
	virtualxid,
	transactionid::text,
	classid::regclass,
	objid,
	objsubid,   
	string_agg(   
	'Pid: '||case when pid is null then 'NULL' else pid::text end||chr(10)||   
	'Lock_Granted: '||case when granted is null then 'NULL' else granted::text end||' , Mode: '||case when mode is null then 'NULL' else mode::text end||' , FastPath: '||case when fastpath is null then 'NULL' else fastpath::text end||' , VirtualTransaction: '||case when virtualtransaction is null then 'NULL' else virtualtransaction::text end||' , Session_State: '||case when state is null then 'NULL' else state::text end||chr(10)||   
	'Username: '||case when usename is null then 'NULL' else usename::text end||' , Database: '||case when datname is null then 'NULL' else datname::text end||' , Client_Addr: '||case when client_addr is null then 'NULL' else client_addr::text end||' , Client_Port: '||case when client_port is null then 'NULL' else client_port::text end||' , Application_Name: '||case when application_name is null then 'NULL' else application_name::text end||chr(10)||    
	'Xact_Start: '||case when xact_start is null then 'NULL' else xact_start::text end||' , Query_Start: '||case when query_start is null then 'NULL' else query_start::text end||' , Xact_Elapse: '||case when (now()-xact_start) is null then 'NULL' else (now()-xact_start)::text end||' , Query_Elapse: '||case when (now()-query_start) is null then 'NULL' else (now()-query_start)::text end||chr(10)||    
	'SQL (Current SQL in Transaction): '||chr(10)||  
	case when query is null then 'NULL' else query::text end,    
	chr(10)||'--------'||chr(10)    
	order by    
	  ( case mode    
	    when 'INVALID' then 0   
	    when 'AccessShareLock' then 1   
	    when 'RowShareLock' then 2   
	    when 'RowExclusiveLock' then 3   
	    when 'ShareUpdateExclusiveLock' then 4   
	    when 'ShareLock' then 5   
	    when 'ShareRowExclusiveLock' then 6   
	    when 'ExclusiveLock' then 7   
	    when 'AccessExclusiveLock' then 8   
	    else 0   
	  end  ) desc,   
	  (case when granted then 0 else 1 end)  
	) as lock_conflict  
from unionall   
group by   
locktype,datname,relation,page,tuple,virtualxid,transactionid::text,classid,objid,objsubid; 
```