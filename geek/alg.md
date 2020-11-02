1. 数组与链表的区别:
    数组支持随机读取
    数组为什么从下标0 开始
        其实一开始就是一段连续的内存空间来分配内存的而 内存的下标也就是分配起始位置的偏移量
    GC [refer](https://cloud.tencent.com/developer/article/1072602) 
       [refer](https://github.com/KeKe-Li/For-learning-Go-Tutorial/blob/master/src/spec/02.0.md)
        1. refer counting 引用计数法
            频繁的更新引用计数器、循环引用问题
        2. mark and swap 标记清除 golang && java 都会使用这种方式
            无法避免stw 的情况的出现: 系统时刻在运行会不断的产生新的对象，mark 阶段会变得永无止境，只有stw，mark才是准确没有遗漏的。golang gc 将mark阶段分拆为两个阶段，第一个阶段，mark和用户程序并行运行，第二个阶段，stw，从root，re-scan 并mark 用户并发时候新产生的对象。最后进行swap
        3. generation 分代收集
    
2. LRU: 单链表 热点数据放在头部 [refer](https://github.com/golang/groupcache/blob/master/lru/lru.go) [refer](https://github.com/hashicorp/golang-lru)
   单链表存储字符串，判断字符串是否是回文
   空间O(1) 时间O(1)
   1. 快慢指针找到中心点
   2. 慢指针反转链表
   3. 从中心点扩散判断
   
3. 递归最终推荐人
    问题：
    1. 防止递归太深，导致内存溢出
    2. 防止重复计算
    3. 如果递归产生了环，如何避免无限的递归 __important__ 判断单链表回环问题
    
4. 两个有序的链表合并 归并排序

5. 二分查找
    求一个数的平方根，精确到小数点后六位
删除链表倒数第 n 个结点
求链表的中间结点    
```php
单链表的反转
function revertLinkedList($node) {
        if($node == null || $node->next == null) return ;
        $p1 = $node;
        $p2 = $node->next;
        $p3 = null;
        while($p2 != null) {
            $p3 = $p2->next;
            $p2->next = $p1;
            $p1 = $p2;
            $p2 = $p3;
        }
    }
    
链表中环的检测
function isCircle($node) {
        $slowNode = $quickNode = $node;
        while($quickNode  != null) {
            $slowNode = $slowNode->next;
            if($quickNode->next == null) return false;
            $quickNode = $quickNode->next->next;
            if($slowNode->equals($quickNode)) return true;
        }
    }
```
4. 链表的实际应用
阻塞队列: 生产者 消费者，保证一致性原子消费 compare and swap;惊群问题 
    线程安全的队列成为 并发队列，具体的开源解决方案

4.1 二分法解决 平方根精度问题 记录两个步长之间的 距离要小于精度
4.2 skiplist 插入节点: 
    
5. 快速排序实现topk 问题 && 插入排序问题 && 桶排序 基数排序
6. golang 中 sort 排序相关 解法

7. 朋友关系图
    邻接矩阵 对称存储 无向图 空间换时间
    邻接链表 存储每个顶点出发指向的 顶点链表
    粉丝关系图是一种稀疏矩阵 很浪费空间 所以用邻接矩阵 表示 能够快速查找关注的 人
    求粉丝需要逆向的邻接链表
    如果需要对链表中做快速搜索 可以 使用跳表 红黑树结构
    
8. bfs && dfs find path s->t  [参考](https://mp.weixin.qq.com/s/0BUBhSqmJJxlI_TISsO9xQ)
    都用邻接链表 存储关系
    bfs: 
        queue[i] = linkedList //和Btree的按层关系一致 queue 维护每一层
        isVisit[j] bool // 表示最远可到达 是否已经到达过 只有没有到达过的 节点有资格 被加入 path 和子孙节点的层次入队列
        pre[j] = i // path i -> j
        
    dfs: 回溯
    
10. 字符串匹配的问题
   bm 算法
   kmp 算法
11. tria && ac 自动机 

12. 分治思想(divide and conquer)
    分治是一种思想，递归是一种编码技巧.分治可以用递归来实现
    分治的条件:
        1.问题可以分解成为规模相同的子问题
        2.子问题之间没有关联
        3.分解到一定程度的时候的终止条件
        4.合并子问题的解
    分治的应用:
        1. 考察有序度指标
        2. 最近点对问题 
        3. n*n 两个矩阵相乘 最快速的方式
        
    感悟: 创新并非离我们很远，创新的源泉来自对事物本质的认识。无数优秀架构设计的思想来源都是基础的数据结构和算法，这本身就是算法的一个魅力所在。


13. 回溯
    八皇后问题[参考描述](https://juejin.im/post/5accdb236fb9a028bb195562)
        1. dfs
        2. bfs
    华容道
    
    01背包问题
    
14. 二分查找变形
    一个很容易写错的题目
    给定已经排好序的数组
    ```markdown
        sorted := []int{0, 1, 2, 3, 3, 3, 4, 6}
    	target := 3
	    return pos
    ```
    解题思路:
    设定 low,high,mid
    其中 mid = (high-low) >> 1 + low // 防止取中点时候的溢出；移位操作效率高
    1. 查找第一个等于给定值
            1. mid对应值 大于 target，说明target在mid的左侧，取 high = mid -1 
            2. mid对应值 小于 target，取mid右侧， 取 low = mid + 1 
            3. mid对应值 等于 target，若 mid - 1 != targe 返回mid (考虑边界)
            
    2. 查找第最后一个给定值
            1. mid对应值 大于 target，说明target在mid的左侧，取 high = mid -1 
            2. mid对应值 小于 target，取mid右侧， 取 low = mid + 1 
            3. mid对应值 等于 target，若 mid + 1 != targe 返回mid (考虑边界)
            
    3. 查找第一个大于等于给定值
            1. mid对应值 大于等于 target，
                说明target在mid的左侧或者当前的mid就是target，
                再判断 mid==0(最左边界) 或者 mid -1 对应的值 小于 target 返回 mid
                否则 说明 target 在左边 high = mid -1 
            2. mid对应值 小于 target，target在右侧，取mid右侧， 取 low = mid + 1
             
    4. 查找最后一个小于等于给定值 (与3刚好相反)
            1. mid对应值 小于等于 target，
                说明target在mid的右侧或者当前的mid就是target，
                再判断 mid== len -1 (最右边界) 或者 mid+1 对应的值 大于 target 返回 mid
                否则 说明 target 在右边 low = mid+1 
            2. mid对应值 大于 target，target在左侧，取mid左侧， 取 high = mid-1
            
15. 图的存储
    1. 邻接矩阵: p[i][j] = k 表示定点 i -> j 的路径是否可达，可达的代价是多少
    2. 邻接链表  p[i] = linkedlist i 顶点连接着其他可达顶点 与 BFS 广度优先算法
        当linkedlist 结构很长的时候 就可以使用其他的复杂方式 二叉平衡树

16. 复杂度分
    1. 代码执行行数 && 最长运行次数的代码
    2. 最坏,最好，平均复杂度
        未知的数据集中寻找的代价
       平均复杂度:
        均摊每种可能
    3. 结合代码结构的方式 && 数据的分布
        预测

17. 队列 && 并发队列 && freelock queue 
    一切脱离场景的 都是无稽之谈
    

18. BFS && DFS
    模板:
```python
    def bfs(node, visited): 
        if not node: 
            return 
        queue = []
        visited.add(node)
        queue.push(node) 
        while !queue.is_empty():
            next_node = queue.pop()
            visited.add(next_node)
            queue.push(next_node.childern())
```

```python
    def dfs(node, visited):
       if not node:
            return 
       visited.add(node)
       for childern in node.get_childern():
            if !visited.is_visit():
                visited.add(childern)
                dfs(childern, visited)
       
```
实战:
    按层遍历:
        bfs:
        
        ```python
            def print_by_level(node, level):
                    
        ```

```go
    func printByLevel(q){
    	sz := len(q)
    	for i:= 0 ; i< sz i ++ {
    	  print(q.Pop())
    	  is q.left != nil {
    	  	q.Push(q.left)
    	  }
    	  is q.right != nil {
    	  	q.Push(q.right)
    	  }
    	}
    	print("\n")
    	printByLevel(q)
    }
```

1. 如何理解最长公共子串的长度，将最好的结果传递 并 挤兑到最后一个 M*N 
2. 代码的编译依赖关系:
    邻接矩阵: 全局扫描被依赖的关系， a -> b / c 例如a 是所有bc的前驱 那么a的 依赖度是2 先解决依赖度为 0 的 
3. 拓扑排序