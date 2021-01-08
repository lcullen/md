相关问题:
    TCP提供了一种字节流服务，而收发双方都不保持记录的边界，应用程序应该如何提供他们自己的记录标识呢？ 业务系统如何知道 消息已经结束


1. ICMP
    查询报文类型:
    差错报文类型:
2. socket 套接字 [什么是套接字](https://www.cnblogs.com/dolphinx/p/3460545.html)
    网络中的两个进程需要通信，用socket 间接的表示了 网络中的进程，是一种应用程序的interface 抽象
    套接字是一个结构体，拥有Protocol_Family(协议族) 和Address_Family(地址族)
    Address_Family: ipv4 ipv6
    Protocol_Family: 
3. socket 在tcp 的 3次握手连接 的变化
    + bind(), listen() 会返回 一个socketfd, listen() 函数的作用是把 主动的socket可发送连接请求的socket 变为被动接受请求的socket
    + accept() 的时候会返回socket 连接成功的fd，用于返回客户端, 内核会维护这个连接fd 直到连接断开
    
    + 第一次握手: client.connect() 客户端阻塞等待返回
    + 第二次握手: server.accept() server 接受请求，在内核中创建一个连接fd 等到连接close() 的时候close() fd
    + 第三次握手: client.sendmsg() 这个时候 server 要writemsg
    为什么要三次握手: (基本至少要握手三次的原因: 各自确认自己的收发能力是ok的)
    client view 的前两次的握手可以保证自己和服务端的收发能力是ok， server view 在前两次只能确保自己的接受能力是ok的，第三次握手能够让服务端确认自己的发送能力是ok的 
    为什么要四次挥手: ()
    
4. writemsg / sendmsg [面向socket编程](https://time.geekbang.org/column/article/116043) [golang 实现](https://tonybai.com/2015/11/17/tcp-programming-in-golang/)
    + send
    + write
    + writemsg
    
    重新回顾tcp 3次握手中到底发生了什么事情
    状态的变化:
        服务端视角: close -> listen -> sync_rcved -> established
            close
                + 创建一个套接字 得到一个文件描述符 fd
                + bind fd with addr info
                + 服务端listen(fd, backlog) 
            listen
                + backlog 解读: 操作系统在listen的时候会根据 __backlog__ 去初始化两个队列 [拓展阅读:阿里中间件团队博客：关于TCP 半连接队列和全连接队列](http://jm.taobao.org/2017/05/25/525-1/) [gorefer](https://blog.csdn.net/yangguosb/article/details/90644683)
                    + sync queue
                    + ack queue
                    >> 一个小插曲: 
                        使用netstat命令可以查看当前连接的状态
                        使用nc -l port 命令 nc addr 能够模仿连接
            sync_rcved 
                    + sync queue: 用来保存客户端第一次握手的请求队列, 服务端回复相应的ack 和自己的sync 序列号 __第二次握手发起__
            establish
                    + ack queue: 当客户端回复了服务端的sync会把此时的已经建立的连接放入这个队列中，等待应用程序取走当前的连接 __第三次握手__
         
        客户端视角: [线头阻塞 head-of-line blocking](https://github.com/bagder/http2-explained/blob/master/zh/part2.md)
        
    2. TIME_WAIT 
        1. TCP 连接中主动断开一方 会有这个状态
        2. 持续时间是 2 * MSL (max segement lifetime) 
            1*MSL: 等待最近一个被动断开放的数据包过期
            1*MSL: 等待主动断开方 确认(ack) 被动断开的 fin. ack 的过期 
            总共是2次等待
        3. 太多会引发的问题：
            端口被沾满
            address already in bind 
                time_wait 阶段的seq 始终要小于最新的 msg 的seq
                    so as timestamp
            [refer1]关于tcp_tw_reuse和SO_REUSEADDR的区别，可以概括为：tcp_tw_reuse是为了缩短time_wait的时间，避免出现大量的time_wait链接而占用系统资源，解决的是accept后的问题；SO_REUSEADDR是为了解决time_wait状态带来的端口占用问题，以及支持同一个port对应多个ip，解决的是bind时的问题。
            [refer2 惊群]SO_REUSEPORT用在多个不同的socket监听在同一个端口上，这种情况比较罕见，容易出现所谓的"惊群"现象。当然，如果用的好，也可以解决一些特定场景的问题。
5. 本地套接字:
    基本上很tcp 和udp 套接字类似，但是服务端是通过有权限的对 获取绝对路径的文件 进行监听 实现的
    
6. 关闭连接  [refer](https://time.geekbang.org/column/article/126126)
    1. FIN
        Statu: FIN-WAIT-1 主动关闭发送 FIN package： the socket is closed and the connection is shutting dwon 主动方第一次握手
    2. RST
        出现的问题：
            在一个RST的套接字继续写数据: broken pipe 
        SIGPIPE信号：默认行为是终止进程，进程需要捕获并处理
            引发SIGPIPE信号的时机：当一个进程向已经收到RST的套接字进行写入数据的时候，内核会发送SIGPIPE信号，要求终止进程，写操作返回Broken pipe 错误
        Connect Reset By Peer            
    3. 套接字计数器：
        每一个进程都有拥有一个socket 引用
        close(): 会对引用减一，其他的进程还是可以访问socket, 暴力关闭 收到请求后返回 RST
        shutdown(): 直接导致socket 不可用，并对外发送fin package，等待当前消息处理完成(主动关闭方发送fin package 并等待被动关闭方返回fin 确认关闭)
    4. so_liner
        close 时 是否要等待buffer 中的数据push 完成
7. keep-alive
    TCP keep-alive
    APP keep-alive [refer](https://technologyconversations.com/2015/09/08/service-discovery-zookeeper-vs-etcd-vs-consul/)
    Golang [refer](https://draveness.me/golang/concurrency/golang-timer.html)
    Http Connection: keep-alive 中的意思， 就是 服务端永远不关闭此链接 由客户端关闭

8. 发送窗口和接受窗口：[参考tcp的流控实现方式refer](https://coolshell.cn/articles/11609.html)
    消费者和生产者模型: 消费能力控制生产能力
    TCP 需要有自我牺牲的精神, 当网络拥塞发生的时候主动让道, 具体的体现为RTT 时间变大
    
9. 拥塞控制
    解决的问题: 雪崩式拥塞
    原因: 大量package 涌入网络中，处理速度慢，滞留在网络中

10. 重传
    1. 指数退避算法
    2. 快速重传， 重复收到相同的ack的时候要通过 ack option 中找到当前ack 对应的已经收到的序列号 和未收到的序列号(sack)
    3. 超时时间设置:
        低通滤波器: 采样80% RTT + 最新的20% RTT = avg(RTT)
11. win 
    数据包中的win=?字段是发送端告诉接收端 我当前最大的接受窗口，当win=0 的时候就让对方停止发送数,对方会有零窗口试探包试探 
        试探包使用当前最大 seq - 1  len = 0 的package 等待对方确认
    有等待的地方就有攻击
12. 滑动窗口(避免接收方的拥塞): 流量控制 是由接收方来告诉发送方 当前接收方的能力
    拥塞窗口(避免网络上的拥塞): 当前发送方发送的能力，
        慢启动: 一开始的拥塞窗口很小，然后每次经过一个RTT cwnd * 2,直到 超过阈值，超过后开始 cwnd = cwnd + 1/cwnd
            发生超时重传的时候 cwnd 回到阈值水平
        快重传: 和10.2是一致的

13. 延迟确认(磨叽姐):
    要解决的问题: 减少不必要的ack 回复
    当client 发送多次seq 并且 server 要对每一次seq 进行ack，多次seq到达server 组合成完整的msg时候， server 要 response ，这个时候又要response ack, 可以把前面的seq ack 合并为 response ack， 当然有timeout

14. hold住哥 [refer mysql中的 组提交 kafka client fire and forget]
    nagle 算法: 要解决的问题是 减少小包在网络中的传输
    优势: 实现简单 劣势: 应用层无感知被hold住了，如果这个时候os down，或者本来就是想发送一个小包 又被hold了(对时延比较敏感的应用 不适合)
    具体的实现方式或者说是send条件
    1. 当前的 发送队列中不存在 等待ack 的seq，
    2. 当前的 等待发送队列的大小 达到了 mss max sequence size
    3. 最多能hold多久
    
当 13 meet 14 时延会非常的高    

15. before UDP send package 
        try to connect server for testing arriveable icmp

16. TCP的可靠性
    只是针对TCP层的可靠， 并不负责应用层的可靠性
    reset by peer && broken pipe  一个是tcp 错， 一个是管道错, 一个rst 一个是signal
3. 信息安全
    + DNS 劫持
        传统的DNS 解析过程 client -> local_hosp_map -> 运营商local_dns (不具备权威) -> 13台DNS server 进行递归解析直到找到自己负责的对应域名，如果找不到 最后由顶级root dns 负责指派 能处理这个dns 解析的server 
            dns 解析分为 
                迭代解析 dns server 无法解析当前dns 就告诉client direct to ask another target dnsserver
                递归 dns server 充当客户端
        1. 串改目标域名: 两个服务通过相互的公网域名进行同行就会产生安全问题
        2. LocalDNS: 解析的是运营商本地的旧的DNS 内容
    + 解决方案
        HTTPDNS 基于http协议 通过http 获取目标域名对应的用户真实的出口ip
    
    + 存储
        进行端到端的加密，端利用对方的 公钥对消息进行加密，接收放通过自己的私钥进行解密，中间的存储层存储的也是加密之后的信息

//-=-=--=-
1. select: 程序疯狂 span 用select 扫描所有的 fd，等fd 满足条件后，处理并且 fd_set fd，等待应用程序span 被处理的 fd
    event
   poll: 对fd 的数量有限制，普遍是1024
   
   以上本质上是监听所有的句柄

2. select/poll -> epoll 过渡的原因: fd的限制
      select: 
        1. 程序要轮询当前的fd数组看事件是否满足
        2. 处理完满足事件的fd 要重新reset 自己感兴趣的事件类型
      poll: 
        1. 还是要轮询 但是拥有了自动扩容的机制
        2. 标记自己感兴趣的事件类型
      epoll:
      
      epoll 为什么 fd 的增长 对其效率没有影响？ 
        epoll_wait 在不同 triggered 返回值:
            level triggered 条件触发: 每次wait 都会 通知 没有处理的 event
        edge triggered 边缘触发
      
      拆解epoll: 核心组件,在内核注册对fd的事件,事件发生内核通过回调函数找到对应的epoll实例,并查看当前用户所关心的事件, 并加入到用户空间的事件完成的queue队列中, 并唤醒用户进程
            epoll 中有一个 红黑树,每个节点是 epitem,
        + 如果当当前的需要被监听的fd 通过epoll_ctl 加入到 epoll 的 红黑树中
        + 成为一个epitem节点,并会产生一个 epoll_entry，
        + 通过这个epoll_entry 能够找到对应的 epoll fd的实例和该实例对应的 epitem.
        + 将每个fd 都加入到rbtree 并注册自己的ep_call_back 处理func
        + 当内核满足当前fd的时候会反向调用 已经注册的ep_call_back , 并找到当前的epoll 实例句柄
        + epoll 实例 filter 对应的时间 如果是用户关注的时间则加到 就绪队列中等待用户 获取
          如果是level trigger 的话 会重新反复的放到就绪队列中 等待用户的获取
        

3. 网络中发包 IP头的变更Mac地址的变更条件:
    A -> B 不同的ip网段不同IP 每次改变目标MAC 地址 而不改变IP 地址。
    A -> B 不同的ip网段相同IP 改变MAC IP

//-----
GRPC server 侧源码解读
比较难理解的一点:
    把单次请求的request 封装成为 setting frame, header frame 、data frame 
    在header frame 用对应的serv 处理请求， 等待data framer的到来，data frame是 blocked的进入recvbuffer 中，recvbuffer 再阻塞的get
    
4. HTTP2 /HTTP1.X 
    1. I/O 多路复用 同一个http 连接上可以 被不同的frame 复用
    2. 二进制传输 使用数据frame
    