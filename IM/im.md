1. 通过视屏id 等用户点击的时候再进行下载

#== 传统聊天系统的 基础能力
    接入层:
        协议先行
        保持连接
        session控制
        消息推送
2. logic_ack vs http_ack
    业务层:
        im_server send msg 的时候把没有ack的消息放入 unconfirm cron queue 中 如果收到client ack 则从queue中删除
        client 端维护 ver 保证消息的唯一性
    tcp/ip 层
        tcp 的可靠传输只是针对于tcp层，而对于应用层
3. 时序基准 [分布式唯一id](https://juejin.im/post/5a7f9176f265da4e721c73a8)
    1. snowfake 时间回拨问题
        snowfake : 41bit时间戳 + 10 bit工作机器id + 12bit序列号
    2. left-segment [refer](https://tech.meituan.com/2017/04/21/mt-leaf.html)
        如果发生时间回拨的问题 看时间相差多少 如果可以容忍 则等待时间追赶，如果不行，则更换可用server_id
    3. 基于 ring_buffer 解决方案[refer](https://github.com/baidu/uid-generator/blob/master/README.zh_cn.md)
        结合db
    4. sequential consistency (顺序一致性)
    5. causal consistency (因果一致性)
4. DNS 解析相关的 安全相关的放在 http.md 中

5. 未读数的设计
    并发的时候要考虑到原子操作 

6. 心跳检测
    NAT network address translation (运营商内网ip 和外网ip 的一个映射表，使内网ip能够与外网通信)
    连接保活
    1. TCP 层的keepalive
    2. 应用层的keepavlie 实现方式
        client : 发送消息之前 判断 上次心跳是否在范围内，否就发送heartbeat 检测是否通, 否则断开重连
        server : ~                                                          ~ 断开连接
