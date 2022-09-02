## Custom your form builder
---
```ruby
class BootstrapFormBuilder < ActionView::Helpers::FormBuilder
delegate :content_tag, :javascript_tag, to: :@template

alias_method :native_text_field,:text_field
script = %Q(
         $(document).ready(function () {
             initBootstrapDateTimePicker("#{options[:id]}",#{options[:minView]});
         });
      )
def xx
  native_text_field(name, *args) + javascript_tag(script)
end
```

解释：
1. @template(目前个人的理解应该时当前的view对象)
2. delegate 代理方法，后面可以直接使用
3. 由于复写了原来的text_field,在其他方法需要用到这个方法时有变化，alias_method可解决这个问题
4. 有些前段插件需要js的支持，可以后台约定生成，页面上就只要写方法就好了，十分简约，将javascript_tag(js)添加进构建的html即可

## Custom form builder 多语言问题
---
```ruby
label(name, class: 'col-sm-3 control-label')
```
builder中label方法会自动去寻找locales中的多语言设定
```
zh-CN:
  activerecord:
    attributes:
      customer:
        phone: '手机'
        name: '姓名'
    models:
      customer: '客户'
```
大概遵循上面这种格式，关于模型名称改写什么，还有在namespace情况下该怎么写？
1. 没有namespace可以直接小些model名
2. 有namespace，可以像这样 irm／customer
3. 通过 *Model.model_name.i18n_key* 这个方法可以直接看到正确的结果。

## Form tag 也想要这么方便的使用怎么办？
---
> form_for可以指定自己写的builder，但是form_tag不可以，那怎么办呢？我的方式是写全局的helper来覆盖原有的helper

```ruby
class BootstrapFormTagHelper
  include ActionView::Helpers::FormTagHepler

  alias_method :native_text_field_tag, :text_field_tag

  def text_field_tag
    ...
  end
end
```
但是这样写也有不好的地方：
1. 这样写就可以在用自己定义的同名方法来构建页面，alias_method是为了保留原生的方法，但是方法名不一样了。
2. 没有form_for那么灵活，form_for可以定义多个builder来适应不同的场景，但是这种方式却不可以做到。如果想要做到适应多种情况，可能要定义另外一套方法了，又要多一套方法。