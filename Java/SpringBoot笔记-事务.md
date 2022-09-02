### 事务传播行为
---

> 事务传播行为是为了解决业务层方法之间互相调用的事务问题. 当事务方法被另一个事务方法调用时, 必须指定事务应该如何传播. 例如: 方法可能继续在现有事务中运行, 也可能开启一个新事务, 并在自己的事务中运行.

```java
public enum Propagation {
    //支持当前事务，如果当前没有事务，就新建一个事务，这是最常用的选择
    REQUIRED(0),
    
    //支持当前事务，如果当前没有事务，就以非事务方式执行。 
    SUPPORTS(1),
    
    //支持当前事务，如果当前没有事务，就抛出异常。 
    MANDATORY(2),
    
    //新建事务，如果当前存在事务，把当前事务挂起。 
    REQUIRES_NEW(3),
    
    //以非事务方式执行操作，如果当前存在事务，就把当前事务挂起。
    NOT_SUPPORTED(4),
    
    //以非事务方式执行，如果当前存在事务，则抛出异常。 
    NEVER(5),
    
    //如果当前存在事务，则在嵌套事务内执行。如果当前没有事务，则进行与PROPAGATION_REQUIRED类似的操作。
    NESTED(6);
}
```

1. PROPAGATION_REQUIRED: 默认的事务传播类型, 如果当前存在事务, 则加入该事务; 如果当前没有事务, 则创建一个新的事务.
2. PROPAGATION_REQUIRES_NEW: 创建一个新的事务, 如果当前存在事务, 则把当前事务挂起. 也就是说不管外部方法是否开启事务, PROPAGATION_REQUIRES_NEW 修饰的内部方法会新开启自己的事务, 且开启的事务相互独立, 互不干扰.
3. PROPAGATION_NESTED: 如果当前存在事务, 则创建一个事务作为当前事务的嵌套事务来运行. 如果当前没有事务, 则该取值等价于TransactionDefinition.PROPAGATION_REQUIRED。
4. PROPAGATION_MANDATORY: 如果当前存在事务, 则加入该事务. 如果当前没有事务, 则抛出异常.(mandatory: 强制性)
5. PROPAGATION_SUPPORTS: 如果当前存在事务, 则加入该事务. 如果当前没有事务, 则以非事务的方式继续运行.
6. PROPAGATION_NOT_SUPPORTED: 以非事务方式运行, 如果当前存在事务, 则把当前事务挂起.
7. PROPAGATION_NEVER: 以非事务方式运行, 如果当前存在事务, 则抛出异常.

### 事务中的隔离级别
---
1. ISOLATION_DEFAULT: 使用后端数据库默认的隔离级别, MySQL 默认采用的 REPEATABLE_READ 隔离级别 Oracle 默认采用的 READ_COMMITTED 隔离级别.
2. ISOLATION_READ_UNCOMMITTED: 最低的隔离级别, 使用这个隔离级别很少, 因为它允许读取尚未提交的数据变更, 可能会导致脏读 幻读或不可重复读
3. ISOLATION_READ_COMMITTED: 允许读取并发事务已经提交的数据, 可以阻止脏读, 但是幻读或不可重复读仍有可能发生
4. ISOLATION_REPEATABLE_READ: 对同一字段的多次读取结果都是一致的, 除非数据是被本身事务自己所修改, 可以阻止脏读和不可重复读, 但幻读仍有可能发生
5. ISOLATION_SERIALIZABLE: 最高的隔离级别, 完全服从 ACID 的隔离级别. 所有的事务依次逐个执行, 这样事务之间就完全不可能产生干扰, 也就是说, 该级别可以防止脏读 不可重复读以及幻读. 但是这将严重影响程序的性能, 通常情况下也不会用到该级别

### @Transactional
1. 当 @Transactional 注解作用于类上时, 该类的所有 public 方法将都具有该类型的事务属性, 同时, 我们也可以在方法级别使用该标注来覆盖类级别的定义.
2. 在 @Transactional 注解中如果不配置rollbackFor属性,那么事务只会在遇到RuntimeException的时候才会回滚, 加上rollbackFor=Exception.class,可以让事务在遇到非运行时异常时也回滚