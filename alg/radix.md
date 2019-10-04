package RADIX

/*
	基数排序 从个位对每个元素 拉链法 排序 重复 十位操作
	step 1. 先确定 最大数
	step 2
*/
func RadixSort(arr []int) []int {
	size := len(arr)
	if size <= 1 {
		return arr
	}
	max := findMax(arr)
	digit := 1
	for max/digit > 0 {
		bucket := make([]int, 10)
		semSorted := make([]int, size)
		for i := 0; i < size; i++ {
			bucket[(arr[i]/digit)%10]++
		}
		//占位
		for i := 1; i < size; i++ {
			bucket[i] += bucket[i-1]
		}

		for i := 0; i < size; i++ {
			single := (arr[i] / digit) % 10
			bucket[single]--
			semSorted[bucket[single]] = arr[i]
		}

		for i := 0; i < size; i++ {
			arr[i] = semSorted[i]
		}
		digit *= 10
	}
	return []int{}
}

func findMax(arr []int) int {
	if len(arr) == 0 {
		return 0 //actually it should panic
	}
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
