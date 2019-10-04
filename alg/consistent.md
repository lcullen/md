CAP (consistence available partition)
1. Raft
    要解决的问题:
    + 选举leader 防止会出现脑裂(有两个leader)
        具体实施策略
        * 角色与对应职责(状态与状态维护)
            - follower(低层): 所有节点init的时候的状态
            
                * route write request to leader
                * 维护一个term(相当于mvcc,但是term 越大层级越高) 和 one vote ticket
                    - vote 时机(receive vote request): self.not_yet_vote && vote_request.term >= self.term
                    
                * 维护 heart_beat_timer (rand between 150~300 ms:防止并发发起leader选举,leader选举采用first in first become leader 的策略)
                * heart_beat_timer run out: term += 1 and become candidate 
                
            - candidate(中层):follower=>leader的中间状态
                tips: 如果没有candidate,follower会有比较大的概率处于无限制的竞选成为leader的状态
                * sendRequestVote() 
                
            - leader(高层): 
                * 处理写请求(会在日志接下来的日志复制详细讲)
                * 周期性发送心跳
                
    + 日志复制(参考TCC 2PC) 
        - leader 维护当前 commitLogIndex
        - stepCandidate()
        - stepLeader()
        
    + 安全性