## Obejct
---
> Object 类是一个特殊的类，是所有类的父类

```java
// native 方法, 用于返回当前运行时对象的 Class 对象, 使用了 final 关键字修饰, 故不允许子类重写
public final native Class<?> getClass()


//native 方法, 用于返回对象的哈希码, 主要使用在哈希表中, 比如 JDK 中的HashMap。
public native int hashCode()

// 用于比较 2 个对象的内存地址是否相等, String 类对该方法进行了重写以用于比较字符串的值是否相等
public boolean equals(Object obj)

//naitive 方法，用于创建并返回当前对象的一份拷贝
protected native Object clone() throws CloneNotSupportedException

// 返回类的名字实例的哈希码的 16 进制的字符串. 建议 Object 所有的子类都重写这个方法
public String toString()

// native 方法, 并且不能重写. 唤醒一个在此对象监视器上等待的线程(监视器相当于就是锁的概念). 如果有多个线程在等待只会任意唤醒一个
public final native void notify()

// native 方法, 并且不能重写. 跟 notify 一样, 唯一的区别就是会唤醒在此对象监视器上等待的所有线程, 而不是一个线程
public final native void notifyAll()

// native方法, 并且不能重写. 暂停线程的执行. 注意: sleep 方法没有释放锁, 而 wait 方法释放了锁, timeout 是等待时间
public final native void wait(long timeout) throws InterruptedException

// 多了 nanos 参数, 这个参数表示额外时间(以毫微秒为单位,范围是 0-999999). 所以超时的时间还需要加上 nanos 毫秒
public final void wait(long timeout, int nanos) throws InterruptedException

// 跟之前的2个wait方法一样, 只不过该方法一直等待, 没有超时时间这个概念
public final void wait() throws InterruptedException

// 实例被垃圾回收器回收的时候触发的操作
protected void finalize() throws Throwable { }
```

## 关键字
---
> 访问控制

1. public: 表示紧跟其后的成员可以被任何人引用
2. private: 表示紧跟其后的成员除了类型创建者和类型内部的方法, 任何人都不可引用
3. protected: protected关键字与private效果相当, 差别仅在于继承的类可以访问protected成员

> 类 方法 变量修饰符

### abstract
---
类和类之间如果具有相同的特征, 将这些共同的特征提取出来, 形成的就是抽象类. 类本身是不存在的, 所以抽象类无法实例化对象. 类到对象是实例化，对象到类是抽象。

1. abstract只能修饰类和方法, 不能修饰变量
2. 声明抽象方法不可写出大括号
3. 抽象类也是类, 属于引用数据类型
4. 类和类之间如果具有相同的特征, 将这些共同的特征提取出来, 形成的就是抽象类
5. 抽象类无法实例化对象, 抽象类就是用来被继承的, 子类可以是抽象类, 也可以是非抽象类。
6. 非抽象子类继承了抽象父类, 必须重写父类的抽象方法, 可以有自己的方法，但不能是抽象方法
7. 抽象子类继承了抽象父类, 不需要重写父类的抽象方法, 可以有自己的方法，但必须是抽象方法
8. abstract不能和final, private, static联合使用, 只能单独使用, 或者联合public abstract使用, 一般就是public abstract联合使用。
9. 抽象类虽然无法实例化, 但是抽象类有构造方法, 这个构造方法是供子类使用的。
10. 抽象类中可以有实例方法, 不一定有抽象方法，但是抽象方法一定得出现在抽象类中。

```java
public abstract class Bird extends Animal {
    public Bird() {
        super();
    }
    
    public abstract void doSome();
}
```

### native: JNI-Java Native Interface
---
```java
public class HelloJNI {
    //native 关键字告诉 JVM 调用的是该方法在外部定义
    private native void helloJNI();
 
    static{
        System.loadLibrary("helloJNI");//载入本地库
    }
    public static void main(String[] args) {
        HelloJNI jni = new HelloJNI();
        jni.helloJNI();
    }
}
```

### final: 最终 不可改变的
---
1. 修饰一个类: 当前这个类不能有任何的子类
2. 修饰一个方法: 这个方法就是最终方法, 也就是不能被覆盖重写
3. 修饰一个局部变量: 对于基本类型来说, 不可变说的是变量当中的数据不可改变, 对于引用类型来说，不可变说的是变量当中的地址值不可改变
4. 修饰一个成员变量: 同样不可变


### static
---
> static方法就是没有this的方法. 在static方法内部不能调用非静态方法, 反过来是可以的. 而且可以在没有创建任何对象的前提下, 仅仅通过类本身来调用static方法. 这实际上正是static方法的主要用途, 很多时候会将一些只需要进行一次的初始化操作都放在static代码块中进行

### strictfp
---
> strictfp实际上是对浮点类型做精确运算的一个关键字, 实际意思是FP-strictfp, Java中的浮点类型有float和double两种, 当这两种类型的数字进行运算时Java虚拟机会根据自己的规则进行预算和表达, 这种运算方式是虚拟机底部自动完成的, 有时候得到的结果不是很满意. 该关键字就是为了能够声明, 让浮点类型的数据按照javaEE的规范进行编译和运算, 这样就能得到更加准确的浮点运算的正确结果.

```java
// 该关键字可以修饰在接口、类或者是方法上面
// 该关键字可以修饰的接口上但是却不能修饰在接口中的一个方法上面这是Java对它做的约束规定. 也可以修饰的整个类上面, 但是该类中的构造函数却不能用该关键字修饰

strictfp interface  JieKou {xxxx}

public strictfp class Demo {xxxxx}

strictfp void JiSuan() {xxxxxx}
```

### synchronized
---
> synchronized关键字解决了多个线程之间的资源同步性, synchronized关键字保证了它修饰的方法或者代码块任意时刻只有一个线程在访问.

### transient
---
`对于 transient 修饰的成员变量, 在类的实例对象的序列化处理过程中会被忽略. 因此, transient变量不会贯穿对象的序列化和反序列化, 生命周期仅存于调用者的内存中而不会写到磁盘里进行持久化. (mybatis: 类里的字段不想映射到表里用这个字段)`

### volatile
---
> volatile 是 Java 中的关键字, 是一个变量修饰符, 被用来修饰会被不同线程访问和修改的变量


## Exception
---
> 在 Java 中, 所有的异常都有一个共同的祖先 java.lang 包中的 Throwable 类. Throwable 类有两个重要的子类:

1. Exception: 程序本身可以处理的异常, 可以通过 catch 来进行捕获. Exception 又可以分为 Checked Exception(受检查异常 必须处理) 和 Unchecked Exception(不受检查异常 可以不处理).
2. Error: 属于程序无法处理的错误, 不建议通过 catch 来进行捕获不建议通过catch捕获. 例如 Java 虚拟机运行错误(Virtual MachineError), 虚拟机内存不够错误 (OutOfMemoryError), 类定义错误(NoClassDefFoundError）等. 这些异常发生时, Java 虚拟机（JVM）一般会选择线程终止.

> Checked Exception 和 Unchecked Exception 有什么区别?

1. Checked Exception 即 受检查异常, Java 代码在编译过程中, 如果受检查异常没有被 catch或者throws 关键字处理的话, 就没办法通过编译.
2. Unchecked Exception 即 不受检查异常, Java 代码在编译过程中, 我们即使不处理不受检查异常也可以正常通过编译.

> RuntimeException及其子类都统称为非受检查异常, 常见的有:

1. NullPointerException: 空指针错误
2. IllegalArgumentException: 参数错误比如方法入参类型错误
3. NumberFormatException: IllegalArgumentException的子类, 字符串转换为数字格式错误
4. ArrayIndexOutOfBoundsException: 数组越界错误
5. ClassCastException: 类型转换错误
6. ArithmeticException: 算术错误
7. SecurityException: 安全错误比如权限不够
8. UnsupportedOperationException: 不支持的操作错误比如重复创建同一用户

> Throwable 类常用方法有哪些？

1. String getMessage(): 返回异常发生时的简要描述
2. String toString(): 返回异常发生时的详细信息
3. String getLocalizedMessage(): 返回异常对象的本地化信息
4. void printStackTrace(): 在控制台上打印 Throwable 对象封装的异常信息

#### try-catch-finally
> 不要在 finally 语句块中使用 return, 当 try 语句和 finally 语句中都有 return 语句时, try 语句块中的 return 语句会被忽略. 这是因为 try 语句中的 return 返回值会先被暂存在一个本地变量中, 当执行到 finally 语句中的 return 之后, 这个本地变量的值就变为了 finally 语句中的 return 返回值

1. try块: 用于捕获异常, 其后可接零个或多个 catch 块, 如果没有 catch 块, 则必须跟一个 finally 块
2. catch块: 用于处理 try 捕获到的异常
3. finally块: 无论是否捕获或处理异常, finally 块里的语句都会被执行. 当在 try 块或 catch 块中遇到 return 语句时, finally 语句块将在方法返回之前被执行


#### 如何使用 try-with-resources 代替try-catch-finally?
> 面对必须要关闭的资源, 我们总是应该优先使用 try-with-resources 而不是try-finally. 随之产生的代码更简短, 更清晰, 产生的异常对我们也更有用. try-with-resources语句让我们更容易编写必须要关闭的资源的代码, 若采用try-finally则几乎做不到这点.

1. 适用范围(资源的定义): 任何实现 java.lang.AutoCloseable或者 java.io.Closeable 的对象
2. 关闭资源和 finally 块的执行顺序: 在 try-with-resources 语句中, 任何 catch 或 finally 块在声明的 资源 关闭后 运行

```java
// try-catch-finally
Scanner scanner = null;
try {
    scanner = new Scanner(new File("D://read.txt"));
    while (scanner.hasNext()) {
        System.out.println(scanner.nextLine());
    }
} catch (FileNotFoundException e) {
    e.printStackTrace();
} finally {
    if (scanner != null) {
        scanner.close();
    }
}

// try-with-resources
// 通过使用分号分隔, 可以在try-with-resources块中声明多个资源
try (Scanner scanner = new Scanner(new File("test.txt"))) {
    while (scanner.hasNext()) {
        System.out.println(scanner.nextLine());
    }
} catch (FileNotFoundException fnfe) {
    fnfe.printStackTrace();
}
```

## 反射
---
1. 在运行时判断任意一个对象所属的类
2. 在运行时构造任意一个类的对象
3. 在运行时判断任意一个类所具有的成员变量和方法
4. 在运行时获取泛型信息
5. 在运行时调用任意一个对象的成员变量和方法
6. 在运行时处理注解
7. 生成动态代理


## 范型
```java
// 泛型类
//在实例化泛型类时，必须指定T的具体类型
public class Generic<T>{

    private T key;

    public Generic(T key) {
        this.key = key;
    }

    public T getKey(){
        return key;
    }
}

// 泛型接口
public interface Generator<T> {
    public T method();
}

// 实现泛型接口，不指定类型
class GeneratorImpl<T> implements Generator<T>{
    @Override
    public T method() {
        return null;
    }
}

// 实现泛型接口, 指定类型
class GeneratorImpl<T> implements Generator<String>{
    @Override
    public String method() {
        return "hello";
    }
}


// 泛型方法
// 类在实例化时才能真正的传递类型参数，由于静态方法的加载先于类的实例化，也就是说类中的泛型还没有传递真正的类型参数，静态的方法的加载就已经完成了，所以静态泛型方法是没有办法使用类上声明的泛型的。只能使用自己声明的 <E>
public static < E > void printArray( E[] inputArray )
{
     for ( E element : inputArray ){
        System.out.printf( "%s ", element );
     }
     System.out.println();
}
```


#### extends通配符
---
```java
// 范型类
class Pair<T> {
    private T first;
    private T last;
    public Pair(T first, T last) {
        this.first = first;
        this.last = last;
    }
    public T getFirst() {
        return first;
    }
    public T getLast() {
        return last;
    }
}

public class Main {  
	public static void main(String[] args) {
    Pair<Integer> p = new Pair<>(123, 456);
    int n = add(p);
    System.out.println(n);
  }

  // 设置extends通配符，Pari<Integer>类型可传入Pair<Number>
  static int add(Pair<? extends Number> p) { 
    Number first = p.getFirst();
    Number last = p.getLast();
    return first.intValue() + last.intValue();
  }
}
```

1. 这种使用 <? extends Number> 的泛型定义称之为上界通配符（Upper Bounds Wildcards), 即把泛型类型T的上界限定在Number了
2. 使用extends通配符表示可以读, 不能写

#### super通配符
---
```
// Pair<? super Integer>表示，方法参数接受所有泛型类型为Integer或Integer父类的Pair类型
void set(Pair<? super Integer> p, Integer first, Integer last) {
    p.setFirst(first);
    p.setLast(last);
}
```

1. 限定类型
2. 使用extends通配符表示只能写, 不能读

##### 何时使用extends, 何时使用super? (PECS原则: Producer Extends Consumer Super)
---

1. 如果需要返回T, 它是生产者(Producer), 使用extends通配符
2. 如果需要写入T, 它是消费者(Consumer), 要使用super通配符

#### 无限定通配符
---
```java
void sample(Pair<?> p) {
}
```

因为<?>通配符既没有extends, 也没有super, 因此:

1. 不允许调用set(T)方法并传入引用(null除外)
2. 不允许调用T get()方法并获取T引用(只能获取Object引用)
3. 换句话说, 既不能读, 也不能写, 那只能做一些null判断

#### 擦拭法(Type Erasure)
---
> Java范型实现方式是擦拭法, 所谓擦拭法是指, 虚拟机对泛型其实一无所知, 所有的工作都是编译器做的. Java的泛型是由编译器在编译时实行的, 编译器内部永远把所有类型T视为Object处理, 但是, 在需要转型的时候, 编译器会根据T的类型自动为我们实行安全地强制转型。

1. 编译器把类型<T>视为Object
2. 编译器根据<T>实现安全的强制转型

使用泛型的时候, 我们编写的代码也是编译器看到的代码:
```java
Pair<String> p = new Pair<>("Hello", "world");
String first = p.getFirst();
String last = p.getLast();
```

而虚拟机执行的代码并没有泛型：
```java
Pair p = new Pair("Hello", "world");
String first = (String) p.getFirst();
String last = (String) p.getLast();
```

> 擦拭法决定了泛型<T>:

1. 不能是基本类型, 例如: int
2. 不能获取带泛型类型的Class, 例如：Pair.class
3. 不能判断带泛型类型的类型, 例如: x instanceof Pair
4. 不能实例化T类型, 例如: new T()
5. 泛型方法要防止重复定义方法, 例如: public boolean equals(T obj)
6. 子类可以获取父类的泛型类型<T>

## Optional
> 设计意图

1. Optional 是用来作为方法返回值的
2. Optional 是为了清晰地表达返回值中没有结果的可能性
3. 且如果直接返回 null 很可能导致调用端产生错误(尤其是NullPointerException)

> 使用TIPS

1. Optional 是用来作为方法返回值的, 不要滥用 Optional API
2. 不要使用Optional作为Java Bean实例域的类型, 因为 Optional 没有实现Serializable接口(不可序列化)
3. 不要使用 Optional 作为类构造器参数
4. 不要使用 Optional 作为Java Bean Setter方法的参数
5. 不要使用Optional作为方法参数的类型
6. 不要在集合中使用 Optional 类
7. 不要把容器类型（包括 List, Set, Map, 数组, Stream 甚至 Optional ）包装在Optional中
8. Optional 是为了清晰地表达返回值中没有结果的可能性, 不要给Optional变量赋值 null
9. 确保Optional内有值才能调用 get() 方法
10. 使用 equals 而不是 == 来比较 Optional 的值, Optional 的 equals 方法已经实现了内部值比较