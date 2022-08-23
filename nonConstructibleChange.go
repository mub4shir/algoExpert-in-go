package main

import (
	"fmt"
	"sort"
)

func main() {

	coins := []int{5, 7, 1, 1, 2, 3, 22}
	fmt.Println(nonConstructibleChange(coins))

}

// O(nlog(n)) time | O(1) space
func nonConstructibleChange(coins []int) int {
	currentCreatedChange := 0
	sort.Ints(coins)
	for _, coin := range coins {
		if coin > currentCreatedChange+1 {
			return currentCreatedChange + 1
		}
		currentCreatedChange += coin

	}
	return currentCreatedChange + 1
}
