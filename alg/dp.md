1. 01 bag 
n 件物品 , w[] 对应的重量, v[] 对应的值, c 对应空间
```go
    package main
    import "math"
    func bag(n,c int, w,v []int) int {
        common := make([][]int, n+1)	//对应 n c
        capCommon := make([]int, c)
        for j , _ := range capCommon{
            if w[n] < j {
            	capCommon[j] = w[n]
            } else {
            	capCommon[j] = 0
            }
        }
        common[n] = capCommon
        
        for i:= n-1 ; i >0 ; i -- {
        	capEach := make([]int, c)
        	for j , _ := range capEach {
        		if w[j] > j {
        			capEach[j] = common[i+1][j]
        		} else {
        			capEach[j] = int(math.Max(float64(common[i+1][c-w[j]] + v[j]), float64(common[i+1][j])))
        		}
        		common[i] = capEach
        	}
        }
        return common[1][n] 	
    }
```
2. 回文子串
```php 中心拓展法
    //维护一个数组RL[pos] 表示以pos 为中心的 回文串半径长度  
    function longestP(s) {
        $str = implode("#",explode(s, ""));
        $maxLen = $maxRight = $pos = 0 ;
        $RL = [];
        for $i = 0; $i < count($str); $i ++ {
            if($i < $maxRight) {
                $RL[$i] = max($RL[$pos * 2 - $i], maxRight-$i); 
            }
            while(($i+$RL[$i]) < count($str) 
                 && ($i-$RL[$i] > 0) 
                 && str[$i+$RL[$i] +1] == str[$i-$RL[$i] -1]) 
            {
                 $RL[$i] ++;
            }
            if($i + $RL[$i] -1 > maxRight) {
                $pos = $i;
                $maxRight = $i + $RL[$i];
            }
            $maxLen = $maxLen > $RL[$i] ? $maxLen:$RL[$i];
        }
        return $maxLen;
    }
```

[好文](https://segmentfault.com/a/1190000003914228)
```dp
    s[j][i] 表示从j 到i是回文，那么s[j+1][i-1] 也是回文
    function longest(s) {
        $p = [][];
        for($i=0;$i<count(s),$i++) {
            $j = $i
            while($j>=0) {
                if($s[$i] == $s[$j] && (i-j<2 || $P[j+1][i-1])){
                    $p[j][i] = true;
                }
            }
        }
    }

```

3. 不同子序列
https://www.jianshu.com/p/c0935f9d5b5d 
https://zhuanlan.zhihu.com/p/28934358

```php
    //common[i][j] t.sub(0,i) s.sub(0,j) 
    function distinctNum($s,$t) {
        if(count($t) > count($s)) return 0
        $common =[]; //
        foreach($s as $k => $v) {
            $common[$k][0] = 1;
        }
        
        foreach($t as $k => $v) {
            $common[0][$k] = 0;
        }
        
        for($i=1;$i<count($s); $i ++) {
            for($j=1;$j<count(j);$j++) {
                if($s[$i] == $t[$j]) {
                    $common[$i][$j] = $common[$i-1][$j-1] + $common[$i-1][$j]
                }else{
                    $common[$i][$j] = $common[$i-1]                
                }
            }
        }
        return $common[count($t)][count($s)]
    }
```

4. 最大子序列和

```php

```

5. [最小路径和](https://blog.csdn.net/qq_27703417/article/details/70981769)
```php
    function lessSum(arr Array, s int, t int) { //arr 二维数组
        //初始化两边的 到达路劲权重
        common[0][0] = w[0][0]
        for i=1; i<s; i++ {
            common[0][i] = common[0][i-1] + w[0][i]
        }
        for j=1; j<t; j ++ {
            common[j][0] = common[j-1][0] + w[j][0]
        }
        for i=1;i<s;i++ {
            for j=1;j<t;j++ {
                common[j][i] = min(common[j-1][i],common[j][i-1]) + w[j][i] 
            }
        }
        return common[s-1][t-1]
    }
```

6. 不同的二叉搜索树

```php

```

7. 正则表达式匹配[refer](https://hk029.gitbooks.io/leetbook/%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92/010.%20Regular%20Expression%20Matching/010.%20Regular%20Expression%20Matching.html)



8. 最大子序和
如果a[i]是负数，那么它不可能代表最优序列的起点，因为任何包含a[i]的作为起点的子序列都可以通过使用a[i+1]作为起点得到改进。类似的，任何负的子序列也不可能是最优子序列的前缀（原理相同）
```gotemplate
    func(nums []int) int {
        thisSum, max := 0, 0
        for _, i := range nums {
            thisSum += i
            if thisSum > max {
                thisSum = max
            }
            
            if thisSum <= 0 {
                thisSum = 0
            }
        }
        return max
    }
```
9. 乘积最大子序列
    * 前n - 1位的乘积最小子序列的乘积 * A[n]最大
    * 前n - 1位的乘积最大子序列的乘积 * A[n]最大
    * 前n - 1位的乘积最大子序列和乘积最小子序列的乘积 * A[n]都不是最大，而A[n]本身最大
    f_max(n) = max(f_max(n - 1) * A[n], f_min(n - 1) * A[n], A[n])
    f_min(n) = min(f_max(n - 1) * A[n], f_min(n - 1) * A[n], A[n])
    
    dp[i][0] 第i个乘积结尾最小
    dp[i][1] 第i个乘积结尾最大
     
    ```gotemplate
       func maxProduct(nums []int) int {
           if len(nums) == 0 {
               return 0 
           }
           dp := make([]int, len(nums))
           dp[0][0] = nums[0]
           dp[0][1] = nums[0]
           max := nums[0]
           for i:= 1; i < len(nums); i ++ {
               dp[i][0] = min(dp[i-1][0] * nums[i], dp[i-1][1] * nums[i], num[i])
               dp[i][1] = max(dp[i-1][0] * nums[i], dp[i-1][1] * nums[i], num[i])
               if dp[i][1] > max {
                   max = dp[i][1] 
               }   
           }
       }
    ```
    ```php
        class Solution {
            function maxProduct($nums) {
              if(count($nums) == 1 ) return $nums[0];
                    $dp = [];
                    $dp[0][0] = $dp[0][1] = $nums[0];
                    $max = $nums[0];
                    for ($i=1;$i<count($nums);$i++) {
                        $dp[$i][0] = min($dp[$i-1][0] * $nums[$i], $dp[$i-1][1] * $nums[$i], $nums[$i]);
                        $dp[$i][1] = max($dp[$i-1][0] * $nums[$i], $dp[$i-1][1] * $nums[$i], $nums[$i]);
                        $max = max($dp[$i][1], $max);
                    }
                    return $max;
            }
        }
    ```
10. 打家劫舍 I[](http://mzorro.me/2016/03/15/leetcode-house-robber/)
首先是问题分解，假设当前已经肆虐过了前i个房子(0...i-1)，且rob[i]是抢劫了下标为i的房子时的最大收益，pass[i]是不抢劫下标为i的房子时的最大收益，那么可以得到状态转移方程：
rob[i] = nums[i] + pass[i-1]
pass[i] = max(rob[i-1], pass[i-1])

双头指针 i 向右移动 每移动一次准备好 robCur or passCur的 最大当前抢劫值
```php
 function rob($nums) {
        if(count($nums) == 0) return 0;
        if(count($nums) == 1) return $nums[0];
        $lastRob = $nums[0]; $lasPass = 0;
        for($i=1;$i<count($nums);$i++) {
            $robCur = $lasPass + $nums[$i];
            $passCur = max($lasPass, $lastRob);
            $lasPass = $passCur;
            $lastRob = $robCur;
        }
        return max($lasPass,$lastRob);
    }
```
##一是只要遇到字符串的子序列或配准问题首先考虑动态规划DP，二是只要遇到需要求出所有可能情况首先考虑用递归##
关于回溯算法和递归之间的关系
##我们常说的回溯法，其实是一种算法思想，而这种算法思想主要通过递归来实现，并不是说两者等同##
11. 打家劫舍 II
11. 打家劫舍 III 
12. 分割回文串
13. Leetcode 93：[复原IP地址](https://blog.csdn.net/qq_17550379/article/details/82460013)
对于一个已经完成多少，还需完成同样的操作的多少 适合用递归


