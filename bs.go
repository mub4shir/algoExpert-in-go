package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Buffer(nil, 1e9)
	_out, _ := os.Create("user.out")
	out := bufio.NewWriter(_out)
next:
	for in.Scan() {
		s := string(in.Bytes())
		in.Scan()
		tar, _ := strconv.Atoi(string(in.Bytes()))
		for i, _i := -1, 1; _i < len(s); _i++ {
			i++
			_neg := false
			if s[_i] == '-' {
				_i++
				_neg = true
			}
			v := int(s[_i] & 15)
			for _i++; s[_i]&15 < 10; _i++ {
				v = v*10 + int(s[_i]&15)
			}
			if _neg {
				v = -v
			}
			if v == tar {
				fmt.Fprintln(out, i)
				continue next
			}
			if v > tar {
				break
			}
		}
		fmt.Fprintln(out, -1)
	}
	out.Flush()
	os.Exit(0)
}

func search([]int, int) (_ int) { return }
func mySearch(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		m := i + (j-i)/2
		if target == nums[m] {
			return m
		} else if nums[m] > target {
			j = m - 1
		} else {
			i = m + 1
		}
	}
	return -1
}
