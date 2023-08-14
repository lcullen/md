# 分布式一致性算法

## ZAB

## Paxos
角色: proposer, acceptor, learner

basic paxos 过程: proposer 先发出提案编号 [v] -> acceptor 返回编号是否ok 并且承诺不再接收 比v小的
-> proposer 得到大多数ok 发送 [v, k] -> acceptor 再次校验是否ok

multi-paxos 将acceptor 改为强leader 类型: 读写都经过master node

## Raft (强leader)
选举: 关键词 term, candidate, 
    
日志复制: [wal](https://www.codedump.info/post/20210628-etcd-wal/)
    boltdb:
成员变更: 

## 分布式 强一致性 

两阶段提交: redirect mysql 两阶段提交
TCC: 


