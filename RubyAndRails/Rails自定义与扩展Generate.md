## Thor Rails generate生成器之自定义生成器
---
>需求1: rails项目主目录下有个文件夹叫 *modules*， 里面有好多文件夹，每个文件夹的结构个rails项目很像，也有app(controller,views,model),config,lib..,
需要写一个 *generate生成器* 来生成modules下的这么一个文件夹结构，比如 *rails generate module=cmi*
参考：http://guides.ruby-china.org/generators.html

### 生成生成器
---
```shell
# rails generate 提供的生成 生成器 的命令
rails generate generator module
=>
# 生成器文件夹
create  lib/generators/module
# 生成器主要代码文件
create  lib/generators/module/module_generator.rb
# 用途说明文件
create  lib/generators/module/USAGE
# 模板文件位置
create  lib/generators/module/templates
```

### 自定义生成器
---
```ruby
class ModuleGenerator < Rails::Generators::NamedBase
  source_root File.expand_path('../templates', __FILE__)

  # module_name :01_framework_fwk
  # module_real_name :fwk
  # module_path :modules/01_framework_fwk
  attr_reader :module_name, :module_real_name, :module_path

  def initialize(*args)
    super
    @module_name = file_name.underscore
    @module_real_name = module_name.split("_").last
    @module_path = "modules/#{module_name}"
  end

  def create_module_folders
    #　生成文件夹
    empty_directory "#{module_path}"
    empty_directory "#{module_path}/app"
    empty_directory "#{module_path}/app/controllers"
    empty_directory "#{module_path}/app/helpers"
    empty_directory "#{module_path}/app/models"
    empty_directory "#{module_path}/app/views"
    # 省略一部分
    # 复制文件
    template 'routes.rb',   "#{module_path}/config/routes.rb"
    template 'en.yml',      "#{module_path}/config/locales/en.yml"
    template 'zh.yml',      "#{module_path}/config/locales/zh.yml"
  end
end
```

## Thor Rails generate生成器之扩展原生生成器
---
>需求2: example: rails generate controller xxx 会在app/controller下生成文件，现在需要给这个系列命令添加一个参数 --module=xxx,
添加这个参数之后就会在modules文件夹下的指定文件夹添加controller文件。

### generate command extend
---
>这样rails generate 就支持 --module选项了.

```ruby
# lib/rails/generators/named_base.rb
# 添加选项
class_option :module, :type => :string, :default => '', :desc => "choose which module to place the files"

# 替换方法
no_tasks do
  def template(source, *args, &block)
    # TODO don't know why have this judge temporary
    unless source.eql?('migration.rb')
      args.each_with_index do |arg,index|
        args[index] = "#{get_module}#{arg}"
      end
    end
    inside_template do
      super
    end
  end
end

# 添加方法
def get_module
  if !options[:module].present?
    ''
  else
    if Rails.application.config.fwk.module_mapping[options[:module]]
      "modules/#{Rails.application.config.fwk.module_mapping[options[:module]]}/"
    else
      ''
    end
  end
end
```

### controller command expand
---
> rails g controller 将在指定的module文件夹下的对应位置生成文件

```ruby
# lib/rails/generators/erb/controller/controller_generator.rb
require 'rails/generators/erb'

module Erb # :nodoc:
  module Generators # :nodoc:
    class ControllerGenerator < Base # :nodoc:
      argument :actions, type: :array, default: [], banner: "action action"

      def copy_view_files
        # 替换文件生成位置
        base_path = File.join("#{get_module}app/views", class_path, file_name)
        empty_directory base_path

        actions.each do |action|
          @action = action
          formats.each do |format|
            @path = File.join(base_path, filename_with_extensions(action, format))
            template filename_with_extensions(:view, format), @path
          end
        end
      end
    end
  end
end
```

### migration command expand
---
> rails g migration 文件位置发生相同模式的改变

```ruby
# lib/generators/active_record/migration/migration_generator.rb
# 重写方法
def create_migration_file
  set_local_assigns!
  validate_file_name!
  #　替换路径
  migration_template "migration.rb", "#{get_module}db/migrate/#{file_name}.rb"
end
```

### model command expand
---
> rails g model model和migration的文件位置发生相同模式的改变

```ruby
# lib/rails/generators/active_record/model/model_generator.rb
# 重写方法
def create_migration_file
  return unless options[:migration] && options[:parent].nil?
  attributes.each { |a| a.attr_options.delete(:index) if a.reference? && !a.has_index? } if options[:indexes] == false
  # 替换路径
  migration_template "../../migration/templates/create_table_migration.rb", "#{get_module}db/migrate/#{file_name}.rb"
end
```

## 扩展rails 脚手架
---
>需求3: example: rails g scaffold HighScore name:string --module=irm，执行之后，按照上面的规则，将对应的文件添加到modules对应文件夹下，并添加一些必要的模板文件，简而言之就是
让rails g scaffold 这个命令支持--module选项

### 让 scaffold 支持 --module 参数
```ruby
# lib/rails/generators/erb/scaffold/scaffold_generator.rb
# 添加对 --module 参数的支持
def create_root_folder
  empty_directory File.join("#{get_module}app/views", controller_file_path)
end

protected

# 确定模板页面
 def available_views
   %w(index edit show new get_data)
 end
```

### 添加模板文件
---
```ruby
# lib/templates/erb/scaffold/new.html.erb
# 这里可以参考一下源码:lib/templates/rails/scaffold_controller/controller.rb
class <%= controller_class_name %>Controller < <%= controller_class_name.include?('::') == true ? "#{controller_class_name.split('::').first}::" : ''  %>ApplicationController
before_action :set_<%= file_name %>, only: [:show, :edit, :update, :destroy]

  def index
    @<%= plural_file_name %> = <%= file_name.camelize %>.desc('_id')
    @<%= plural_file_name %> = @<%= plural_file_name %>.paginate(page: params[:page], per_page: 30)
  end

  private

  def set_<%= file_name %>
@<%= file_name %> = <%= orm_class.find(file_name.camelize, "params[:id]") %>
  end
end
```

## 总结
---
1. 这里的方式是完全替换原来相应的generator
2. 另一种方式: 通过打开类的方式，send(:include),也可以实现功能，不过这种方式代码不好组织，但这种方式是覆盖而非替换.