## GROUP AND HAVING USAGE
>having是分组（group by）后的筛选条件，分组后的数据组内再筛选
where则是在分组前筛选

```ruby
Answer.where("count > 10").group(:user_id).having("count(user_id) > 20")
```
```sql
select 
user_id,
count(*) as number
from
answers
where count > 10
group by user_id
having count(user_id) > 20
```

## GROUP AND COUNT
```ruby
Order.group(:status, :name).count
```
```sql
SELECT COUNT (*) AS count_all, status AS status
FROM "orders"
GROUP BY status
```

## Range Conditions
```ruby
Client.where(created_at: (Time.now.midnight - 1.day)..Time.now.midnight)
```


## LEFT OUTER JOIN
```ruby
Author.left_outer_joins(:posts).distinct.select('authors.*, COUNT(posts.*) AS posts_count').group('authors.id')
```

## EXISTS?
```ruby
Client.exists?(1)
Client.exists?(id: [1,2,3])
Client.exists?(name: ['John', 'Sergei'])
```

## EXPLAIN
```ruby
User.where(status: 1).explain
```

## 查询除ID外重复的记录