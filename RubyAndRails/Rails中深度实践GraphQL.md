## 关于GraphQL的一些概念
---
1. graphql的使用过程是：描述数据 - 请求数据 - 返回数据（JSON）
2. graphql有两个大佬，一个是schema，一个是query。
3. schema是一个容器，它规定了我们可以查什么，而query则确定我们要查什么。
4. schema是类似database的存在，但是并不是database。它具有权限方面的限制，一个用户对应一个schema。
5. GraphQL queries begin from root types: query, mutation, and subscription.
6. Queries:从API获取特定的数据。将查询设置为只读，就像Rest里面的Get，但是查询不仅是Get。
7. Mutations: 突变，对API数据的修改，比如：create、update、destroy。
8. Types: 类型，用于定义数据类型，或者定义Rails模型model，类型包括根据Queries/Mutations中的请求相应数据的字段和函数。
9. Fields: 字段，表示给定类型的属性
10. Functions: 方法、功能，给上面的字段提供数据

## 开始使用
---
```ruby
# Gemfile
gem 'graphql'
```
```shell
rails g graphql:install
```

## Schem
---
> schema包含所有都Types和Fields
```ruby
class PersonalPracticeSchema < GraphQL::Schema
  mutation(Types::MutationType)
  query(Types::QueryType)

  # Opt in to the new runtime (default in future graphql-ruby versions)
  use GraphQL::Execution::Interpreter
  use GraphQL::Analysis::AST

  # Add built-in connections for pagination
  use GraphQL::Pagination::Connections
end
```

## Declare Types
---
```ruby
# 指定类型, 还有是否为空(不可省略)
module Types
  class ArticleType < Types::BaseObject
    graphql_name 'Article'

    implements GraphQL::Types::Relay::Node
    global_id_field :id

    field :id, ID, null: false
    field :title, String, null: false
    field :body, String, null: false
    field :like_account, Integer, null: false
    field :comment_account, Integer, null: false
    field :user, Types::UserType, '作者', null: false, preload: :user
    field :comments, [Types::CommentType], null: true, preload: :comments
    field :created_at, GraphQL::Types::ISO8601Date, null: true
    field :updated_at, GraphQL::Types::ISO8601Date, null: true
  end
end
```

## N+1问题
---
```ruby
gem 'batch-loader'

BatchLoader::GraphQL.for(object.id).batch(default_value: false) do |items, loader|
  collection = Department.joins(:managers).where(users: { id: items }).group('users.id').count
  collection.each do |k, v|
    loader.call(k, true) if v.positive?
  end
end
```

## Enum类型
---
```ruby
class Types::Enum::RoleCategory < Types::BaseEnum
  graphql_name 'RoleCategory'
  description '角色类型'

  build_from(Role.enum_hash(:category))
end
```

## Arguments
---
```ruby
module Resolvers
  class Roles < Resolvers::BaseResolver
    type Types::RoleType.pagination_type, null: false

    argument :name, String, '名称', as: :name_cont, required: false
    argument :category, Types::Enum::RoleCategory, '类型', required: false, default_value: 'permission'

    def safe_resolve(order: nil, **args)
      q = args.slice(:name_cont)
      roles = Role.ransack(q).result
      roles = roles.where(category: args[:category])

      roles = roles.order(order)
      roles
    end
  end
end
```

## Union
---
```ruby
class Types::UserType < Types::BaseObject
  graphql_name 'User'
  description 'user'

  implements GraphQL::Types::Relay::Node
  global_id_field :id

  field :auditable, Types::Union::Auditable, '审批', null: false, preload: :auditable
end

class Types::Union::Auditable < Types::BaseUnion
  description <<-DESC
    审批类型(订单1 | 订单2| 订单3 | 订单4 | 订单5)
  DESC
  
  # 所有可能到类型
  possible_types Types::PurchaseType, Types::ShipmentPlanType

  def self.resolve_type(object, _context)
    "Types::#{object.class.name}Type".constantize
  end
end
```

```json
// 请求Union类型 ``` on Tyoe{columns}
 auditable{
        ... on PurchaseType{
          id
          sn
        }
        ... on ShipmentPlan{
          id
          sn
          
        }
```

## Global ID
---
```ruby
class Types::UserType < Types::BaseObject
  graphql_name 'User'
  description 'user'

  implements GraphQL::Types::Relay::Node
  global_id_field :id
end

# Return global id
def node_id
  ErpSchema.id_from_object(self, nil, {})
end

# Load object by global id
# 依赖 object_from_id
argument :id, ID, 'User ID', required: true, loads: Types::UserType, as: :user
def id_from_object(object, type_definition, ctx)
  GraphQL::Schema::UniqueWithinType.encode(object.class.name, object.id)
end

def object_from_id(id, ctx)
  class_name, item_id = GraphQL::Schema::UniqueWithinType.decode(id)
  Object.const_get(class_name).find(item_id)
end

# Restore id
argument :id, ID, 'User ID', required: true, restore: true
```

## Restore ID
---
```ruby
argument :id, ID, 'ID', required: true, restore: true
module Types
  class BaseArgument < GraphQL::Schema::Argument
    def initialize(arg_name = nil, type_expr = nil, desc = nil, restore: false, **options)
      super(arg_name, type_expr, desc, **options)
      @prepare = RestoreId if restore
    end

    class RestoreId
      attr_reader :gid, :context

      class << self
        def call(gid, context)
          new(gid, context).call
        end
      end

      def initialize(gid, context)
        @gid = gid
        @context = context
      end

      def call
        if gid.is_a?(Array)
          id = gid.map { |it| restore(it) }
          return id
        end

        restore
      end

      private

      def restore(id = gid)
        return unless id

        GraphQL::Schema::UniqueWithinType.decode(id).last
      end
    end
  end
end
```

## TypeScript 不为空
---
```ruby
def edge_type
  @edge_type ||= begin
                   edge_name = graphql_name + "Edge"
                   node_type_class = self
                   Class.new(edge_type_class) do
                     graphql_name(edge_name)
                     node_type(node_type_class, null: false)
                   end
                 end
end

def connection_type
  @connection_type ||= begin
                         conn_name = graphql_name + "Connection"
                         edge_type_class = edge_type
                         Class.new(connection_type_class) do
                           graphql_name(conn_name)
                           edge_type(edge_type_class, node_nullable: false)
                         end
                       end
end

module GraphQL
  module Types
    module Relay
      class BaseConnection < Types::Relay::BaseObject
        class << self
          def edge_type(edge_type_class, edge_class: GraphQL::Relay::Edge, node_type: edge_type_class.node_type, nodes_field: true, node_nullable: true)
            node_type_name = node_type.graphql_name

            @node_type = node_type
            @edge_type = edge_type_class
            @edge_class = edge_class

            field :edges, [edge_type_class],
                  null: false,
                  description: "A list of edges.",
                  edge_class: edge_class

            define_nodes_field(node_nullable) if nodes_field

            description("The connection type for #{node_type_name}.")
          end
        end
      end
    end
  end
end
```

## 为枚举类型添加说明
---
```ruby
module Types
  module Enums
    class PlayerPositionEnumType < Types::BaseEnum
      description '选手位置'
      # 生成 value(key, description)
      build_from(Player.enum_hash(:position))
    end
  end
end


# typed: strong
class ApplicationRecord < ActiveRecord::Base
  self.abstract_class = true

  class << self
    def cached_enum_hash
      @cached_enum_hash ||= {}.with_indifferent_access
    end

    def enum_hash(attribute)
      # 生成Graphql Enum所需要的中英文数据
      cached_enum_hash[attribute] ||= begin
                                        keys = public_send(attribute.to_s.pluralize).keys
                                        keys.each_with_object({}) do |k, m|
                                          v = human_attribute_name("#{attribute}.#{k}", locale: 'zh')
                                          m[k] = v
                                        end
                                      end
    end
  end
end

module Types
  class BaseEnum < GraphQL::Schema::Enum
    class << self
      def build_from(collection)
        case collection
        when Array
          collection.each { |k| value(k, k) }
        when Hash
          collection.each { |k, v| value(k, v) }
        else
          raise 'Unsupported'
        end
      end

      # 当接口输出Enum的时候可以在那个Type后面调用 object_type 方法以生成文档
      # 前端也可以拿到对应的值与说明
      def object_type
        @object_type ||= begin
                           enum = self
                           name = "#{graphql_name}WithLocale"

                           Class.new(::Types::BaseObject) do
                             graphql_name(name)

                             field :value, enum, '枚举值', null: true

                             def value
                               object
                             end

                             field :description, String, '描述', null: true

                             def description
                               enum = self.class.fields['value'].type
                               enum.values[value].description
                             end
                           end
                         end
      end
    end
  end
end
```

## 集成Ransack
---
```ruby
module RansackSupport
  extend ActiveSupport::Concern

  class ArgumentConfig
    attr_reader :items

    def initialize
      @items = {}.with_indifferent_access
    end

    def add(name, des, options)
      exp ||= options[:exp] || name
      @items[name] = { exp: exp, des: des }
    end
  end

  class_methods do
    def ransack(name = :q, description = 'RANSACK FILTER')
      config = ::RansackSupport::ArgumentConfig.new
      yield config

      type_name = graphql_name.to_s.camelize
      prefix = "#{type_name}#{name.to_s.camelize}"

      input = build_ransack_input(prefix, config)
      build_ransack_argument(name, input, description)
    end

    def build_ransack_input(prefix, config)
      enum_name = "#{prefix}Enum"
      input_name = "#{prefix}Input"

      enum = Class.new(::Types::BaseEnum) do
        graphql_name enum_name
        config.items.map do |name, options|
          value(name, options[:des], value: options[:exp])
        end
      end

      Class.new(::Types::BaseInputObject) do
        graphql_name input_name
        argument :name, enum, description: '参数名', required: true
        argument :value, GraphQL::Types::JSON, description: '参数值', required: false
      end
    end

    def build_ransack_argument(name, input, description, object = self)
      object.argument name, [input], description, required: false, prepare: ->(args, _ctx) do
        args&.each_with_object({}.with_indifferent_access) do |it, m|
          k = it.name

          v = it.value
          m[k] = v
        end
      end
    end
  end
end
```

## Graphql Pagination Connection
---
使用方式
```ruby
# api_schema.rb
use GraphQL::Pagination::Connections

# graphql type
# 默认方式
field :users, Types::UserType.connection_type, null: false
```

请求示例
```json
query employees($rql: JSON) {
  employees(first: 5, after: "WyIyMDIwLTA5LTI4IDA4OjEwOjM2Ljg0NTMyNTAwMCBVVEMiLDE3XQ") {
    nodes {
      id
      name
      nickName
    }
    pageInfo{
      endCursor
      startCursor
    }
  }
}
```

自己实现类似游标分页的效果
```ruby
field :featured_comments, CommentType.connection_type do
  # Add an argument:
  argument :since, String, required: false
end

def featured_comments(since: nil)
  comments = object.comments.featured
  if since
    comments = comments.where("created_at >= ?", since)
  end
  comments
end
```

## Custom Relay Connection
---
可以添加额外想要返回的数据, 比如 total_count 之类的, 也可以是其他针对返回列表的统计数据
```ruby
class Types::PostEdgeType < GraphQL::Types::Relay::BaseEdge
  node_type(PostType)
end

class Types::PostConnectionWithTotalCountType < GraphQL::Types::Relay::BaseConnection
  edge_type(PostEdgeType)

  field :total_count, Integer, null: false
  def total_count
    object.items.size
  end
end
```

## GraphQL PRO Pagination
---
使用Pro版本之后可以实现真正的 cursor 分页
```ruby
use GraphQL::Pagination::Connections
connections.add(ActiveRecord::Relation, GraphQL::Pro::PostgresStableRelationConnection)
```

## 集成分页
---
Gem: graphql-pagination
```ruby
type Stafftools::Types::Games::OptionType.collection_type, null: true
```

## Graphql错误处理
---
可以在 schema 中 resuce 一些常见的异常
在公用的resolver文件中定义add_error, 在 Graphql 默认的 error 中返回错误
```ruby
# api_schema.rb
def self.unauthorized_object(error)
  raise GraphQL::ExecutionError.new('permission denied', extensions: { code: 'Unauthorized', status: 401 })
end

rescue_from(ActiveRecord::RecordInvalid, ActiveRecord::RecordNotUnique) do |err, _obj, _args, _ctx, _field|
  raise GraphQL::ExecutionError.new(err.message, extensions: { code: 'Unsupported Media Type', status: 415 })
end

rescue_from(ActiveRecord::RecordNotFound) do |err, _obj, _args, _ctx, _field|
  raise GraphQL::ExecutionError.new(err.message, extensions: { code: 'Not Found', status: 404 })
end

# resolver support
def add_error(message, ext = {})
  error = GraphQL::ExecutionError.new(message, extensions: ext)
  context.add_error(error)
end

def check_permission
  unless current_employee.admin
    raise GraphQL::ExecutionError.new('permission denied', extensions: { code: 'FORBIDDEN', status: 403 })
  end
end
```