## GROUP AND HAVING USAGE
---
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
---
```ruby
Order.group(:status, :name).count
```
```sql
SELECT COUNT (*) AS count_all, status AS status
FROM "orders"
GROUP BY status
```

## GROUP AND SUM
---
```ruby
# sum多个字段 返回一个数组
all_data.group('mini_game_results.mini_game_id').pluck('mini_game_id, sum(mini_game_results.spent), sum(mini_game_results.repay)')
```

## Range Conditions
---
```ruby
Client.where(created_at: (Time.now.midnight - 1.day)..Time.now.midnight)
```


## Query Method(preload, includes, left_outer_joins, eager_load)
---
```ruby
# preload
# select "users".* FROM "users"
# select from articles where user_id = [] 
User.preload(:articles)

# includes
# select "users".* FROM "users"
# select from articles where user_id = []
User.includes(:articles)
User.includes(articles: [:comments])
User.includes(:articles, :comments)
User.includes(:articles, :comments).where(comments: {user_id: User-ID})

# eager load
# select users.*, articles.* from users left out join articles on user.id = articles.id
User.eager_load(:articles)
User.eager_load(articles: [:comments])

# includes reference
# includes + references 的效果类似于 eager_load， 但是他比 eager_load 更灵活
# select users.*, articles.* from users left out join articles on user.id = articles.id
User.includes(:articles).references(:articles)

# joins
# select comments.* from comments inner join users on users.id = comments.id inner join articles on articles.id = comments.article_id
Comment.joins(:user, :article)

# left outer joins
Author.left_outer_joins(:posts).distinct.select('authors.*, COUNT(posts.*) AS posts_count').group('authors.id')
```

## EXISTS?
---
```ruby
Client.exists?(1)
Client.exists?(id: [1,2,3])
Client.exists?(name: ['John', 'Sergei'])
```

## EXPLAIN
---
```ruby
User.where(status: 1).explain
```