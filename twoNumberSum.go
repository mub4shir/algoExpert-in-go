package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, -4, 8, 11, 1, -1, 6}
	target := 10
	res := twoNumberSum(arr, target)
	res2 := twoNumberSum2(arr, target)
	res3 := twoNumberSum3(arr, target)

	for _, value := range res {
		fmt.Printf("%d ", value)
	}
	fmt.Println()
	for _, value := range res2 {
		fmt.Printf("%d ", value)
	}
	fmt.Println()
	for _, value := range res3 {
		fmt.Printf("%d ", value)
	}
}

// O(nlog(n)) time | O(1) space
func twoNumberSum(arr []int, target int) []int {
	sort.Ints(arr)
	left, right := 0, len(arr)-1
	for left < right {
		currentSum := arr[left] + arr[right]
		if currentSum == target {
			return []int{arr[left], arr[right]}
		} else if currentSum < target {
			left += 1
		} else {
			right += 1
		}
	}
	return []int{}
}

// O(n) time | O(n) space
func twoNumberSum2(arr []int, target int) []int {
	nums := map[int]bool{}
	for _, num := range arr {
		potentialMatch := target - num
		if _, found := nums[potentialMatch]; found {
			result := []int{potentialMatch, num}
			sort.Ints(result)
			return result
		} else {
			nums[num] = true
		}
	}
	return []int{}
}

// O(n^2) time | O(1) space
func twoNumberSum3(arr []int, target int) []int {
	for i := 0; i < len(arr)-1; i++ {
		firstNumber := arr[i]
		for j := i + 1; j < len(arr); j++ {
			secondNumber := arr[j]
			if firstNumber+secondNumber == target {
				result := []int{firstNumber, secondNumber}
				sort.Ints(result)
				return result
			}
		}
	}
	return []int{}
}
