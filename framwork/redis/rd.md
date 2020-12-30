sds 的字符串预分配
    free->len -> buf 有一个表头

1. string
    + raw
    + embstr
    
2. list 发布订阅的 功能也用了 链表
    + quicklist
    + ziplist 压缩列表 连续的内存分配
        
3. hash 
    + 渐进式 rehash 策略 
        同时保持两个 hash table 同时查询，逐渐整合成一个 hashtable
    + hash func design
        

4. set

5. zset 
    模式: dict + skiplist
    dict 满足set的查找特性
    skiplist 满足排序特性 + 更具范围获取值 (红黑树做不到这一点)
    skiplist 在插入node的步骤 [refer](https://www.jianshu.com/p/58bab10b7ab9)

    intset
        如何解释 intset 
    
应用
1. 分布式锁
    共享内存:
    细节:
        + setnx && expireAt 
        + 原子操作 setNx && expireAt
        + 集群环境下 刚加锁的实例挂了，来不及同步给其他实例
            redlock能够保证大多数原则(大多数的 实例加上了锁)
2. 延时队列
    + blpop 长时间的阻塞redis server 会自动断
    + 使用 zset score 代表 执行时间 value 代表执行任务
    思考 延时队列和定时任务的 golang 各自实现方式 
    todo 
    
3. 位图
    + 打点 setbit
4. 简单的限流策略

5. bloom filter: 通俗的话来讲,如果我(filter)好像看见过你(元素)
    空间效率高:
    检查元素的存在: 存在,可能存在,一定不存在
    主要原理对元素进行N次hash求值得到N个相对于bitarr的位置并在对应的位置上置1

6. expire callback 过期回调策略 [strategy](https://stackoverflow.com/questions/13174615/how-to-get-callback-when-key-expires-in-redis)
    

###redis 持久化
1. 


###IO 模型相关####
IO请求步骤:
1. 进程向内核发起IO请求
2. 阶段1 内核接受请求将磁盘数据copy到内核态 (真正的  I/O 发生的 地方)
3. 阶段2 内核态copy 数据到 用户态 
4. 通知进程 唤起进程

IO 模型
1. 同步阻塞
2. 同步非阻塞
3. 异步阻塞
4. 异步非阻塞
5. 事件驱动

IO多路复用的背景:
    进程池: 将一个连接分配给一个 一个进程。处理完连接后返还进程
    如果一个连接占用一个进程，并且进程阻塞在IO上，进程就挂起，可用进程资源减少
    所以将线程池的每个线程 划分成两部分:
        reactor: 管理连接,每个连接都可以被 accept
        handler: 处理请求,每个连接中的 具体业务逻辑交给handler 处理
        
Copy On Write 
eg: 在容器中添加元素。
    不直接在容器中添加元素，而是新建一个容器，将老容器指向新容器中的元素
    
=======
redis 事件模型
1. 文件事件
    通过IO多路复用模型 dispatcher 监听套接字对象产生的事件，并分发对应的 handler 进行处理
    客户端的每次操作都会产生一个套接字对象的文件事件，连接，读，写，关闭 操作时，产生 连接事件，读事件，写事件，关闭事件，dispatcher 找到对应事件 handler 分发处理事件。
    
    具体监听步骤:
        redis init 监听 && 预关联 AE_READABLE EVENT && 连接应答 HANDLER 
        当一个客户端 connect 的时候产生 AE_READABLE EVENT 通过dispatch 找到  AE_READABLE HANDLER 处理连接事件，此时 dispatch 监听客户端套接字和套接字状态，预关联AE_READABLE EVENT 和 命令处理器 HANDLER; 当客户端发送命令的产生 AE_READABLE 找到命令处理器 进行处理; 执行命令 处理器预关联 AE_WRITEABLE 与 命令回复处理器; 当客户端套接字 尝试获取回复时， 产生AR_WRITEABLE EVENT 找到处理区进行处理... 
2. 时间事件
    
=======
过期策略:
redis 会将所有设置了过期时间的 key 放入到map 中
1. 定时扫描策略(集中处理每秒扫描10次 )
    * 每次拿出20个key,删除过期key
    * 如果过期key 超过1/4 重复上面的步骤 
    如果同一时间过期的key 太多 会导致客户端请求访问超时,可以设置确保整个过程不超过25ms
    但是如果客户端的超时时间为10ms 还是会当做访问超时,所以要尽量避免大量的key同一个时间过期
    
2. 惰性删除
    和延迟加载一样,用到了这个key，查看是否过期，过期就删除
3. key 过期 怎么解决 数据一致性的问题
    aof rdb 怎么处理的
    
    
======= redis 集群配置管理 =======
[出处1掘金](https://juejin.im/book/5afc2e5f6fb9a07a9b362527/section/5b029e77f265da0b9f409688)
[出处2](https://github.com/CodisLabs/codis/issues/1141)
1. codis
    通常使用 1024 个 slot 每一个slot对应一个redis 实例
    codis proxy 中 所有的 slot 和redis 实例之间的 映射关系 用 zk or etcd 同步
    对于每个 key 在proxy 中寻找 redis 实例的过程
     
    hash = crc32(command.key)
    slot_index = hash % 1024
    redis = slots[slot_index].redis
    redis.do(command)
    
    扩容 和 缩容
    扩容: SLOTSSCAN 将对应slot 中 可以进行迁移 如果请求有新的key 落在 当前迁移 slot 中 直接放在新的redis 节点
        增加一个group，在group添加运行一个redis-server，将slot迁移到这个group上
    缩容: 将待缩容的group上的slot迁移到别的server上，然后去掉这个group
    
2. 

=======哨兵 
哨兵的作用是: __维护__ 客户端 到 服务端的 __配置文件__
    客户端订阅 哨兵 switcher-master 消息
        当消息发生的时候 客户端自己维护 redis的连接 信息
    哨兵的定时任务:
        检测 master 节点 info
        检测 slave 节点 info
        检测 其他 从节点 info
    哨兵选举 raft (监控master 节点的最有机会成为 leader哨兵)
    leader 哨兵 故障转移
        选取健康的 优先级高 日志长度最大的 slave 作为master 节点 通知其他节点 更换 master，将有故障的master 标记为从节点
    思考: 哨兵的故障转移、数据还是会丢失、监控slave
 

=======红黑树和跳跃表 
[refer掘金](https://juejin.im/book/5afc2e5f6fb9a07a9b362527/section/5b5ac63d5188256255299d9c)
[referGeek](https://time.geekbang.org/column/article/42896)
1. level <= 64
    从level 1 按概率 持续 获得可到达的最大 level 层级
2. 插入: 根据 level + down 的方式 找到 插入位置 
   更新: redis 使用删除 + 重新插入的方式 代替更新
   level的更新:

3. 相比于 红黑树
    可读性 实现难易程度
    rank 使用 span 跨度计算出当前的 排名
    range 获得范围内的 数据

4. 跳表中如果score的值都一样 效率会 退化到 n，所以还会用value 去比较

======散列 hash dict
装载因子
开放寻址: 如果遇到hash冲突,顺序查找空位，并插入
    线性探测: hash(key) + 1 如上
    二次探测: hash(key) + n^2
拉链法: 每个数组的位置是链表的表头 装载同hash
    当链表长度太长的时候 可以动态的 改变数据结构 用 红黑树代替链表

```c
type dict {
    type
    private
    ht[2] //copy on write
    rehash
} 
```
scale:
redis 中的 rehash 通过 ht[0].used  2^? 来scale ht[1] 所需要的大小
1. bgsave 正在执行的时候 loadfactor 能够大一点再进行rehash 避免不必要的空间和时间复杂度 loadfactor > 0.5 
2. bgsave 不执行的时候 loadfactor > 1 进行rehash scale
3. rehash 的时候用的 hash func && add little salt
    分而治之: 每当hashtable 发生crud的时候 把当前rehash对应的k-v进行迁移，迁移完后rehashidx ++,
4. loadfactor 合理设置

=======常见问题
缓存一致性

缓存穿透/击穿
缓存雪崩
热点数据集中失效

======== redis 的事务 [zpop](https://cloud.tencent.com/developer/article/1189074)
1. 是通过 乐观锁处理
    watch key
        
2. 多个命令发生的时候如何处理回滚
redis 的multi 多条操作，前面几条操作成功，最后一条失败，无法做已经执行成功语句的 回滚。比如
127.0.0.1:6379> get locked
"1"
127.0.0.1:6379> watch locked
OK
127.0.0.1:6379> multi
OK
127.0.0.1:6379> incr locked
QUEUED
127.0.0.1:6379> lpop list_key
QUEUED
127.0.0.1:6379> incr locked
QUEUED
127.0.0.1:6379> exec
1) (integer) 2
2) (nil)
3) (integer) 3
127.0.0.1:6379> get locked
"3"
上面的语句 incr 会执行成功，但是lpop 是失败的


redis:
多往往内部实现 && 背景 && 容错 && 高可用设计的方式去思考 为什么要这么设计
分为几种架构模式:
    一主多从
        1. 高可用的保证 拓扑架构图的设计
            replication 
        2. 数据一致性的实现 (不能保证强一致性):
            * 不同于mysqldb redis只是为了做cache 而不是为了实现数据的一致性
            * 有较弱的一致性 rdb 快照 && aof 命令式的可追加
    cluster:
        * 分布式锁 面临的问题 和解决方案: [refer](https://dbaplus.cn/news-159-3080-1.html)
            + 原子操作 
            + TTL 和 expireTime
                如果发生了GC 问题 或者其他网络延迟的问题， client 在set lock的时间都要大于 expiretime 导致 client 无法知道到底有没有锁成功
                解决方案是 加入定时器的机制 如果client 在expire的时间内无法返回报错 由客户自己处理，还有一种解决方案是 token， 从redis 那版本号， 低于这个版本号的就无法set， basic 算法吧这个是
        * 数据分片怎么做的:
    sentinel:
        拓扑结构 和一主多从的区别

缓存更新、穿透、雪崩:
    cache aside
    page cache fync to disk
    夹带bloom_filter 进行过滤
             
     
为什么那么快，并发又可以那么的高:
    1. redis 
        redis 是基于事件模型驱动的程序, 事件分为
        文件事件: 文件时间基于reactor模型和 epoll I/O 多路复用, 分别监听socket的accept && readable && writable 当事件得到满足就交由对应的处理器
        时间事件: 周期性任务,aof,rdb, 
  
