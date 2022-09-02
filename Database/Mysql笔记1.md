## DELETE与TRUNCATE
---
1. 不带where参数的delete语句可以删除mysql表中所有内容，使用truncate table也可以清空mysql表中所有内容。
2. truncate效率上比delete快，但truncate删除后不记录mysql日志，不可以恢复数据。
delete的效果有点像将mysql表中所有记录一条一条删除到删完，
而truncate相当于保留mysql表的结构，重新创建了这个表，所有的状态都相当于新表。

## 允许远程登录
---
1. vim /etc/mysql/my.cnf找到bind-address = 127.0.0.1
     注释掉这行，如：#bind-address = 127.0.0.1
     或者改为： bind-address = 0.0.0.0
2. 授权用户能进行远程连接
```sql
grant all privileges on *.* to root@"%" identified by "password" with grant option;
flush privileges;
```
第一行命令解释如下:
第一个*代表数据库名；第二个*代表表名。这里的意思是所有数据库里的所有表都授权给用户。
root：授予root账号。 “%”：表示授权的用户IP可以指定，这里代表任意的IP地址都能访问MySQL数据库。
“password”：分配账号对应的密码，这里密码自己替换成 你的mysql root帐号密码。
第二行命令是刷新权限信息，也即是让我们所作的设置马上生效。
*  /etc/init.d/mysqld restart：重启mysql

**开放3306端口号：**
首先采用vi编辑器打开 /etc/sysconfig/iptables，
```shell
# 在-A INPUT -j REJECT --reject-with icmp-host-prohibited
之前加入以下代码：
-A INPUT -m state --state NEW -m tcp -p tcp --dport 3306 -j ACCEPT
```

```shell
# 保存并退出vi编辑器，然后执行以下命令：
service iptables restart
sudo ufw enable 3306
```

## Mysql do not know root password
---
```shell
#/etc/mysql/my.cnf
[mysqld]
skip-grant-tables
```
之后重启mysql，可以无密码进入，设置密码后，删除skip-grant-tables,重启mysql

## 快速use database
---
```shell
mysql> use dbname
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A
```
use database:会预读数据库的表信息，如果数据库有很多表会很慢
use database -A 可以直接切换数据库。
my.cnf中有一个配置为：disable_auto_rehash

## mysql导入导出数据更快的一种方式
---
[Fastest Way To Load Data In MySQL](http://shopperplus.github.io/blog/2014/11/08/fastest-way-to-load-data-in-mysql.html)
```sql
-- When loading a table from a text file, use LOAD DATA INFILE. This is usually 20 times faster than using INSERT statements
LOAD DATA LOCAL INFILE 'x.txt'
REPLACE INTO TABLE product_sale_facts FIELDS TERMINATED BY ',' (`id`,`date_id`,`order_id`,`product_id`,`address_id`,`unit_price`,`purchase_price`,`gross_profit`,`quantity`,`channel_id`,`gift`)
```

在web应用里，经常有导出数据东需求，这种方式速度上会更加快一些。
```sql
# 直接导出到文件
select * from t_apps where created>'2012-07-02 00:00:00' into outfile /tmp/apps.csv
```


## [Use MySQL stream for large datasets](https://ruby-china.org/topics/27829)
---
```ruby
# 添加stream选项
result = client.query('SELECT id, email FROM shopperplus_customers', :stream => true)
```

## 查看数据库大小
```sql
use information_schema;

-- 查看各个数据库大小
SELECT
	concat(round(sum(DATA_LENGTH / 1024 / 1024), 2), 'MB') AS data
FROM
	TABLES;

-- 查看数据库各个表大小
SELECT
	table_schema AS '数据库',
	table_name AS '表名',
	table_rows AS '记录数',
	truncate(data_length / 1024 / 1024, 2) AS '数据容量(MB)',
	truncate(index_length / 1024 / 1024, 2) AS '索引容量(MB)'
FROM
	information_schema.tables
WHERE
	table_schema = 'irm_prod_sec'
ORDER BY
	data_length DESC,
	index_length DESC;
```

## mysqldump 导出 导入
---
如果有报错: `Unknown table 'COLUMN_STATISTICS' in information_schema (1109)` 需要加上参数 `--column-statistics=0`
```shell
mysqldump -h 10.0.101.20
          -u root
          -p irm_prod_sec
          --ignore-table=irm_prod_sec.sug_product_configs_copy2
          --ignore-table=irm_prod_sec.sug_product_configs_copy1 
          --ignore-table=irm_prod_sec.sug_sap_product_config_tmp 
          --ignore-table=irm_prod_sec.sessions 
          --ignore-table=irm_prod_sec.irm_events
          --ignore-table=irm_prod_sec.sug_sap_product_tmp
          --ignore-table=irm_prod_sec.irm_oauth_tokens
          --column-statistics=0 verbose=true > database.sql
          
mysql -u -p
show databases;
use db;
source path/db.sql
```

## Mysql忘记密码
---
```shell
cd /usr/local/mysql/bin

./mysqld_safe --skip-grant-tables
./mysql
FLUSH PRIVILEGES; 
ALTER USER 'root'@'localhost' IDENTIFIED BY '你的新密码';
```

## Mysql从一个很大的表中删除大批量数据
---
```sql
1. create copy table
2. insert into selet(no need to delete data)
3. remove original table
4. rename copy table to original table name
````