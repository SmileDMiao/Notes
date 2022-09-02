## IOC
---
> IoC(Inverse of Control:控制反转) 是一种设计思想, 而不是一个具体的技术实现. IoC 的思想就是将原本在程序中手动创建对象的控制权, 交由 Spring 框架来管理

1. 控制: 指的是对象创建(实例化, 管理)的权力
2. 反转: 控制权交给外部环境(Spring 框架, IoC 容器)

> 将对象之间的相互依赖关系交给 IoC 容器来管理, 并由 IoC 容器完成对象的注入. 这样可以很大程度上简化应用的开发, 把应用从复杂的依赖关系中解放出来. IoC 容器就像是一个工厂一样, 当我们需要创建一个对象的时候, 只需要配置好配置文件/注解即可, 完全不用考虑对象是如何被创建出来的。

#### Spring Bean
---
> 简单来说, Bean 代指的就是那些被 IoC 容器所管理的对象. 我们需要告诉 IoC 容器帮助我们管理哪些对象, 这个是通过配置元数据来定义的. 配置元数据可以是 XML 文件, 注解或者 Java 配置类.

> 将一个类声明为 Bean 的注解有哪些?

1. @Component: 通用的注解, 可标注任意类为 Spring 组件. 如果一个 Bean 不知道属于哪个层, 可以使用@Component 注解标注。
2. @Repository: 对应持久层即 Dao 层, 主要用于数据库相关操作
3. @Service: 对应服务层, 主要涉及一些复杂的逻辑, 需要用到 Dao 层
4. @Controller: 对应 Spring MVC 控制层, 主要用户接受用户请求并调用 Service 层返回数据给前端页面

> 注入 Bean 的注解有哪些？

1. @Autowired
2. @Resource
3. @Inject

> @AUtowired和@Resource?

1. @Autowired 是 Spring 提供的注解, @Resource 是 JDK 提供的注解
2. Autowired 默认的注入方式为byType(根据类型进行匹配), @Resource默认注入方式为 byName(根据名称进行匹配)
3. 当一个接口存在多个实现类的情况下, @Autowired 和@Resource都需要通过名称才能正确匹配到对应的 Bean. Autowired 可以通过 @Qualifier 注解来显示指定名称, @Resource可以通过 name 属性来显示指定名称

> @Component 和 @Bean 的区别是什么？

1. @Component 注解作用于类,  而@Bean注解作用于方法
2. @Component通常是通过类路径扫描来自动侦测以及自动装配到 Spring 容器中 (我们可以使用 @ComponentScan 注解定义要扫描的路径从中找出标识了需要装配的类自动装配到 Spring 的 bean 容器中). @Bean 注解通常是我们在标有该注解的方法中定义产生这个 bean, @Bean告诉了 Spring 这是某个类的实例, 当我需要用它的时候还给我.
3. @Bean 注解比 @Component 注解的自定义性更强, 而且很多地方我们只能通过 @Bean 注解来注册 bean. 比如当我们引用第三方库中的类需要装配到 Spring容器时, 则只能通过 @Bean来实现。

> Bean 的生命周期

1. Bean 容器找到配置文件中 Spring Bean 的定义
2. Bean 容器利用 Java Reflection API 创建一个 Bean 的实例
3. 如果涉及到一些属性值 利用 set()方法设置一些属性值
4. 如果 Bean 实现了 BeanNameAware 接口, 调用 setBeanName()方法，传入 Bean 的名字
5. 如果 Bean 实现了 BeanClassLoaderAware 接口, 调用 setBeanClassLoader()方法, 传入 ClassLoader对象的实例
6. 如果 Bean 实现了 BeanFactoryAware 接口, 调用 setBeanFactory()方法, 传入 BeanFactory对象的实例
7. 与上面的类似，如果实现了其他 *.Aware接口, 就调用相应的方法
8. 如果有和加载这个 Bean 的 Spring 容器相关的 BeanPostProcessor 对象, 执行postProcessBeforeInitialization() 方法
9. 如果 Bean 实现了InitializingBean接口, 执行afterPropertiesSet()方法
10. 如果 Bean 在配置文件中的定义包含 init-method 属性, 执行指定的方法
11. 如果有和加载这个 Bean 的 Spring 容器相关的 BeanPostProcessor 对象, 执行postProcessAfterInitialization() 方法
12. 当要销毁 Bean 的时候, 如果 Bean 实现了 DisposableBean 接口, 执行 destroy() 方法
13. 当要销毁 Bean 的时候, 如果 Bean 在配置文件中的定义包含 destroy-method 属性, 执行指定的方法

## AOP
---
> AOP(Aspect-Oriented Programming:面向切面编程)能够将那些与业务无关, 却为业务模块所共同调用的逻辑或责任(例如事务处理, 日志管理, 权限控制等)封装起来, 便于减少系统的重复代码, 降低模块间的耦合度, 并有利于未来的可拓展性和可维护性.

> 术语

1. Target: 目标, 被通知的对象
2. Proxy:	代理, 向目标对象应用通知之后创建的代理对象
3. JoinPoint:	连接点, 目标对象的所属类中, 定义的所有方法均为连接点
4. Pointcut: 切入点, 被切面拦截 / 增强的连接点 (切入点一定是连接点, 连接点不一定是切入点)
5. Advice: 通知, 增强的逻辑 / 代码, 也即拦截到目标对象的连接点之后要做的事情
6. Aspect: 切面, 切入点(Pointcut)+通知(Advice)
7. Weaving: 织入,	将通知应用到目标对象, 进而生成代理对象的过程动作

> AspectJ 定义的通知类型有哪些？

1. Before: 前置通知, 目标对象的方法调用之前触发
2. After: 后置通知, 目标对象的方法调用之后触发
3. AfterReturning: 返回通知, 目标对象的方法调用完成, 在返回结果值之后触发
4. AfterThrowing: 异常通知, 目标对象的方法运行中抛出 / 触发异常后触发, AfterReturning 和 AfterThrowing 两者互斥. 如果方法调用成功无异常, 则会有返回值. 如果方法抛出了异常, 则不会有返回值
5. Around: 环绕通知, 编程式控制目标对象的方法调用. 环绕通知是所有通知类型中可操作范围最大的一种, 因为它可以直接拿到目标对象, 以及要执行的方法, 所以环绕通知可以任意的在目标对象的方法调用前后搞事, 甚至不调用目标对象的方法


> 多个切面的执行顺序如何控制？

`通常使用@Order 注解直接定义切面顺序`