### XML映射文件标签
---
1. select: 映射查询语句
2. insert: 映射插入语句
3. delete: 映射删除语句
4. update: 映射更新语句
5. resultMap: 描述如何从数据库结果集中加载对象, 是最复杂也是最强大的元素
6. cache: 该命名空间的缓存配置
7. sql: 可被其它语句引用的可重用语句块

#### resultType
> 期望从这条语句中返回结果的类全限定名或别名. 如果返回的是集合, 那应该设置为集合包含的类型, 而不是集合本身的类型. resultType 和 resultMap 之间只能同时使用一个.

##### resultMap
> 对外部 resultMap 的命名引用. resultType 和 resultMap 之间只能同时使用一个。

#### resultSets
> 这个设置仅适用于多结果集的情况, 它将列出语句执行后返回的结果集并赋予每个结果集一个名称, 多个名称之间以逗号分隔.

#### useGeneratedKeys
> 仅适用于 insert 和 update, 这会令 MyBatis 使用 JDBC 的 getGeneratedKeys 方法来取出由数据库内部生成的主键

#### keyProperty
> 仅适用于 insert 和 update 指定能够唯一识别对象的属性, MyBatis 会使用 getGeneratedKeys 的返回值或 insert 语句的 selectKey 子元素设置它的值, 默认值: 未设置（unset)


### `#{}` 和 `${}`
---
1. `${}` 是 Properties 文件中的变量占位符, 它可以用于标签属性值和 sql 内部, 属于静态文本替换, 比如${driver}会被静态替换为com.mysql.jdbc. Driver
2. `#{}` 是 sql 的参数占位符, MyBatis 会将 sql 中的#{}替换为 ? 号, 在 sql 执行前会使用 PreparedStatement 的参数设置方法, 按序给 sql 的 ? 号占位符设置参数值

### Mapper接口方法参数不同时, 时候可以重载?
---
> 可以重载, 但是 Mybatis 的 XML 里面的 ID 不允许重复

```java
public interface StuMapper {

	List<Student> getAllStu();

	List<Student> getAllStu(@Param("id") Integer id);
}
```

```xml
<select id="getAllStu" resultType="com.pojo.Student">
 		select * from student
		<where>
			<if test="id != null">
				id = #{id}
			</if>
		</where>
</select>
```

### 拦截器(Plugin)
---
> mybatis 插件在四大组件 (Executor, StatementHandler, ParameterHandler, ResultSetHandler) 处提供了简单易用的插件扩展机制. mybatis 对持久层的操作就是借助于四大核心对象.
> myBatis 所允许拦截的方法如下: 

1. 执行器: Executor (update, query, commit, rollback 等方法)
2. SQL语法构建器: StatementHandler (prepare, parameterize, batch, update, query 等方法)
3. 参数处理器: ParameterHandler (getParameterObject, setParameters 方法)
4. 结果集处理器: ResultSetHandler (handleResultSets, handleOutputParameters 等方法)

> Mybatis 插件接口: Interceptor

1. Intercept方法: 插件的核心方法
2. plugin方法: 生成target的代理对象
3. setProperties方法: 传递插件所需参数

```java
// plugin方法是拦截器用于封装目标对象的, 通过该方法我们可以返回目标对象本身, 也可以返回一个它的代理.
// 当返回的是代理的时候我们可以对其中的方法进行拦截来调用intercept方法, 当然也可以调用其他方法.
// setProperties方法是用于在Mybatis配置文件中指定一些属性的。
public interface Interceptor {

  Object intercept(Invocation invocation) throws Throwable;

  Object plugin(Object target);

  void setProperties(Properties properties);

}
```

```java
@Component("MybatisAuditInterceptor")
@Intercepts({@Signature(method = "update", type = Executor.class, args = {MappedStatement.class, Object.class})})
public class MybatisAuditInterceptor implements Interceptor {
    @Override
    public Object intercept(Invocation invocation) throws Throwable {
    }
}

```

```xml
<plugins>
    <plugin interceptor="com.knight.javaPractice.Initializer.MybatisAuditInterceptor">
    </plugin>
</plugins>
```

### MyBatis的XML映射文件中, 不同的XML映射文件, id 是否可以重复?
---
> 不同的XML映射文件, 如果配置了 namespace, 那么 id 可以重复. 如果没有配置 namespace, 那么 id 不能重复. 原因就是 namespace+id 是作为 Map<String, MappedStatement> 的 key 使用的, 如果没有 namespace, 就剩下 id, id 重复会导致数据互相覆盖.


### 参数, 字符串替换
---
> 当 SQL 语句中的元数据(如表名或列名)是动态生成的时候, 字符串替换将会非常有用. 如果你想 select 一个表任意一列的数据时:

```java
// 多个方法
@Select("select * from user where id = #{id}")
User findById(@Param("id") long id);

@Select("select * from user where name = #{name}")
User findByName(@Param("name") String name);

@Select("select * from user where email = #{email}")
User findByEmail(@Param("email") String email);

// 一个方法
// 其中 ${column} 会被直接替换, 而 #{value} 会使用 ? 预处理
@Select("select * from user where ${column} = #{value}")
User findByColumn(@Param("column") String column, @Param("value") String value);

User userOfId1 = userMapper.findByColumn("id", 1L);
```

### 执行器
---
> 作用范围: Executor 的这些特点, 都严格限制在 SqlSession 生命周期范围内

1. SimpleExecutor: 每执行一次 update 或 select, 就开启一个 Statement 对象, 用完立刻关闭 Statement 对象
2. ReuseExecutor: 执行 update 或 select, 以 sql 作为 key 查找 Statement 对象, 存在就使用, 不存在就创建, 用完后, 不关闭 Statement 对象, 而是放置于 Map<String, Statement>内, 供下一次使用. 简言之, 就是重复使用 Statement 对象.
3. BatchExecutor: 执行 update(没有 select, JDBC 批处理不支持 select), 将所有 sql 都添加到批处理中(addBatch()), 等待统一执行(executeBatch()), 它缓存了多个 Statement 对象, 每个 Statement 对象都是 addBatch()完毕后, 等待逐一执行 executeBatch()批处理。与 JDBC 批处理相同。

```java
@Test
public void testBatchExecutor() throws SQLException {
    // 通过factory.openSession().getConnection()实例化JdbcTransaction ，用于构建BatchExecutor
    jdbcTransaction = new JdbcTransaction(factory.openSession().getConnection());

    // 实例化BatchExecutor
    BatchExecutor executor = new BatchExecutor(configuration, jdbcTransaction);

    // 映射SQL
    ms = configuration.getMappedStatement("com.artisan.UserMapper.updateById");

    Map map = new HashMap();
    map.put("arg0",1);
    map.put("arg1","222");

    // 调用doUpdate
    executor.doUpdate(ms,map);
    executor.doUpdate(ms,map);

    // 刷新
    executor.doFlushStatements(false);
    // 提交  否则不生效
    executor.commit(true);
 
}
```

### 游标查询
---
```java
public void test1() throws IOException {
      List<xxxBean> list= new ArrayList<>();
      SqlSession sqlSession =sqlSessionTemplate.getSqlSessionFactory().openSession();
      //90137L表示查询的参数
      Cursor<HdfsPathRecord> cursor = sqlSession.selectCursor(xxxDao.class.getName() + ".selectByAppId",90137L);
     Iterator iter = cursor.iterator();
      while (iter.hasNext()) {
        xxxBean obj = (xxxBean) iter.next();
       //TODO:接下来可以根据obj执行实际逻辑或者将obj放到一个list中
      }
      cursor.close();
      sqlSession.close();
}
```