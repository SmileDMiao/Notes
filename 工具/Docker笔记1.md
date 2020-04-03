基本概念
----

---

### 镜像Image

1. 操作系统分为内核和用户空间。对于 Linux 而言，内核启动后，会挂载 root 文件系统为其提供用户空间支持。而 Docker 镜像，就相当于是一个 root 文件系统。Docker镜像是一个特殊的文件系统，除了提供容器运行时所需的程序、库、资源、配置等文件外，还包含了一些为运行时准备的一些配置参数（如匿名卷、环境变量、用户等）。
2. 镜像不包含任何动态数据，其内容在构建之后也不会被改变。

### 容器Container

1. 镜像Image和容器Container的关系，就像是面向对象程序设计中的 类 和 实例 一样，镜像是静态的定义，容器是镜像运行时的实体。容器可以被创建、启动、停止、删除、暂停等。
2. 容器的实质是进程，但与直接在宿主执行的进程不同，容器进程运行于属于自己的独立的 命名空间。因此容器可以拥有自己的 root 文件系统、自己的网络配置、自己的进程空间，甚至自己的用户 ID 空间。容器内的进程是运行在一个隔离的环境里，使用起来，就好像是在一个独立于宿主的系统下操作一样。这种特性使得容器封装的应用比直接在宿主运行更加安全。
3. 按照 Docker 最佳实践的要求，容器不应该向其存储层内写入任何数据，容器存储层要保持无状态化。所有的文件写入操作，都应该使用数据卷 Volume 或者绑定宿主目录，在这些位置的读写会跳过容器存储层，直接对宿主或网络存储发生读写，其性能和稳定性更高。

Dockerfile reference
--------------------

---

**FROM: 指定基础镜像**

除了选择现有镜像为基础镜像外，Docker 还存在一个特殊的镜像，名为 scratch。这个镜像是虚拟的概念，并不实际存在，它表示一个空白的镜像。

**RUN:执行命令**

1. Dockerfile 中每一个指令都会建立一层，RUN 也不例外。每一个 RUN 的行为，就和刚才我们手工建立镜像的过程一样：新建立一层，在其上执行这些命令，执行结束后，commit 这一层的修改，构成新的镜像。
2. 有时候需要添加清理工作的命令，删除了为了编译构建所需要的软件，清理了所有下载、展开的文件，并且还清理了 apt 缓存文件。这是很重要的一步，我们之前说过，镜像是多层存储，每一层的东西并不会在下一层被删除，会一直跟随着镜像。因此镜像构建时，一定要确保每一层只添加真正需要添加的东西，任何无关的东西都应该清理掉。

**COPY: 从构建上下文目录中 \<源路径\> 的文件/目录复制到新的一层的镜像内的 \<目标路径\> 位置。**

1. \<目标路径\> 可以是容器内的绝对路径，也可以是相对于工作目录的相对路径（工作目录可以用 WORKDIR 指令来指定）。目标路径不需要事先创建，如果目录不存在会在复制文件前先行创建缺失目录。
2. 使用 COPY 指令，源文件的各种元数据都会保留。比如读、写、执行权限、文件变更时间等。

**ADD: **

**类似COPY，但还可以做多一点事情:**

`ADD http://example.com/big.tar.xz /usr/src/things/​`

1. 解压压缩文件并把它们添加到镜像中
2. 从 url 拷贝文件到镜像中

**CMD: 指定默认的容器主进程的启动命令**

1. 容器中的应用都应该以前台执行，而不是像虚拟机、物理机里面那样，用 systemd 去启动后台服务，容器内没有后台服务的概念。
2. 比如使用 service nginx start 命令，则是希望 upstart 来以后台守护进程形式启动 nginx 服务。而刚才说了 CMD service nginx start 会被理解为 CMD [ "sh", "-c", "service nginx start"]，因此主进程实际上是 sh。那么当 service nginx start 命令结束后，sh 也就结束了，sh 作为主进程退出了，自然就会令容器退出。正确的做法是直接执行 nginx 可执行文件，并且要求以前台形式运行。比如：CMD ["nginx", "-g", "daemon off;”]

**ENTRYPOINT: 入口点**

1. ENTRYPOINT 的目的和 CMD 一样，都是在指定容器启动程序及参数。ENTRYPOINT 在运行时也可以替代，不过比 CMD 要略显繁琐，需要通过 docker run 的参数 --entrypoint 来指定。
2. 当指定了 ENTRYPOINT 后，CMD 的含义就发生了改变，不再是直接的运行其命令，而是将 CMD 的内容作为参数传给 ENTRYPOINT 指令，换句话说实际执行时，将变为：
3. \<ENTRYPOINT\> "\<CMD\>”

**ENV: 设置环境变量
**

ENV key value

ENV key1=value1 key2=value2

**ARG: 构建参数**

构建参数和 ENV 的效果一样，都是设置环境变量。所不同的是，ARG 所设置的构建环境的环境变量，在将来容器运行时是不会存在这些环境变量的。但是不要因此就使用 ARG 保存密码之类的信息，因为 docker history 还是可以看到所有值的。

**VOLUME: 定义匿名卷**

1. 为了防止运行时用户忘记将动态文件所保存目录挂载为卷，在 Dockerfile 中，我们可以事先指定某些目录挂载为匿名卷，这样在运行时如果用户不指定挂载，其应用也可以正常运行，不会向容器存储层写入大量数据。
2. VOLUME /data 这里的 /data 目录就会在运行时自动挂载为匿名卷，任何向 /data 中写入的信息都不会记录进容器存储层，从而保证了容器存储层的无状态化。

**EXPOSE: 声明端口**

1. EXPOSE \<端口1\> [\<端口2\>] [\<port\>/\<protocol\>]
2. EXPOSE 指令是声明运行时容器提供服务端口。

**WORKDIR: 指定工作目录**

1. WORKDIR \<工作目录路径\>。
2. 使用 WORKDIR 指令可以来指定工作目录（或者称为当前目录），以后各层的当前目录就被改为指定的目录，如该目录不存在，WORKDIR 会帮你建立目录。

**USER 指定当前用户**

1. USER 指令和 WORKDIR 相似，都是改变环境状态并影响以后的层。WORKDIR 是改变工作目录，USER 则是改变之后层的执行 RUN, CMD 以及 ENTRYPOINT 这类命令的身份。
2. USER 只是帮助你切换到指定用户而已，这个用户必须是事先建立好的，否则无法切换。

**HEALTHCHECK: 健康检查**

1. HEALTHCHECK 指令是告诉 Docker 应该如何进行判断容器的状态是否正常
2. 和 CMD, ENTRYPOINT 一样，HEALTHCHECK 只可以出现一次，如果写了多个，只有最后一个生效。

**ONBUILD**

当以当前镜像为基础镜像，去构建下一级镜像的时候才会被执行。

Docker Command Line Interface
-----------------------------

---

### docker commit

### 根据容器的更改创建新的镜像

### docker build: 构建镜像

### **docker build [OPTIONS] PATH | URL | -**

当我们进行镜像构建的时候，并非所有定制都会通过 RUN 指令完成，经常会需要将一些本地文件复制进镜像，比如通过 COPY 指令、ADD 指令等。而 docker build 命令构建镜像，其实并非在本地构建，而是在服务端，也就是 Docker 引擎中构建的。那么在这种客户端/服务端的架构中，如何才能让服务端获得本地文件呢？这就引入了上下文的概念。当构建的时候，用户会指定构建镜像上下文的路径，docker build 命令得知这个路径后，会将路径下的所有内容打包，然后上传给 Docker 引擎。这样 Docker 引擎收到这个上下文包后，展开就会获得构建镜像所需的一切文件。

### docker attach: 进入容器

如果从这个 stdin 中 exit，会导致容器的停止。

### docker exec

进入容器：` docker exec -it container -id /bin/bash`

### 导入导出容器

`docker export  container-id >  container-id.tar`

`docker import remote-location-url`

`cat ubuntu.tar | docker import - ubuntu:new`

### docker volume: 数据卷

数据卷 是一个可供一个或多个容器使用的特殊目录，它绕过 UFS，可以提供很多有用的特性:

1. 数据卷 可以在容器之间共享和重用
2. 对 数据卷 的修改会立马生效
3. 对 数据卷 的更新，不会影响镜像
4. 数据卷 默认会一直存在，即使容器被删除

创建: docker volume create volume-name

查看: docker volume inspect volume-name

列表: docker volume ls

移除本地无用的数据卷: docker volume prune

删除: docker volume rm

### docker network

docker network connect[disconnect create inspect ls prune rm]

不同的容器使用同一个网络可以使容器互联，创建网络:

docker network create -d bridde my-net -d 参数指定 Docker 网络类型，有 bridge overlay。

运行一个容器并连接到新建的 my-net 网络: docker run -it --rm --name busybox1 --network my-net busybox sh

### docker port container

查看端口映射情况

### docker run: 运行容器

```
docker run [OPTIONS] IMAGE [COMMAND] [ARG...]
docker run --name test debian​​​​​​​​ //给container命名
docker run -d //后台运行
docker run --rm //退出时自动删除​​
docker run -P //随机映射一个 49000~49900 的端口到内部容器开放的网络端口
docker run -p 3000:5000 //本地3000端口映射到容器5000端口
docker run --read-only //Mount the container’s root filesystem as read only
docker run --networke=my-net //使用network
docker run -i //以交互模式运行容器，通常与 -t 同时使用
docker run -t //为容器重新分配一个伪输入终端，通常与 -i 同时使用
docker run --mount type=volume(bind,tmpfs) source(src)=my-vol target(destination,dst)=/webapp,readonly(导致绑定装入以只读方式装入容器中) //挂载数据卷到容器
```

###