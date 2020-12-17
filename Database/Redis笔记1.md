### 数据结构
1. String
2. Hash
3. List: 字符串列表,按照插入顺序排序
4. Set: string 类型的无序集合。
5. SortedSet: string类型元素的有序集合,且不允许重复的成员
6. HyperLogLog
7. Geo
8. Pub/Sub
9. Redis-Module(BloomFilter RedisSearch Redis-ML Reedis-cell(限流 ))

---
### 记录一些常用到的命令
```shell
# 管理命令
# 内存
redis-cli info
redis-cli info memory
redis-cli dbsize
# 查看key情况
redis-cli --bigkeys
#访问情况
redis-cli  info |grep ops
redis-cli info clients
# 统计
redis-cli --stat
# 慢查询
redis-cli SLOWLOG RESET
redis-cli SLOWLOG GET
# 性能测试
redis-benchmark [option] [option value] 
```

```shell
# 获取rdb存储路径
CONFIG GET dir
# 切换数据库
SELECT index
# 删除指定key
DEL key/keys
# 检查key是否存在
EXISTS key
# 清空某个数据库的数据
FLUSHDB
# 删除所有数据库数据
FLUSHALL
# 搜索key
scan 0 match pattern count 1000
# 锁
# 先拿setnx来争抢锁，抢到之后，再用expire给锁加一个过期时间防止锁忘记了释放。
# set指令有非常复杂的参数，可以同时把setnx和expire合成一条指令来用的！
```

**List**
```shell
# list size
llen key
# 原子性地返回并移除存储在 source 的列表的最后一个元素（列表尾部元素）， 并把该元素放入存储在 destination 的列表的第一个元素位置（列表头部）[可靠队列]
RPOPLPUSH source destination
# 通过索引获取列表元素
LINDEX key -1
LRANGE key -3 -1
```

**HyperLogLog**
统计非重复数据
```shell
PFADD login-1 A B C D E F A B C
PFADD login-2 G H I J K A B C
PFCOUNT login-1 => 6
PFMERGE login-1-2 login-1 login-2
```

---
### Redis事物
MULTI  EXEC  DISCARD 和 WATCH 是 Redis 事务的基础。
事务可以一次执行多个命令, 并且带有以下两个重要的保证:
+ 事务是一个单独的隔离操作：事务中的所有命令都会序列化、按顺序地执行。事务在执行的过程中，不会被其他客户端发送来的命令请求所打断。
+ 事务是一个原子操作：事务中的命令要么全部被执行，要么全部都不执行。

1. MULTI 命令用于开启一个事务，它总是返回 OK
2. MULTI 执行之后， 客户端可以继续向服务器发送任意多条命令， 这些命令不会立即被执行， 而是被放到一个队列中， 当 EXEC 命令被调用时， 所有队列中的命令才会被执行。
3. 另一方面， 通过调用 DISCARD ， 客户端可以清空事务队列， 并放弃执行事务。
4. WATCH 命令可以为 Redis 事务提供 check-and-set （CAS）乐观锁 行为。被 WATCH 的键会被监视，并会发觉这些键是否被改动过了。 如果有至少一个被监视的键在 EXEC 执行之前被修改了， 那么整个事务都会被取消
5. 至于那些在 EXEC 命令执行之后所产生的错误， 并没有对它们进行特别处理： 即使事务中有某个/某些命令在执行时产生了错误， 事务中的其他命令仍然会继续执行。

---
### keys and scan
keys命令可以扫出指定模式的key列表。Redis的单线程的。keys指令会导致线程阻塞一段时间，线上服务会停顿，直到指令执行完毕，服务才能恢复。scan指令可以无阻塞的提取出指定模式的key列表，但是会有一定的重复概率，在客户端做一次去重就可以了，但是整体所花费的时间会比直接用keys指令长。但是对于 SCAN 这类增量式迭代命令来说， 因为在对键进行增量式迭代的过程中， 键可能会被修改， 所以增量式迭代命令只能对被返回的元素提供有限的保证 。

---
### 用redis做异步队列
一般使用list结构作为队列，rpush生产消息，lpop消费消息。当lpop没有消息的时候，要适当sleep一会再重试。或者使用blpop，在没有消息的时候，它会阻塞住直到消息到来。

---
### Pub/Sub
使用pub/sub主题订阅者模式，可以实现 1:N 的消息队列。在消费者下线的情况下，生产的消息会丢失，得使用专业的消息队列。

---
### Redis延时队列
使用sortedset，拿时间戳作为score，消息内容作为key调用zadd来生产消息，消费者用zrangebyscore指令获取N秒之前的数据轮询进行处理。

---
### Redis持久化，主从数据交互
RDB做镜像全量持久化，AOF做增量持久化。可以把RDB理解为一整个表全量的数据，AOF理解为每次操作的日志。
RDB: 继续提供服务，只有当有人修改当前内存数据时，才去复制被修改的内存页，用于生成快照
1. RDB:父进程fork子进程去处理保存工作，父进程没有磁盘IO操作。子进程创建后，父子进程共享数据段，父进程继续提供读写服务，写脏的页面数据会逐渐和子进程分离开来。(fork and copy-on-write)
2. 在恢复大量数据时比AOF快。
3. 只有一个文件非常紧凑，非常适合备份
4. 故障时会丢失数据
5. fork子进程可能会耗时较多

AOF
1. 使用AOF持久化会让Redis变得非常耐久，默认fsync一秒一次，fsync会在后台线程执行
2. AOF文件是一个只进行追加操作的日志文件，Redis 可以在 AOF 文件体积变得过大时，自动地在后台对 AOF 进行重写， 重写后的新 AOF 文件包含了恢复当前数据集所需的最小命令集合。 整个重写操作是绝对安全的
3. AOF 文件有序地保存了对数据库执行的所有写入操作， 这些写入操作以 Redis 协议的格式保存， 因此 AOF 文件的内容非常容易被人读懂
4. 对于相同的数据集来说，AOF 文件的体积通常要大于 RDB 文件的体积。

```shell
# 同步保存RDB
SAVE
# 不阻塞保存RDB
BGSAVE
# 对AOF文件重建
BGREWRITEAOF
# 对AOF文件进行修复
redis-check-aof --fix 
```

### Redis Mode
#### master-slave:
![IMAGE](resources/BDC00442FEA13660EA479B911A16BED5.jpg =410x346)
```shell
# 以守护进程的方式运行
daemonize yes
port 6379
# 绑定的主机地址
bind 0.0.0.0
# 指定当本机为slave服务时，设置master服务的IP地址及端口，在redis启动的时候他会自动跟master进行数据同步
slaveof 127.0.0.1 6380
# 当master设置了密码保护时，slave服务连接master的密码
masterauth <master-password>
```

#### 哨兵模式(Sentinel)
Redis Sentinal着眼于高可用，在master宕机时会自动将slave提升为master，继续提供服务。
Redis Cluster着眼于扩展性，在单个redis内存不足时，使用Cluster进行分片存储。

Sentinel着眼于高可用，哨兵模式集成了主从模式的有点，同时有了哨兵监控节点还能保证当 Master 节点离线后，哨兵监控节点会把 Slave 节点切换为 Master 节点，保证服务可用。
哨兵模式是在主从模式的基础上增加了哨兵监控节点，最简单的哨兵模式需要一个 Master、一个 Slave 、三个哨兵监控节点。
![IMAGE](resources/68B581BE6877FAE78A917B91E9DBFABF.jpg =425x386)
```shell
# redis-sentinel redis-sentinel.conf
# 复制三份，改下端口
bind 0.0.0.0
port 26379
# 这个2代表，当集群中有2个sentinel认为master死了时，才能真正认为该master已经不可用了。
sentinel monitor mymaster 10.201.12.66 6380 2
# 指定主节点应答哨兵sentinel的最大时间间隔，超过这个时间，哨兵主观上认为主节点下线，默认30秒 
# Default is 30 seconds.
sentinel down-after-milliseconds mymaster 30000
# 指定了在发生failover主备切换时，最多可以有多少个slave同时对新的master进行同步。这个数字越小，完成failover所需的时间就越长；反之，但是如果这个数字越大，就意味着越多的slave因为replication而不可用。可以通过将这个值设为1，来保证每次只有一个slave，处于不能处理命令请求的状态。
# sentinel parallel-syncs <master-name> <numslaves>
sentinel parallel-syncs mymaster 1
```

```shell
# 连接到sentinel
redis-cli -p 26379
# 查看sentinel
info sentinel
# 查看当前主节点
SENTINEL get-master-addr-by-name mymaster
# 强制当前sentinel执行故障转移(模拟故障)
sentinel failover mymaster
```

#### Redis Cluster
Redis-Cluster 是 Redis 官方的一个 高可用 解决方案，Cluster 中的 Redis 共有 2^14（16384） 个 slot 槽。创建 Cluster 后，槽 会 平均分配 到每个 Redis 节点上。
![IMAGE](resources/D7AD3E164BDB2CBF2C7778EAC5D77AA2.jpg =560x365)

```shell
redis-cli --cluster help
# 创建集群
redis-cli --cluster create 127.0.0.1:6379 127.0.0.1:6380 127.0.0.1:7002 127.0.0.1:6381 --cluster-replicas 1
```

---
### Redis的内存回收
Redis过期策略
1. 定期删除: 每隔一段时间,扫描Redis中过期key字典,并清除部分过期的key。
2. 惰性删除: 当访问一个key时,才判断该key是否过期，过期则删除。

Redis淘汰策略:
当内存不足以容纳新写入数据时，新写入操作会报错,内存使用到达maxmemory上限时触发内存淘汰策略:
1. 在键空间中，移除最近最少使用的 key（这个是最常用的）
2. 在键空间中，随机移除某个 key。
3. 在设置了过期时间的键空间中，移除最近最少使用的 key。
4. 在设置了过期时间的键空间中，随机移除某个 key。
5. 在设置了过期时间的键空间中，有更早过期时间的 key 优先移除。

---
### 缓存击穿
缓存击穿表示恶意用户模拟请求很多缓存中不存在的数据，由于缓存中都没有，导致这些请求短时间内直接落在了数据库上，导致数据库异常。
BloomFilter
```shell
git clone git://github.com/RedisLabsModules/rebloom
cd rebloom
make
#redis.conf
loadmodule /path/to/rebloom.so
# redis-cli
BF.ADD bloomFilter foo
BF.EXISTS bloomFilter foo
```

---
### 缓存雪崩
缓存在同一时间内大量键过期（失效），接着来的一大波请求瞬间都落在了数据库中导致连接异常。
我们一般需要在时间上加一个随机值，使得过期时间分散一些。