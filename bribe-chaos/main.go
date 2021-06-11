package main

import (
	"fmt"
)

/*
   It is New Year's Day and people are in line for the Wonderland rollercoaster
   ride. Each person wears a sticker indicating their initial position in the
   queue from 1 to n. Any person can bribe the person directly in front of
   them to swap positions, but they still wear their original sticker.
   One person can bribe at most two others.

   Determine the minimum number of bribes that took place to get to a given queue
   order. Print the number of bribes, or, if anyone has bribed more than two
   people, print Too chaotic.
*/

func minimumBribes(q []int) {
	moves := 0
	for i := len(q) - 1; i >= 0; i-- {
		if q[i]-(i+1) > 2 {
			fmt.Println("Too chaotic")
			return
		}

		j := q[i] - 2
		if j < 0 {
			j = 0
		}

		for ; j < i; j++ {
			if q[j] > q[i] {
				moves++
			}
		}
	}
	fmt.Println(moves)
}

func main() {
	//q := []int{2, 1, 5, 3, 4}
	//q := []int{2, 5, 1, 3, 4}
	//q := []int{5, 1, 2, 3, 7, 8, 6, 4}
	q := []int{1, 2, 5, 3, 7, 8, 6, 4}
	//q := []int{1, 2, 5, 3, 4, 7, 8, 6}
	minimumBribes(q)
}
