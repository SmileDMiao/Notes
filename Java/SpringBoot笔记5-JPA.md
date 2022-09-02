### Entity
---
1. @Entity: 声明一个类对应一个数据库实体
2. @Table: 设置表名
3. @ID: 声明一个字段为主键
4. @GeneratedValue: 直接使用 JPA 内置提供的四种主键生成策略来指定主键生成策略
5. @Column: 声明字段
6. @Transient: 指定不持久化特定字段
7. @Enumerated: 枚举
8. @EnableJpaAuditing: 开启 JPA 审计功能
9. @Modifying: 注解提示 JPA 该操作是修改操作, 注意还要配合@Transactional注解使用
10. @MappedSuperclass: 映射父类, 就是用来标识父类实体类
11. @EntityListeners: 注解于Entity上, 关联event callback的实现类

### 关联关系
---
> 关系声明

1. @OneToOne: 声明一对一关系
2. @OneToMany: 声明一对多关系
3. @ManyToOne: 声明多对一关系
4. @ManyToMany: 声明多对多关系

```java
public class User extends BaseModel {

    // @JoinTable注解的name属性指定关联表的名字, joinColumns指定外键的名字，关联到关系维护端(User)
    @JsonIgnoreProperties(value = {"hibernateLazyInitializer", "handler"})
    @ManyToOne(targetEntity = Role.class, fetch = FetchType.LAZY)
    @JoinColumn(name = "role_id", referencedColumnName = "id")
    private Role role;

    // inverseJoinColumns指定外键的名字 要关联的关系被维护端(permission)
    @JsonIgnoreProperties(value = {"hibernateLazyInitializer", "handler"})
    @ManyToMany(fetch = FetchType.LAZY)
    @JoinTable(
            name = "role_permissions",
            joinColumns = @JoinColumn(name = "id", referencedColumnName = "role_id"),
            inverseJoinColumns = @JoinColumn(name = "permission_id", referencedColumnName = "id")
    )
}
```

###  Repository接口
---
> Repository 接口是 Spring Data 的一个核心接口, 它不提供任何方法, 是一个空接口, 即是一个标记接口. 开发者需要在自己定义的接口中声明需要的方法. Spring Data可以让我们只定义接口, 只要遵循Spring Data的规范, 就无需写实现类. 若我们定义的接口继承了Repository, 则该接口会被IOC容器识别为一个Repository Bean, 纳入到IOC容器中, 进而可以在该接口中定义满足一定规范的方法. 实际上在IOC容器中放的是该接口的实现类, 只不过spring帮我们实现了, 实际上它是一个代理.

`Repository 的子接口`
1. Repository: 仅仅是一个标识, 表明任何继承它的均为仓库接口类.
2. CrudRepository: 继承Repository, 实现了一组CRUD相关的方法.
3. PagingAndSortingRepository: 继承 CrudRepository, 实现了一组分页排序相关的方法.
4. JpaRepository: 继承PagingAndSortingRepository, 实现一组JPA规范相关的方法.
5. 自定义的 xxxRepository: 需要继承 JpaRepository, 这样的XxxxRepository接口就具备了通用的数据访问控制层的能力.
6. JpaSpecificationExecutor: 不属于Repository体系, 实现一组JPACriteria 查询相关的方法.

### 条件查询
---
```java
public interface UserRepository extends JpaRepository<User, Long>, JpaSpecificationExecutor<User> {

    public List<User> findByUsername(String name);

    public List<User> findByUsernameLike(String name);

}

public Page<User> findByCondition(Integer page, Integer size, String sort_name, String username, String phone, String email) {
        Sort sort = Sort.by(Sort.Direction.DESC, sort_name);
        Pageable pageable = PageRequest.of(1, 10, sort);

        return userRepository.findAll((root, criteriaQuery, criteriaBuilder) -> {
            List<Predicate> predicates = new ArrayList<Predicate>();

            if (!StrUtil.hasBlank(username)) {
                predicates.add(criteriaBuilder.equal(root.get("username"), username));
            }

            if (!StrUtil.hasBlank(phone)) {
                predicates.add(criteriaBuilder.like(root.get("phone"), phone));
            }

            if (!StrUtil.hasBlank(email)) {
                predicates.add(criteriaBuilder.equal(root.get("email"), email));
            }

            return criteriaQuery.where(predicates.toArray(new Predicate[predicates.size()])).getRestriction();
        }, pageable);
}

userRepository.findAll();
userRepository.findById(1L);
userRepository.findByUsername("Hello");
userRepository.findByUsernameLike("Hello");
```
### 排序&分页
---
```java
Sort sort = Sort.by(Sort.Direction.DESC, sort_name);
Pageable pageable = PageRequest.of(0, 10, sort)
```

### JOIN查询
---
```java
userService.jpaFindByJoin(1, 10, "createdAt", "malzahar", "aa");
public Page<User> jpaFindByJoin(Integer page, Integer size, String sort_name, String username, String roleName) {
      Sort sort = Sort.by(Sort.Direction.DESC, sort_name);
      Pageable pageable = PageRequest.of(0, 10, sort);

      return userRepository.findAll((root, criteriaQuery, criteriaBuilder) -> {
          List<Predicate> predicates = new ArrayList<>();

          Join<User, Role> roleJoin = root.join("role", JoinType.LEFT);


          if (!StrUtil.hasBlank(username)) {
              predicates.add(criteriaBuilder.equal(root.get("username"), username));
          }

          if (!StrUtil.hasBlank(roleName)) {
              predicates.add(criteriaBuilder.equal(roleJoin.get("name"), roleName));
          }

          return criteriaQuery.where(predicates.toArray(new Predicate[0])).getRestriction();
      }, pageable);
}
```

### Entity Callback
1. @PrePersist:	可以理解为新增之前的回调方法(before create)
2. @PostPersist: 可以理解为在保存到数据库之后进行调用(after save)
3. @PreRemote:	可以理解为在操作删除方法之前调用(before destroy)
4. @PostRemote: 可以理解为在删除方法操作之后调用(after destroy)
5. @PreUpdate:	可以理解为在变化储存到数据库之前调用(before save)
6. @PostUpdate:	在实体更新之后调用(after save)
7. @PostLoad: 在实体从 DB 加载到程序里面之后回调(after read)

### N+1
---
1. @NamedEntityGraph
2. @EntityGraph

```java
@NamedEntityGraph(name = "user.role", attributeNodes = {@NamedAttributeNode("role")})
public class User extends BaseEntity {
    @JsonIgnoreProperties(value = {"hibernateLazyInitializer", "handler"})
    @ManyToOne(targetEntity = Role.class, fetch = FetchType.LAZY)
    @JoinColumn(name = "role_id", referencedColumnName = "id")
    @ToString.Exclude
    private Role role;
}

public interface UserRepository extends JpaRepository<User, Long>, JpaSpecificationExecutor<User> {

    @EntityGraph(value = "user.role")
    List<User> findByUsername(String name);
}
```

### 注解 方法约定
---
![IMAGE](resources/3CC921ADE68860D27C8CF3C4D5C454AD.jpg =740x665)
![IMAGE](resources/7D3769C505675FC64887768BBE116AD8.jpg =747x171)
![IMAGE](resources/9D483B1A22219289D07B1260F7B9EDE7.jpg =747x441)
![IMAGE](resources/893D4378FF5E0430224696802246AF89.jpg =747x677)