package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	k := 2
	fmt.Println(k >> 1)
	for i, j := 0, 1; i < j; {
		math.Abs()
	}
}

type CQueue struct {
	in, out []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.in = append([]int{value}, this.in...)
}

func (this *CQueue) DeleteHead() int {
	if len(this.out) == 0 {
		this.transIn2Out()
	}
	if len(this.out) == 0 {
		return -1
	}

	val := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return val
}

func (this *CQueue) transIn2Out() {
	for len(this.in) > 0 {
		this.out = append(this.out, this.in[len(this.in)-1])
		this.in = this.in[:len(this.in)-1]
	}
}

// ==== jianzhi 30
type MinStack struct {
	stack, ms []int
}

/** initialize your data structure here. */
func Constructorx() MinStack {
	return MinStack{
		ms: []int{math.MaxInt},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	this.ms = append(this.ms, min(this.Min(), x))
}

func (this *MinStack) Pop() {
	this.ms = this.ms[:len(this.ms)-1]
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.ms[len(this.ms)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

//====findKth

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if allLen := len(nums1) + len(nums2); allLen%2 == 1 {
		return (float64(findKth(nums1, nums2, allLen>>1)) + float64(findKth(nums1, nums2, allLen>>1+1))) / 2
	} else {
		return float64(findKth(nums1, nums2, allLen>>1))
	}
}

func findKth(num1, num2 []int, k int) int {
	if len(num1) == 0 {
		return num2[k-1]
	}
	if len(num2) == 0 {
		return num1[k-1]
	}

	if k == 1 {
		return min(num1[0], num2[0])
	}

	kHalf := k >> 1

	if num1[kHalf] < num2[kHalf] {
		num1 = num1[kHalf:]

	} else {
		num2 = num2[kHalf:]
	}
	return findKth(num1, num2, k-kHalf)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ===next num
func nextPermutation(nums []int) {

	i := len(nums) - 2
	for i >= 0 && nums[i] > nums[i+1] {
		i--
	}

	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[i] < nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	for x, y := i+1, len(nums)-1; x < y; {
		nums[x], nums[y] = nums[y], nums[x]
		x++
		y--
	}
}

func searchRange(nums []int, target int) []int {
	return []int{findFirst(nums, target), findLast(nums, target)}
}

func findFirst(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)>>1
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] == target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	if nums[low] == target {
		return low
	}
	return -1
}

func findLast(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)>>1
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	if nums[low] == target {
		return low
	}
	return -1
}

func combinationSum2(candidates []int, target int) (res [][]int) {
	sort.Ints(candidates)
	cntMap := map[int]int{}
	for _, v := range candidates {
		cntMap[v] += 1
	}

	path := []int{}
	var dfs = func(int, int) {}
	dfs = func(idx, rest int) {
		if idx == len(candidates) || candidates[idx] > rest {
			return
		}

		if rest == 0 {
			res = append(res, [][]int{path}...)
			return
		}
		dfs(idx+1, rest)

		for tryCnt := 1; tryCnt <= cntMap[candidates[idx]] && tryCnt*candidates[idx] < rest; tryCnt++ {
			for i := tryCnt; i > 0; i-- {
				path = append(path, candidates[idx])
			}
			dfs(idx+tryCnt, rest-tryCnt*candidates[idx])
			path = path[:len(path)-tryCnt]
		}
	}
	dfs(0, target)
	return
}
