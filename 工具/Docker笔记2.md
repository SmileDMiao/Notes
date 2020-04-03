## Docker Compose

### 模版文件
**build:**
指定 Dockerfile 所在文件夹的路径。你也可以使用 context 指令指定 Dockerfile 所在文件夹的路径。使用 dockerfile 指令指定 Dockerfile 文件名。

**depends_on:**
解决容器的依赖、启动先后的问题。先启动依赖服务。

**ports:**
暴露端口信息。使用宿主端口：容器端口 (HOST:CONTAINER) 格式

**volumes:**
数据卷所挂载路径设置.

**restart:**
指定容器退出后的重启策略为始终重启，该命令对保持服务始终运行十分有效，在生产环境中推荐配置为 always 或者 unless-stopped

**environment:**
设置环境变量。你可以使用数组或字典两种格式。
```shell
environment:
  RACK_ENV: development
  SESSION_SECRET:

environment:
  - RACK_ENV=development
  - SESSION_SECRET
```