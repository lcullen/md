data intensive design
1. reliability 可靠性
    + 功能正确
    + 可容错
    + 资源可用
    + 授权功能
    
    * 一句话总结这个可靠性：可持续正确运行，及时错误发生
    * 通常很严重的bug是因为没有处理号 error handler

2. scalability 延展性

3. 数据应用场景
    + 存储
    + 缓存
    + 搜索
    + 流式处理
    + 批处理
4. 描述系统的可拓展性
    + 对负载的描述：
        每秒的读写量是多少
    + 程序性能的描述
        吞吐
        平均响应时间 [常用的统计方式 refer](https://prometheus.io/)
    + 如何提升系统的可拓展性
        水平拓展 垂直拓展
5. 可维护性
    解决系统的复杂度: 通过更加高层次的抽象

###数据模型和查询语言
    -- 语言的边界就是思想的边界
6. 网络模型
    1. 对比层次模型:
        层次模型一个很古老的模型 就是一棵树
        网络模型是一个节点可以有多个父节点(图)
7. 文档模型
    1. 无模式: 没有固定的schema
    2. 读时模式: 只有在读取到的时候能感知到模式
    3. 写时模式: 在写的时候就要符合模式 两者之间的关系是 动态和静态的关系
    
8. facebook() 的时间线设计
    1. 描述负载
        用户行为发生对系统的压力
            write:
                avg: 4.6k/per_second 
                top: 12k/per_second
            read:
                pull: 
                    avg: big_sql
                    top:
                push:
                    avg: 4.6k/per_second * follower
                    top: 12k/per_secode * follower
    2. 描述性能
        SLA: 
    3. 应对负载的方法:
        1. 加机器
        2. 加机器

9. 存储 && 索引的发展史
    1. 文本存储k-v
        ```bash
           #!/bin/bash
           db_set() {
               echo "$1,$2" >> db_file.txt
           }
           db_get() {
               egrep "^$1" db_file.txt | sed -e "s/^$1,//" | tail -n 1
           }
        ```
        弊端: 
            * file_sort && condition filtering [refer mysql file_sort && condition filtering] 解答 mysql对于已经使用索引进行 filter 完的数据要使用猜测的方式 过滤完 condition，这个时候可能需要回表查找
            * 全表扫描
    2. HashMap
        1. 内存Map    
            记录了当前key 对应的 value 在磁盘文件的偏移量 [refer——kafka的日志偏移设计](https://juejin.im/book/5c7d270ff265da2d89634e9e/section/5ca6fae751882543e70d2402)
        2. 随机读
            * 标记删除
            * 灾备
            * 校验和
            * 并发控制
                * 顺序写入
                * 
     3. sstable sort string table
        关键字 leveldb
     4. LSM tree 放弃读的能力， 提供高性能写入
        内存可以通过B树排序 与 磁盘文件 使用归并排序整合
        
        B tree
            对比sstable更加适合用页的方式管理
            从root 节点 index =》  到磁盘 =》 加载到内存修改 =》 flush back 磁盘
            * 如果页太大， 需要分裂的时候发生了奔溃 会出现孤儿 页
            * 并发控制 lock
            * 防崩溃韧性: wal
            Copy On Write技术实现原理：
            fork()之后，kernel把父进程中所有的内存页的权限都设为read-only，然后子进程的地址空间指向父进程。当父子进程都只读内存时，相安无事。当其中某个进程写内存时，CPU硬件检测到内存页是read-only的，于是触发页异常中断（page-fault），陷入kernel的一个中断例程。中断例程中，kernel就会把触发的异常的页复制一份，于是父子进程各自持有独立的一份。
            
            作者：Java3y
            链接：https://juejin.im/post/5bd96bcaf265da396b72f855
            来源：掘金
            著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
//==== 编码
1.  为了解决将进程中的数据 编码到 介质中
    读模式: 
    写模式:
2. 传统的编码格式:
    json: 
        优势: 可读性
        劣势: 太占空间 && 数字精度的问题
        
    Protocol Buffers && Thrift 大致组织的形式: 
        1. 二进制的组织形式: type + flag + length + data_body
            修改前3项的 任何一项 都会影响 data_body 的解析
        2. 兼容性问题: ??
            旧代码读取新代码产生的数据: 
                新加的 字段标签不能被老代码识别
            新代码读取旧代码产生的数据:
                老数据根据新的schema 产生的消息 必须是可选的 或者 有默认值
    Avro 很适合当做大数据的 迁移schema
3. Restful && SOAP  
    1. restful :使用URL来标识资源，并使用HTTP功能进行缓存控制，身份验证和内容类型协商
    2. soap: 

//=====分布式
延时
一致性[容错]
拓展

1. 复制 是为了什么？ 为了高可用，为了可拓展处理更多的请求，为了低延时(距离物理位置更加进的先访问)   
    q:如何复制&&复制的载体【3种，语义复制(statement), 行复制:(binlog), 基于触发器，以及各种方式的优劣】，复制的方式【异步同步】,
        q: 优劣分别是 语义不清晰和环境相关 如 update where now(),current_time()
    q:以及复制的时候会发生的问题
        从库失效 追赶问题 和下面的 一致性问题
        
    1. 引发的数据不一致该如何解决
        1. 时序一致性: 分布式id
            以及将命令分等级
        1. 程序应用逻辑上处理:
            * 读写都走主库
            * uid hash 一致性到同一个数据中心
        分布式数据库可能还会产生的问题：
            一致前缀读 所有的数据被写入到不同的分片中，如何保证每个写入是有顺序的标志
            
   复制的几种场景:
    1. 单主
    2. 多主         
    3. 无主
    2. 读写的仲裁法定人数
        1. $w + $r > $n  可以保证高可用
        为高可用所做出的的 trade off
        1. 如果当前节点写入失败 如何保证所有节点的回滚
        2. 并发的读和写 如何保证读到的是最新的
        3. 时序一致性
2. 分区:
    不同的分区有各自的leader 每个节点中的分区可能是 leader 也可能是其他分区的 follower
    数据倾斜 和热点问题 
    分区索引:
        用主键进行hash分区 但是也会有特殊事件的 热点问题
        次级索引:
            文档索引，每一个分区维护自己的 次级索引
        最左前缀匹配原则:
            选择第一个字段为等值，后一个字段为范围查询
3. 事务:
    ACID:
    
4. 分布式:
    在不可靠的组件中建立一个可靠的系统；
    网络分区的检测:
        1. 主机监控当前进程: 当前宿主机通知其他节点，当前进程崩溃
            主机监控当前进程的方式: node exporter
        2. 心跳 etcd的raft协议
        3. response 因为网络的原因而造成的超时 要主要应用级别的可重试
    时钟:
        1. 时钟: 出现回拨问题
        2. 单调钟: 会改变频率 进行同步，但是始终都是单调的 
    依赖时钟: AAB 问题
    进程挂起: 获取本地时间后 进程处于挂起状态
    term 令牌 防止脑裂无效命令
    
5. 一致性共识
    线性一致【读已之写,新鲜度】 强一致性
        如何保证强一致性: 共识算法
        代价: CAP 在P的情况下要满足C, 那么只能牺牲A

7  tx事务
    1. 隔离级别mysql的隔离级别实现方式

8. 分布式

9. cas 也能够解决 ABA 问题:
    
    