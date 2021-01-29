

补充:
elastic search 基础能力
    1. index = database, type = table , doc = record , fields = column
    2. HA: 一个index 主从分片 && replica  主分片被确立的时候无法修改。 和search:  的路由有关系, 路由到具体的主分片上. 并且搜索当前主分片的所有副本拿到最新的数值
    3.  
    
GRPC客户端的负载均衡策略:
    谈谈GRPC的负载均衡策略，grpc官方并未提供服务发现注册的功能实现，但是为不同语言的gRPC代码实现提供了可扩展的命名解析和负载均衡接口。其基本原理是：服务启动后grpc客户端向命名服务器发出名称解析请求，名称会解析为一个或多个ip地址，每个ip地址会标识它是服务器地址还是负载均衡地址，以及标识要使用的那个客户端的负载均衡策略或服务配置。客户端实例化负载均衡策略，如果解析返回的地址是负载均衡器的地址，则客户端将使用扩展的负载均衡策略，反之客户端使用服务器配置请求的负载均衡策略。负载均衡策略为每个服务器地址创建一个子通道，当有rpc请求时，负载均衡策略决定那个子通道即grpc服务器将接收请求，当可用服务器为空时客户端的请求将被阻塞。
    这种方式好处是灵活，支持扩展，可以扩展满足自己需求的策略，缺点是需要自己集成，需要一定的工作量，对技术人员有一定的要求。


程序设计:[refer](https://github.com/donnemartin/system-design-primer/blob/master/README-zh-Hans.md)
I. 定时器的设置:
    1. redis 时间事件设计思路 : 
        
    2. kafka 延迟任务的设计思路: 
    
    3. goland timer的实现思路:
    
II. 秒杀场景 
    1. 只有同一个竞态资源(目标: 减少无效请求) 
        随之而来的问题: 
            * 带宽吃满: 流量不一定要真正的打到后端, 使尽可能的减少请求
            * 机器负载升高， 影响机器上其他的服务, 秒杀服务必须分开部署 分开部署服务
        解决方案: 
            *  取巧的使用随机值, 只有部分用户的流量能极小概率的 打到后端
    2. 不同的竞态资源 12306: 
        * 分流处理

* 对关键词进行自我的解释, 
* 出现关键词的历史原因是什么， 是为的解决什么问题

III. Feed [refer](https://www.jianshu.com/p/990a9316656a) [refer](https://www.bookstack.cn/read/ddia/spilt.3.ch1.md)
    场景分析: 
    读多写少的场景
    1. 用户关系: 邻接矩阵
    2. 消息的推拉: 读写性能的分析
        写入: memory cache 只写入最近活跃用户的队列中 并且只保留最近30天的数据 
        拉取: 拉取巨星数据并合并到当前页面中
    3. 
    总结一下: 业务 
IV. 短链的生成 [refer](https://hufangyun.com/2017/short-url/)

V. 规则引擎的设计: 

VI. oop 呼叫中心:


项目经验: 
I. 
    1. rule-engin 核心思想: 解析字符串算术表达式: 
        * 添加最小子有状态的单位: key, val, optype
        * 将最小单位通过 逻辑运算符 连接: && || ! 得到逻辑表达式, 解析表达式 使用调度场算法 [refer](https://liam.page/2016/12/14/Shunting-Yard-Algorithm/)
        * 将逻辑表达式逗号分隔连接 形成优先级队列
        ```golang
            opWight := map[string]int{
                "and" : 10,
                "and not": 20,
                "or": 5,
            }
            
            postReg := ""
            stackReg := []string{}
            for _, v := range s {
                curS := string(v)
                if curWight, ok := opWight[curS]; ok {
                    for len(stackReg) > 0 && opWight[stackReg[tail-1]] >  curWight {
                        postReg += stackReg[len(stackReg)-1]
                        stackReg = stackReg[:len(stackReg)-1]
                    }
                    stackReg = append(stackReg, curS)
                } else {
                    postReg += curS
                }
            }
            
            numStack := []int{}
            ret := 0
            for _, v := range postReg {
                curS := string(v)
                if n, err := strconv.Atoi(curS); err == nil {
                    numStack = append(numStack, n)
                } else {
                    n1 := numStack[len(numStack)-1]
                    n2 := numStack[len(numStack)-2]
                    op := curS // + - * /
                    switch case : 处理
                    numStack = append(numStack, n1 op n2 )
                }
            }
        ```
        
    2. strategy-engin
        * 添加最小执行体的定义
        * 根据满足的规则配置对应执行体的执行顺序
        * 返回结果
架构模式: cpu 密集型，读多写少的应用 

II. 
    1. fsm: 结合业务场景寻找开源的解决方案
       改造: 
        1. 参考zap log 中的sync.Pool的使用:  动态缩容当前的对象管理
        2. 参考GRPC option 的入参方式
        
    
跨系统数据一致性的问题:
or 跨数据中心的数据一致性问题:
    1. __避免冲突__ 由客户端处理路由写入到自己的 "家庭中心" 数据, 如果用户地址发生改变, 需要考虑用户的双写操作
    2. __处理冲突__ 
        * 版本号 LWW
        * 
    
    
处理线上紧急问题的能力:     
    * 监控异常业务数据 
    * 

[超强无锁队列的实现](https://zhuanlan.zhihu.com/p/24432607)
无锁队列Golang的实现: (中心思想 自己锁自己 原子操作) 暂时不会做 [refer](https://coolshell.cn/articles/8239.html)
```go
    package freelockqueue
    import (
    	"sync/atomic"
    	"unsafe"
    )
    type Block struct {
        p unsafe.Pointer
        next *Block
    }
    type Queuex struct {
        head, tail *Block
        lenght int32
    } 
    
    type LockFree interface {
        Enqueue(interface{}) bool
        Dequeue() interface{}
    }
    
    func (q *Queuex)Enqueue(data interface{}) bool {
        b := &Block{p: unsafe.Pointer(&data)}
        p := q.tail.next
        for {
        	for p.next != nil {
        		p = p.next
        	}
        	if ok := atomic.CompareAndSwapPointer(&q.tail.p, nil, b.p); ok {
        		break
        	}
        }
        atomic.CompareAndSwapPointer(&q.tail.p, nil, b.p) 
    }

```

II 自我认知 和企业文化的认知 
    * 优点: 
            + 对自己负责的事情 负责到底 owner 意识 [数据产出 && 程序代码设计和实施], 
            + 团队协作, 自驱力: 对系统健康度的调优 
            + 不满足于现状
            + 追求完整性和细节
    * 缺点:
            + 喜欢追求细节导致项目／作业未能按期完成。通过时间管理能力改变工作方式，先完成框架再改善细节得以解决
                可能会浪费一些时间。针对这一不足，我会经常对比一下我的工作现状和职业目标，确保自己没有把太多时间花费在某一个项目或不必要的工作上
            + 不知如何拒绝，同事要求帮忙一概揽下，影响自身工作进度。
                通过多任务处理能力设定优先顺序，以该优先顺序表向求助同事展示自己手上工作，并给其一个自己在何时可以给予帮助的时间估计，让求助人自行决定是否求助，问题解决