es
学习过程[参考](https://gitee.com/geektime-geekbang/geektime-ELK)

1. 关于es 的落盘过程: [refer](https://time.geekbang.org/course/detail/100030501-112076)
2. process:
    1. query
    2. fetch
潜在的问题:
    1. 性能问题
        深度分页: [refer](https://blog.csdn.net/u011228889/article/details/79760167)
            分页太深的解决方案
    2. 相关性算分


部署方式:
    hot & warm 与shard
    roles 角色分布 [refer](https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-node.html)
1. master-eligible node ~ master node
    control the cluster: 
        * 创建/删除 mapping
        * 集群健康检查
        * 指定shard 分配到固定node: 
            分配策略:
                *
2. ... 省略 看doc                
                 