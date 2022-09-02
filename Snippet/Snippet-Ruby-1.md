## Rails4 reset_pk_sequence!
---
```ruby
module ActiveRecord
  module ConnectionAdapters
    module PostgreSQL
      module SchemaStatements
        # Resets the sequence of a table's primary key to the maximum value.
        def reset_pk_sequence!(table, pk = nil, sequence = nil) #:nodoc:
          unless pk and sequence
            default_pk, default_sequence = pk_and_sequence_for(table)

            pk ||= default_pk
            sequence ||= default_sequence
          end

          if @logger && pk && !sequence
            @logger.warn "#{table} has primary key #{pk} with no default sequence"
          end

          if pk && sequence
            quoted_sequence = quote_table_name(sequence)
            max_pk = select_value("SELECT MAX(#{quote_column_name pk}) FROM #{quote_table_name(table)}")
            if max_pk.nil?
              if postgresql_version >= 100000
                minvalue = select_value("SELECT seqmin FROM pg_sequence WHERE seqrelid = #{quote(quoted_sequence)}::regclass")
              else
                minvalue = select_value("SELECT min_value FROM #{quoted_sequence}")
              end
            end

            select_value <<-end_sql, 'SCHEMA'
              SELECT setval(#{quote(quoted_sequence)}, #{max_pk ? max_pk : minvalue}, #{max_pk ? true : false})
            end_sql
          end
        end
      end
    end
  end
end
```

## ActiveRecord and Pry console connect DB
---
```ruby
require 'active_record'

require 'pry-rails'

ActiveRecord::Base.establish_connection(
  adapter: 'postgresql',
  host: '52.83.173.15',
  port: 5432,
  username: 'pas',
  password: 'eRpoKCB91j4NL2F',
  database: 'pas_staging'
)

class InsuranceBenefitPlanAttachment < ActiveRecord::Base
  self.table_name = 'insurance_benefit_plan_attachment'
end

Pry.start
```

## 自定义路由用户名密码
---
```ruby
with_exception_auth =
  lambda do |app|
    Rack::Builder.new do
      use Rack::Auth::Basic do |username, password|
        ActiveSupport::SecurityUtils.secure_compare(::Digest::SHA256.hexdigest(username), ::Digest::SHA256.hexdigest("user_name")) &
          ActiveSupport::SecurityUtils.secure_compare(::Digest::SHA256.hexdigest(password), ::Digest::SHA256.hexdigest("password"))
      end
      run app
    end
  end

mount with_exception_auth.call(ExceptionTrack::Engine), at: "exception-track"
```

## 数字转换成rmb大写中文
---
```ruby
def number_to_capital_zh(n)
  cNum = ["零","壹","贰","叁","肆","伍","陆","柒","捌","玖","-","-","万","仟","佰","拾","亿","仟","佰","拾","万","仟","佰","拾","元","角","分"]
  cCha = [['零元','零拾','零佰','零仟','零万','零亿','亿万','零零零','零零','零万','零亿','亿万','零元'], [ '元','零','零','零','万','亿','亿','零','零','万','亿','亿','元']]

  i = 0
  sNum = ""
  sTemp = ""
  result = ""
  tmp = ("%.0f" % (n.abs.to_f * 100)).to_i
  return '零' if tmp == 0
  raise '整数部分加二位小数长度不能大于15' if tmp.to_s.size > 15
  sNum = tmp.to_s.rjust(15, ' ')

  for i in 0..14
    stemp = sNum.slice(i, 1)
    if stemp == ' '
      next
    else
      result += cNum[stemp.to_i] + cNum[i + 12];
    end
  end

  for m in 0..12
    result.gsub!(cCha[0][m], cCha[1][m])
  end

  if result.index('零分').nil? # 没有分时, 零角改成零
    result.gsub!('零角','零')
  else
    if result.index('零角').nil? # 有没有分有角时, 后面加"整"
      result += '整'
    else
      result.gsub!('零角', '整')
    end
  end

  result.gsub!('零分', '')
  "#{n < 0 ? "负" : ""}#{result}"
end
```

## 季度的开始和结束
---
```ruby
year = 2020
quarter = 2
start_date = Date.new(year, quarter * 3 - 2).beginning_of_day
end_date = Date.new(year, quarter * 3, -1).end_of_day
```

## 001-999
---
```ruby
NUMBER = 1
"%03d" % (NUMBER+1)

(NUMBER+1).to_s.rjust(3, '0')
```

## CSV文件导出数据
---
```ruby
require 'csv'

file = "/tcv/hospital_contain_chinese.csv"
hospital_translations = Hospital::Translation.where(locale: "en")
headers = ["ID", "HospitalID", "Locale", "Name", "OffcialName", "Address", "Description"]

CSV.open(file, 'w', write_headers: true, headers: headers) do |writer|
  hospital_translations.each do |hs|
    writer << [hs.id, hs.hospital_id, hs.locale, hs.name, hs.official_name, hs.address, hs.description]
  end
end
```

## XLSX转成CSV
---
---
```ruby
xlsx = Roo::Spreadsheet.open('hospital_missing_en.xlsx')
csv = xlsx.to_csv

File.open("hospital_missing_en.csv", "w") do |file|
  file << csv
end
```

## CSV read
---
```ruby
path = Rails.root.join('lib/tasks/data_import/datas/hospital_en_name_staging.csv')

CSV.foreach(path, liberal_parsing: true, headers: :false) do |row|
  id = row["ID"]
  zh_name = row["Chinese Name"]
  en_name = row["English Name"]
end
```
## Ruby Local file server
```ruby
ruby -run -ehttpd . -p8000
```