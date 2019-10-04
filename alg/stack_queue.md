1. DFS & BFS

2. stack 逆序函数
```gotemplate
    //无额外空间，利用函数的递归
    // 自我遇到的问题是 容易把 利用非递归的方式 实现递归 方式 总是想要用 while
    
    
    func getBottom (s *Stack) *Item { //有返回的 递归
        item := s.Pop()
        if s.IsEmpty() {
            return item
        }else {
            last := getBottom(s) //阻塞
            s.push(item)
            return last 
        }
    }
    
    func Revers(s *Stack) {//无返回的递归
        if s.IsEmpty() {
            return
        }
        last := getBottom(s)
        Revers(s)
        s.Push(last)
    }
```

2. 给stack 排序，可以申请一个额外 stack 空间
```gotemplate
    var help Stack
    func sortStack(s *Stack) {
        for !s.IsEmpty() {
            for s.Peek() < help.Peek() {
                help.Push(s.Pop) 
            }
            if s.IsEmpty() {
                break
            }
            for help.Peek() < s.Peek() {
                s.Push(help.Pop())
            }
        }
        for !help.IsEmpty() {
            s.Push(help.Pop())
        }
    }
```
queue 
3. 生成窗口最大值数组
 一个固定大小的窗口 每次移动一位得到窗口中的 最大值 形成的 最大值数组
 
思路: 
窗口中 position 小于最大值的 position 都不可能成为最大值 
只有新加入窗口的元素能成为最大值
并且把比他小的都干掉

业务场景:
    按照优先级排队,
    每个队员有不同的权重
    机会窗口有限从左到右滑动
    得到最大权重队员
```gotemplate
    //双向队列
    type IntQueue interface {
        Push(int) 
        PopLast()
        PopFirst() int
        PeekLast() int
        PeekFirst() int
        IsEmpty() bool
    }
    
    func getMaxWindowArr(arr []int, winLen int) []int{
        //base validate ...
        var iq IntQueue
        
        record := make([]int, len(arr) - winLen)
        
        for i:= 0; i< len(arr); i ++ {
           for !iq.isEmpty() && arr[iq.PeekLast()] < arr[i] {
                iq.PopLast()
           }  
           iq.Push(i) //push 的是 index
           
           if iq.PeekFirst() == i - winLen {
                iq.PopFirst()
           }
           
           if i >= winLen {
                record[i-winLen] = iq.PeekFirst()
           }
        }
    }
```

4. 将搜索二叉树转换成双向链表
思路：中序遍历放入队列中

5. 环形单链表约瑟夫问题
思路: