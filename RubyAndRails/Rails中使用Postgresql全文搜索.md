## 利用pg数据库做全文搜索
[PostgreSQL 的全文检索系统之中文支持](https://www.rails365.net/articles/postgresql-de-quan-wen-jian-suo-xi-tong-zhi-zhong-wen-zhi-chi-san)

由于pg自带的全文搜索不支持中文分词，所以需要做一些准备工作：
1. 安装scws
2. 安装zhparser
3. 安装pg扩展
4. 添加文本搜索配置

照着参考资料做基本上是可以完成的，
自己遇到的一点问题就在 文本搜索配置上，自定义的文本搜索配置是关联到数据库上的，在进入pg数据库命令行的之后，必须要先切换数据库，之后再创建文本搜索配置。
```sql
\c database_name :连接到哪个数据库
\dF :查看所有的文本搜索配置.
\dFp :列出所有的文本搜索分词器
```

* 新建的文本搜索配置名称在后面使用pg_search的时候回用到。

* 上面要做的完成之后就可以利用pg_search来完成全文搜索的实现了。

## pg_search 两种方式-单表搜索
pg_search的github上的文档介绍的还是很详细的，仔细看完照着测试基本没有问题，这里只是记录一个基本用法,以及一些参数的含义。
更详细用法[参见](https://github.com/Casecommons/pg_search)
```ruby
include PgSearch
pg_search_scope :chinese_search,                  #搜索方法名称,使用时model调用
                :against => [:title, :body],      #搜索的字段
#                :against => {
#                    :title => 'A',               #使用权重,title的优先级比body要高
#                    :body => 'B'
#                  },
                :associated_against => {
                    :comments => [:body]          #关联搜索,article has many comments
                  },
                :using => {
                    tsearch: {
                        dictionary: 'zhcnsearch', #使用哪种粉本搜索配置(之前创建的)
                        :prefix => true           #是否使用前缀
                    }
                }
```

## pg_search 两种方式-多表搜索
创建文档表
```ruby
$ rails g pg_search:migration:multisearch
$ rake db:migrate
```
多表搜索的options（这里写到了config/initializer文件夹下）
```ruby
PgSearch.multisearch_options = {
    :using => {
        tsearch: {
            dictionary: 'zhcnsearch',   #使用哪种粉本搜索配置(之前创建的)
            :prefix => true,            #是否使用前缀
            :highlight => {             #关键词高亮，在匹配内容的关键词的首尾加上以下标签
                :start_sel => '<em>',   #可以用css控制显示颜色，或者使用p标签加粗
                :stop_sel => '</em>'
            }
        }
    }
}
```
model中添加搜索内容
```
include PgSearch
multisearchable :against => [:title, :body]
```

多表搜索时，其实搜索的还是一个表，不过在model中配置了多表搜索，每当新建更新，删除都会有一个回调来同步到文档表中。

搜索方法示例：
controller:
```ruby
@result = PgSearch.multisearch(params[:search]).with_pg_search_highlight
```
view:
```ruby
@result.each do |item|
  <%= raw item.pg_search_highlight %>
end
```