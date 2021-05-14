package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {

	var nums = []int{2, 7, 11, 15}
	target := 2

	got := twoSum(nums, target)
	want := []int{0}

	if got != int {
		t.Errorf("twoSum(nums, target) \n got: %v \n want: \n%v", got, want)
	}
}
