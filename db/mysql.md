1. mysql 
    为什么要用B+
    1.二叉树 的退化会变成单向链表
    2.B数 多路搜索(引用背景 文件索引) 的退化会变成数组,所有的 数据 都在路径节点中。
        如果要加载一颗很大的数到内存，能够通过加载单个节点得到单个节点的数据后再换
    3.B+ 所有的数据节点都存储在叶子节点，并且相互链接成有序链表。
LOG
2. WAL(write ahead log) 日志先行 磁盘IO太慢
    写日志并更新缓存
    redo log (物理层)循环顺序写磁盘
    保证ACID 的D(持久化)
    innodb_flush_log_at_trx_commit 0:每隔1秒master thread flush redo 到磁盘​1:立刻将redo log flush 到磁盘
2:redo log 仅在 内存中不刷入磁盘​造成相关mysql的抖动问题​
可配置文件大小
    + c checkpoint 更新位置
    + w writepos 日志位置
    + w追上c时,必须push所有的log否则覆盖
binlog(物理层)追加写磁盘
    undo log(逻辑层)
    uncommit:回滚段中的都是反向操作sql
    committed: 删除回滚段sql
两阶段提交: 我要保证通过binlog 来恢复数据是 可靠的 【binlog 是可靠的 不多不少的!】
    保持binlog redolog 一致性
INDEX
    FORCE_INDEX
    选错索引
    不断删除和添加数据的场景，优化器综合考虑磁盘io等情况后选择错误
    
change buffer 在索引中的使用
普通索引 update
唯一索引 update
[锁](https://mp.weixin.qq.com/s/wGOxro3uShp2q5w97azx5A)
1. MDL 元数据锁
    + 维护表元数据的数据一致性
    + 解决事务隔离
    + 解决数据复制
        Waiting for table metadata lock[故障](http://www.ywnds.com/?p=10091)总之，alter table的语句是很危险的（其实他的危险其实是未提交事物或者长事务导致的），在操作之前最好确认对要操作的表没有任何进行中的操作、没有未提交事务、也没有显式事务中的报错语句。如果有alter table的维护任务，在无人监管的时候运行，最好通过lock_wait_timeout设置好超时时间，避免长时间的metedata锁等待。​
行锁 (innodb) 提高并发量

2.  行锁
    + Record Lock: 锁住的是唯一索引
    + 幻读:同样的sql执行两次结果不同
    + Gap LOCK 间隙锁 开区间, 不包括row 本身，避免多个事务对同一个范围多次插入出现幻读 
    + Next-Key Lock （**RR隔离级别**） 包括row 本身，避免多个事务对同一个范围多次插入出现幻读
        具体的案例: a < b < c
            + 唯一索引等值查询 b
                * 已经存在行b 退化为 record lock b
                * 不存在行b lock: (a, c] 但是排除最后一个不相等值 (a, c) 退化成为gap-lock
            + 非唯一索引等值查询 b (说明:非唯一索引无法保证只有一行数据，所以要扫描到第一个不相等的值后才停止扫描)
                * 已经存在行,对应lock : 原本锁定 (a ,b],  (b, c], 因为b != c 退化为(a ,b],  (b, c)
                * 为找到行lock: (a,c]   
            + 唯一索引范围查询 a<= ? < c 
                原本锁定(?, a] (a, c] 因为是唯一索引(?, a] 退化为 record-lock ,整体 lock: [a,c]
            + 非唯一索引范围查询
                无法退化 锁定(?, a] (a, c]
        其他案例:
        limit 对非唯一索引加锁的 影响
        + 非唯一索引寻找对象的 时候 会持续遍历到 第一个不满足条件的 对象 其中遍历过的 对象都会加锁
        + 加上limit 条件后 能够减少 扫描行， 减少lock
        
3. 表锁
    + 意向锁
        暗示了一个事务中 接下来更加细粒度的锁。通常对细粒度的对象进行上锁的前提是，粗对象已经被上锁
TRANS
[隔离级别](https://mp.weixin.qq.com/s/gjt9WdyTQRzx-hr_qIz_mw)
RUC

RR,事务在第一个Read操作时,会建立Read View 始终只有一个记录
RC下,事务在每次Read操作时,都会建立Read View 读取最新的记录

SERILIZE

4. mysql 两阶段提交
    - binlog 写入流程 
        1. binlog cache: 每个tread 单独一份cache 
        2. binlog page: 写入binlog file 
        3. binlog disk: 写入硬盘
        
        由sync_binlog 参数控制page,disk的写入时机
        sync_binlog = 0 : 不写入page:宕机丢失数据
        sync_binlog = 1 : 每次事务提交调用write写入page,调用fsync写入disk: IO变高
        sync_binlog = N : 每次事务提交调用write写入page,N次事务提交后调用fsync写入disk
        
    - redolog 写入流程
        1. redolog buffer: 多个线程公用一个buffer池
        2. redolog page
        3. redolog disk
        
        innodb_flush_log_at_trx_commit
        0: 每次提交只停留在buffer
        1: 每次提交直接写入disk
        2: 每次提交写入page
        
        tips: 
            innodb_log_buffer_size 占满一半的时候后台线程主动落盘
            innodb_flush_log_at_trx_commit=1，B 事务 会提交 A事务 buffer 中还没有落盘的log
        
    - 双 1: sync_binlog = 1 && innodb_flush_log_at_trx_commit = 1 保证数据不丢失
        
    - binlog && redolog 两阶段提交流程
    
        begin:
          redolog.Prepare (redolog.write page)
          binlog.write page
          redolog.fsync
          binlog.fsync
          redolog.Commit
        commit
        
    - 组提交概念: 
        多个 tx 并发的 时候，第一个tx 当做leader, 当第一个tx binlog redolog fsync 的时候,顺带提交后来的log
        
        * 延迟同步时间: binlog_group_commit_sync_delay 等待一段时间后 一起提交
        * 延迟同步次数: binlog_group_commit_sync_no_delay_count 等待一组count 后 一起提交
                 
5. 主从
    主库A校验从库B的连接请求
    维持长连接并发送从B发过来的binlog位置节点开始发送binlog
    从库B io_thread 线程维持与A的通讯 读取binlog 生成中转日志 relay log
    从库B sql_thread 读取relay log 开始同步 binlog
    
    binlog的三种形式(协议) binlog_format
    1. statement (声明): 
       使用 sql 原语同步 
       缺点: sql 中含有  limit 语句被认为是unsafe,容易发生主从数据不一致
       优点: 节省binlog空间 
    2. row:
        记录真实的主键id,对主键id所在行进行操作
        优点: 数据一致
        确定: 太多的binlog 造成io 比较高
    3. mixed(row + statement) 
        结合 row + statement 的优点
        
    延迟
    1. 物理硬件配置低
    2. DDL
    3. 从库慢查询
    
    高可用策略 cap 中的 c a
        可靠性:
        可用性:
    
    主备切换中的知识点:
    
      位点: MASTER_LOG_FILE + MASTER_LOG_POSITION
        1.文件名 + 位置偏移量
        从库B change master 时要发送 原 master 位点信息给新 master
        2. 位点偏移不精准
        3. 为了防止数据丢失，重放几个可能执行过的事务
            并且 使用slave_skip_errors 1032,1062 跳过duplicate,找不到数据的错误
            
      GTID: 每个提交了的事务都对应一个唯一的GTID
        1.从库B发送连接相关请求参数不再发送位点信息而是 master_auto_position=1
        2.主库A收到B的GTID集合发送不在B中的GTID对应的事务binlog
        
6. 修复数据
    1. 修复行: 使用flashback 工具追回(这里可以回顾一下 mysql undolog相关内容)
        倒序 反向操作binlog 
        预防：sql_safe_updates = on; sql 审计
    2. 修复库表: 
            使用全量备份(每天一备) + 增量日志形式 + 跳过误操作的 GTID binlog
        或者有一种延迟备份的方式:指定一个从库延迟 CHANGE MASTER TO MASTER_DELAY = N 备份，
        在时间N内发现故障能够使用延迟备份库恢复
        预防：定期全量备份; 定期演练; 账号权限分离
    
7. kill 指令
    首先 innodb_thread_concurrency 表示查询中的线程 不包括等待中的线程
       
    1. 改变线程状态
    2. 发送信号给线程
    如果线程处于block,改变状态无法感知,只能通过发送信号通知线程退出    
    3. 线程运行到每个状态的时候,会触发不同的操作(状态埋点)
    
8. join 
    eg. table a N行数据， table b M行数据 N << M
    eg. table a straight join b on a.b_id = b.id
    
    join with index
    1. index nested loop join: 
        小表a 作为驱动表会进行全表扫描 
        大表b 只做树搜索 log.时间复杂度  N + N*logM
    
    join without index
    2. nested loop join(join_buffer_size > N) 时间复杂度 N*M 
    3. block nested loop join(join_buffer_size < N)
        N/join_buffer_size * M * N
        
    join buffer size too small
    因为 join_buffer 不够大，需要对被驱动表做多次全表扫描，也就造成了“长事务”。除了老师上节课提到的导致undo log 不能被回收，导致回滚段空间膨胀问题，还会出现：1. 长期占用DML锁，引发DDL拿不到锁堵慢连接池； 2. SQL执行socket_timeout超时后业务接口重复发起，导致实例IO负载上升出现雪崩；3. 实例异常后，DBA kill SQL因繁杂的回滚执行时间过长，不能快速恢复可用；4. 如果业务采用select *作为结果集返回，极大可能出现网络拥堵，整体拖慢服务端的处理；5. 冷数据污染buffer pool，block nested-loop多次扫描，其中间隔很有可能超过1s，从而污染到lru 头部，影响整体的查询体验。

9. 临时表

[比较综合的总结](https://mp.weixin.qq.com/s?__biz=MzA3MTUzOTcxOQ==&mid=2452966046&idx=1&sn=252ced63d40fe3f0c62f7858b032cc4c&scene=21#wechat_redirect)

10. 自增主键
    1. 自增主键用完
        禁止插入
    2. 自增主键的分配
        1. 自增主键不支持回退，为了高并发 
        2. 多个会话插入自增的时候，显示个binlog 如果使用binlog_formate = statement 容易产生数据不一致的问题
        3. 自增id批量插入的分配策略，2^M
    3. 自增主键的存储
        + 5.7 及之前是存储在内存中，重启后计算最大的值 表结构中的AUTO_INCREMENT=?? 是代表下一个id
        + 5.8 之后是存在redo log 中 重启后直接读redolog

11. 数据优化器
    优化器会选择成本最低的方式去查询，所以会先解析得到可能用到的索引 并计算 每种使用方式的成本，最终选择成本低的方式执行
    1. cpu 成本
    2. io 成本
    mysql 会设置 index_dive_limit 来决定 当索引使用in 的时候最大的 in的大小，如果超过最大的in 就会使用索引估算的方式进行：估算in大概会扫描多少行
    show index from table 中的有个字段是叫做 cardinality 基数(没有重复的数字) 用这个唯一性来和总的row 来计算 平均的数 
    3. join 成本计算
    innerjoin: 优化器可以通过选择最低成本交换驱动和被驱动的顺序
    left/right join
    
    cost = 驱动表成本 + 驱动表扇出数 * 被驱动表成本 

12. mysql 的统计数据
    1. 存储方式：
        + 持久化到硬盘
        + 内存方式
         
13. innodb 数据存储的方式
    行记录:
        compact(压缩和协议): 
            元数据(组成方式):
                * column的所占用大小: 根据存入时候真实的数据和 column 类型计算出来
                * null column: null 是一种特殊的类型
                    是一个bool数组 倒序的记录当前row 对应的 null 情况
                * 记录头信息: 
                    + 是否被删除
                    + 当前所在B+ tree的位置信息和类型
                        数据类型
                        最小值
                    + next_row 指针
            数据:
                存放真实的数据
                组成部分 row_id, trans_id, roll_back_id, col1, col2 ....
        思考:
            当发生更新的时候 元数据怎么发生改变, 如果varchar 长度发生变化
        [refer vs Kafka 的存储实现方式]
        redundant:
    
    page 16k:
        * Page_Header  
            记录页面的初始化信息， page address slot num 
        * File_Header
            当前page check_sum [refer tcp check_sum] 
            sequence num
            last modify
        * Page Directer
            将user_recoder分为多个组，
            组成: slot1, slot2, slot3
            slot 的值每个组最大recorder的偏移量， 
            user_recoder : n_owned 字段表示每个组的成员量

MyISAM:
    如何理解MyISAM的索引组织形式       