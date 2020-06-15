## TSVECTOR AND TSQUERY
> to_tsvector() 把一个字符串转换为tsvector一个tsvector是一个标准词位的有序列表(sorted list), 标准词位(distinct lexeme)就是说把同一单词的各种变型体都被标准化相同的。数字表示词位在原始字符串中的位置，比如“man"出现在第6和15的位置上.
标准化过程几乎总是把大写字母换成小写的, 也经常移除后缀(比如英语中的s,es和ing等) 这样可以搜索同一个字的各种变体,而不是乏味地输入所有可能的变体。
```sql
-- string content
SELECT
	orders.id,
	orders.sn || ' ' || orders.currency || ' ' || users.name || ' ' || coalesce((string_agg(order_items.unit, ' ')), '') AS document
FROM
	orders
	LEFT JOIN users ON users.id = orders.creator_id
	LEFT JOIN order_items ON order_items.order_id = orders.id
GROUP BY
	orders.id,
	users.id
	
-- TO_TSVECTOR转换格式
SELECT
	to_tsvector(orders.sn) || ' ' || to_tsvector(orders.currency) || ' ' || to_tsvector(users.name) || ' ' || to_tsvector(coalesce((string_agg(order_items.unit, ' ')), '')) AS document
FROM
	orders
	LEFT JOIN users ON users.id = orders.creator_id
	LEFT JOIN order_items ON order_items.order_id = orders.id
GROUP BY
	orders.id,
	users.id

-- TO_TSQUERY
-- tsquery存储了要搜索的词位，可以使用&（与）、|（或）和!（非）逻辑操作符。可以使用圆括号给操作符分组。
SELECT to_tsvector('It''s kind of fun to do the impossible') @@ to_tsquery('impossible');
SELECT to_tsvector('If the facts don't fit the theory, change the facts') @@ to_tsquery('! fact');
SELECT to_tsvector('If the facts don''t fit the theory, change the facts') @@ to_tsquery('theory & !fact');
SELECT to_tsvector('If the facts don''t fit the theory, change the facts.') @@ to_tsquery('fiction | theory');
SELECT to_tsvector('If the facts don''t fit the theory, change the facts.') @@ to_tsquery('theo:*');

-- 查询文档
SELECT
oid, osn FROM(
  SELECT
    orders.id as oid, orders.sn as osn
  	to_tsvector(orders.sn) || ' ' || to_tsvector(orders.currency) || ' ' || to_tsvector(users.name) || ' ' || to_tsvector(coalesce((string_agg(order_items.unit, ' ')), '')) AS document
  FROM
  	orders
  	LEFT JOIN users ON users.id = orders.creator_id
  	LEFT JOIN order_items ON order_items.order_id = orders.id
  GROUP BY
  	orders.id,
  	users.id
) p_search
WHERE p_search.document @@ to_tsquery('123456 & miao')
```

## SETWEIGHT
```sql
-- SETWEIGHT设置权重
SELECT
pid, p_title FROM (
  SELECT post.id as pid,
    post.title as p_title,
    setweight(to_tsvector(post.language::regconfig, post.title), 'A') || 
    setweight(to_tsvector(post.language::regconfig, post.content), 'B') ||
    setweight(to_tsvector('simple', author.name), 'C') ||
    setweight(to_tsvector('simple', coalesce(string_agg(tag.name, ' '))), 'B') as document  FROM post  JOIN author ON author.id = post.author_id  JOIN posts_tags ON posts_tags.post_id = posts_tags.tag_id  JOIN tag ON tag.id = posts_tags.tag_id  GROUP BY post.id, author.id) p_search
WHERE p_search.document @@ to_tsquery('english', 'Endangered & Species')ORDER BY ts_rank(p_search.document, to_tsquery('english', 'Endangered & Species')) DESC;
```

## TS_RANK
```sql
-- 相关性
SELECT ts_rank(to_tsvector('This is an example of document'), to_tsquery('example | document')) as relevancy;
```

## 索引
```sql
CREATE INDEX index_article ON post 
USING gin(setweight(to_tsvector(language, title),'A') || setweight(to_tsvector(language, content), 'B'))
```

## 中文分词
```shell
brew install scws
git clone https://github.com/amutu/zhparser.git
SCWS_HOME=/usr/local make && make install
```
```sql
-- 验证
select * from pg_available_extensions where name = 'zhparser';
-- 添加配置
CREATE TEXT SEARCH CONFIGURATION parser_name (PARSER = zhparser);
-- 设置分词规则
ALTER TEXT SEARCH CONFIGURATION parser_name ADD MAPPING FOR n,v,a,i,e,l,j WITH simple; 
```

## Triggers For Automatic Updates
tsvector_update_trigger
tsvector_update_trigger_column