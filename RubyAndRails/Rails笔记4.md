## 字符串转时间:
---
20120514144424 转成 2012-05-14 14:44:24
```ruby
DateTime.parse('20120514144424').strftime('%Y-%m-%d %H:%M:%S')
=> "2012-05-14 14:44:24"
"20120514144424".to_time.strftime('%Y-%m-%d %H:%M:%S')
=> "2012-05-14 14:44:24"
"20120514144424".to_datetime.strftime('%Y-%m-%d %H:%M:%S')
=> "2012-05-14 14:44:24"
"20120514144424".to_date
=> Mon, 14 May 2012
# Timestamp类型转换
DateTime.strptime('253402185600000','%Q').to_time

# year: 年 quarter: 季度
# 哪一年哪个季度的开始和结束日期
start_date = Date.new(year, quarter * 3 - 2)
end_date = Date.new(year, quarter * 3, -1)
```

---
## 时间对象的显示转换
```ruby
time = Time.now
=> 2016-07-18 02:58:22 +0800
time.to_formatted_s(:time)
=> "02:58"
time.to_formatted_s(:db)
=> "2016-07-18 02:58:22"
time.to_formatted_s(:number)
=> "20160718025822"
time.to_formatted_s(:short)
=> "18 Jul 02:58"
time.to_formatted_s(:long)
=> "July 18, 2016 02:58"
time.to_formatted_s(:long_ordinal)
=> "July 18th, 2016 02:58"
time.to_formatted_s(:rfc822) # GMT
=> "Mon, 18 Jul 2016 02:58:22 +0800"
```

## Gem: counter_culture
```sql
# 这种写法会导致 rake db:create 报错
counter_culture :user, column_name: 'spent_amount', delta_column: 'spent', column_names: {
      Prediction::Order.checked => 'spent_amount',
    }

# 这种写法 rake db:create 不会报错
counter_culture :user, column_name: 'spent_amount', delta_column: 'spent', column_names: {
      ["prediction_orders.status = ?", 2] => 'spent_amount',
    }
```

## [SQL注入](https://rails-sqli.org/)
---
```ruby
attr = a = "1) OR 1=1--"
USer.delete_all("id= #{attr}")

## Rails DEBUG
---
```ruby
# 开启SQL LOG
ActiveRecord::Base.logger = Logger.new(STDOUT)
# 关闭SQL LOG
ActiveRecord::Base.logger = nil
# 忽略statement timeout
ActiveRecord::Base.connection.execute('SET statement_timeout TO 0')
```

## 导出的CSV文件用EXCEL打开乱码
---
生成csv文件的方式
```ruby
File.open(@file_path, 'w', encoding: Encoding::ASCII_8BIT) do |f|
  charges.copy_to do |line|
    f.write line
  end
end
```
用这种方式生成普通文本编辑器打开是没有乱码的, 或是其他的软件numbers/LibreOfice, 但是Excel打开就是有乱码
需要在文件头部添加对应到BOM(Byte Order Mark)
是Unicode(编码方案包括UTF-8,UTF-16,UTF-32等)的字节顺序标记，它有三个作用：
1. 说明字节序：big-endian和little-endian两种，UTF-16和UTF-32都有这两种字节序(utf-16be,utf-16le,utf-32be,utf-32le)。
2. 说明字符流属于Unicode编码
3. 说明字符流是哪一种Unicode编码方式

| 编码方式 | BOM |
| :----: | :----: |
| UTF-8 | EF BB BF |
| UTF-16(BE) | FE FF |
| UTF-16(LE) | FF FE  |
| UTF-32(BE) | 00 00 FE FF |
| UTF-32(LE) | FF FE 00 00 |

添加BOM
```ruby
File.open(@file_path, 'w', encoding: Encoding::ASCII_8BIT) do |f|
  f.write("\xEF\xBB\xBF")
  charges.copy_to do |line|
    f.write line
  end
end
```