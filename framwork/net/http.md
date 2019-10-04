1. ICMP
    查询报文类型:
    差错报文类型:
2. socket 套接字 [什么是套接字](https://www.cnblogs.com/dolphinx/p/3460545.html)
    网络中的两个进程需要通信，用socket 间接的表示了 网络中的进程，是一种应用程序的interface 抽象
    套接字是一个结构体，拥有Protocol_Family(协议族) 和Address_Family(地址族)
    Address_Family: ipv4 ipv6
    Protocol_Family: 
3. socket 在tcp 的 3次握手连接 的变化
    + bind(), listen() 会返回 一个socketfd listen() 函数的左右是把 主动的socket可发送连接请求的socket 变为被动接受请求的socket
    + accept() 的时候会返回socket 连接成功的fd，用于返回客户端, 内核会维护这个连接fd 直到连接断开
    
    + 第一次握手: client.connect() 客户端阻塞等待返回
    + 第二次握手: server.accept() server 接受请求，在内核中创建一个连接fd 等到连接close() 的时候close() fd
    + 第三次握手: client.sendmsg() 这个时候 server 要writemsg
4. writemsg / sendmsg
    
    

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