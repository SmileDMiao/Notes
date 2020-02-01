## EXPLAIN AND ANALYZE
```sql
EXPLAIN: 查看sql的执行计划
OPTIONS:
    ANALYZE [ boolean ] 是否真正执行
    VERBOSE [ boolean ] 显示详细信息
    COSTS [ boolean ] 显示代价信息
    BUFFERS [ boolean ] 显示缓存信息
    TIMING [ boolean ] 显示时间信息
    FORMAT { TEXT | XML | JSON | YAML } 输出格式，默认为 text

EXPLAIN (ANALYZE, TIMING, VERBOSE, BUFFERS, COSTS, FORMAT JSON) SELECT * FROM USERS;
```

```shell
QUERY PLAN:
# (PG估算成本cost启动成本..结束成本，rows:估计输出行数，width:平均字段长度大小)
# (实际开始时间..结束时间， rows: 实际输出行数，loops: 该节点循环次数)
 Limit  (cost=0.00..1790.57 rows=10 width=9) (actual time=318.470..5722.171 rows=1 loops=1)
   Output: info
   # shared hit:命中缓存数，read:要进行IO读取的数
   Buffers: shared hit=1439 read=52616
   ->  Seq Scan on public.tb  (cost=0.00..179057.19 rows=1000 width=9) (actual time=318.469..5722.169 rows=1 loops=1)
         Output: info
         Filter: (tb.info ~ '^376823'::text)
         Rows Removed by Filter: 9999999
         Buffers: shared hit=1439 read=52616
 # 计划时间
 Planning time: 0.153 ms
 # 实际执行时间
 Execution time: 5722.193 ms
(10 rows)
Time: 5722.690 ms (00:05.723)
```

```json
                                         QUERY PLAN
-----------------------------------------------------------------------------
 [
   {
     "Plan": { 
       "Node Type": "Limit",
       "Parallel Aware": false,
       "Startup Cost": 0.43, //PG估算成本:cost启动成本
       "Total Cost": 0.52, //PG估算成本:结束成本
       "Plan Rows": 10, //估计输出行数
       "Plan Width": 9, //平均字段长度大小
       "Actual Startup Time": 0.030, // 实际开始时间
       "Actual Total Time": 0.032, // 实际结束时间
       "Actual Rows": 1, // 实际输出行数
       "Actual Loops": 1, //该节点循环次数
       "Output": ["info"],
       "Shared Hit Blocks": 4, //命中缓存数
       "Shared Read Blocks": 0, //要进行IO读取的数
       "Shared Dirtied Blocks": 0,
       "Shared Written Blocks": 0,
       "Local Hit Blocks": 0,
       "Local Read Blocks": 0,
       "Local Dirtied Blocks": 0,
       "Local Written Blocks": 0,
       "Temp Read Blocks": 0,
       "Temp Written Blocks": 0,
       "Plans": [
         {
           "Node Type": "Index Only Scan",
           "Parent Relationship": "Outer",
           "Parallel Aware": false,
           "Scan Direction": "Forward",
           "Index Name": "idx_tb1", // 命中了索引的名称
           "Relation Name": "tb",
           "Schema": "public",
           "Alias": "tb",
           "Startup Cost": 0.43,
           "Total Cost": 8.46,
           "Plan Rows": 1000,
           "Plan Width": 9,
           "Actual Startup Time": 0.030,
           "Actual Total Time": 0.031,
           "Actual Rows": 1,
           "Actual Loops": 1,
           "Output": ["info"],
           "Index Cond": "((tb.info ~>=~ '376823'::text) AND (tb.info ~<~ '376824'::text))",
           "Rows Removed by Index Recheck": 0,
           "Filter": "(tb.info ~ '^376823'::text)",
           "Rows Removed by Filter": 0,
           "Heap Fetches": 1,
           "Shared Hit Blocks": 4,
           "Shared Read Blocks": 0,
           "Shared Dirtied Blocks": 0,
           "Shared Written Blocks": 0,
           "Local Hit Blocks": 0,
           "Local Read Blocks": 0,
           "Local Dirtied Blocks": 0,
           "Local Written Blocks": 0,
           "Temp Read Blocks": 0,
           "Temp Written Blocks": 0
         }
       ]
     },
     "Planning Time": 0.288,
     "Triggers": [],
     "Execution Time": 0.059
   }
 ]
```