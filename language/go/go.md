
======
atomic

不同的go routine 上下文的切换 可能导致非原子性操作出现
0. add，cas，load，store，swap
1. add
    传入参数为被操作数的指针，差量要避免溢出问题
2. cas

==============
Context 
 用于实现 一对多 goroutine的 协作流程 控制goroutine
   context 包含了四个用于繁衍Context值的函数
   1. WithCancel
   2. WithDeadLine
   3. WithTimeout
   以上三种都会自动脱离father Ctx 并且关闭 doneCh 用的是深度优先策略
   4. WithValue

==============
CSP
//communicating sequential processes
4. CSP 模型[参考](https://draveness.me/golang-channel)
    channel 是多 goroutine 之间的沟通桥梁
    结构体: 
    ```gotemplate
    type hchan struct {
    	qcount   uint //当前 channel 中元素的个数
    	dataqsiz uint // 循环队列的 长度
    	buf      unsafe.Pointer // 缓冲区大小
    	elemsize uint16 //元素大小
    	closed   uint32 // 关闭字段 flag
    	elemtype *_type //元素类型
    	sendx    uint  // 一个循环 数组
    	recvx    uint  // 当前 接收 数据的 位置，和消费数据的 位置
    	recvq    waitq //接收 goroutine 被阻塞队列
    	sendq    waitq //发送 goroutine 被阻塞队列
    
    	lock mutex // 对队列的操作 要加锁
    }
    ```
    
    1. 创建
        根据make 参数 创建对应的 hchan 结构体
        
    2. 接收
        1. 当前 channel 创建的时候 没有 buffer池(也可以不发生阻塞行为)
            这里涉及到 如果已经有 goroutine B 向这个channel 发送 数据，
                那么这个goroutine 会发生一次 goroutine调度。B 会把自己的 goroutine pending 起来
                并把 B 放入 hchan 的 sendq 队列中
            goroutine A 从 channel 中读数据，会直接把 B 的数据拷到A的 指定 地址，此时B 又发生 调度 将 B 放入P 的 next 可执行队列中
            
            除了上述B已经发生pending的状态外，如果没有数据到达，A 将发生pending 调度， 同时A 被阻塞
        2. 有buffer
            2.1. buffer 中没有数据， A 放入 recvq 队列 pending
            2.2. buffer 中有数据, 发生数据拷贝 并且 更新recvx 位置
            
        __提问__: 可放入 chan 中的数据结构 有哪些？ func ? 
    3. 发送(基本上和接受是差不多的)
    4. 关闭
        __提问__: 当前  sendq 和recvq 任何一方还有等待操作，如果直接被关闭会发生什么？
        
goroutine 暴涨:
    内存吃紧，GC, __Golang调度器__ 负担 
进程 线程 协程
1. 进程是资源的隔离单位
    演化进度:
        * 出现了 epoll 和 poll select io 多路复用的模型
            但是没有解决 数据拷贝(内核态和用户态的拷贝)和 上下文切换的成本问题, 只是减少了线程数
2. 线程是进程里面的多任务执行体, 线程间的调度也是在内核空间发生的
    * 空间成本
    * 时间成本
3. 
实战应用[refer](https://mp.weixin.qq.com/s/CGLWKawX7qTaMqIMWqmptA)


======Container && Sort
Heap (从最后一个叶子节点的 父节点开始 down 进行排序)

工业级的系统排序
    时间复杂度: 可以考虑使用 O(n^2)
    原地排序: 插入排序
    if O(n^2) < O(knlogn + c) : 插入排序
    else set recursion length 
        if recursion length == 0 
            use heapSort
        else 
            递归
go 中的sort 
    1. if hi - lo > 12
         define maxDepth = 2*lgn
         use quickSort(lo,hi,maxDepth) {
            if maxDepth == 0 then use heapSort(lo,hi) and return
            maxDepth -- 
            ml, mh := partition() {
                pivot = getPivotBy4Couple()
                return ml, mh // ml,mh 是重复pivot的起始结束边界
            }
            use quickSort(lo,mi)
            use quickSort(mh,hi)
         }
    2. if hi - lo between 1 and 12
        use gap = 6 shell sort
        use insertSort again
    
======panic && defer==== [参考](https://draveness.me/golang-panic-recover)
1. panic 是一个linked 结构体，当panic 发生的时候 会循坏调用当前 goroutine defer，如果defer 中有关键字 recover
    goroutine 就能 拯救 否则 defer 调用完之后 直接 crash goroutine
2. defer 结构体也是一个通过linked的 结构，内置了 defer 需要调用的func
   defer 不是原子的 [refer](https://segmentfault.com/a/1190000006823652)
        先对return xxx 进行 xxx 的赋值 然后 再执行 defer 语句， 然后再真正的返回

=====go concurrency == 
1. 并发的副作用
2. 基本的同步原语
    1. mutex
    


I. go GMP 调度 [refer](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/)
goroutine 是较线程 更加轻量级(占用的上下文空间更加的小切换更加的容易)的用户态 协程, 添加goroutine 提高了程序的并发能力
但是应用程序要进行对goroutine的调度和协调, 避免出现饿的饿死, 旱的旱死. [tony bai](https://tonybai.com/2020/03/21/illustrated-tales-of-go-runtime-scheduler/)
G = goroutine
P = goroutine的运行空间
M = 真实的程序线程数
    * 会遇到的问题: 饥饿
    * M 、G 的阻塞: M的系统阻塞 G的chan 阻塞
    * 公平性:  每个超过运行10ms 都会被标记为可抢占的 sysymon 后台协程可以直接强制使超过20ms的G被移除 (goroutine中自带了runtimetimer计时器 time.sleep 用的就是这个)
    
II go channel 底层实现 or 数据结构 [refer](https://codeburst.io/diving-deep-into-the-golang-channels-549fd4ed21a8)
[用cond实现一个channel](https://time.geekbang.org/column/article/96994)
    channel 为的是解决 并发过程中 竞态资源的问题, 用communication 来代替share data
    其实channel 是屏蔽了lock 的细节: 锁的粒度和锁的时间问题, 把锁的控制都放在了channel 中了
    如何用sync.Cond 实现一个channel ? 
    核心结构
    1. sendq 和 rcvq

2. 无锁队列的实现 
   * sync.Waitgroup 的实现方式
   * Blpop 的实现方式
    
3. timer 的实现细节: [refer](https://www.cyhone.com/articles/analysis-of-golang-timer/)
    runtime中全局维护 64bucket, 每个bucket 都是一个四岔最小堆, 堆顶是距离现在最近的时间


II. go pprof 
    1. 

基础:
1. 接口
    1. 空接口(eface) 和其他接口(iface)
        * 空接口 里面会有 2个变量分别用来表示 当前的类型 和当前的类型数据
        * 有方法的 接口
    2. 空结构体: 没有成员变量 不占用内存, 一些场景并不想占用内存的场景, 向chan 发送信号
        * chan 当做信号的使用 chan<-struct{}{}
        * map val 是空结构体 当做set 使用 map[key]struct{}
        * 都是同一个内存地址 var s struct{}  var b struct{}  &b == &a
        
    3. interface 的比较
        只有当动态值类型一样 && 动态值一样  
2. 传参:
    值类型     复制的是值类型
    引用类型的 复制的是指针
3. equal 相等 判断相等 判等
   1. 

4. chan
    1. chan 的三种操作 && 3中状态
        1. 读、写、close
        2. closed、active、nil
        
    2. chan 的个人理解
         * 无buffer 能够当做同步使用, 有buffer 可以当做异步
         * 并发场景的使用chan 实际上也是用了 hchan 中对recvq 和senq read write 做锁操作的，
            把锁的内部细微实现进行了包装 而不用交由用户去处理lock 和 unlock的方式 
            
5. timer 各种设计的区别 和选型问题  [golang timer](http://xiaorui.cc/archives/6483)
    1. timer 和当前的协程 挂钩匹配

6. slice 的了解程度:
    是对数组的封装,指向了底层的数组, 并保持了当前的 容量和大小    
7. copy: 深度拷贝互不影响
8. map && sync.Map 
    * map 哪些类型能作为key， 可比较的类型都能作为key
        底层的hash 冲突怎么解决的: 
        扩容: copy on write
    * syc.map 使用readyonly 和 dirty read 来提高map并发能力, 所有的读先落到 read 中， 如果没有找到则只能加锁读取dirty段
        当miss 的数量达到一定的程度， 说明总是dirty 读取， 需要提升dirty 段到read 中 防止过度的miss 
9. log pkg: 自带的log.output 是线性安全的
10. http pkg: 路由的实现

11. heap  && stack  && 队列的 内部实现
     1. 在container pkg 中有 heap 的实现
         
     2. 

10 . go build 和 go install 分别发生了什么

11. pprof的 __性能调优__: [refer](https://www.bilibili.com/video/BV1iA411i7Nt?from=search&seid=10157390214658424535)
    * go tool
    * go test -bench=. -benchtime=3s -run=none
    *  [pprof和benchmark的结合](https://my.oschina.net/solate/blog/3034188)

12. go rune 类型是什么 具备int32 的所有能力, 但是又是一种能够支持字符串的类型
    * 

4. 对于新的go泛型的理解
    * 什么是go泛型
    * 为什么需要go泛型
    * demo
    
13. 实际场景使用chan [refer](https://segmentfault.com/a/1190000017958702) [qiyingyiji](https://lessisbetter.site/2019/06/09/golang-first-class-function/#%E7%89%88%E6%9C%AC3)


如何理解unsafe.point
[refer](https://www.cnblogs.com/qcrao-2018/p/10964692.html)



内存模型 memory model [refer](https://go.dev/ref/mem)
happens before : 读写成为一对，写在前，读在后，并且中间没有任何的写操作
1. 
   1. 编译器会导致内存变量的重排
   2. 读写不能保证一致性 需要 加上同步元语
2. 天然 happens before
   1. channel
   2. lock

go tools:
1. race
2. gcflags
3. mock tools


====
逃逸分析: 
1. 构建变量，用最安全的方式分配内存
   1. 堆上分配的指针不能指向栈
   2. 栈上的指针不能指向已经被回收的栈内存
2. goroutine 的分配发生在栈上，默认4k，可以自动扩缩
3. 

---曹春辉--
1. GMP 调度组件&&调度计算
    * 程序的局部性: runnext 去解掉
    * 生产者和消费者模型
        生产
            * 优先级 runnext(指针)-> local(数组) -> global(链表)
            * 新生成的goroutine 优先级大于 老的，优先塞入 runnext，如果有值，剔除老的，将老的放入local，localarr 如果已经满256，剔除一半 生成batch 放入global
            * runnext 是为了解决程序的局部性，可以往cpu core 缓存扩散的方向去考虑
        消费(四大法宝)
            * schedule
            * runtime.execute
            * runtime.gogo
            * runtime.goexit
            --
            首先从runnext获取 -> local 随机%取一个
        消费的时候遇到阻塞问题怎么解:
            * chan,read,select,sleep,lock
            
2. 编译
    refer: golang.design/gossa 查看编译过程 https://golang.design/gossa?id=1d5bd23e-caa6-11ec-adb4-0242ac16000d
    或者godbolt.org
    go tool compile -S ./helle.go | grep "hello.go:5" //stringstoslicebyte
    https://github.com/x-motemen/gore
    参数调用规约: https://github.com/acodercat/function-call-principle
    
    goyacc grammer.y 语法树相关
 
 3. 常用数据结构
    * channel
        * 数据copy 是怎么操作的  
        * 并发安全
            * chansend,chanrecv 都成对的加锁
        * gopark goready
        
    * timer 四叉堆 老版本:
        -> 分片锁 多个四叉堆
 4. system call
    * 复习 函数调用规约
    
 5. 作业:
    * tryLock with retry timeout
    * dead lock
    * 时间轮，timer cpu high
    * https://books.studygolang.com/gopl-zh/ 互斥锁
 
 6. 内存分配与垃圾回收
    路线: 
        * 内存从哪里来
        * 垃圾从哪里去
    GC
        * 标记对象怎么来的
            worker buffer 往p.gcwoker 里面推送
            mutator 写屏障中推送当前对象
        * 标记对象到哪里去
    栈分配和堆分配 有什么区别，为什么要区分，或者逻辑化这种
        * 栈分配是 轻量级的操作，只移动栈针，和赋值，最后返回的时候直接释放
        * 堆分配: -gcflags="m"
            * 逃逸分析: 所有的场景分析 -> http://github.com/golang/go/tree/master/test  escape_xx.go
    内存分配
        * 顺序分配
        * 链式分配: 如何对内存切块 (分级分配方式)
            一下是内存碎片的产生和优化
            first-fit
            next-fit
            best-fit            
    Go 内存分配 使用分级分配 && mmap的方式获取内存快
        arena block 为64m 内存块
            mcache 绑定p 中
            mcenter 中心化 缓存
            mheap arena 向操作系统 申请
        分配方式:
            tiny 多级缓存
            small
            large

    垃圾回收:
        * 语义垃圾: slice 缩容
        * 语法垃圾:       
            * 触发的方式: 用户主动，后台，每次内存的分配
            * gcworder 去处理标记 问题, 推送到本地的        
            __gopark__  __goready__ 需要确认
            * cpu 低于25%
            * 所接受的挑战:
                * cpu 控制 
                * 三色标记对象漏标:  (写屏障)
                    * 强弱三色不变性 


===框架层面需要理解的
 * 皮裤套棉裤
 * 路由怎么找=》 
    * 字典树 trie
    * radix tree