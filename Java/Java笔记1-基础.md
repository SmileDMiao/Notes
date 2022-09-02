## 重载和重写?
---
1. 重载就是同样的一个方法能够根据输入数据的不同, 做出不同的处理
2. 重写就是当子类继承自父类的相同方法, 输入数据一样, 但要做出有别于父类的响应时, 你就要覆盖父类方法

![IMAGE](resources/A37DD3ED0CC203D5D6C32357ABCD7C1A.jpg =737x256)

## 基本类型 包装类
---
![IMAGE](resources/9476AB06E271B96261316BE8CA8C1FD9.jpg =709x332)

1. Java有八种基本数据类型: byte, short, int, long, float, double, boolean, char
2. Java为其提供了8种对应的包装类: Byte, Short, Integer, Long, Float, Double, Boolean, Character
3. Java 里使用 long 类型的数据一定要在数值后面加上 L, 否则将作为整型解析
4. char a = 'h' char :单引号, String a = "hello" :双引号

> 基本类型和包装类型的区别?

1. 成员变量包装类型不赋值就是 null, 而基本类型有默认值且不是 null
2. 包装类型可用于泛型, 而基本类型不可以
3. 基本数据类型的局部变量存放在 Java 虚拟机栈中的局部变量表中, 基本数据类型的成员变量(未被static修饰)存放在 Java 虚拟机的堆中. 包装类型属于对象类型, 我们知道几乎所有对象实例都存在于堆中
4. 相比于对象类型，基本数据类型占用的空间非常小
5. 所有整型包装类对象之间值的比较, 全部使用 equals 方法比较
6. 两种浮点数类型的包装类 Float Double 并没有实现缓存机制

> 包装类型的缓存机制了解么?

```java
public static Integer valueOf(int i) {
    if (i >= IntegerCache.low && i <= IntegerCache.high)
        return IntegerCache.cache[i + (-IntegerCache.low)];
    return new Integer(i);
}
private static class IntegerCache {
    static final int low = -128;
    static final int high;
    static {
        // high value may be configured by property
        int h = 127;
    }
}
```

> 自动装箱与拆箱了解吗？原理是什么？

1. 装箱: 将基本类型用它们对应的引用类型包装起来
2. 拆箱: 将包装类型转换为基本数据类型
3. 装箱其实就是调用了 包装类的 valueOf()方法
4. 拆箱其实就是调用了 xxxValue()方法

## 为什么 Java 中只有值传递？
---
>Java中将实参传递给方法(或函数)的方式是 值传递

1. 如果参数是基本类型的话，很简单，传递的就是基本类型的字面量值的拷贝，会创建副本。
2. 如果参数是引用类型，传递的就是实参所引用的对象在堆中地址值的拷贝，同样也会创建副本。

```java
// 交换方法不影响最后输出
public static void main(String[] args) {
    int num1 = 10;
    int num2 = 20;
    swap(num1, num2);
    System.out.println("num1 = " + num1);
    System.out.println("num2 = " + num2);
}

public static void swap(int a, int b) {
    int temp = a;
    a = b;
    b = temp;
    System.out.println("a = " + a);
    System.out.println("b = " + b);
}


// 这里传递的还是值, 不过, 这个值是实参的地址罢了.
// 也就是说 change 方法的参数拷贝的是 arr(实参)的地址, 它和 arr 指向的是同一个数组对象. 这也就说明了为什么方法内部对形参的修改会影响到实参
public static void main(String[] args) {
      int[] arr = { 1, 2, 3, 4, 5 };
      System.out.println(arr[0]);
      change(arr);
      System.out.println(arr[0]);
	}

	public static void change(int[] array) {
      // 将数组的第一个元素变为0
      array[0] = 0;
	}
	
	
	public class Person {
    private String name;
}

// swap 方法的参数 person1 和 person2 只是拷贝的实参 xiaoZhang 和 xiaoLi 的地址
// person1 和 person2 的互换只是拷贝的两个地址的互换罢了, 并不会影响到实参 xiaoZhang 和 xiaoLi 
public static void main(String[] args) {
    Person xiaoZhang = new Person("小张");
    Person xiaoLi = new Person("小李");
    swap(xiaoZhang, xiaoLi);
    System.out.println("xiaoZhang:" + xiaoZhang.getName());
    System.out.println("xiaoLi:" + xiaoLi.getName());
}

public static void swap(Person person1, Person person2) {
    Person temp = person1;
    person1 = person2;
    person2 = temp;
    System.out.println("person1:" + person1.getName());
    System.out.println("person2:" + person2.getName());
}
```

## String StringBuffer StringBuilder 的区别?
---

1. String 是不可变的
2. StringBuilder 与 StringBuffer 都继承自 AbstractStringBuilder 类, 在 AbstractStringBuilder 中也是使用字符数组保存字符串, 不过没有使用 final 和 private 关键字修饰, 最关键的是这个 AbstractStringBuilder 类还提供了很多修改字符串的方法比如 append 方法.
3. String 中的对象是不可变的, 也就可以理解为常量, 线程安全
4. AbstractStringBuilder 是 StringBuilder 与 StringBuffer 的公共父类, 定义了一些字符串的基本操作, 如 expandCapacity append insert indexOf 等公共方法. StringBuffer 对方法加了同步锁或者对调用的方法加了同步锁, 所以是线程安全的。StringBuilder 并没有对方法进行加同步锁, 所以是非线程安全的
5. 每次对 String 类型进行改变的时候, 都会生成一个新的 String 对象, 然后将指针指向新的 String 对象.
6. StringBuffer 每次都会对 StringBuffer 对象本身进行操作, 而不是生成新的对象并改变对象引用. 相同情况下使用 StringBuilder 相比使用 StringBuffer 仅能获得 10%~15% 左右的性能提升, 但却要冒多线程不安全的风险.

#### == 和 equals() 的区别?
---
> 因为 Java 只有值传递, 所以, 对于 == 来说, 不管是比较基本数据类型, 还是引用数据类型的变量, 其本质比较的都是值, 只是引用类型变量存的值是对象的地址

1. 对于基本数据类型来说, == 比较的是值
2. 对于引用数据类型来说, == 比较的是对象的内存地址
3. String 中的 equals 方法是被重写过的, 比较的是 String 字符串的值是否相等. Object 的 equals 方法是比较的对象的内存地址.

#### 字符串拼接
---
> 字符串对象通过 "+" 的字符串拼接方式, 实际上是通过 StringBuilder 调用 append() 方法实现的, 拼接完成之后调用 toString() 得到一个 String 对象. 但是在循环内使用 "+" 进行字符串的拼接的话, 编译器不会创建单个 StringBuilder 以复用, 会导致创建过多的 StringBuilder 对象.

```java
// 创建过多SpringBuilder
String[] arr = {"he", "llo", "world"};
String s = "";
for (int i = 0; i < arr.length; i++) {
    s += arr[i];
}
System.out.println(s);

// 只创建一个SpringBuilder
String[] arr = {"he", "llo", "world"};
StringBuilder s = new StringBuilder();
for (String value : arr) {
    s.append(value);
}
System.out.println(s);
```

#### 运行时常量池
---
> Class 文件中除了有类的版本 字段 方法 接口等描述信息外，还有用于存放编译期生成的各种字面量（Literal）和符号引用（Symbolic Reference）的 常量池表(Constant Pool Table). 字面量是源代码中的固定值的表示法, 即通过字面我们就能知道其值的含义. 字面量包括整数 浮点数和字符串字面量, 符号引用包括类符号引用 字段符号引用 方法符号引用和接口方法符号引用. 常量池表会在类加载后存放到方法区的运行时常量池中. 运行时常量池的功能类似于传统编程语言的符号表, 尽管它包含了比典型符号表更广泛的数据. 既然运行时常量池是方法区的一部分, 自然受到方法区内存的限制, 当常量池无法再申请到内存时会抛出 OutOfMemoryError 错误.

#### 字符串常量池
---
> 字符串常量池 是 JVM 为了提升性能和减少内存消耗针对字符串(String 类)专门开辟的一块区域, 主要目的是为了避免字符串的重复创建.

1. HotSpot 虚拟机中字符串常量池的实现是 src/hotspot/share/classfile/stringTable.cpp
2. StringTable 本质上就是一个HashSet<String> 容量为 StringTableSize(可以通过 -XX:StringTableSize 参数来设置)
3. StringTable 中保存的是字符串对象的引用, 字符串对象的引用指向堆中的字符串对象.

#### `String s1 = new String("Hello")` 这句话创建了几个字符串对象？
---
1. 如果字符串常量池中不存在字符串对象 "Hello" 的引用, 那么会在堆中创建 2 个字符串对象
2. class文件中 ldc 命令用于判断字符串常量池中是否保存了对应的字符串对象的引用, 如果保存了的话直接返回, 如果没有保存的话, 会在堆中创建对应的字符串对象并将该字符串对象的引用保存到字符串常量池中.
3. 如果字符串常量池中已存在字符串对象 "Hello" 的引用, 则只会在堆中创建 1 个字符串对象

#### intern 方法有什么作用?
---
> String.intern() 是一个 native(本地)方法, 其作用是将指定的字符串对象的引用保存在字符串常量池中

1. 如果字符串常量池中保存了对应的字符串对象的引用, 就直接返回该引用.
2. 如果字符串常量池中没有保存了对应的字符串对象的引用, 那就在常量池中创建一个指向该字符串对象的引用并返回.



## 注解
---
> 注解(Annotation)是放在Java源码的类, 方法, 字段, 参数前的一种特殊"注释". 注释会被编译器直接忽略, 注解则可以被编译器打包进入class文件, 因此, 注解是一种用作标注的 "元数据". 从JVM的角度看, 注解本身对代码逻辑没有任何影响, 如何使用注解完全由工具决定.

**Java的注解可以分为三类**
1. 第一类是由编译器使用的注解, 例如@Override: 让编译器检查该方法是否正确地实现了覆写, @SuppressWarnings：告诉编译器忽略此处代码产生的警告. 这类注解不会被编译进入.class文件, 它们在编译后就被编译器扔掉了.
2. 第二类是由工具处理.class文件使用的注解, 比如有些工具会在加载class的时候, 对class做动态修改, 实现一些特殊的功能. 这类注解会被编译进入.class文件, 但加载结束后并不会存在于内存中. 这类注解只被一些底层库使用, 一般我们不必自己处理.
3. 第三类是在程序运行期能够读取的注解, 它们在加载后一直存在于JVM中, 这也是最常用的注解. 例如, 一个配置了@PostConstruct的方法会在调用构造方法后自动被调用(这是Java代码读取该注解实现的功能, JVM并不会识别该注解)

```java
// Java语言使用 @interface语法来定义注解(Annotation), 它的格式如下:
public @interface Report {
    int type() default 0;
    String level() default "info";
    String value() default "";
}
```

#### 元注解
---
> 有一些注解可以修饰其他注解, 这些注解就称为元注解(meta annotation). Java标准库已经定义了一些元注解, 我们只需要使用元注解, 通常不需要自己去编写元注解. 其中, 必须设置@Target和@Retention, @Retention一般设置为RUNTIME, 因为我们自定义的注解通常要求在运行期读取. 一般情况下, 不必写@Inherited和@Repeatable

**@Target**
> 使用@Target可以定义Annotation能够被应用于源码的哪些位置:

1. 类或接口: ElementType.TYPE
2. 字段: ElementType.FIELD
3. 方法: ElementType.METHOD
4. 构造方法: ElementType.CONSTRUCTOR
5. 方法参数: ElementType.PARAMETER

**@Retention**
> 另一个重要的元注解@Retention定义了Annotation的生命周期:

1. 仅编译期: RetentionPolicy.SOURCE
2. 仅class文件: RetentionPolicy.CLASS
3. 运行期: RetentionPolicy.RUNTIME
4. 如果@Retention不存在, 则该Annotation默认为CLASS. 因为通常我们自定义的Annotation都是RUNTIME, 务必要加上@Retention(RetentionPolicy.RUNTIME)这个元注解

**@Repeatable**
> 使用@Repeatable这个元注解可以定义Annotation是否可重复. 经过@Repeatable修饰后, 在某个类型声明处, 就可以添加多个@Report注解

**@Inherited**
> 使用@Inherited定义子类是否可继承父类定义的Annotation. @Inherited仅针对@Target(ElementType.TYPE)类型的annotation有效, 并且仅针对class的继承, 对interface的继承无效.

#### 处理注解
---
> Java的注解本身对代码逻辑没有任何影响. 因为注解定义后也是一种class, 所有的注解都继承自java.lang.annotation.Annotation, 读取注解, 需要使用反射API.

```java
// @Range注解, 我们希望用它来定义一个String字段的规则: 字段长度满足@Range的参数定义
@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.FIELD)
public @interface Range {
    int min() default 0;
    int max() default 255;
}

// 在某个JavaBean中, 我们可以使用该注解
public class Person {
    @Range(min=1, max=20)
    public String name;

    @Range(max=10)
    public String city;
}
```

> 定义了注解, 本身对程序逻辑没有任何影响. 我们必须自己编写代码来使用注解。这里, 我们编写一个Person实例的检查方法, 它可以检查Person实例的String字段长度是否满足@Range的定义. 这样一来, 我们通过@Range注解, 配合check()方法, 就可以完成Person实例的检查. 注意检查逻辑完全是我们自己编写的, JVM不会自动给注解添加任何额外的逻辑.

```java
void check(Person person) throws IllegalArgumentException, ReflectiveOperationException {
    // 遍历所有Field:
    for (Field field : person.getClass().getFields()) {
        // 获取Field定义的@Range:
        Range range = field.getAnnotation(Range.class);
        // 如果@Range存在:
        if (range != null) {
            // 获取Field的值:
            Object value = field.get(person);
            // 如果值是String:
            if (value instanceof String) {
                String s = (String) value;
                // 判断值是否满足@Range的min/max:
                if (s.length() < range.min() || s.length() > range.max()) {
                    throw new IllegalArgumentException("Invalid field: " + field.getName());
                }
            }
        }
    }
}
```