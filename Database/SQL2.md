## JSON QUERY
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
#### LOWER
```sql
select * from users where lower(name) = lower('ABCD')
```

#### ILIKE
```sql
select * from uses where name Ilike 'ABCD'
```

#### CITEXT:不区分大小写的类型
```sql
CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;
ALTER TABLE users ALTER COLUMN username TYPE citext;
```

#### POSIX
```sql
SELECT id FROM groups WHERE name ~* 'adm'
```

## LATERAL
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
+ A和B的数据类型、值不完全相同则返回FALSE。
+ A和B的数据类型、值都相同返回TURE。
+ 将空值视为相同。