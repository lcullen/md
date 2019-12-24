prometheus


SLA server level agreement 服务等级协议
公平性: 取代平均值的不公平 
(定义)用类似中位数的概念来表示服务
    * 百分位  50分位对应中位数数值 => 99分位 对应响应时间的中位数

1. 时序数据库
2. 应用
   背景 
   1.持续增长的网络io流量
   2. cpu 的使用率
        sum(increase(node_cpu(mode="{idle/user}[1m]"))) by (instance)
    1. counter 记录持续性增长指标
        func
            increase(key[duration]) 在duration中key的增量、
            rate(key[duration])  每秒钟产生的变化 -> 变化平率很大的 很敏感的数据
            #rate(increase(key[duration])) 增率 
            sum() by()
    2.  gague 
        topk(k, key) 取出最高的前k，的key
        topk 只能用在gague 或者使用 rate 和 increase 包裹的 counter 中， topk 只能查看瞬时的值
            使用场景: alter 
        count() 很mysql 中的count() 一致
            场景:
                当cpu 达到 count() > num 的时候就报警
3. 运维软件
    screen && daemonize        
4. server 启动
    * 时间校验 nptupdate
    * 日志目录
        prometheus/data 历史数据
        wal/近期的数据 冷数据 重启备份
        内存 近数据
     * 配置文件 prometheus.yml
        global:
        job_name: 
            target server:port 
     * node_export
        