kafka 
I. 如何做到高吞吐
    1. broker:(存储的方式 && I/O):
        I/O 对应的调度算法:
            a.FIFO / CF
        a. 顺序写入: 替代随机写入，减少了寻址的时间 __磁盘顺序写的性能会颠覆三观,甚至比内存的随机写入性能都要高__ [LSM&&SSTable](https://juejin.cn/post/6844903863758094343) __每次大规模的搜索会使用bloomfilter处理__ 包括LSM 查找是否存在的key, []
            选址方式(kafak索引的组织方式): 每一个文件夹下都会有三种文件格式的数据 *.log *.timestamp *.index
                * offset: 每一个*.index 文件都是以 relative offset 作为文件名 进行二分搜索找到对应的*.index file 锁定到真正的物理offset
                * timestamp: 和二级索引类似 timestamp => offset， offset=> address
        b. 页缓存:
            一次性fsync 到磁盘， 组提交方式 [refer](http://www.yangchengec.cn/setup/498.html)
        c. 零拷贝 : 一般模式下read(socket) write(socket) 都是用户态行为 整个过程都需要 内核和用户态的参与，所有进行了4次拷贝
    2. 客户端
        producer: 
            a. 压缩算法:
            b. 发送策略: 
        consumer:
            a. 保存对于broker fetch request 的元数据保存
        

II. 高可用是如何保证的 __CAP__ 中间件的选型方式:基本可用&&最终一致性 : 
        基本可用:__级别调整:any,one,quorum__: 削峰、过载保护、服务降级
    副本机制[分布式事务的解决方案:还会涉及到乐观锁和悲观锁的选型问题]: 每一个topic_partition 能够有多个replication 对应 
    副本机制中的 角色选择套件 kafka 选择的是 强leader(partition侧视角) 和多主(topic侧视角)的混合 
        同步侧的推拉分析 
    1. 强leader(一主多从) : 写性能的瓶颈问题如何解决
        * raft : 
            a. 两阶段提交 优化为 一个阶段的提交: app <-> (ask/answer) leader <-> half answer (ask/answer) follower; 然后将每次的commit 信息append 在heart beat or msg 中 
        * kafka的副本同步问题和关键点:  
            a. 日志的同步方式(推和拉?) 有推有拉: [日志写入过程和延迟设计](http://www.justdojava.com/2019/12/11/kafka-replication-request/);
            b. HW && ISR(变更条件和时机):  leader 需要收到当前isr 中所有的副本成功提交之后 才返回客户端成功写入的response
            c. 高吞吐还和数据分片有关系 kafka 能够对topic 进行partition的数据分片的方式
        * leader 切换数据丢失的问题: 
    2. 无主(一定有一个coordinator): 所有的node 既不是leader 也不是副本能够接收所有的读写请求， 由协调器决定读写仲裁
        由coordinator来决定写入和读取 用于仲裁当前是否是写入成功，和能够读取到最新的值
        多个节点数据不一致的问题
    3. 多主(多活):
       强leader 模型存在写入瓶颈
    
 III. 数据可靠性
    1. 数据丢失的问题:
        a. leader 切换导致的数据丢失问题:
            node 重启 会根据当前所在的的副本截断高于水位的日志，此时从leader fetch msg 的时候 leader 也发生了故障，导致当前完成重启的节点当做了新的leader，产生数据丢失
            
IV. 特殊应用 延时队列的设计 
       

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
  

golang 基础
 + context 的作用
 + 并发包里面有哪些
 + chan
 + sort 里面工业级别的算法
 + go routine 的调度:
    GPM: 
    

    

mysql:
    innodb
    + 索引
        方向: 可以把record 分为 普通用户数据 目录项数据(为的是找到更小的页) 
        * 为什么要使用索引
        * 索引的存储方式
    MyISAM
    
网络: 
I. HTTP2 /HTTP1.X 
    1. I/O 多路复用 同一个http 连接上可以 被不同的frame 复用
    2. 二进制传输 使用数据frame
II. HTTP
III TCP: 
    1. 可靠性是指啥: 
IV. Headline Blocking 问题

微服务:
I. 限流策略和设计:大量的请求导致系统过载
    GRPC有这种能力么？
    1. 限流: 在滑动窗口内最大的请求数量, 参考rrt 时间内的最多请求数量, 如果超过了当前的数量 处理方式: 
    2. 熔断: 在滑动窗口内最大的失败阈值是多少
    3. health:
    4. slb: 带权优先级队列
    以上这些都是依赖于采样打点分析 也是可以参考tcp 的流控去实现的 [refer](https://coolshell.cn/articles/11609.html)
        

补充:
elastic search 基础能力
    1. index = database, type = table , doc = record , fields = column
    2. HA: 一个index 主从分片 && replica  主分片被确立的时候无法修改。 和search:  的路由有关系, 路由到具体的主分片上. 并且搜索当前主分片的所有副本拿到最新的数值
    3.  
    
GRPC客户端的负载均衡策略:
    谈谈GRPC的负载均衡策略，grpc官方并未提供服务发现注册的功能实现，但是为不同语言的gRPC代码实现提供了可扩展的命名解析和负载均衡接口。其基本原理是：服务启动后grpc客户端向命名服务器发出名称解析请求，名称会解析为一个或多个ip地址，每个ip地址会标识它是服务器地址还是负载均衡地址，以及标识要使用的那个客户端的负载均衡策略或服务配置。客户端实例化负载均衡策略，如果解析返回的地址是负载均衡器的地址，则客户端将使用扩展的负载均衡策略，反之客户端使用服务器配置请求的负载均衡策略。负载均衡策略为每个服务器地址创建一个子通道，当有rpc请求时，负载均衡策略决定那个子通道即grpc服务器将接收请求，当可用服务器为空时客户端的请求将被阻塞。
    这种方式好处是灵活，支持扩展，可以扩展满足自己需求的策略，缺点是需要自己集成，需要一定的工作量，对技术人员有一定的要求。
程序设计:
I. 定时器的设置:
    1. redis 时间事件设计思路
    2. kafka 延迟任务的设计思路
    3. goland timer的实现思路:
    
II. 秒杀场景