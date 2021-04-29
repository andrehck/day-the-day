package main

import "fmt"

func main() {
	var nums = []int{2, 7, 11, 15}
	target := 2
	twoSum(nums, target)
	fmt.Println((twoSum(nums, target)))
}

func twoSum(nums []int, target int) []int {
	var hash = make(map[int]int, 0)
	for index, num := range nums {
		if value, ok := hash[target-num]; ok {
			return []int{index, value}
		} else {
			hash[num] = index
		}
	}
	return nums

}
