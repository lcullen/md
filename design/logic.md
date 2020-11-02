1. 如何解决高并发写入的问题 [refer 12306](https://github.com/GuoZhaoran/spikeSystem?utm_source=gold_browser_extension)
    * 是否每个请求都是有效的 [refer issue](https://github.com/GuoZhaoran/spikeSystem/issues/4)
        发放token 机制: 根据 session hash 当前qps 1% 放行
    * 吞吐量:
        先返回结果，异步写入
    * 分布式数据一致性:
    
2. 高并发下的库存问题 [refer](https://mp.weixin.qq.com/s/BPvWE4TCHx_X7pEbSNCQFg)
    面临的问题:
        * double request : 分布式过期锁
        * 数据一致性 库存的一致性
        * 秒杀 
    方案:
        1. 最终一致性问题 (如果使用redis 存储库存的信息)
            数据的最终一致性的手段:
                * 直接邮寄
                * 反熵
                * 传播谣言
        2. 
