package main

import (
	"fmt"
)

func sort(arr []int) {
	for j := 1; j < len(arr); j++ {
		k := arr[j]
		i := j - 1
		for i >= 0 && arr[i] > k {
			arr[i+1] = arr[i]
			i = i - 1
		}
		arr[i+1] = k
	}
}

func main() {
	nums := []int{31, 41, 59, 26, 41, 58}
	fmt.Println(nums)
	sort(nums)
	fmt.Println(nums)
}
