## B-Tree
支持排序，支持大于、小于、等于、大于或等于、小于或等于的搜索。

## Hash
hash索引存储的是被索引字段VALUE的哈希值，只支持等值查询。hash索引特别适用于字段VALUE非常长(不适合b-tree索引，因为b-tree一个PAGE至少要存储3个ENTRY，所以不支持特别长的VALUE)的场景
```sql
create index hash-index on table t_hash using hash (text)
```

## PG插件：PG_TRGM
提供了两个索引: GIST ; GIN
```sql
CREATE TABLE test_trgm (t text);

CREATE INDEX trgm_idx ON test_trgm USING GIST (t gist_trgm_ops);

CREATE INDEX trgm_idx ON test_trgm USING GIN (t gin_trgm_ops);
```
```ruby
enable_extension :pg_trgm
def up
  execute(
    CREATE INDEX CONCURRENTLY customer_names_on_last_name_idx on table
     USING gin(last_name gin_trgm_ops pattern);
end

def down
  execute(
    DROP INDEX customer_names_on_last_name_idx;)
```

## Gin
gin是倒排索引，存储被索引字段的VALUE或VALUE的元素，以及行号的list或tree。
1. 当需要搜索多值类型内的VALUE时，适合多值类型，例如数组。
2. 当用户的数据比较稀疏时(含有大量空值的情况)，如果要搜索某个VALUE的值，可以适应btree_gin支持普通btree支持的类型
3. 按任意列进行搜索时，gin支持多列展开单独建立索引域

```sql
# 多值类型array
create index gin-index on t_gin using gin (array);
# 稀疏数据搜索
create extension btree_gin;
create index gin_index on t_gin using gin (sparse);
# 多值任意搜索
create index gin-index on t_gin using gin (c1,c2,c3,c4,c5,c6,c7,c8,c9); 
```

## GiST
GiST的意思是通用的搜索树(Generalized Search Tree)。 它是一种平衡树结构的访问方法,在系统中作为一个基本模版,可以使用它实现任意索引模式。B-trees, R-trees和许多其它的索引模式都可以用GiST实现。
Gist索引适用于多维数据类型和集合数据类型，和Btree索引类似，同样适用于其他的数据类型。和Btree索引相比，Gist多字段索引在查询条件中包含索引字段的任何子集都会使用索引扫描，而Btree索引只有查询条件包含第一个索引字段才会使用索引扫描。
Gist索引创建耗时较长，占用空间也比较大。

## BRIN
BRIN 索引是块级索引，有别于B-TREE等索引，BRIN记录并不是以行号为单位记录索引明细，而是记录每个数据块或者每段连续的数据块的统计信息。因此BRIN索引空间占用特别的小，对数据写入、更新、删除的影响也很小。
BRIN属于LOSSLY索引，当被索引列的值与物理存储相关性很强时，BRIN索引的效果非常的好。
例如时序数据，在时间或序列字段创建BRIN索引，进行等值、范围查询时效果很棒。

## RUM
rum 是一个索引插件，由Postgrespro开源，适合全文检索，属于GIN的增强版本。
RUM不仅支持GIN支持的全文检索，还支持计算文本的相似度值，按相似度排序等。同时支持位置匹配。

## BLOOM
bloom索引接口是PostgreSQL基于bloom filter构造的一个索引接口，属于lossy索引，可以收敛结果集(排除绝对不满足条件的结果，剩余的结果里再挑选满足条件的结果)，因此需要二次check，bloom支持任意列组合的等值查询。



## 覆盖索引
只需要在一棵索引树上就能获取SQL所需的所有列数据，无需回表，速度更快。




## BTREE加速模糊搜索
+ 前缀模糊
```sql
-- 索引和查询都需要明确指定collate "C", 不指定模糊查询不会命中索引
create index index_name on table_name(column-name collate "C");
select * from test where info like 'abcd%' collate "C";
"idx_tb" btree (info)
-- 模糊查询精确查询都会命中索引
"idx_tb1" btree(info text_pattern_ops)
```
+ 后缀模糊
```sql
-- 使用反转函数(reverse)索引，可以支持后模糊的查询。
create index idx1 on test(reverse(info) collate "C");
-- 模糊查询: *2514
select * from test where reverse(info) like '4152%';
```

## 左右模糊搜索
pg_trgm索引可以支持前后模糊的查询

1. 如果要高效支持多字节字符(例如中文)数据库lc_ctype不能为"C"
2. 建议输入3个或3个以上字符否则效果不佳
```sql
create extension pg_trgm;      
create table test001(c1 text);
insert into test001 select gen_hanzi(20) from generate_series(1,100000);
create index idx_test001_1 on test001 using gin (c1 gin_trgm_ops);

explain (analyze,verbose,timing,costs,buffers) select * from test001 where c1 like '你%';  
explain (analyze,verbose,timing,costs,buffers) select * from test001 where c1 like '%你好啊%';  
explain (analyze,verbose,timing,costs,buffers)  select * from test001 where c1 like '%三璑筂%' collate "zh_CN";
explain analyze select * from tbl where info ~ 'b05830e';
```