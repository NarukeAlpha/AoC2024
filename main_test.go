package main

import (
	"testing"
)

func TestDay1(t *testing.T) {
	var testLeftArray = []int{3, 4, 2, 1, 3, 3}
	var testRightArray = []int{4, 3, 5, 3, 9, 3}

	t.Run("Day 1 Part 1 test values", func(t *testing.T) {
		dis := Day1(testLeftArray, testRightArray)
		t.Log("Test run value", dis)
	})

}
