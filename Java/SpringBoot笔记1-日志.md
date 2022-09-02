## Logback
---

### 配置项
> contextName

每个logger都关联到logger上下文, 默认上下文名称为 default. 但可以使用contextName标签设置成其他名字, 用于区分不同应用程序的记录

> property

用来定义变量值的标签, property标签有两个属性, name和value. 其中name的值是变量的名称, value的值时变量定义的值. 通过property定义的值会被插入到logger上下文中. 定义变量后, 可以使 "${name}" 来使用变量

> root

根logger, 也是一种logger, 且只有一个level属性

> logger

用来设置某一个包或者具体的某一个类的日志打印级别以及指定appender

> appender

`appender 是一个日志打印的组件, 这里组件里面定义了打印过滤的条件, 打印输出方式, 滚动策略, 编码方式, 打印格式等等. 但是它仅仅是一个打印组件, 如果我们不使用一个 logger 或者 root 的 appender-ref 指定某个具体的 appender 时, 它就没有什么意义. appender让我们的应用知道怎么打, 打印到哪里, 打印成什么样. 而logger则是告诉应用哪些可以这么打. 例如某个类下的日志可以使用这个appender打印或者某个包下的日志可以这么打印.`
1. 负责写日志的组件
2. filter: 是appender里面的子元素. 它作为过滤器存在
3. filter: 执行一个过滤器会有返回DENY, NEUTRAL, ACCEPT三个枚举值中的一个
4. DENY: 日志将立即被抛弃不再经过其他过滤器. NEUTRAL: 有序列表里的下个过滤器过接着处理日志. ACCEPT: 日志会被立即处理, 不再经过剩余过滤器
5. RollingFileAppender: 滚动记录文件, 先将日志记录到指定文件, 当符合某个条件时, 将日志记录到其他文件. 它是FileAppender的子类
6. . ConsoleAppender：把日志添加到控制台
7. . FileAppender：把日志添加到文件

> LEVEL

1. TRACE: 在线调试, 该级别日志, 默认情况下, 既不打印到终端也不输出到文件
2. DEBUG: 终端查看, 在线调试. 该级别日志, 默认情况下会打印到终端输出, 但是不会归档到日志文件. 因此, 一般用于开发者在程序当前启动窗口上, 查看日志流水信息.
3. INFO: 报告程序进度和状态信息, 一般这种信息都是一过性的, 不会大量反复输出
4. WARN: 警告信息, 程序处理中遇到非法数据或者某种可能的错误. 该错误是一过性的, 可恢复的, 不会影响程序继续运行, 程序仍处在正常状态
5. ERROR: 状态错误, 该错误发生后程序仍然可以运行, 但是极有可能运行在某种非正常的状态下, 导致无法完成全部既定的功能
6. ALL: 所有
7. OFF: 关闭

### 输出JSON
---
```xml
<!--encoder: net.logstash.logback.encoder.LoggingEventCompositeJsonEncoder-->
<dependency>
	<groupId>net.logstash.logback</groupId>
	<artifactId>logstash-logback-encoder</artifactId>
	<version>7.2</version>
</dependency>
```

#### TraceID
```java
// MDC
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

@SpringBootConfiguration
public class MyWebMvcConfigurerAdapter implements WebMvcConfigurer {
    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        registry.addInterceptor(new TraceSetup());
    }
}
```

### SQL LOG
---
```xml
<!--spring-boot-data-source-decorator-->
<dependency>
  	<groupId>com.github.gavlyukovskiy</groupId>
	  <artifactId>p6spy-spring-boot-starter</artifactId>
	  <version>1.8.0</version>
</dependency>
```

```java
// 自定义日志
// 默认com.p6spy.engine.spy.appender.SingleLineFormat
public class P6SpyLogger implements MessageFormattingStrategy {
    /**
     * 日志格式
     * @param connectionId 连接id
     * @param now 当前时间
     * @param elapsed 耗时多久
     * @param category 类别
     * @param prepared mybatis带占位符的sql
     * @param sql 占位符换成参数的sql
     * @param url sql连接的 url
     * @return 自定义格式日志
     */
    @Override
    public String formatMessage(int connectionId, String now, long elapsed, String category, String prepared, String sql, String url) {
        return !"".equals(sql.trim()) ? "P6SpyLogger " + LocalDateTime.now() + " | elapsed " + elapsed + "ms | category " + category + " | connection " + connectionId + " | url " + url + " | sql \n" + sql : "";
    }
}

```

### 异步日志: ch.qos.logback.classic.AsyncAppender
---
```xml
<appender name="ALL" class="ch.qos.logback.core.rolling.RollingFileAppender">
    <file>${logDir}/all.log</file>
    <encoder>
        <pattern>${commonPattern}</pattern>
    </encoder>
    <rollingPolicy class="ch.qos.logback.core.rolling.TimeBasedRollingPolicy">
        <FileNamePattern>${logDir}/all.%d{yyyy-MM-dd}.log.zip</FileNamePattern>
        <maxHistory>15</maxHistory>
    </rollingPolicy>
</appender>
<appender name="ASYNC" class="ch.qos.logback.classic.AsyncAppender">
    <includeCallerData>true</includeCallerData>
    <discardingThreshold>-1</discardingThreshold>
    <queueSize>1024</queueSize>
    <appender-ref ref="ALL" />
</appender>
```

### Example Configuration File
---
```xml
<?xml version="1.0" encoding="UTF-8"?>
<configuration>
    <include resource="org/springframework/boot/logging/logback/defaults.xml"/>
    <springProperty scope="context" name="springAppName" source="spring.application.name"/>
    <property name="LOG_FILE_PATH" value="logs/logback/"/>
    <appender name="CONSOLE" class="ch.qos.logback.core.ConsoleAppender">
        <filter class="ch.qos.logback.classic.filter.LevelFilter">
            <level>DEBUG</level>
        </filter>
        <encoder>
            <pattern>%date [%thread] %-5level [%logger{50}] %file:%line - %msg%n</pattern>
            <charset>UTF-8</charset>
        </encoder>
    </appender>

    <appender name="FILE_INFO" class="ch.qos.logback.core.rolling.RollingFileAppender">
        <!--日志名称，如果没有File 属性，那么只会使用FileNamePattern的文件路径规则如果同时有<File>和<FileNamePattern>，那么当天日志是<File>，明天会自动把今天的日志改名为今天的日期。即，<File> 的日志都是当天的。-->
        <!--滚动策略，按照时间滚动 TimeBasedRollingPolicy-->
        <rollingPolicy class="ch.qos.logback.core.rolling.TimeBasedRollingPolicy">
            <!--文件路径,定义了日志的切分方式——把每一天的日志归档到一个文件中,以防止日志填满整个磁盘空间-->
            <FileNamePattern>${LOG_FILE_PATH}/knight-%d{yyyy-MM-dd}-part-%i.log</FileNamePattern>
            <!--只保留最近90天的日志-->
            <maxHistory>90</maxHistory>
            <!--用来指定日志文件的上限大小，那么到了这个值，就会删除旧的日志-->
            <!--<totalSizeCap>1GB</totalSizeCap>-->
            <timeBasedFileNamingAndTriggeringPolicy class="ch.qos.logback.core.rolling.SizeAndTimeBasedFNATP">
                <!-- maxFileSize:这是活动文件的大小，默认值是10MB,本篇设置为1KB，只是为了演示 -->
                <maxFileSize>2MB</maxFileSize>
            </timeBasedFileNamingAndTriggeringPolicy>
        </rollingPolicy>

        <encoder class="net.logstash.logback.encoder.LoggingEventCompositeJsonEncoder">
            <providers>
                <timestamp>
                    <fieldName>timestamp</fieldName>
                    <timeZone>UTC+8</timeZone>
                    <pattern>yyyy-MM-dd HH:mm:ss</pattern>
                </timestamp>
                <pattern>
                    <pattern>
                        {
                        "trace_id": "%X{TraceId}",
                        "app": "${springAppName}",
                        "level": "%level",
                        "pid": "${PID:-}",
                        "thread": "%thread",
                        "class": "%logger",
                        "message": "%message",
                        "stack_trace": "%exception{20}"
                        }
                    </pattern>
                </pattern>
            </providers>
        </encoder>
    </appender>

    <appender name="API_FILE_INFO" class="ch.qos.logback.core.rolling.RollingFileAppender">
        <!--日志名称，如果没有File 属性，那么只会使用FileNamePattern的文件路径规则如果同时有<File>和<FileNamePattern>，那么当天日志是<File>，明天会自动把今天的日志改名为今天的日期。即，<File> 的日志都是当天的。-->
        <!--滚动策略，按照时间滚动 TimeBasedRollingPolicy-->
        <rollingPolicy class="ch.qos.logback.core.rolling.TimeBasedRollingPolicy">
            <!--文件路径,定义了日志的切分方式——把每一天的日志归档到一个文件中,以防止日志填满整个磁盘空间-->
            <FileNamePattern>${LOG_FILE_PATH}/API-%d{yyyy-MM-dd}-part-%i.log</FileNamePattern>
            <!--只保留最近90天的日志-->
            <maxHistory>90</maxHistory>
            <!--用来指定日志文件的上限大小，那么到了这个值，就会删除旧的日志-->
            <!--<totalSizeCap>1GB</totalSizeCap>-->
            <timeBasedFileNamingAndTriggeringPolicy class="ch.qos.logback.core.rolling.SizeAndTimeBasedFNATP">
                <!-- maxFileSize:这是活动文件的大小，默认值是10MB,本篇设置为1KB，只是为了演示 -->
                <maxFileSize>2MB</maxFileSize>
            </timeBasedFileNamingAndTriggeringPolicy>
        </rollingPolicy>

        <encoder class="net.logstash.logback.encoder.LoggingEventCompositeJsonEncoder">
            <providers>
                <timestamp>
                    <fieldName>timestamp</fieldName>
                    <timeZone>UTC+8</timeZone>
                    <pattern>yyyy-MM-dd HH:mm:ss</pattern>
                </timestamp>
                <pattern>
                    <pattern>
                        {
                        "trace_id": "%X{TraceId}",
                        "app": "${springAppName}",
                        "level": "%level",
                        "pid": "${PID:-}",
                        "thread": "%thread",
                        "class": "%logger",
                        "message": "%message",
                        "stack_trace": "%exception{20}"
                        }
                    </pattern>
                </pattern>
            </providers>
        </encoder>
    </appender>

    <!--控制框架输出日志-->
    <logger name="org.slf4j" level="INFO"/>
    <logger name="springfox" level="INFO"/>
    <logger name="io.swagger" level="INFO"/>
    <logger name="org.springframework" level="INFO"/>
    <logger name="org.hibernate.validator" level="INFO"/>

    <logger name="com.knight.javaPractice.controller.concern.AopLog" level="INFO">
        <appender-ref ref="API_FILE_INFO"/>
    </logger>

    <logger name="p6spy" level="INFO">
        <appender-ref ref="API_FILE_INFO"/>
    </logger>


    <root level="INFO">
        <appender-ref ref="CONSOLE"/>
    </root>
</configuration>
```