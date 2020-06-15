## 关于GraphQL的一些概念
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
```ruby
# Gemfile
gem 'graphql'
```
```shell
rails g graphql:install
```

## Schem
schema包含所有都Types和Fields
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
```ruby
class Types::Enum::RoleCategory < Types::BaseEnum
  graphql_name 'RoleCategory'
  description '角色类型'

  build_from(Role.enum_hash(:category))
end
```

## Arguments
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