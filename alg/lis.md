package main
https://www.nowcoder.com/questionTerminal/585d46a1447b4064b749f08c2ab9ce66

/* 写一个时间复杂度尽可能低的程序，求一个一维数组中最长递增子序列的长度。例如，在序列1,-1,2,-3,4,-5,6,-7中，其最长递增子序列的长度为4(1,2,4,6)。
 */

/*
	fun1 : 假想序列增加一个，从头询问当加入这个节点的时候，每个节点的连接加入节点 的 子序列长度
	step 1: 初始化同长度 array 每个位置 值为 1
*/

func LISN2(s []int) int {
	result := 0
	l := len(s)
	b := make([]int, l)
	//序列中 每加入 一个 i 元素
	for i := 0; i < l; i++ {
		//初始化 b
		b[i] = 1
		for j := 0; j < i; j++ {
			if s[j] < s[i] && b[j]+1 >= b[i] {
				b[i] = b[j] + 1
			}

			if result < b[i] {
				result = b[i]
			}
		}

	}
	return result
}

/*
	维护当前最大子序列 的 最小值 a
	fun2 : 同理假设 序列每增加一个 元素 比a 大 那么就能构成新的最长 递增 子序列
	step 1:
*/

2. 寻找两个有序数组中的中位数
    中位数: 一个可以将数分成两部分的数，其中一部分任何一个数大于另一部分
    ```gotemplate
       
    ```

1. 反转链表 三变量 第一个变量不动,在最后返回之前添加到头结点
    ### 解题思路
    双指针问题
    在原始head之前预定义缺省null节点 pre := nil, 当前节点cur := head 每次改变cur 指针的next 指向pre，改变完之后移动pre、cur 指针，
    并且每次先赋值pre 再赋值cur 防止链表断裂
    
    ### 代码
    
    ```golang
    /**
     * Definition for singly-linked list.
     * type ListNode struct {
     *     Val int
     *     Next *ListNode
     * }
     */
    func reverseList(head *ListNode) *ListNode {
        if head = nil {
            return head
        }
        p1 := head
        p2 := head.Next
        for p2 != nil {
            p3 = p2.Next
            p2.Next = p1
            p1 = p2
            p2 = p3
        }
        return p2
    }
    ```
2. 反转链表II (反转第m ~ n)
    ### 解题思路
    双指针问题
    在原始head之前预定义缺省null节点 pre := nil, 当前节点cur := head 每次改变cur 指针的next 指向pre，改变完之后移动pre、cur 指针，
    并且每次先赋值pre 再赋值cur 防止链表断裂
    
    ### 代码
    
    ```golang
    /**
     * Definition for singly-linked list.
     * type ListNode struct {
     *     Val int
     *     Next *ListNode
     * }
     */
    func reverseList(head *ListNode, m, n int) *ListNode {
        if head == nil {
            return head
        }
        pos, dummy := 1, head
        for pos < m  {
            head = head.Next 
            pos ++
        }
        p1 := head.Next
        for p1.Next != nil && pos < n{
            p2 := p1.Next
            head.Next = p2
            p1.Next = p2.Next
            p2.Next = p1
            
            head = p1
            p1 = p2
            pos ++ 
        }
        return  dummy
    }
    ```

3. 寻找两个正序数组中位数

```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

}
//递归法
func findKthNum(num1, num2 []int, k int) {
	if len(num1) == 0 {
		return num2[k-1]	
	} else if len(num2) == 0 {
		return num1[k-1]	
	}
	
	if k == 1 {
		return 	max(num1[0], num2[0])
	}	
	
	km := k >> 2
	if num1[km] < num2[km] {
		num1 = num1[km:]	
	} else {
		num2 = num2[km:]	
	}
	return findKth(num1, num2, k-km)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

```
137:
只出现一次的数字
给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。

你必须设计并实现线性时间复杂度的算法且不使用额外空间来解决此问题。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/single-number-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```go
func singleNumber(nums []int) int {

}
```
