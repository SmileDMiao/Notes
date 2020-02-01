## 安装
我之前java版本低老是有问题，换到java8就好了，总之按照官方文档来就行了，基本不会出错。
启动：在elasticsearch住目录中 bin/elasticsearch
启动多个就是多个节点，port自9200开始递增
node从0开始递增
以守护进程运行: bin/elasticsearch -d

## 概念理解
在Elasticsearch中存储数据的行为就叫做索引
索引（名词） 如上文所述，一个索引(index)就像是传统关系数据库中的数据库，它是相关文档存储的地方，index的复数是indices 或indexes。
索引（动词） 「索引一个文档」表示把一个文档存储到索引（名词）里，以便它可以被检索或者查询。这很像SQL中的INSERT关键字，差别是，如果文档已经存在，新的文档将覆盖旧的文档
倒排索引 传统数据库为特定列增加一个索引，例如B-Tree索引来加速检索。Elasticsearch和Lucene使用一种叫做倒排索引(inverted index)的数据结构来达到相同目的。
shareds: Elasticsearch提供了将索引细分为多个称为碎片的片段的功能。 创建索引时，可以简单地定义所需的分片数。 每个分片本身就是一个功能完整且独立的“索引”，可以在集群中的任何节点上托管。
replicas:Elasticsearch允许您将索引的碎片的一个或多个副本复制到所谓的复制分片，或简写为replicas

对比
```
Relational DB -> Databases -> Tables -> Rows -> Columns
Elasticsearch -> Indices   -> Types  -> Documents -> Fields
```


## elasticsearch-rails
之前使用过PG数据库的全文搜索功能，但本身不支持中文分词，性能上也不是很好，在使用过程中还遇到一些坑，这里使用更为专业的Elasticsearch来做全文搜索。

### Gemfile
```ruby
gem 'elasticsearch-model'
gem 'elasticsearch-rails'
```

### Index设置
```ruby
# 先创建自己需要的index
Article.__elasticsearch__.create_index!(index: 'personal')
class Article
  # set index name and type name
  index_name    'personal'
  document_type 'articles'

  # set data schema
  mapping do
    indexes :title, term_vector: :yes
    indexes :body, term_vector: :yes
  end

  # use article’s title and body
  def as_indexed_json(options={})
    as_json(only: ['title','body'])
  end
end
```

### 导入数据
```ruby
#导入article数据到elasticsearch中(specific index and type)
Article.import(index: 'personal', type: 'articles')
```

### 回调
Gem文档提供了两种回调方式，一种是自动回调，通过include Elasticsearch::Model::Callbacks，还有一种是手动的管理回调，这里我使用手动管理，使用sidekiq异步更新Index
使用Concern的方式以便多model使用
```ruby
module Searchable
  extend ActiveSupport::Concern
  included do
    include Elasticsearch::Model
    # 创建 obj.__elasticsearch__.index_document
    after_commit on: [:create] do
      SearchIndexerJob.perform_later('index', self.class.name, self.id)
    end
    # 更新 obj.__elasticsearch__.update_document
    after_commit on: [:update] do
      SearchIndexerJob.perform_later('update', self.class.name, self.id)
    end
    # 删除 obj.__elasticsearch__.delete_document
    after_commit on: [:destroy] do
      SearchIndexerJob.perform_later('delete', self.class.name, self.id)
    end
  end
end
```

## 搜索
elasticsearch的搜索功能强大，api数量参数同样很多，elasticsearch-rails这个gem提供了单model搜索和
多个model一个搜索的功能，搜索的参数其实两种都差不多，使用默认的很简单，想要复杂的就需要去看elasticsearch的query dsl了。
这里只记录说明使用到的搜索，并使用多表搜索的方式搜索单表两种都好借鉴。

```ruby
# simple_query_string:A query that uses the SimpleQueryParser to parse its context.
# Unlike the regular query_string query, the simple_query_string query will never throw an exception
# query: content of query
# default_operator: The default operator used if no explicit operator is specified.
# minimum_should_match: The minimum number of clauses that must match for a document to be returned
# fields: the column to be searched
# highlight: the match ccontent display highlight
search_params = {
    query: {
        simple_query_string: {
            query: '撒旦 快乐',
            default_operator: 'AND',
            minimum_should_match: '70%',
            fields: %w(title body)
        }
    },
    highlight: {
        pre_tags: ['[em]'],
        post_tags: ['[/em]'],
        fields: {title: {}, body: {}}
    }
}
# 如果是多model搜索，在搜索的model中设置好index，type，在fields中添加搜索的字段名，在下面添加模型名称
@result = Elasticsearch::Model.search(search_params, [Article]).page(params[:page])
```