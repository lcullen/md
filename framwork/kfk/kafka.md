[最强refer](https://blog.csdn.net/u013256816/article/details/71091774) zhixiaosi
**背景[cap是怎么做到的]**
>> for interview

1. kfk 与 zk 的关系
    kfk 自带 zk
    zk 保留 kfk 元数据 && 消费信息
    
2. kafka produce config: 
    produce config
    acks = 0 fire-and-forget
    acks = 1 leader confirm
    acks = all 影响吞吐量

3. kafka 序列化器
    使用 avro + Confluent Schema Registry 的方式实现kafka schema 的定义
    
4. consumer
    consumer 集群之间互相不影响各自的消费
    group consumer 如何分配分区:
        第一个joinGroup的consumer会成为群主，拥有一份组员的所有信息，当一个consumer
        调用joinGroup操作的时候，群主会将新的分配信息交给协调器。
        
        群主的分配策略
        
    再均衡的概念:
        分布式中的心跳 vs kafka 心跳
            kafka
            1. broker 充当consumer cluster 的 coordinator; 其他的broker watch zookeeper/ctrl -> broker_id 的失效 
                ctrl/broker_id 失效的时候 会不会触发惊群效应 ? 如何解决 ? [todo]
                    顺序监听 形成一个监听链路表, 当前链路中的
            2. 
        如果一个consumer 不在向 broker 发送心跳信息，协调器会发生重均衡
    
    线程安全: 
        不同的集群互不影响; 每一个partition 都只有一个集群的 唯一一个消费者 可以消费, 不同集群之间的消费者互相不影响
        
    当前消费者组的状态机:
         preparerebalance /completebalance / empty / stable
         消费者提交问题消息体: group_topic_partition => offset_expire_at 
         rebalance的场景:
         成员: join_req: 请求加入当前的group,上报自己要消费的topic 以及自己的offset (老成员需要上报)
         coordinator:  sync_res: 重新奉陪当前对于topic 的分区安排
         
    
5: 如何解决实时性的要求:
    kafka stream read-process-write
    

6: kafka 首领副本和追随副本(强leader 分布式: 所有的读写都经过leader 分区)
    首领副本：处理客户端的读写请求
    追随副本：同步首领数据，如果收到读写请求就返回报错信息 (kafka客户端自己负责把请求正确的发往首领副本)
    追随副本进入ISR队列的要求:
        1. 与zookeeper 保持会话
        2. 10s内从首领副本获取消息
        3. 10s内从首领副本或去过最新消息
    生产者在可靠的系统中要正确配置
        1. acks
        2. 错误重试机制
            producer 处理broker发回的可重试的错误码
                1. 记录错误
                2. 回调
    消费者在可靠系统中的正确配置
        1. groupid
        2. auto.offset.reset 如果是新的topic 没有设置offset 那么可以选择从latest or earliest 进行消费
        3. enable.auto.commit 
        4. auto.commit.interval.ms 自动提交的时间间隔(频率)
        
kafka broker 处理请求 [refer reactor 网络模式] acceptor -> proccesser 
    1. 客户端的请求
        a. 写入请求
        b. 获取请求
            数据一致性的问题， 是否低于高水位 offset
    2. 分区副本的请求
        a. 复制类型的请求
    3. 控制器发送给分区首领的请求
    [客户端的负载均衡] kafka 客户端必须自己维护对 leader 分区 的元数据信息: 否则发送到非leader 分区时 会接受到错误的响应
        每个broker 都拥有全局的 broker cluster 元数据信息:
             
        
#
    HW: high water 首领副本更新HW的时机准则是 远程副本或本地首领副本更新完LEO的时候 HW = min(所有有副本的LEO)
    LEO: log end offset, 远程副本是 拉取 首领副本的日志同步的时候 更新LEO，首领是收到producer的消息的时候更新本地LEO
    间与HW与LEO之间的是未提交的日志
    每个broker重启之后都会执行截取操作，只截取到HW的日志，所有未提交的日志都会失去
leader 
    epoch
        相当于leader的版本号
        epoch 机制是在 重启截取操作之前 会拉取一次 leader的LEO 
    election 选举:
        1. 通常会研究out of sync replica 里面获取一个副本分区当做首领
            a.  

存储:
    日志: (提高网络传输) [refer redis aof]
        1. 将多条log compress 为一条 log 中的value值，多条被压缩的日志有 relative offset
            [refer mysql page slot] 
        2. 变长字段
    日志索引:
        1. 索引偏移量
            xxx.index:    offset1  => length  offset2=> length
            每一个length 区间都有大量的 record 然后根据差值再去具体找 offset
        2. 时间索引
            类似
    日志清理(多维度):
        1. 基于时间删除:
            借助timestamp.index 找到最大的 lastmodify,标记文件为.delete 文件 
            [refer redis 渐进式rehash]
        2. 基于offset 和基于时间是同一个方式
        3. 日志压缩
            对同一个key的log 进行压实操作，取最新的 一个key当做有效
            操作如下:
                1. 第一次扫描: SkimpyOffsetMap 来保存当前key的最大 offset
                2. 第二次扫描: 对比日志中 相同的key 的offset 是否大于当前 map 的offset        

kafka 不会发生lock 问题，因为每一个consumer 都是单独订阅和消费producer

kafka 的压缩:
    可配置压缩 
    1. 客户端压缩
        发送时的批量压缩
    2. 服务端压缩
kafka的稀疏索引
    1. timestamp 索引
    2. offset 索引
    
日志清理:
   删除:
   1. 根据时间戳: 标记清除，标记当前的日志文件为.delete 后缀， 交给后台delete线程进行异步删除
   2. 基于总的日志大小
   2. 根据offset: 标记需要清理的 最小offset， 小于这个值的所有offset segment 都要被标记为deletable segment
   压缩: 可以对比redis的 aof 或者 RDB
   1. 相同的key LWW: last write win, 日志序列不再联系
   
kafka的fsync 机制 和 mysql的进行类比:
    mysql的落盘机制
        1. redo log 的三个落盘过程
            redo_log_buffer -(group_commit)> page_cache -(fsync)> disk
        2. bin log 的落盘过程
            bin_cache -(fsync)> disk
        结合redo bin log 的两阶段提交
            -(perpare)>redo_log_buffer  -> bin_cache -> redo_log fsync -> bin_log fsync -> (可以忽略)redo_log commit
        3. mysql 的双1 设置 (如何保证数据的一致性 完整性)   

协议:
    1. 其中一项correlation_id, 用于关联请求上下文 对response 进行关联 

延迟设计:
    1. 使用netty 的时间环做延迟计时策略
    2. 延迟任务可以被外部时间提前触发，也就是可以被取消
    3. 如何去各自满足 消费者 && follower 拉取offset的 延迟请求
        1. follower: leader 副本追加了本地日志(涉及事务，一定是未提交的事务)
        2. consumer: 更新lso 作为consumer 的拉取外部事件

broker 担当 Controller 强leader 方式 控制所有的集群事件
    1. first create zookeeper node as leader,others watch node change
    2. 监听所有的zookeeper 文件节点变更事件:
        1. isr
        2. topic
        3. partition

consumer(client.id) 端自定义 分区分配的策略 
    1. single : first custom, second as config
    2. multiple:  group_coordinator,consume_coordinator
        group coordinator 是由对应的 leader 副本所在 
        
TTL:
    
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
    2. 数据复制: 
        1. 为什么没有使用 仲裁的方式(少数服从多数)
            * 仲裁适用于读多写少的服务 ? 
            * 使用isr : 写入的性能取决于 最慢的 follower
        2. 写入时候发生了什么:
             根据生产者不同的参数设定 表现出的 一致性会有所不同
        3. 如何理解epoch: 
            副本更新高水位的时候 奔溃，重启的时候需要从高水位截断日志重新从leader 同步数据
            如果此时 leader 也奔溃, 那么 副本还没有完全同步数据的话将丢失数据，所有加入了纪元号
            从奔溃中恢复的时候 发现当前leader 的纪元号是最大的， 而且自己副本的当前最大纪元也是相同的，则不发生截断行为
IV. 特殊应用 延时队列的设计：
    1. 延时任务并不是直接发送给 用户的topic 而是发送到内部的dely queue 中， 
        等delay queue 中的任务条件得到满足的时候 再发送给用户topic
    2. delay queue 的设计思路

 
