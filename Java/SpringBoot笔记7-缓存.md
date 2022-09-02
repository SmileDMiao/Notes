## Redis配置
---
```yml
spring
  redis:
    host: localhost
    database: 0
    port: 6379
    password:
    timeout: 1000ms
    lettuce:
      pool:
        max-active: 8
        max-wait: -1ms
        max-idle: 8
        min-idle: 0
  cache:
    type: redis
```

## @Cacheable
---
`@Cacheable 注解在方法上, 表示该方法的返回结果是可以缓存的. 也就是说, 该方法的返回结果会放在缓存中, 以便于以后使用相同的参数调用该方法时. 会返回缓存中的值, 而不会实际执行该方法. 对于同一个方法, 如果参数相同, 那么返回结果也是相同的. 如果参数不同, 缓存只能假设结果是不同的, 所以对于同一个方法, 你的程序运行过程中, 使用了多少种参数组合调用过该方法，理论上就会生成多少个缓存的 key.`

### @EnableCaching: 启用缓存

```java
@EnableCaching
@SpringBootApplication
public class MyApplication {
    public static void main(String[] args) {
        SpringApplication.run(MyApplication.class, args);
    }
}
```

### @Cacheable: 这个注解一般用在查询方法上
1. cacheNames/value: 用来指定缓存组件的名字
2. key: 缓存数据时使用的 key, 可以用它来指定. 默认是使用方法参数的值
3. keyGenerator: key 的生成器. key 和 keyGenerator 二选一使用
4. cacheManager: 可以用来指定缓存管理器, 从哪个缓存管理器里面获取缓存
5. condition: 可以用来指定符合条件的情况下才缓存
6. unless: 否定缓存. 当 unless 指定的条件为 true, 方法的返回值就不会被缓存.
7. sync: 是否使用异步模式, true/false

```java
// @Cacheable 提供两个参数来指定缓存名: value cacheNames
@Cacheable("userByName")
List<User> findByUsernameLike(String name);

@CachePut(value="addresses", condition="#customer.name=='Tom'")
public String getAddress(Customer customer) {...}

@CachePut(value="addresses")
public String getAddress(Customer customer) {...}

// 关联多个缓存key
@Cacheable({"menu", "menuById"})
```

### @CacheEvict: 一般用在更新或者删除的方法上
> @CacheEvict是用来标注在需要清除缓存元素的方法或类上的

### @CachePut: 通常用在新增方法上
> 与@Cacheable不同的是使用@CachePut标注的方法在执行前不会去检查缓存中是否存在之前执行过的结果, 而是每次都会执行该方法, 并将执行结果以键值对的形式存入指定的缓存中

```java
@CachePut(value="addresses")
public String getAddress(Customer customer) {...}
```

### @Caching
> @Caching注解可以让我们在一个方法或者类上同时指定多个Spring Cache相关的注解, 其拥有三个属性: cacheable、put和evict, 分别用于指定@Cacheable、@CachePut和@CacheEvict.

```java
@Caching(cacheable = @Cacheable("users"),
         evict = { @CacheEvict("cache2"), @CacheEvict(value = "cache3", allEntries = true) })
```

### @CacheConfig
> 通过@CacheConfig注释, 我们可以将一些缓存配置简化到类级别的单个位置, 这样我们就不必多次声明

```java
@CacheConfig(cacheNames={"addresses"})
public class CustomerDataService {

    @Cacheable
    public String getAddress(Customer customer) {...}
```

### 不支持设置过期时间怎么办?
```java


```


### 序列化
```java
```