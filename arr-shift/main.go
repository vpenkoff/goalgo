package main

import (
	"fmt"
)

/*
 *  A left rotation operation on an array shifts each of the array's elements 1 unit
 *  to the left. For example, if 2 left rotations are performed on array [1, 2, 3, 4, 5],
 *  then the array would become [3, 4, 5, 1, 2]. Note that the lowest index item
 *  moves to the highest index in a rotation. This is called a circular array.
 *
 *  Given an array a of n integers and a number, d, perform d left rotations on the array.
 *  Return the updated array to be printed as a single line of space-separated integers.
 *
 */

func rotateLeft(arr []int, d int) []int {
	n := len(arr)

	for i := 0; i < d; i++ {
		first := arr[0]
		for j := 0; j < n-1; j++ {
			arr[j] = arr[j+1]
		}

		arr[n-1] = first
	}

	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	d := 4

	res := rotateLeft(arr, d)
	fmt.Println(res)
}
