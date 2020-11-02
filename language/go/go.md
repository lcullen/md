
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
    

======golang 中的 ticker 实现原理
1. 

=====go GMP 调度 [refer](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/)


2. 无锁队列的实现 
   * sync.Waitgroup 的实现方式
   * Blpop 的实现方式
    
    
    
    
    
    
    