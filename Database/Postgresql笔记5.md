## 基准测试:pgbench
```sql
create table stocks  
(  
  crt_time timestamp(0),    
  indicator1 float8,    
  indicator2 float8,     
  ...
  indicator10 float8  
);  
create index idx_stocks_sec_time on stocks using brin (crt_time) with (pages_per_range=1); 

do language plpgsql $$  
declare  
  sql text;  
begin  
  for i in 1..3000 loop  
    sql := format('create table %I (like stocks including all)', 'stocks'||lpad(i::text, 6, '0') );  
    execute sql;  
  end loop;  
end;  
$$; 
```

**options**
```shell
pgbench -i [option...] [dbname]
```

-M querymode: 提交查询到服务器使用的协议：
  1. simple：使用简单的查询协议。
  2. extended：使用扩展的查询协议。 
  3. prepared：使用带有预备语句的扩展查询协议。

-n: --no-vacuum(do not run VACUUM during initialization)

-r: 在benchmark完成后报告每个命令的平均每语句延迟(从客户的角度看的执行时间)

-F: fillfactor
用给定的填充因子创建pgbench_accounts、pgbench_tellers 和pgbench_branches表。缺省是100。

-f: 声明一个自定义的脚本文件)

-c: 模拟客户端的数量，也就是并发数据库会话的数量。缺省是1。

-t: number of transactions each client runs (default: 10)

-T: 运行时间(决不相信任何只运行几秒钟的测试。使用-t 或-T选项使测试最少运行几分钟, 尝试多次测试)。-t 和-T是互相排斥的。

-j threads
pgbench中工作线程的数量。在多CPU的机器上使用多个线程会很有帮助。 客户端的数量必须是线程数量的倍数，因为每个线程都有相同数量的客户端会话管理。 缺省是1。

-P: --progress=NUM(show thread progress report every NUM seconds)

-D: --define=VARNAME=VALUE(define variable for use by custom script)

```sql
-- statement has too many arguments (maximum is 9)
\set indicator1 random(1,1000)
insert into stocks000001 values (now(),:indicator1,456,8,4,9,443,380,:52,49,772);


pgbench -M prepared -n -r -P 1 -f ./pg_test1.sql blog
```

TPS: 这三个过程，每秒能够完成N个这三个过程
Tps即每秒处理事务数，包括了
1. 用户请求服务器
2. 服务器自己的内部处理
3. 服务器返回给用户