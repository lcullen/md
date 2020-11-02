[最强refer](https://blog.csdn.net/u013256816/article/details/71091774) zhixiaosi
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
    group consumer 如何分配分区:
        第一个joinGroup的consumer会成为群主，拥有一份组员的所有信息，当一个consumer
        调用joinGroup操作的时候，群主会将新的分配信息交给协调器。
        
        群主的分配策略
        
    再均衡的概念:
        如果一个consumer 不在向 broker 发送心跳信息，协调器会发生重均衡
    
    线程安全: 
    
    
5: 如何解决实时性的要求:
    kafka stream read-process-write
    

6: kafka 首领副本和追随副本
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
kafka broker 处理请求
#
    HW: high water 首领副本更新HW的时机准则是 远程副本或本地首领副本更新完LEO的时候 HW = min(所有有副本的LEO)
    LEO: log end offset, 远程副本是 拉取 首领副本的日志同步的时候 更新LEO，首领是收到producer的消息的时候更新本地LEO
    间与HW与LEO之间的是未提交的日志
    每个broker重启之后都会执行截取操作，只截取到HW的日志，所有未提交的日志都会失去
leader epoch
    相当于leader的版本号
    epoch 机制是在 重启截取操作之前 会拉取一次 leader的LEO 

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
    