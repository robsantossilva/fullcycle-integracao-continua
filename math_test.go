package main

import "testing"

func TestSum(t *testing.T) {
	total := sum(15, 15)
	if total != 30 {
		t.Errorf("Sum result is invalid: Result is: %d, expected value: %d", total, 30)
	}
}
