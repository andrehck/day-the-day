package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	got := diff()
	want := 1

	if got != want {
		t.Errorf("fizzbuzz(3) \n got: %v \n want: \n%v", got, want)
	}
}
