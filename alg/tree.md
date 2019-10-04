## tree

1. 按层号打印tree
    ####打印的顺序是 父节点先访问，子节点后入queue####
    * 首先把当前root节点赋值为本层last节点
    * 当本层最后一个节点被pop 出的时候 保存下一层 nlast 节点为last 节点
    * nlast 节点的记录方式: 每次push queue 的时候 ++
```gotemplate
    func PrintTreeWithLevel(root *Node) {
        queue.Push(root)
        last, nlast := root, root
        for !queue.Empty() {
            node := queue.Pop()
            visit(node)
            if node.Left != nil {
                queue.Push(node.Left)
                nlast = node.Left
            }
            
            if node.Right != nil {
                queue.Push(node.Right)
                nlast = node.Right
            }
            
            if node == last { //表示当前行访问的是最后的一个节点，并且已经假设当前node的下一个行的最后节点已经通过上述Push into Queue
                fmt.Println()
                last = nlast //保存下一行的最后一个节点
            }
        }
    }
```

-----
1. tree序列化与反序列化

题目：首先我们介绍二叉树先序序列化的方式，假设序列化的结果字符串为str，初始时str等于空字符串。先序遍历二叉树，如果遇到空节点，就在str的末尾加上“#!”，“#”表示这个节点为空，节点值不存在，当然你也可以用其他的特殊字符，“!”表示一个值的结束。如果遇到不为空的节点，假设节点值为3，就在str的末尾加上“3!”。现在请你实现树的先序序列化。

```gotemplate
    func Serialize(root *Node) {
        if root == nil {
            serializeArr =  append(serializeArr, "#!")
            return 
        }
        serializeArr = append(serializeArr, fmt.Sprint("%s!", root.Val))
        Serialize(root.Left)
        Serialize(root.Right)
    }
    
    var serializeArr []string 
    
    var point int 
    func BuildTree(s []string) (root *Node){
        if point >= len(s) return nil 
        if s[point] == "#" {
            return nil 
        } else {
            node := &Node { Val: s[point]} 
            node.Left = BuildTree(s, point++)
            node.Right = BuildTree(s, point++)
        }
    }
```

2. 94 