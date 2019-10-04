1. 任务队列的设计 
[参考1](https://mp.weixin.qq.com/s?__biz=MjM5ODYxMDA5OQ==&mid=2651959957&idx=1&sn=a82bb7e8203b20b2a0cb5fc95b7936a5&chksm=bd2d07498a5a8e5f9f8e7b5aeaa5bd8585a0ee4bf470956e7fd0a2b36d132eb46553265f4eaf&scene=21#wechat_redirect)
[参考2](https://mp.weixin.qq.com/s?__biz=MjM5ODYxMDA5OQ==&mid=2651959961&idx=1&sn=afec02c8dc6db9445ce40821b5336736&chksm=bd2d07458a5a8e5314560620c240b1c4cf3bbf801fc0ab524bd5e8aa8b8ef036cf755d7eb0f6&mpshare=1&scene=1&srcid=0316rh7QmkSKJH06XFENtsgw#rd)
    - go machinery  
        延时任务 常规套路
        1. retry 机制 
            什么情况下需要重试:
        2. 任务实时消费 
            list 队列 实时消费 采用blpop 阻塞式方式获取实时任务
        3. 延迟任务 delay task 
            通过将未来执行时间当做 score 进行排序,
            获取zset 中 已经可执行任务
            zset score = future timestamp zrange by (-inf,now]
        4. task args 序列化方式
            启动时预注册map[string]func handler(不用反射:可以跨服务send task)
            序列化之后传参找到对应handler 处理
        
        go machinery feature 
        1. server 是整个machinery的core: 简单分析一下 设计模式 
        
            - 单一职责: 用 __职责__ 或 __变化原因__ 来衡量接口是否设计合理,但是职责和变化原因不可度量,
                __建议:接口一定要做到单一职责，做到只有一个原因引起状态变化,将多职责的接口组合到一个具体的类中__
            
                可提供 publish 功能的 driver machinery 支持 MongoDB Redis Memcached MongoDB provider
                interface __IBroker__ publish && consume 引起队列的变化
                    IBroker 维护自己的 
                        * config 属性(用于connect provider)
                        * publish 
                        * consume(会传入对应的TaskProcessor)
                interface __TaskProcessor__
                将 consume 抽象成为 interface __Worker__
            - 依赖倒置(面向接口编程,也有一种说法叫做design by contract): 模块之间的依赖通过抽象发生,实现类之间通过接口进行依赖;
                * Server 类与 Broker 类 通过 接口 IBroker 解决依赖关系
                * 实现类 worker 实现 TaskProcessor 接口 解决 对应 依赖模块 IBroker.consume方法
                * 使用setter 注入依赖对象 broker,
                
        启发:
        1. 如何做到gracefully stop worker consuming msg
            golang 持有的特性 syc.WaitGroup 
            用一个衍生的promise 模式来阐述 syc.WaitGroup 的作用
        
    - go-work 
      定时任务 HashedWheelTimer driver redis 具体的实现
        
        
        
        
2. 滑动窗口的设计