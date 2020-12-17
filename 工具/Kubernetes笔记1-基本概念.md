`K8s中大部分概念比如Node, Pod, Replication Controller, Service等都可以被看作一种资源对象`

##  Master
集群的控制节点
关键进程
1. kubernetes api server: 提供HTTP Rest接口服务
2. kubernetes controller manager: K8s里所有资源对象的自动化控制中心
3. kubernetes scheduler: 负责资源调度的进程

## Node
除了Master, K8s集群中的其他机器称为Node
kubectl get nodes
kubectl describe node xxxx

关键进程
1. kubelet: 负责Pod对应容器的创建启停任务, 与Master协作实现集群管理
2. kube-proxy: 实现K8s Service的通信与负载均衡
3. docker-engine: Docker引擎, 负责本机容器的创建与管理工作

## Pod
一个容器组, 包含一个pause根容器, 还有其他的业务容器
1. 以Pause容器状态判断整体Pod状态
2. 业务容器共享Pause容器的IP与Volume

## Label
标签
一个Label是一个key=value的键值对, key与value由用户指定。
我们可以对一个资源绑定一个或多个不同的Label来实现对维度的资源分组管理功能

## Replication Controller
定义了一个期望的场景, 即声明某种Pod的副本数量在任意时刻都符合某个预期值
1. Pod期待副本数量
2. 用于筛选目标Pod的Label Selector
3. Pod副本数量小于预期值, 用于创建新Pod的template

## Deployment
可以看作是RC的升级, 两者很相似, 可以随时知道当前Pod部署进度
使用场景
1. 创建Deployment来生成Replica Set并完成Pod副本创建
2. 检查Deployment状态
3. 更新Deployment
4. 回滚之前版本的Deployment
5. 暂停Deployment
6. 扩展Deployment
7. 清理不需要的旧版本

## Horizontal Pod Autoscaler
手动的Kubectl scale可以实现扩容与缩容
HPA是自动扩容
指标
1. CPU
2. 自定义指标

## StatefulSet
本质上来说可以看成Deployment/RC的变种
1. 每个Pod有稳定唯一的网络标识可以用来发现集群内部的其他成员
2. 控制的Pod副本的启停是受控的, 操作第n个pod时前n-1个pod已经是运行且准备就绪的
3. Pod采用稳定的持久化存储卷 

## Service
service定义了一个服务的访问地址, 前端应用通过这个地址访问背后的一组由Pod副本组成的集群实例, service与其后端的Pod副本通过Label Selector来实现无缝对接

Node IP
Pod IP
Cluster IP

## Job
批处理任务通常并行或窜行, 处理完之后整个批处理任务结束。
Job也控制这一组Pod容器
1. Job所控制的Pod副本是短暂运行的, 每个容器运行一次, 当所有的Pod副本都运行结束时, 对应的Job也就结束了。
2. Job控制的Pod副本是能够实现并行计算的

## Volume
存储卷是Pod中能够被多个容器访问的共享目录, Volume被定义在Pod上, 生命周期与Pod相同与容器不相关

## Namespace
命名空间, 很多情况下用于实现多租户的资源隔离

## Annotation
注解信息, 也是key value形式, 但是没有Label那么严格, 只是用户定义的附加信息, 查看使用.

## ConfigMap
配置文件管理
将存储在etcd中的ConfigMap通过Volume映射的方式变成目标Pod内的配置文件, 不管Pod被调度到哪台机器都会完成自动映射, 修改ConfigMap中的key value, 那么映射的配置也会自动更新