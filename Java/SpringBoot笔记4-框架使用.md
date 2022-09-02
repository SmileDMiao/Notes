## SpringBoot注入方式
---
> Field 注入

```java
@Controller
public class HelloController {
    @Autowired
    private AlphaService alphaService;
    @Autowired
    private BetaService betaService;
}
```

> Setter 方法注入

```java
@Controller
public class HelloController {
    private AlphaService alphaService;
    private BetaService betaService;
    
    @Autowired
    public void setAlphaService(AlphaService alphaService) {
        this.alphaService = alphaService;
    }
    @Autowired
    public void setBetaService(BetaService betaService) {
        this.betaService = betaService;
    }
}
```

> 构造器注入 

```java
@Controller
public class HelloController {
    private final AlphaService alphaService;
    private final BetaService betaService;
    
    @Autowired
    public HelloController(AlphaService alphaService, BetaService betaService) {
        this.alphaService = alphaService;
        this.betaService = betaService;
    }
}
```

## Filter(过滤器) And Interceptor(拦截器)
---
> 区别

1. 过滤器和拦截器触发时机不一样, 过滤器是在请求进入容器后, 但请求进入servlet之前进行预处理的. 请求结束返回也是, 是在servlet处理完后, 返回给前端之前.
2. 拦截器可以获取IOC容器中的各个bean, 而过滤器就不行, 因为拦截器是spring提供并管理的, spring的功能可以被拦截器使用, 在拦截器里注入一个service, 可以调用业务逻辑. 而过滤器是JavaEE标准, 只需依赖servlet api, 不需要依赖spring.
3. 过滤器的实现基于回调函数, 而拦截器(代理模式)的实现基于反射
4. Filter是依赖于Servlet容器, 属于Servlet规范的一部分, 而拦截器则是独立存在的, 可以在任何情况下使用.
5. Filter的执行由Servlet容器回调完成, 而拦截器通常通过动态代理(反射)的方式来执行.
6. Filter的生命周期由Servlet容器管理, 而拦截器则可以通过IoC容器来管理, 因此可以通过注入等方式来获取其他Bean的实例, 因此使用会更方便.
7. 过滤器(Filter): 可以拿到原始的http请求, 但是拿不到你请求的控制器和请求控制器中的方法的信息
8. 拦截器(Interceptor): 可以拿到你请求的控制器和方法, 却拿不到请求方法的参数.
9. 切片(Aspect): 可以拿到方法的参数, 但是却拿不到http请求和响应的对象



#### 过滤器
---
> 两种方式的本质都是一样的, 都是去FilterRegistrationBean注册自定义Filter

1. 使用spring boot提供的FilterRegistrationBean注册Filter
2. 使用原生servlet注解定义Filter 

> 方法

1. init方法: 是过滤器的初始化方法, 当web容器创建这个bean的时候就会执行
2. doFilter方法: 是执行过滤的请求的核心, 当客户端请求访问web资源时, 这个时候我们可以拿到request里面的参数, 对数据进行处理后, 通过filterChain方法将请求将请求放行
3. destory方法: 是当web容器中的过滤器实例被销毁时会被执行, 主要作用是释放资源。


```java
public class TimeFilter implements Filter {

    @Override
    public void init(FilterConfig filterConfig) {
        System.out.println("过滤器初始化");
    }

    @Override
    public void doFilter(ServletRequest servletRequest, ServletResponse servletResponse, FilterChain filterChain) throws IOException, ServletException {
        System.out.println("过滤器执行了");
        long start2 = System.currentTimeMillis();
        filterChain.doFilter(servletRequest, servletResponse);
    }

    @Override
    public void destroy() {
        System.out.println("过滤器销毁了");
    }
}
```


#### 拦截器
---
> 说明

1. 如果preHandle方法return true, 则继续后续处理.
2. preHandle是请求执行前执行的
3. postHandler是请求结束执行的, 但只有preHandle方法返回true的时候才会执行
4. afterCompletion是视图渲染完成后才执行, 同样需要preHandle返回true, 该方法通常用于清理资源等工作
```java
@Component
public class TraceSetup implements HandlerInterceptor {

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
        MDC.put("TraceId", UUID.randomUUID().toString());
        return true;
    }

    @Override
    public void afterCompletion(HttpServletRequest request, HttpServletResponse response, Object handler, Exception ex) throws Exception {
        MDC.clear();
    }
}

// 配置
@SpringBootConfiguration
public class MyWebMvcConfigurerAdapter implements WebMvcConfigurer {
    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        registry.addInterceptor(new TraceSetup());
    }
}
```

## 数据库连接池
---
```yml
spring
  datasource:
  driver-class-name: com.mysql.cj.jdbc.Driver
  url: jdbc:mysql://localhost:3306/java
  username: root
  password: 123
  hikari:
    minimum-idle: 5
    maximum-pool-size: 20
    auto-commit: true
    idle-timeout: 30000
    pool-name: SpringBootHikariCP
    max-lifetime: 60000
    connection-timeout: 30000
  druid:
    # 连接池初始化大小
    initial-size: 5
    # 最小空闲连接数
    min-idle: 10
    # 最大连接数
    max-active: 20
    web-stat-filter:
      enabled: true
      # 不统计这些请求数据
      exclusions: "*.js,*.gif,*.jpg,*.png,*.css,*.ico,/druid/*"
    stat-view-servlet:
      enabled: true
      # 访问监控网页的登录用户名和密码
      login-username: druid
      login-password: druid
```

## Actuator
---
```yml
management:
  endpoint:
    health:
      show-details: always
  endpoints:
    web:
      exposure:
        include: '*'
```

```shell
/actuator/health
/actuator/metrics
/actuator/metrics/jvm.memory.max
/actuator/loggers
/actuator/configprops
```