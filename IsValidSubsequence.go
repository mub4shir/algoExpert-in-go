package main

import (
	"fmt"
)

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 10}
	sequence := []int{1, 6, -1, 10}
	fmt.Println(isValidSubsequence(array, sequence))
}

// O(n) time | O(1) space
func isValidSubsequence(array []int, sequence []int) bool {
	arrIdx, seqIdx := 0, 0
	for arrIdx < len(array) && seqIdx < len(sequence) {
		if array[arrIdx] == sequence[seqIdx] {
			seqIdx += 1
		}
		arrIdx += 1
	}
	return seqIdx == len(sequence)

}

// O(n) time | O(1) space
func isValidSubsequence2(array []int, sequence []int) bool {
	seqIdx := 0
	for _, value := range array {
		if seqIdx == len(sequence) {
			break
		}
		if value == sequence[seqIdx] {
			seqIdx += 1
		}
	}
	return seqIdx == len(sequence)
}
