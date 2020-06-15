## 求唯一值去重
```sql
create table sex (sex char(1), otherinfo text);    
create index idx_sex_1 on sex(sex);
  
insert into sex select 'm', generate_series(1,50000000)||'this is test';    
insert into sex select 'w', generate_series(1,50000000)||'this is test'; 

-- 正常做法
select distinct sex from sex;

-- 递归做法
-- 相当于每次区一个min(sex) 放入skip中
with recursive skip as (    
  (    
    select min(t.sex) as sex from sex t where t.sex is not null    
  )    
  union all    
  (    
    select (select min(t.sex) as sex from sex t where t.sex > s.sex and t.sex is not null)     
      from skip s where s.sex is not null   
  )    
)     
select * from skip where sex is not null;  
```

## 求差
```sql
create table a(id int primary key, info text);  
create table b(id int primary key, aid int, crt_time timestamp);  
create index b_aid on b(aid);

insert into a select generate_series(1,1000), md5(random()::text);  
insert into b select generate_series(1,5000000), generate_series(1,500), clock_timestamp();  

-- 我实测其实方式2是比递归查询要快的
-- 方式1
select * from a where id not in (select aid from b);   

-- 方式2
select a.id from a left join b on (a.id=b.aid) where b.* is null;  

-- 方式3
select * from a where id not in   
(  
with recursive skip as (    
  (    
    select min(aid) aid from b where aid is not null    
  )    
  union all    
  (    
    select (select min(aid) aid from b where b.aid > s.aid and b.aid is not null)     
      from skip s where s.aid is not null    
  )    
)     
select aid from skip where aid is not null  
);  
```