1. map_reduce base [refer](https://mp.weixin.qq.com/s/PcHeYksZ6298HncMats_ZQ)
or [refer](http://zhangtielei.com/posts/blog-hadoop-mapred.html)
    + 切片 file 自定义
    + map
    + shuffle 可预定义聚合
    + reduce 自定义
    
2. hdfs
    NameNode
    DataNode
    
3. [在高并发场景下如何优化分布式锁的并发性能](https://mp.weixin.qq.com/s?__biz=MzU0OTk3ODQ3Ng==&mid=2247483926&idx=1&sn=2a796ef514dea15790e45d79d233833e&chksm=fba6ea15ccd1630387b8738a00a8c1dc6ae0c535305ec4d6e3c76d64eff48bf1d47ae0eaea07&scene=21#wechat_redirect)
    解决方案: 分段加锁 将同一个临界资源分成多个: 当其中的一个临界资源
    缺点: 实现复杂
        1. 其中一部分的临界资源 没有达到要求就要 释放锁
    