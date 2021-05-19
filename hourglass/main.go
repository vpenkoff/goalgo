package main

import (
	"fmt"
	"math"
)

/*
 *   Given a 6x6 2D Array:
 *
 *   1 1 1 0 0 0
 *   0 1 0 0 0 0
 *   1 1 1 0 0 0
 *   0 0 0 0 0 0
 *   0 0 0 0 0 0
 *   0 0 0 0 0 0
 *
 *   An hourglass is a subset of values with indices falling in this pattern in the
 *   graphical representation:
 *
 *       a b c
 *         d
 *       e f g
 *
 *   An hourglass sum is the sum of an hourglass' values. Calculate the hourglass
 *   sum for every hourglass, then print the maximum hourglass sum. The array will always be 6x6.
 *
 */

func main() {
	a := [6][6]int{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}

	sum := 0
	max := math.MinInt32

	for k := 0; k < 4; k++ {
		for m := 0; m < 4; m++ {
			for i := k; i < k+3; i++ {
				for j := m; j < m+3 && i != k+1; j++ {
					sum = sum + a[i][j]
				}
			}
			sum = sum + a[k+1][m+1]
			if sum > max {
				max = sum
			}
			sum = 0
		}
	}

	fmt.Printf("Max sum is: %d\n", max)
}
