package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 5, 5, 5, 5, 5, 3, 4, 9}
	fmt.Println(first_and_last(arr, 5))
}

func first_and_last(arr []int, target int) (int, int) {
	sort.Ints(arr)
	for i, v := range arr {
		if v == target {
			start := i
			for (i+1) < len(arr) && arr[i+1] == target {
				i += 1
			}
			return start, i
		}
	}
	return -1, -1
}
