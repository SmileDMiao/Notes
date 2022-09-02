# InfluxDB:时序数据库
---
## 安装
```shell
brew install influxdb
$ influx
Connected to http://localhost:8086 version 0.9
InfluxDB shell 0.9
```

## 命令
```sql
-- 创建数据库
CREATE DATABASE db_name
-- 查看所有数据库
SHOW DATABASES
-- 切换使用某个db
USE db_name
-- 显示所有的数据库
show databases
-- 删除数据库
drop database db_name
-- 显示该数据库中所有的表
show measurements
-- 查看series
show series from measurement_name
-- 查看某个measurement的retention policy
show retention policies on db_name
-- 删除表
drop measurement measurement_name
-- 插入数据
INSERT cpu,host=serverA,region=us_west value=0.64
-- 查询
SELECT * FROM /.*/ LIMIT 1
SELECT * FROM "cpu_load_short"
SELECT * FROM "cpu_load_short" WHERE "value" > 0.9
```

## 配置文件
1. max-values-per-tag = 100000
之前我有遇到数据写入不进的问题，和这个配置有关
The maximum number of tag values allowed per tag key. The default setting is 100000. Change the setting to 0 to allow an unlimited number of tag values per tag key. 

2. max-series-per-database = 1000000
The maximum number of series allowed per database. The default setting is one million. Change the setting to 0 to allow an unlimited number of series per database.

## 基本概念理解
1. database
数据库可以有多个用户、连续查询、保留策略和测量。InfluxDB是一个无模式数据库，它意味着随时都可以添加新的测量、标记和字段。

2. field key
Field keys (butterflies and honeybees) are strings and they store metadata; the field key butterflies tells us that the field values 12-7 refer to butterflies and the field key honeybees tells us that the field values 23-22 refer to, well, honeybees.【官网是这么说的，我理解就是数据库中的字段名】

3. field value
字段值是数据；它们可以是字符串、浮点数、整数或布尔值，并且因为InfluxDB是一个时间序列数据库，所以字段值总是与时间戳相关联。

4. field set
the collection of field-key and field-value pairs make up a field set

5. measurement
我的理解是相当于数据库里的一张表

6. point
a point is the field set in the same series with the same timestamp.【我理解就是表里面的一行数据】

7. retention policy
【就是数据的存储策略】A retention policy describes how long InfluxDB keeps data (DURATION) and how many copies of those data are stored in the cluster (REPLICATION)

8. series
a series is the collection of data that share a retention policy, measurement, and tag set.

9. tag
【相当于数据库中的索引属性】tags are made up of tag keys and tag values. Both tag keys and tag values are stored as strings and record metadata

# Telegraf:数据收集
---
## 安装
```shell
brew install telegraf
brew services start telegraf
```

## 配置
```shell
# 创建一个默认配置的配置文件
telegraf config > telegraf.conf
# 创建一个指定输入输出插件的配置文件
# create a config file with specific input and output plugins
telegraf --input-filter cpu:mem:net:swap --output-filter influxdb:kafka config > telegraf.conf
```

```shell
# 细看配置文件中的几个部分
#
# 全局tag
[global_tags]

# telegraf本身的配置，配置比如收集数据的间隔，发送数据的批次数量之类的
[agent]

# 数据存储到哪里，到时候根据自己的需要细看配置
OUTPUT PLUGINS
[[outputs.influxdb]]

# 收集哪些数据，到时候根据自己的需要细看配置
INPUT PLUGINS
[[inputs.cpu]]

# https://docs.influxdata.com/telegraf/v1.7/concepts/aggregator_processor_plugins/
# 处理插件，转换、处理、过滤数据
PROCESSOR PLUGINS
[[processors.converter]]
# 聚合插件，数据特征聚合
AGGREGATOR PLUGINS
[[aggregators.basicstats]]
```

# Chronograf:图表展示
---
## 安装
```shell
brew install chronograf
```

# Kapacitor:事件警告
---
## 安装
```shell
brew install kapacitor
```

## 配置
**example config**
```shell
kapacitord config
kapacitord config > kapacitor.generated.conf
```