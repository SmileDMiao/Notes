#### 1. @SpringBootApplication
> 我们可以把 @SpringBootApplication看作是 @Configuration, @EnableAutoConfiguration, @ComponentScan 注解的集合.


1. @EnableAutoConfiguration:  启用 SpringBoot 的自动配置机制
2. @ComponentScan:  扫描被@Component(@Service, @Controller)注解的 bean, 注解默认会扫描该类所在的包下所有的类
3. @Configuration: 允许在 Spring 上下文中注册额外的 bean 或导入其他配置类


### 2. Spring Bean 相关
#### @Autowired
> 自动导入对象到类中, 被注入进的类同样要被 Spring 容器管理比如: Service 类注入到 Controller 类中

```java
@Service
public class UserService {
  ......
}

@RestController
@RequestMapping("/users")
public class UserController {
   @Autowired
   private UserService userService;
   ......
}
```

#### @Primary
> 赋予bean更高的优先级

```java
@Primary
@Bean
public RedisCacheManager ttlCacheManager(RedisConnectionFactory redisConnectionFactory) {
    return new TtlRedisCacheManager(RedisCacheWriter.lockingRedisCacheWriter(redisConnectionFactory),
            this.getRedisCacheConfigurationWithTtl(60));
}
```

#### @Component @Repository @Service @Controller
> 我们一般使用 @Autowired 注解让 Spring 容器帮我们自动装配 bean. 要想把类标识成可用于 @Autowired 注解自动装配的 bean 的类, 可以采用以下注解实现:

1. @Component: 通用的注解, 可标注任意类为 Spring 组件. 如果一个 Bean 不知道属于哪个层, 可以使用@Component 注解标注.
2. @Repository: 对应持久层即 Dao 层, 主要用于数据库相关操作
3. @Service: 对应服务层, 主要涉及一些复杂的逻辑, 需要用到 Dao 层。
4. @Controller: 对应 Spring MVC 控制层, 主要用户接受用户请求并调用 Service 层返回数据给前端页面.

#### @RestController
> @RestController注解是@Controller和@ResponseBody的合集, 表示这是个控制器 bean, 并且是将函数的返回值直 接填入 HTTP 响应体中, 是 REST 风格的控制器
单独使用 @Controller 不加 @ResponseBody的话一般使用在要返回一个视图的情况, 这种情况属于比较传统的 Spring MVC 的应用, 对应于前后端不分离的情况. @Controller +@ResponseBody 返回 JSON 或 XML 形式数据

#### @Scope
> 声明 Spring Bean 的作用域

```java
@Bean
@Scope("singleton")
public Person personSingleton() {
    return new Person();
}
```

1. singleton: 唯一 bean 实例, Spring 中的 bean 默认都是单例的。
2. prototype: 每次请求都会创建一个新的 bean 实例.
3. request: 每一次 HTTP 请求都会产生一个新的 bean, 该 bean 仅在当前 HTTP request 内有效.
4. session: 每一次 HTTP 请求都会产生一个新的 bean, 该 bean 仅在当前 HTTP session 内有效.

#### @Configuration
> 一般用来声明配置类, 可以使用 @Component注解替代, 不过使用Configuration注解声明配置类更加语义化.

```java
@Configuration
public class AppConfig {
    @Bean
    public TransferService transferService() {
        return new TransferServiceImpl();
    }

}
```

### 3. 处理常见的 HTTP 请求类型
> @GetMapping @PostMapping @PutMapping @DeleteMapping @PatchMapping

1. @GetMapping("users") 等价于@RequestMapping(value="/users",method=RequestMethod.GET)
2. @PostMapping("users") 等价于@RequestMapping(value="/users",method=RequestMethod.POST)
3. @PutMapping("/users/{userId}") 等价于@RequestMapping(value="/users/{userId}",method=RequestMethod.PUT)
4. @DeleteMapping("/users/{userId}")等价于@RequestMapping(value="/users/{userId}",method=RequestMethod.DELETE)
5. @PatchMapping("/profile")

### 4. 前后端传值
####  @PathVariable  @RequestParam
1. @PathVariable用于获取路径参数
2. @RequestParam用于获取查询参数

#### @RequestBody
> 用于读取 Request 请求(可能是 POST,PUT,DELETE,GET 请求)的 body 部分并且Content-Type 为 application/json 格式的数据, 接收到数据之后会自动将数据绑定到 Java 对象上去. 系统会使用HttpMessageConverter或者自定义的HttpMessageConverter将请求的 body 中的 json 字符串转换为 java 对象. 一个请求方法只可以有一个@RequestBody, 但是可以有多个@RequestParam和@PathVariable

### 5. 读取配置信息
```yml
wuhan2020: 2020年初武汉爆发了新型冠状病毒，疫情严重，但是，我相信一切都会过去！武汉加油！中国加油！

my-profile:
  name: Guide哥
  email: koushuangbwcx@163.com

library:
  location: 湖北武汉加油中国加油
  books:
    - name: 天才基本法
      description: 二十二岁的林朝夕在父亲确诊阿尔茨海默病这天，得知自己暗恋多年的校园男神裴之即将出国深造的消息——对方考取的学校，恰是父亲当年为她放弃的那所。
    - name: 时间的秩序
      description: 为什么我们记得过去，而非未来？时间“流逝”意味着什么？是我们存在于时间之内，还是时间存在于我们之中？卡洛·罗韦利用诗意的文字，邀请我们思考这一亘古难题——时间的本质。
    - name: 了不起的我
      description: 如何养成一个新习惯？如何让心智变得更成熟？如何拥有高质量的关系？ 如何走出人生的艰难时刻？
```
#### @value(常用)
> 使用 @Value("${property}") 读取比较简单的配置信息

```java
@Value("${wuhan2020}")
String wuhan2020;
```

#### @ConfigurationProperties(常用)
> 通过@ConfigurationProperties读取配置信息并与 bean 绑定

```java
@Component
@ConfigurationProperties(prefix = "library")
class LibraryProperties {
    @NotEmpty
    private String location;
    private List<Book> books;

    @Setter
    @Getter
    @ToString
    static class Book {
        String name;
        String description;
    }
  省略getter/setter
  ......
}
```

#### PropertySource(不常用)
> @PropertySource读取指定 properties 文件

```java
@Component
@PropertySource("classpath:website.properties")

class WebSite {
    @Value("${url}")
    private String url;

  省略getter/setter
  ......
}
```

### 6. 参数校验Validation(javax.validation.constraints)

> @Valid 和 @Validated

1. @Valid注解: 是Bean Validation 所定义，可以添加在普通方法, 构造方法, 方法参数, 方法返回, 成员变量上, 表示它们需要进行约束校验
2. @Validated注解: 是 Spring Validation 锁定义, 可以添加在类, 方法参数, 普通方法上, 表示它们需要进行约束校验. 同时, @Validated 有 value 属性, 支持分组校验.

> 空和非空检查

1. @NotBlank: 只能用于字符串不为 null, 并且字符串 #trim() 以后 length 要大于 0
2. @NotEmpty: 集合对象的元素不为 0, 即集合不为空，也可以用于字符串不为 null
3. @NotNull: 不能为 null
4. @Null: 必须为 null

> 数值检查

1. @DecimalMax(value) ：被注释的元素必须是一个数字，其值必须小于等于指定的最大值
2. @DecimalMin(value) ：被注释的元素必须是一个数字，其值必须大于等于指定的最小值
3. @Digits(integer, fraction): 被注释的元素必须是一个数字，其值必须在可接受的范围内
4. @Positive: 判断正数
5. @PositiveOrZero: 判断正数或 0
6. @Max(value): 该字段的值只能小于或等于该值
7. @Min(value): 该字段的值只能大于或等于该值
8. @Negative: 判断负数
9. @NegativeOrZero: 判断负数或 0

> Boolean 值检查

1. @AssertFalse: 被注释的元素必须为 true 
2. @AssertTrue: 被注释的元素必须为 false

> 长度检查

1. @Size(max, min): 检查该字段的 size 是否在 min 和 max 之间, 可以是字符串 数组 集合 Map 等

> 日期检查

1. @Future: 被注释的元素必须是一个将来的日期
2. @FutureOrPresent: 判断日期是否是将来或现在日期
3. @Past: 检查该字段的日期是在过去
4. @PastOrPresent: 判断日期是否是过去或现在日期
其它检查
5. @Email ：被注释的元素必须是电子邮箱地址。
6. @Pattern(value) ：被注释的元素必须符合指定的正则表达式

### 7. 全局处理 Controller 层异常
1. @ControllerAdvice: 注解定义全局异常处理类
2. @ExceptionHandler: 注解声明异常处理方法

### 8. 事务 @Transactional
> 我们知道 Exception 分为运行时异常 RuntimeException和非运行时异常. 在@Transactional注解中如果不配置rollbackFor属性, 那么事物只会在遇到RuntimeException的时候才会回滚, 加上rollbackFor=Exception.class, 可以让事物在遇到非运行时异常时也回滚, @Transactional 注解一般用在可以作用在类或者方法上.

1. 作用于类: 当把@Transactional 注解放在类上时, 表示所有该类的public 方法都配置相同的事务属性信息.
2. 作用于方法: 当类配置了@Transactional, 方法也配置了@Transactional，方法的事务会覆盖类的事务配置信息.

```java
@Transactional(rollbackFor = Exception.class)
public void save() {
  ......
}
```

### 9. Json 数据处理

#### @JsonIgnoreProperties
> 作用在类上用于过滤掉特定字段不返回或者不解析.

#### @JsonIgnore
> 一般用于类的属性上, 作用和上面的@JsonIgnoreProperties 一样.

```java
//生成json时将userRoles属性过滤
@JsonIgnoreProperties({"userRoles"})
public class User {

    private String userName;
    private String fullName;
    private String password;
    @JsonIgnore
    private List<UserRole> userRoles = new ArrayList<>();
}
```

#### @JsonFormat
> 一般用来格式化 json 数据

```java
@JsonFormat(shape=JsonFormat.Shape.STRING, pattern="yyyy-MM-dd'T'HH:mm:ss.SSS'Z'", timezone="GMT")
private Date date;
```

#### @JsonUnwrapped
> 扁平化对象


### 10. 测试相关
1. @ActiveProfiles: 一般作用于测试类上, 用于声明生效的 Spring 配置文件
2. @Test声明一个方法为测试方法
3. @Transactional被声明的测试方法的数据会回滚, 避免污染测试数据
4. @WithMockUser Spring Security 提供的, 用来模拟一个真实用户, 并且可以赋予权限
```java
@SpringBootTest(webEnvironment = RANDOM_PORT)
@ActiveProfiles("test")
@Slf4j
public abstract class TestBase {
  ......
}

@Test
@Transactional
@WithMockUser(username = "user-id-18163138155", authorities = "ROLE_TEACHER")
void should_import_student_success() throws Exception {
    ......
}
```

### 11. Lombok
1. @NonNull: 不能为空, 如果为空会直接抛出NullPointerException
2. @Getter|@Setter: 直接生产成员变量的set和get方法
3. @ToString: 覆盖默认的toString()方法, 将呈现类的基本信息
4. @EqualsAndHashCode: 只用于类, 覆盖默认的equals和hashCode
5. @NoArgsConstructor: 会生成一个无参的构造函数
6. @RequiredArgsConstructor: 使用该注解会对final和@NonNull字段生成对应参数的构造函数
7. @AllArgsConstructor: 生成了全参数的构造函数
8. @Builder: 如果你熟悉建造者模式