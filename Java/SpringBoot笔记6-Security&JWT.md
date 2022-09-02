## 入口配置
---
```java
public class SecurityConfig {
    SecurityFilterChain filterChain(HttpSecurity httpSecurity) throws Exception {
            ExpressionUrlAuthorizationConfigurer<HttpSecurity>.ExpressionInterceptUrlRegistry registry = httpSecurity.authorizeRequests();
    
            // 不需要保护的资源路径允许访问
            for (String url : ignoreUrlConfig.getUrls()) {
                registry.antMatchers(url).permitAll();
            }
    
            // 允许跨域请求的OPTIONS请求
            registry.antMatchers(HttpMethod.OPTIONS).permitAll();
    
            // 由于使用的是JWT，我们这里不需要csrf
            // 基于token，所以不需要session
            // 除上面外的所有请求全部需要鉴权认证
            registry.and()
                    .csrf()
                    .disable()
                    .sessionManagement()
                    .sessionCreationPolicy(SessionCreationPolicy.STATELESS)
                    .and()
                    .authorizeRequests()
                    .anyRequest()
                    .authenticated()
                    .and()
                    .exceptionHandling()
                    // 权限不足处理
                    .accessDeniedHandler(requestAccessDeniedHandler)
                    // 认证失败处理
                    .authenticationEntryPoint(requestUnauthorized)
                    .and()
                    // Token Validate
                    .addFilterBefore(jwtAuthenticationTokenFilter, UsernamePasswordAuthenticationFilter.class)
                    // Permission Check
                    .addFilterBefore(permissionSecurityFilter, FilterSecurityInterceptor.class);
    
            return httpSecurity.build();
    }
}
```

## Spring Security 封装的用户信息 权限信息
---
```java
public class CurrentDetails implements UserDetails {
    private final User user;

    private final List<Permission> permissions;

    // 用户信息 权限信息
    public CurrentDetails(User user, List<Permission> permissions) {
        this.user = user;
        this.permissions = permissions;
    }

    // 将用户权限信息转换为 Spring Security中的权限信息
    @Override
    public Collection<? extends GrantedAuthority> getAuthorities() {
        return permissions.stream()
                .map(permission -> new SimpleGrantedAuthority(permission.getName()))
                .collect(Collectors.toList());
    }
}
```

## Token Filter
---
```java
public class JwtAuthenticationTokenFilter extends OncePerRequestFilter {

    @Override
    protected void doFilterInternal(HttpServletRequest request,
                                    HttpServletResponse response,
                                    FilterChain chain) throws ServletException, IOException {
        String authHeader = request.getHeader(this.tokenHeader);

        if (authHeader != null && authHeader.startsWith(this.tokenHead)) {

            if (username != null && SecurityContextHolder.getContext().getAuthentication() == null) {
                UsernamePasswordAuthenticationToken authentication = new UsernamePasswordAuthenticationToken(userDetails, null, userDetails.getAuthorities());
                authentication.setDetails(new WebAuthenticationDetailsSource().buildDetails(request));

                // 全局注入角色权限信息和登录用户基本信息
                SecurityContextHolder.getContext().setAuthentication(authentication);
            }
        }

        chain.doFilter(request, response);
    }
}
```

## 权限检查
---
1. PermissionSecurityMetadataSource: 请求访问需要的权限
2. PermissionSecurityFilter: 调用PermissionAccessDecisionManager
3. PermissionAccessDecisionManager: 判断权限

```java
public class PermissionSecurityMetadataSource implements FilterInvocationSecurityMetadataSource {

    @Override
    public Collection<ConfigAttribute> getAttributes(Object o) throws IllegalArgumentException {
        List<ConfigAttribute> configAttributes = new ArrayList<>();

        // 获取当前访问的路径
        String url = ((FilterInvocation) o).getRequestUrl();
        String path = URLUtil.getPath(url);
        PathMatcher pathMatcher = new AntPathMatcher();

        // 获取所有权限
        List<Permission> permissions = permissionMapper.selectList();

        // 转换成 Spring Security 权限
        Map<String, ConfigAttribute> configAttributeMap = new ConcurrentHashMap<>();
        for (Permission permission : permissions) {
            configAttributeMap.put(permission.getUrl(), new org.springframework.security.access.SecurityConfig(permission.getName()));
        }

        // 获取访问该路径所需资源
        for (String pattern : configAttributeMap.keySet()) {
            if (pathMatcher.match(pattern, path)) {
                configAttributes.add(configAttributeMap.get(pattern));
            }
        }

        // 为空就不走 PermissionAccessDecisionManager#decide方法了
        configAttributes.add(new org.springframework.security.access.SecurityConfig("Fake Permission"));

        return configAttributes;
    }
}
```

```java
public class PermissionAccessDecisionManager implements AccessDecisionManager {
    @Override
    public void decide(Authentication authentication,
                       Object object,
                       Collection<ConfigAttribute> configAttributes) throws AccessDeniedException, InsufficientAuthenticationException {
        // 当接口未被配置资源时直接放行
        if (CollUtil.isEmpty(configAttributes)) {
            return;
        }

        for (ConfigAttribute configAttribute : configAttributes) {
            // 将访问所需资源和用户拥有资源进行比对
            String needAuthority = configAttribute.getAttribute();

            for (GrantedAuthority grantedAuthority : authentication.getAuthorities()) {
                if (needAuthority.trim().equals(grantedAuthority.getAuthority())) {
                    return;
                }
            }
        }
        throw new AccessDeniedException("抱歉，您没有访问权限");
    }
}
```

## TIPS
---
1. accessDeniedHandler: 当用户已经登陆且没有权限才会被调用
2. getAttributes: 返回空就不会调用decide方法