data intensive design
1. reliability 可靠性
    + 功能正确
    + 可容错
    + 资源可用
    + 授权功能
    
    * 一句话总结这个可靠性：可持续正确运行，及时错误发生
    * 通常很严重的bug是因为没有处理号 error handler

2. scalability 延展性

3. 数据应用场景
    + 存储
    + 缓存
    + 搜索
    + 流式处理
    + 批处理
4. 描述系统的可拓展性
    + 对负载的描述：
        每秒的读写量是多少
    + 程序性能的描述
        吞吐
        平均响应时间 [常用的统计方式 refer](https://prometheus.io/)
    + 如何提升系统的可拓展性
        水平拓展 垂直拓展
5. 可维护性
    解决系统的复杂度: 通过更加高层次的抽象

###数据模型和查询语言
    -- 语言的边界就是思想的边界
6. 网络模型
    1. 对比层次模型:
        层次模型一个很古老的模型 就是一棵树
        网络模型是一个节点可以有多个父节点(图)
7. 文档模型
    1. 无模式: 没有固定的schema
    2. 读时模式: 只有在读取到的时候能感知到模式
    3. 写时模式: 在写的时候就要符合模式 两者之间的关系是 动态和静态的关系
    
8. facebook() 的时间线设计
    1. 描述负载
        用户行为发生对系统的压力
            write:
                avg: 4.6k/per_second 
                top: 12k/per_second
            read:
                pull: 
                    avg: big_sql
                    top:
                push:
                    avg: 4.6k/per_second * follower
                    top: 12k/per_secode * follower
    2. 描述性能
        SLA: 
    3. 应对负载的方法:
        1. 加机器
        2. 加机器

9. 存储 && 索引的发展史
    1. 文本存储k-v
        ```bash
           #!/bin/bash
           db_set() {
               echo "$1,$2" >> db_file.txt
           }
           db_get() {
               egrep "^$1" db_file.txt | sed -e "s/^$1,//" | tail -n 1
           }
        ```
        弊端: 
            * file_sort && condition filtering [refer mysql file_sort && condition filtering] 解答 mysql对于已经使用索引进行 filter 完的数据要使用猜测的方式 过滤完 condition，这个时候可能需要回表查找
            *  