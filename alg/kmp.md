package main

import "fmt"

/*
	str
	pattern
step 1. 构建 pattern 自我映射关系 需要额外的空间记录每个字符下标 route := make([]int,len(pattern))
		* 两个指针 j,i 初始化的时候 j=0,i=1
		* i++ 直至 i== len(pattern)边界 或者 pattern[i] = pattern[j], 此时 route[i] = j+1
		* i 继续++ 如果 此时 pattern[i] != pattern[j]; 直至 pattern[i] == pattern[route[--j]] 或者 j=0 ; 如果找到相等 route[i] = j + 1
step 2.
*/
func main() {
	str := "abcd"
	pattern := "bc"
	if len(pattern) == KMP(str, pattern) {
		fmt.Println("有")
	} else {
		fmt.Println("没有")
	}
}

func KMP(str, pattern string) int {
	//flag 表示pattern 的位置
	flag, routeCooked := 0, NextJ(pattern)
	for i := 0; i < len(str); {
		if str[i] == pattern[flag] {
			i++
			flag++
			continue
		}
		//不相等
		//退回到上一个节点比较 如果到头了 i 前进
		if flag = routeCooked[flag]; flag == 0 {
			i++
		}
	}
	return flag
}

var route []int

func NextJ(pattern string) []int {
	route = make([]int, len(pattern))

	if len(pattern) <= 1 {
		return route
	}
	for j, i := 0, 1; i < len(pattern); {
		if pattern[i] == pattern[j] {
			route[i] = j + 1
			j++
			i++
			continue
		}
		//j 后退一步 获取 route 中的值 赋值给
		if j--; j > 0 {
			j = route[j]
		} else {
			//j 退不动的时候 i 前进
			i++
		}
	}
	return route
}
