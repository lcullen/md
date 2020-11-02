1. ABCDE 指定 i=2 转换为 DEABC

```gotemplate
    //局部逆序 整体逆序 活用部分逆序的方式
    func Revers(str []string, point int) []string{
      for i:=0,j:=point; i<=j;i++,j-- {
        str[i], str[j] := str[j], str[i] 
      }  
      
      for i:=point+1,j:=len(arr)-1; i<=j;i++,j-- {
        str[i], str[j] := str[j], str[i] 
      }  
      
      for i:=01,j:=len(arr)-1; i<=j;i++,j-- {
             str[i], str[j] := str[j], str[i] 
      } 
    } 
```

2. 拼接最小字典序 --字符串数组
```gotemplate
//将整个 str arr 排序，单独将每个元素看成 一个元素 比较大小的依据是 str1 + str2 > str2 + str1 
//表示 compare 实际上是一个 排序方式

    func GerMinDictionary(str []string) string{
         
    }
    
    func QuicSort(str []string) []string {
        for 
    }
    
    func Partition (str []string, i, j int) {
        initV :=  str[i]
        i0 := i
        for i < j {
            for compare(initV < str[j]) {
                j -- 
            }
            swap(i,j)
            
            for compare(initV > str[i]) {
                i ++
            }    
            swap(i, j)
        }
        swap(i0,i)
        return i;
    }
    func compare(str1, str2 string) {
        return str1+str2 > str2+str1
    }
```
3. 最长无重复子串

brute
```gotemplate
    //eg. abcd    return abcd 4
    // abcb       return abc 3 
    //brute
    maxLen := 0
    for i:=1;i<len(strs);i++ {
        curMaxLen := 0
        for j:=i-1,cur:=strs[i];j > 0; j-- {
           if cur == strs[j] {
                break
           } 
           curMaxLen ++
        } 
        if curMaxLen >= maxLen {
            maxLen = curMaxLen
        }
    }
```
DP
```gotemplate
/*
    链接：https://www.nowcoder.com/questionTerminal/5947ddcc17cb4f09909efa7342780048
"滑动窗口" 
    比方说 abcabccc 当你右边扫描到abca的时候你得把第一个a删掉得到bca，
    然后"窗口"继续向右滑动，每当加到一个新char的时候，左边检查有无重复的char，
    然后如果没有重复的就正常添加，
    有重复的话就左边扔掉一部分（从最左到重复char这段扔掉），在这个过程中记录最大窗口长度
*/
    func SlideWindow (str string) (maxLen int){
        lastLoc := make(map[string]int,0)
        idx,maxLen := -1, maxLen
        for i,char := range str {
            if loc,ok := lastLoc[char]; ok {
                idx := i + 1
            }
            lastLoc[char] = i;
            if i - idx > maxLen {
                    maxLen = i - idx
            } 
        }
        return maxLen
    }
```

3. [最长回文](https://segmentfault.com/a/1190000003914228)
```gotemplate
   func longestPalindrome(s string) {
   	str := strings.Join(strings.Split(s, ""), "#")
   	str = "#" + str + "#"
   	maxRight, pos := 0, 0
   	RL := make([]int, len(str))
   	for i := 0; i < len(str); i++ {
   		if i <= maxRight { //
   			RL[i] = int(math.Min(float64(RL[pos*2-i]), float64(maxRight-i)))
   		}
   		for i+RL[i] < len(str) && i-RL[i] >= 0 && str[i+RL[i]] == str[i-RL[i]] {
   			RL[i]++
   		}
   		if RL[i]+i > maxRight {
   			maxRight = RL[i] + i
   			pos = i
   		}
   	}
   }
```

4. 最大连续子串和
思路 每次新进来的 一个元素要么加入已有的 数组中 要么新起一个新数组并且第一个元素是他
```gotemplate
     func maxSubArray(nums []int) (maxSum int) {
        curSum := 0
        if len(nums) == 0 {
           return 
        }        
        for _, v := range nums {
            if curSum + v > v {
                curSum += v
            }
            if curSum + v > maxSum {
                maxSum = curSum + v
            }
        }
        return 
     }
```

5 不同的子序列（leetcode hard）动态规划
https://www.jianshu.com/p/c0935f9d5b5d 
https://zhuanlan.zhihu.com/p/28934358 
当s 增加一位的时候到第i位的时候，不管此时的s的i位字符和 t的j位字符 是否相等， f(i,j) = f(i-1, j)
但是 如果 i-1 == j 的时候， j 又可以前进一位 和i 进行比较 
```gotemplate
    func numDistinct(s, t string) int {
        common := make(map[int][]int)
        for i:=0; i< len(s); i++ {
            common[i][0] = 1
        }
        for i:=0; i < len(s);i ++ {
            for j := 0; j < len(t);j ++ {
                if s[i] != t[j] {
                    common[i][j] = common[i-1][j]
                }else {
                    common[i][j] = common[i-1][j-1] + common[i-1][j]
                }
            }
        }
        return common[len(s)][len(t)]
    }

```

