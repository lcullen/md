1. rpc (通讯框架、通讯协议、序列化方式)
    1. 建立、管理连接
        形式:
        + http 
            基于tcp协议、三次握手连接 grpc 是基于http 协议的
        + socket
        
        存活检测: 
        断连重试:  
    2. 处理请求
    3. 传输协议
    4. 数据序列化反序列化

back grpc 如何做到上述 client 端 
get conn process: 
1. init all dialOptions 
    比较常用的 dpt:
        + unaryInt
        + scChan 用于异步lb 配置 最终会落实到 balancer 配置
            + load balance  如何与 connection 同步信息
                1. lb 异步监听 etcd watcher 如果发生变动 维护自身address slice
                2. 同时维护全量地址表到 自己的 addrCh 通道 并暴露接口 单向通道通知 connection (单向通道用于签名表示只读，或者只写)
                    tips: 通过chan 解耦 lb 与 connection，cc 通过hang 住 lb的 通知来更新 addConn slice
                    
                3. 通过 上述的 address 创建每一个 addressconnection 通过tcp 进行连接 获得transport
                4. 监控transport 
                
        + backoffStrategy 重试机制:
        + health check:
2. 责任链模式: invoke()
    概念:
    1. UnaryClientInterceptor 
        ```go
           //返回一个 interceptor 的函数
           		return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
           			buildChain := func(current grpc.UnaryClientInterceptor, next grpc.UnaryInvoker) grpc.UnaryInvoker {
           				return func(currentCtx context.Context, currentMethod string, currentReq, currentRepl interface{}, currentConn *grpc.ClientConn, currentOpts ...grpc.CallOption) error {
           					return current(currentCtx, currentMethod, currentReq, currentRepl, currentConn, next, currentOpts...)
           				}
           			}
           			chain := invoker
           			//包装每一个 interceptor 为上一个 interceptor 的 invoker
           			// 为每一个 invoker 添加上一个 invoker
           			for i := len(interceptors) - 1; i >= 0; i-- {
           				chain = buildChain(interceptors[i], chain)
           			}
           			//返回最终可顺序执行的 interceptor 函数
           			return chain(ctx, method, req, reply, cc, opts...)
           		}
        ```

server 端:
0. 服务启动并注册服务 等待请求的到来
1. 从context中获取请求的基本信息
    * 获取trace 结构体 并重新包装 含有trace的context
2. incoming stream: grpc 中 使用stream 来代表请求
    * 请求打过来 从已经register 好的map 中获取服务信息 如果没有找到服务 使用 unknownStreamDesc
    * 找到对应服务 如果找到了 普通 method 使用 processUnaryRPC 返回;如果是stream 方法 使用 processStreamingRPC 返回
         
=========
服务治理
1. 

微服务:
I. 限流策略和设计:大量的请求导致系统过载
    GRPC有这种能力么？
    1. 限流: 在滑动窗口内最大的请求数量, 参考rrt 时间内的最多请求数量, 如果超过了当前的数量 处理方式: 
    2. 熔断: 在滑动窗口内最大的失败阈值是多少
    3. health:
    4. slb: 带权优先级队列
    以上这些都是依赖于采样打点分析 也是可以参考tcp 的流控去实现的 [refer](https://coolshell.cn/articles/11609.html)
        
