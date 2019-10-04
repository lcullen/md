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
         
        