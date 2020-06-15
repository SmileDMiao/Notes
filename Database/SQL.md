## COALESCE空值处理.
```sql
-- 计算用户的客户数; 用户的客户规则设置数量; 查出客户数>规则数的用户
-- customer_rules.rules(JSON)
SELECT
	"users".id,
	users.name,
	count(customers.id) AS count,
	COALESCE(customer_rules.rules->>'limit',20::text)
FROM
	"users"
	LEFT OUTER JOIN "user_associations" ON "user_associations"."user_id" = "users"."id"
	AND "user_associations"."name" = 'visible_users'
	AND "user_associations"."record_type" = 'Customer'
	LEFT OUTER JOIN "customers" ON "customers"."company_id" = '22b54c59-3b95-4864-9292-4ab433d51182'
	AND "customers"."id" = "user_associations"."record_id"
	LEFT OUTER JOIN customer_rules ON users.id::text = ANY (customer_rules.user_ids)
	AND customer_rules.rule_type = 0
WHERE
	"users"."company_id" = '22b54c59-3b95-4864-9292-4ab433d51182'
GROUP BY
	users.id,
	customer_rules.id
HAVING count(customers.id)::INTEGER <= COALESCE((customer_rules.rules->>'limit')::integer,20)
```

## DISTINCT ON:实现从每个分组中取最...的一条数据
```sql
-- 查出最近多长时间内(没有orders或者samplings或者quotations)的客户
-- DISTINCT ON: 查出用户的最新一条单据
-- 查出用户最新的一条orders/samplings/quotations与customer关联, 创建时间小于设置时间或为空即期望结果
SELECT
	"customers".*
FROM
	"customers"
	LEFT OUTER JOIN ( SELECT DISTINCT ON (customer_id)
		*
	FROM
		orders
	ORDER BY
		customer_id,
		created_at DESC) o ON o.customer_id = customers.id
	LEFT OUTER JOIN ( SELECT DISTINCT ON (customer_id)
		*
	FROM
		samplings
	ORDER BY
		customer_id,
		created_at DESC) s ON s.customer_id = customers.id
	LEFT OUTER JOIN ( SELECT DISTINCT ON (customer_id)
		*
	FROM
		quotations
	ORDER BY
		customer_id,
		created_at DESC) q ON q.customer_id = customers.id
WHERE
	"customers"."company_id" = '22b54c59-3b95-4864-9292-4ab433d51182'
	AND(o.created_at < '2020-05-25 10:58:01 +0800'
		OR o.created_at IS NULL
		OR s.created_at < '2020-05-25 10:58:01 +0800'
		OR s.created_at IS NULL
		OR q.created_at < '2020-05-25 10:58:01 +0800'
		OR s.created_at IS NULL)
```