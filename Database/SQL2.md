## JSON QUERY
---
```ruby
# payload: [{"kind"=>"person"}]
Segment.where("payload @> ?", [{kind: "person"}].to_json)

# data: {"interest"=>["music", "movies", "programming"]}
Segment.where("data @> ?",  {"interest": ["music", "movies", "programming"]}.to_json)
Segment.where("data #>> '{interest, 1}' = 'movies' ")
Segment.where("jsonb_array_length(data->'interest') > 1")
Segment.where("data->'interest' ? :value", value: "movies") 
Segment.where("data -> 'interest' ? :value", value: ['programming'])

# data: {"customers"=>[{:name=>"david"}]}
Segment.where("data #> '{customers,0}' ->> 'name' = 'david' ")
Segment.where("data @> ?",  {"customers": [{"name": "david"}]}.to_json)
Segment.where("data -> 'customers' @> '[{\"name\": \"david\"}]'")
Segment.where(" data -> 'customers' @> ?", [{name: "david"}].to_json)

# data: {"uid"=>"5", "blog"=>"recode"}
Segment.where("data @> ?", {uid: '5'}.to_json)
Segment.where("data ->> 'blog' = 'recode'")
Segment.where("data ->> 'blog' = ?", "recode")
Segment.where("data ? :key", :key => 'uid')
Segment.where("data -> :key LIKE :value", :key => 'blog', :value => "%recode%")

# tags: ["dele, jones", "solomon"]
# get a single tag
Segment.where("'solomon' = ANY (tags)")
# which segments are tagged with 'solomon'
Segment.where('? = ANY (tags)', 'solomon')
# which segments are not tagged with 'solomon'
Segment.where('? != ALL (tags)', 'solomon')
Segment.where('NOT (? = ANY (tags))', 'solomon')
# multiple tags
Segment.where("tags @> ARRAY[?]::varchar[]", ["dele, jones", "solomon"])
# tags with 3 items
Segment.where("array_length(tags, 1) >= 3")
```

## 不区分大小写的查询
---
#### LOWER
```sql
select * from users where lower(name) = lower('ABCD')
```

#### ILIKE
---
```sql
select * from uses where name Ilike 'ABCD'
```

#### CITEXT:不区分大小写的类型
---
```sql
CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;
ALTER TABLE users ALTER COLUMN username TYPE citext;
```

#### POSIX
---
```sql
SELECT id FROM groups WHERE name ~* 'adm'
```

## LATERAL
---
LATERAL连接更像是相关子查询，而不是普通子查询。LATERAL连接右边的函数或子查询必须为其左边的每一行计算一次 - 就像相关的子查询一样 - 而普通的子查询(表表达式)只被评估一次。 
### 带有LATERAL的SQL的计算步骤
1. 逐行提取被 lateral 子句关联（引用）的FROM或JOIN中的ITEM（也叫source table）的记录(s)中的column(s)
2. 使用以上提取的column(s), 关联计算lateral子句中的ITEM
3. lateral的计算结果row(s)，与所有from, join ITEM(s)正常的进行JOIN计算
4. 从1到3开始循环，直到所有source table的行都取尽。

例子1
查找出用户最近的5条事件
```sql
select d.id, to_json(array_agg(da.event_type)) activities
from developers d
join lateral (
  select event_type, developer_id
  from activities
  where developer_id=d.id
  limit 5
) da on d.id=da.developer_id
group by d.id
order by d.id;
```

## IS NOT DISTINCT FROM
---
+ A和B的数据类型、值不完全相同则返回FALSE。
+ A和B的数据类型、值都相同返回TURE。
+ 将空值视为相同。

## 判读字符串中是否包含中文
---
```sql
select 'hello' ~ '[\u2e80-\ua4cf]|[\uf900-\ufaff]|[\ufe30-\ufe4f]';
```

## SQL JOIN
---
![IMAGE](resources/2D1E8253CE95D287293CEE08C55E671A.jpg =843x888)

## SQL 在 Join条件中使用 LIKE
---
```sql
LEFT JOIN PERSON ON PERSON.name LIKE concat('%', Users.name,'%')
```

## SQL Update From SELECT and JOIN
---
```sql
UPDATE
 Persons pe,
 (
  SELECT
  record.uNo,
  TO_DAYS(NOW()) - TO_DAYS(max(record.cTime)) AS aaa
  FROM
   record
  GROUP BY
   record.uNo) t SET pe.rday = t.aaa
WHERE
 pe.uNo = t.uNo;
 
 UPDATE
 Persons pe
 inner join
 (
  SELECT
  record.uNo,
  TO_DAYS(NOW()) - TO_DAYS(max(record.cTime)) AS aaa
  FROM
   record
  GROUP BY
   record.uNo) t  on t.uNo = pe.uNo SET pe.rday = t.aaa
```

## SQL WHERE 两列组合条件
---
```ruby
where("scp.id is null and ( (sug_spare_operations.use_type = 'DOA' AND ssra.id IS NOT NULL) or (sug_spare_operations.use_type != 'DOA' AND ssra.id IS NULL))")
```

## 判断日期是否是周末
---
```sql
-- 星期日为1，星期一为2，星期六为7
 SELECT DAYNAME('2012-12-01'), DAYOFWEEK('2012-12-01');
```

## SUM IF
---
```sql
select sum(if(qty > 0, qty, 0)) as total_qty from inventory_product group by product_id
```