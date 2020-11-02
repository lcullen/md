1. LB
    1. DNS [refer](https://www.ruanyifeng.com/blog/2016/06/dns.html)
        dig domain.com
        分级查询 && 有缓存
    2. 网路层: Linux visual server
        IPVS调度器的三种策略:
            * 重写请求的目标IP地址为真实服务器地址 【递归】
            * 重写请求的目标IP地址为真实服务器地址 && 真实服务器请求直接返回给客户端 【传递】
            * 重写请求的目标MAC地址为真实服务器地址 && 真实服务器请求直接返回给客户端 【传递】         
    3. 应用层
        NG
        
2. 业务状态
    1. 存储中间件: groupcache
        * define getter_func after miss
        * distribute by http peers

3. HA
    1. 故障转移fail over
        raft Paxos
    2. 超时 time out
        重视程度、一开始不重视、代码的臭味道、百分位控制
    3. 降级
    4. 限流
        