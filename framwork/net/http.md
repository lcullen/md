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
                + backlog 解读: 操作系统在listen的时候会根据 __backlog__ 去初始化两个队列 [拓展阅读:阿里中间件团队博客：关于TCP 半连接队列和全连接队列](http://jm.taobao.org/2017/05/25/525-1/)
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
            1*MSL: 等待主动断开方 确认(ack) 被动断开放的 fin. ack 的过期 
            总共是2次等待
        3. 太多会引发的问题：
            端口被沾满 
        
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