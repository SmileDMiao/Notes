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

---
## DISTINCT ON:实现从每个分组中取最...的一条数据(最近N天没有创建订单的客户)
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

另外一种思路
```sql
-- LEFT JOIN(这个时间段内有创建订单的客户)
-- WHERE(过滤掉有创建订单的客户剩下的就是没有创建的)
SELECT
	"customers".* 
FROM
	"customers"
	LEFT JOIN (
	SELECT
		customer_id AS filter_customer_id 
	FROM
		"orders" 
	WHERE
		"orders"."company_id" = 8 
		AND "orders"."deleted_at" IS NULL 
		AND ( created_at > '2020-04-09 09:37:48.248227' ) 
	GROUP BY
		1 
	HAVING
		( COUNT ( * ) > 0 ) 
	) AS orphan_order_filter ON orphan_order_filter.filter_customer_id = customers.ID 
WHERE
	"customers"."deleted_at" IS NULL 
	AND "customers"."company_id" = 8 
	AND ( customers.created_at < '2020-04-09 09:37:48.297993' ) 
	AND (orphan_order_filter.filter_customer_id IS NULL)
```
实际比较下来，后一种方式比distinct on略快，cost要小一些


---
## 聚合函数string_agg 和 array_agg json_agg
```sql
CREATE TABLE city(
	country character varying(64),
	city character varying(64)
);

INSERT INTO city VALUES
('中国','台北'),
('中国','香港'),
('中国','上海'),
('日本','东京'),
('日本','大阪');

SELECT country,string_agg(city,',') FROM city GROUP BY country;
SELECT country,array_agg(city) FROM city GROUP BY country;

COALESCE(json_agg(json_build_object('dye_vat', vat_num, 'batch_cloth_id', batch_cloths.id)) FILTER (WHERE dye_vats.id IS NOT NULL), '{}') AS dye_vat_strict
```

---
## filter子句
```sql
select 
  stu_id, 
  count(*), 
  count(*) filter (where score<60) as "不及格"
from sherry.agg_filter
group by stu_id
```

---
## order by null放在后面
```sql
order by result DESC NULLS LAST
```

---
## Sum not repect columns
1. user has many turnovers(流水表)
2. user has many sale_monts(月销售目标表)(1: 100, 2: 200)(只有两条数据)
3. 计算4，5月的总流水和销售目标
4. 利用数据之间的联系, 算出销售目标
```sql
SELECT
	users.id,
	sum(turnovers.amount),
	sum(sale_months.amount) * count(DISTINCT sale_months.id) / count(turnovers.user_id) AS expected
FROM
	users
	LEFT JOIN turnovers ON turnovers.user_id = users.id
	LEFT JOIN sale_months ON sale_months.user_id = users.id
		AND sale_months.month in('2020-4', '2020-5')
	GROUP BY
		users.id
```

---
## Group by time with day
1. turnovers: 流水表 | recieve_at(2020-10-01 12:03:45)
2. 时间是包含时分秒的, 希望以天分组, 计算每天的流水
3. 流水的日期可能不是连续的, 比如6月5号这边压根没有流水, 但是希望输出连续的日期, 以及每天的流水
4. date_trunc函数format日期
```sql
-- '1 day'::interval 每天
-- '1 month'::inteverl 每月
-- '3 month'::inteval 每季度
SELECT
	date_trunc('day', dd)::date,
	COALESCE(pp.money, 0)
FROM
	generate_series('2020-06-01'::timestamp, '2020-06-30'::timestamp, '1 day'::interval) dd
	LEFT JOIN (
		SELECT
			date_trunc('day', turnovers.receive_at::TIMESTAMPTZ AT TIME ZONE '+08:00'::INTERVAL) AS t_time,
			sum(amount) AS money
		FROM
			turnovers
		GROUP BY
			1) pp ON pp.t_time::timestamp = dd::timestamp
```

---
## DATE_TRUNC AND AT TIME ZONE
```sql
SELECT
	accounts_users.id,
	accounts_users.created_at "数据库时间",
	accounts_users.created_at at time zone 'utc' AS "Time Zone 创建时间",
	date_trunc('second', accounts_users.created_at::TIMESTAMPTZ + INTERVAL '8 hour') "DATE TRUNC 创建时间"
FROM
	accounts_users
WHERE
	accounts_users.id = 1;
```