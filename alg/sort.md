1. 桶排序-计数排序

2. 桶排序-基数排序

3. 几乎有序的长数组,选择对应的排序算法
    * 构建最小堆:从最后一个节点的父节点开始构建，从右往左 依次构建自己在子堆 for i:=n1/2-1; i >= 0; i--
```go
  func ReSort(arr []int) []int{
    var gap = 10
  	n := len(arr)
  	for k:=10;k < n ;k ++ {
  		toBeSortArr := arr[k-gap:k]
  		//heapsort 
  		//参考goland container.heap 
  		//从最后一个节点的父节点 从上到下down 操作
  		n1 := len(toBeSortArr)
  		for i:=n1/2-1; i >= 0; i-- {
  			for {
  				j1 := i*2 + 1  //left
  				j := j1
  				if j1 >= n || j1 < 0 {
  					break
  				}
  
  				if toBeSortArr[j1] < toBeSortArr[j1+1] {
  					j = j1 + 1
  				}
  
  				if toBeSortArr[i] > toBeSortArr[j] {
  					break
  				}
  				swap(toBeSortArr, i, j)
  				i = j
  			}
  		}
  	}
  }
```

4. 把两个有序的数组合并成为一个有序的数组,其中的一个数组的长度为两个数组的和
```gotemplate
    func MergeArr(arr1,arr2 []int) []int {
        biggerLen := len(arr1) - 1
        smalleLen := len(arr2) 
        
        biggerFlag := biggerLen - smalleLen 
        smalleFlag := smalleLen - 1
        for biggerFlag > 0 && smalleFlag > 0 {
           for arr[biggerFlag] > arr[smalleFlag] && biggerFlag >= 0{
                arr[biggerLen] = arr[biggerFlag]
                biggerLen -- 
                biggerFlag --
           }  
           
           for arr[biggerLen] < arr[smalleFlag] && smalleFlag >= 0{
                arr[biggerLen] = arr[smalleFlag]
                biggerLen --
                smalleFlag --
           }
        }
        
        return arr1
    }
```
5. 荷兰国旗问题 3 种数，0，1，2. 将0放在最左边,1中间,2最右边

```gotemplate

```

6. 一个总体有序，部分无序的求需要排序的 最短子数组长度


7. 给定一组数字,求排完序之后相邻两数之间的最大差值

```gotemplate
    //基本的解题思路是桶排序，划分成n等块 再加一个空桶用来找出空桶前后的最大最小值的差距
    var buckt []int
    
    func getMinGap(arr []int) int {
        if len(arr) <= 1 {
            return 0
        }
        buckt = make([]struct{[]int}, len(arr))
        min, max := getMinMax(arr)
        bucktGap := (max-min)/(len(arr)-1)
        for i:=0;i<len(arr);i++{
            j := arr[i]/bucktGag
            if v, ok := buckt[j]; ok {
               v = append(v, arr[i])
            }else {
                buckt[j] = []string{[]int{arr[i]}}
            }
        }
         
        for i := 0; i <= len(arr) ;i++ {
            if len(buck[i]) == 0 {
                return max(buck[i-1])- min(buck[i+1])
            }
        } 
        return 0
    }
```
8. quick sort
```php
    function quickSort($arr, $begin, $end) {
        if(empty($arr) || $begin+1 >= $end) return $arr;
        $init = $arr[$begin];
        $initPos = $begin;
        $begin ++
        while($begin < $end) {
            while($arr[$begin] <= $init) {
                $begin ++;
            }
            swith($arr, $begin, $initPos);
            while($arr[$end]>= $init) {
                $end --;
            }
            switch($arr, $end, $initPos);
        }
        switch($arr, $end, $initPos);
        return $end;
    }
```