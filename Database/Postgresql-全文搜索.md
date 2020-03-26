## tsvector and tsquery
tsvector类型表示一个为文本搜索优化的形式下的文档。
```sql
SELECT 'a fat cat sat on a mat and ate a fat rat'::tsvector;
-- 含空白或标点的词位
SELECT $$the lexeme '    ' contains spaces$$::tsvector;

SELECT to_tsvector('english', 'The Fat Rats');
```

tsquery类型表示一个文本查询。
```sql
-- 一个tsquery中的词位可以被标注为*来指定前缀匹配
SELECT 'super:*'::tsquery;
-- 一个tsquery中的词位可以被标注一个或多个权重字母，这将限制它们只能和具有那些权重之一的tsvector词位相匹配
SELECT 'fat:ab & cat'::tsquery;

select * from test where to_tsvector(‘english’,name) @@ to_tsquery(‘english’,‘1_tans’)
```

## zhparser
zhparser是pg数据库的中文全文搜索扩展,基于SCWS
```sql
-- 创建新的文本搜索配置
CREATE TEXT SEARCH CONFIGURATION

SELECT * FROM ts_parse('zhparser', '保障房资金压力');
```

## RUM
[RUM](https://github.com/postgrespro/rum)