package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	array := []int{-7, -5, -4, 3, 6, 8, 9}
	res := sortedSquaredArray1(array)
	fmt.Println(res)
	res1 := sortedSquaredArray(array)
	fmt.Println(res1)

}

// O(nlog(n)) time | O(n) space
func sortedSquaredArray1(arr []int) []int {
	size := len(arr)
	res := make([]int, size)
	j := 0
	for _, i := range arr {
		res[j] = i * i
		j += 1
	}
	sort.Ints(res)
	return res
}

// O(n) time | O(n) space
func sortedSquaredArray(arr []int) []int {
	res := make([]int, len(arr))
	start, end := 0, len(arr)-1
	for i := len(arr) - 1; i >= 0; i-- {
		if int64(math.Abs(float64(arr[start]))) > int64(math.Abs(float64(arr[end]))) {
			res[i] = arr[start] * arr[start]
			start += 1
		} else {
			res[i] = arr[end] * arr[end]
			end -= 1
		}
	}
	return res
}
