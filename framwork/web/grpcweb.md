GRPC vs restful 
1. what
    web client 端能过像调用本地函数一样调用远程方法
2. why
3. how

GRPC 一些小的知识点和技巧

+ 单reactor 多goroutine 管理连接处理请求 
    accept 是一个阻塞方法用于 接受连接，每接受连接后 启用 goroutine 处理请求 

+ server 端维持一个 waitGroup 可以用作 gracefully shutdown
    等待所有的请求处理完毕 才真正的 关闭整个server

+ GRPC 基本流程
    + server
        如何路由找到服务(启动时初始化全局路由处理映射)
        1. 根据proto注册对应服务,放入全局map[route]*UnaryServerInterceptorHandler，等到request的时候可以根据路径 找到对应handler
            两种服务描述 
            * unary rpc 
            * stream rpc 可以当做双工，直到 eof 错误发生，才表示结束
            UnaryServerInterceptor 拦截器模式 
         在实例化server，new server 时可以自定义拦截器
         
    + client 实现进程内的负载均衡策略 使用客户端维护 选择可用连接
        1. 初始化 clientConnect cc 的时候 一并 使用各种 Builder 外包装Wrapper 的的方式 初始化了
            接受当前addressConn，当前addressConn 发现没有 balancer 为空 默认使用 pickFirst lb
        2. lb 是如何更新新加进来的 addrConn 
        3. lb 是如何选择当前使用哪个连接
        4. 当故障发生时 lb 是如何剔除 不可用连接的
    + GRPC-Web
    
Go web [http](http://cizixs.com/2016/08/17/golang-http-server-side/) 


1. 关于责任链路pattern:
   * 